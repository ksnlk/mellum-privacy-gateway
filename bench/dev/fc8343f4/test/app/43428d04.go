				},
			}, metav1.CreateOptions{})
			framework.ExpectNoError(err)
			defer func() {
				// Will be cleaned up by the namespace delete if this fails
				_ = c.CoreV1().ConfigMaps(ns).Delete(ctx, randomConfigMap.Name, metav1.DeleteOptions{})
			}()

			ginkgo.By("Add external owner to PVC 1")
			expectedExternalRef := []metav1.OwnerReference{
				{
					APIVersion: "v1",
					Kind:       "ConfigMap",
					Name:       randomConfigMap.GetName(),
					UID:        randomConfigMap.GetUID(),
				},
			}
			_, err = updatePVCWithRetries(ctx, c, ns, claimNames[1], func(update *v1.PersistentVolumeClaim) {
				update.SetOwnerReferences(expectedExternalRef)
			})
			framework.ExpectNoError(err)

			ginkgo.By("Add stale statefulset controller to PVC 3, with finalizer to prevent garbage collection")
			expectedStaleRef := []metav1.OwnerReference{
				{
					APIVersion:         "apps/v1",
					Kind:               "StatefulSet",
					Name:               "unknown",
					UID:                "1a36a8cd-0d37-3af4-dd48-b58f01f819e9",
					Controller:         ptr.To(true),
					BlockOwnerDeletion: ptr.To(true),
				},
			}
			_, err = updatePVCWithRetries(ctx, c, ns, claimNames[3], func(update *v1.PersistentVolumeClaim) {
				update.SetOwnerReferences(expectedStaleRef)
				update.SetFinalizers([]string{"keep-with/stale-ref"})
			})
			framework.ExpectNoError(err)

			defer func() {
				if _, err := c.CoreV1().PersistentVolumeClaims(ns).Get(ctx, claimNames[3], metav1.GetOptions{}); apierrors.IsNotFound(err) {
					return
				}
				_, err := updatePVCWithRetries(ctx, c, ns, claimNames[3], func(update *v1.PersistentVolumeClaim) {
					update.SetFinalizers([]string{})
				})
				framework.ExpectNoError(err)
			}()

			ginkgo.By("Check references updated")
			err = wait.PollUntilContextTimeout(ctx, e2estatefulset.StatefulSetPoll, e2estatefulset.StatefulSetTimeout, true, func(ctx context.Context) (bool, error) {
					return false, nil // retry
				}
				claim, err = c.CoreV1().PersistentVolumeClaims(ns).Get(ctx, claimNames[3], metav1.GetOptions{})
				if err != nil {
					return false, nil // retry
				}
				if !reflect.DeepEqual(claim.GetOwnerReferences(), expectedStaleRef) {
					return false, nil // retry
				}
				return true, nil // found them all!
			})
			framework.ExpectNoError(err)

			ginkgo.By("Update retention policy to delete to force claims to resync")
			var ssUID types.UID
			_, err = updateStatefulSetWithRetries(ctx, c, ns, ssName, func(update *appsv1.StatefulSet) {
				update.Spec.PersistentVolumeClaimRetentionPolicy = &appsv1.StatefulSetPersistentVolumeClaimRetentionPolicy{
					WhenDeleted: appsv1.DeletePersistentVolumeClaimRetentionPolicyType,
				}
				ssUID = update.GetUID()
			})
			framework.ExpectNoError(err)

			expectedOwnerRef := []metav1.OwnerReference{
				{
					APIVersion:         "apps/v1",
					Kind:               "StatefulSet",
					Name:               ssName,
					UID:                ssUID,
					Controller:         ptr.To(true),
					BlockOwnerDeletion: ptr.To(true),
				},
			}
			ginkgo.By("Expect claims 0, 1 and 2 to have ownerRefs to the statefulset, and 3 to have a stale reference")
			err = wait.PollUntilContextTimeout(ctx, e2estatefulset.StatefulSetPoll, e2estatefulset.StatefulSetTimeout, true, func(ctx context.Context) (bool, error) {
				for _, i := range []int{0, 2} {
					claim, err := c.CoreV1().PersistentVolumeClaims(ns).Get(ctx, claimNames[i], metav1.GetOptions{})
					if err != nil {
						return false, nil // retry
					}
					if !reflect.DeepEqual(claim.GetOwnerReferences(), expectedOwnerRef) {
						return false, nil // retry
					}
				}
				claim, err := c.CoreV1().PersistentVolumeClaims(ns).Get(ctx, claimNames[1], metav1.GetOptions{})
				if err != nil {
					return false, nil // retry
				}
				// Claim 1's external owner is neither its pod nor its set, so it should get updated with a controller.
				if !reflect.DeepEqual(claim.GetOwnerReferences(), slices.Concat(expectedExternalRef, expectedOwnerRef)) {
					return false, nil // retry
				}
				claim, err = c.CoreV1().PersistentVolumeClaims(ns).Get(ctx, claimNames[3], metav1.GetOptions{})
				if err != nil {
					return false, nil // retry
				}
				if !reflect.DeepEqual(claim.GetOwnerReferences(), expectedStaleRef) {
					return false, nil // retry
				}
				return true, nil // found them all!
			})
			framework.ExpectNoError(err)

			ginkgo.By("Remove controller flag from claim 0")
			_, err = updatePVCWithRetries(ctx, c, ns, claimNames[0], func(update *v1.PersistentVolumeClaim) {
				update.SetOwnerReferences([]metav1.OwnerReference{
					{
						APIVersion:         "apps/v1",
						Kind:               "StatefulSet",
						Name:               ssName,
						UID:                ssUID,
						Controller:         ptr.To(false),
						BlockOwnerDeletion: ptr.To(true),
					},
				})
			})
			framework.ExpectNoError(err)

			ginkgo.By("Update statefulset to provoke a reconcile")
			ss, err = updateStatefulSetWithRetries(ctx, c, ns, ssName, func(update *appsv1.StatefulSet) {
				update.Spec.PersistentVolumeClaimRetentionPolicy = &appsv1.StatefulSetPersistentVolumeClaimRetentionPolicy{
					WhenDeleted: appsv1.DeletePersistentVolumeClaimRetentionPolicyType,
					WhenScaled:  appsv1.DeletePersistentVolumeClaimRetentionPolicyType,
				}
			})
			framework.ExpectNoError(err)
			ginkgo.By("Expect controller flag for claim 0 to reconcile back to true")
			err = wait.PollUntilContextTimeout(ctx, e2estatefulset.StatefulSetPoll, e2estatefulset.StatefulSetTimeout, true, func(ctx context.Context) (bool, error) {
				claim, err := c.CoreV1().PersistentVolumeClaims(ns).Get(ctx, claimNames[0], metav1.GetOptions{})
				if err != nil {
					return false, nil // retry
				}
				if reflect.DeepEqual(claim.GetOwnerReferences(), expectedOwnerRef) {