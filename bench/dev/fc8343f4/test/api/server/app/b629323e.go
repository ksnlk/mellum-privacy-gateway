	}

	updatedBytes, err := json.MarshalIndent(updatedObject, "\t", "\t")
	if err != nil {
		t.Fatalf("Failed to marshal updated object: %v", err)
	}

	if createdAccessor.GetResourceVersion() != updatedAccessor.GetResourceVersion() {
		t.Fatalf("Expected same resource version to be %v but got: %v\nold object:\n%v\nnew object:\n%v",
			createdAccessor.GetResourceVersion(),
			updatedAccessor.GetResourceVersion(),
			string(createdBytes),
			string(updatedBytes),
		)
	}
}

// TestNoOpApplyWithEmptyMap
func TestNoOpApplyWithEmptyMap(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	deploymentName := "no-op"
	deploymentsResource := "deployments"
	deploymentBytes := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "` + deploymentName + `",
			"labels": {
				"app": "nginx"
			}
		},
		"spec": {
			"replicas": 1,
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"annotations": {},
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [{
						"name": "nginx",
						"image": "nginx:1.14.2",
		Param("fieldManager", "apply_test").
		Resource("configmaps").
		Name("myconfig").
		Body(ssBytes).
		Do(context.TODO()).
		Get()
	if err != nil {
		t.Fatalf("Failed to create object: %v", err)
	}

	require.NoError(t, err)
	require.Equal(t, obj, newObj)
}

// TestNoSemanticUpdateAppleSameResourceVersion makes sure that APPLY requests which makes no semantic changes
// will not change the resource version (no write to etcd is done)
//
// Some of the non-semantic changes are:
// - Applying an atomic struct that removes a default
// - Changing Quantity or other fields that are normalized
func TestNoSemanticUpdateApplySameResourceVersion(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	ssBytes := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "StatefulSet",
		"metadata": {
			"name": "nginx",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"serviceName": "nginx",
			"selector": { "matchLabels": {"app": "nginx"}},
			"template": {
				"metadata": {
					"labels": {"app": "nginx"}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx",
						"resources": {
							"limits": {"memory": "2048Mi"}
						}
					}]
				}
			},
			"volumeClaimTemplates": [{
				"metadata": {"name": "nginx"},
				"spec": {
		Do(context.TODO()).
		Get()
	if err != nil {
		t.Fatalf("Failed to create object: %v", err)
	}
	rvApplied, err := getRV(obj)
	if err != nil {
		t.Fatalf("Failed to get RV: %v", err)
	}
	if rvApplied != rvCreated {
		t.Fatal("ResourceVersion changed after apply")
	}
}

// TestNoSemanticUpdateAppleSameResourceVersion makes sure that PUT requests which makes no semantic changes
// will not change the resource version (no write to etcd is done)
//
// Some of the non-semantic changes are:
// - Applying an atomic struct that removes a default
// - Changing Quantity or other fields that are normalized
func TestNoSemanticUpdatePutSameResourceVersion(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	ssBytes := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "StatefulSet",
		"metadata": {
			"name": "nginx",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"serviceName": "nginx",
			"selector": { "matchLabels": {"app": "nginx"}},
			"template": {
				"metadata": {
					"labels": {"app": "nginx"}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx",
						"resources": {
							"limits": {"memory": "2048Mi"}
						}
					}]
				}
			},
			"volumeClaimTemplates": [{
				"metadata": {"name": "nginx"},
				"spec": {
			"apiVersion": "v1",
			"kind": "Pod",
			"metadata": {
				"name": "test-pod-uid",
				"uid":  "22d84327-3e2b-00d9-58d2-c9a6bfd00452"
			},
			"spec": {
				"containers": [{
					"name":  "test-container",
					"image": "test-image"
				}]
			}
		}`)).
		Do(context.TODO()).
		Get()
	if !apierrors.IsConflict(err) {
		t.Fatalf("Expected conflict error but got: %v", err)
	}
}

func TestApplyUpdateApplyConflictForced(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	obj := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"replicas": 3,
			"selector": {
				"matchLabels": {
					 "app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest"
					}]
				}
			}
		Do(context.TODO()).
		Get()
	if err == nil {
		t.Fatalf("Apply should fail to create without fieldManager")
	}

	_, err = client.CoreV1().RESTClient().Patch(types.ApplyPatchType).
		Namespace("default").
		Resource("configmaps").
		Name("test-cm").
		Param("fieldManager", "apply_test").
		Body(obj).
		Do(context.TODO()).
		Get()
	if err != nil {
		t.Fatalf("Apply failed to create with fieldManager: %v", err)
	}
}

// TestApplyRemoveContainerPort removes a container port from a deployment
func TestApplyRemoveContainerPort(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	obj := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"replicas": 3,
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest",
						"ports": [{
							"containerPort": 80,
							"protocol": "TCP"
						}]
					}]
				}
			}
		}
	}`)

	_, err := client.CoreV1().RESTClient().Patch(types.ApplyPatchType).
		AbsPath("/apis/apps/v1").
		Namespace("default").
		Resource("deployments").
		Name("deployment").
		Param("fieldManager", "apply_test").
		Body(obj).Do(context.TODO()).Get()
	if err != nil {
		t.Fatalf("Failed to create object using Apply patch: %v", err)
	}

	obj = []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"replicas": 3,
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest"
					}]
				}
			}
		Resource("deployments").
		Name("deployment").
		Param("fieldManager", "apply_test").
		Body(obj).Do(context.TODO()).Get()
	if err != nil {
		t.Fatalf("Failed to remove container port using Apply patch: %v", err)
	}

	deployment, err := client.AppsV1().Deployments("default").Get(context.TODO(), "deployment", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to retrieve object: %v", err)
	}

	if len(deployment.Spec.Template.Spec.Containers[0].Ports) > 0 {
		t.Fatalf("Expected no container ports but got: %v, object: \n%#v", deployment.Spec.Template.Spec.Containers[0].Ports, deployment)
	}
}

// TestApplyFailsWithVersionMismatch ensures that a version mismatch between the
// patch object and the live object will error
func TestApplyFailsWithVersionMismatch(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	obj := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"replicas": 3,
			"selector": {
				"matchLabels": {
					 "app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest"
					}]
				}
			}
		AbsPath("/apis/apps/v1").
		Namespace("default").
		Resource("deployments").
		Name("deployment").
		Param("fieldManager", "apply_test").
		Body([]byte(obj)).Do(context.TODO()).Get()
	if err == nil {
		t.Fatalf("Expecting to get version mismatch when applying object")
	}
	status, ok := err.(*apierrors.StatusError)
	if !ok {
		t.Fatalf("Expecting to get version mismatch as API error")
	}
	if status.Status().Code != http.StatusBadRequest {
		t.Fatalf("expected status code to be %d but was %d", http.StatusBadRequest, status.Status().Code)
	}
}

// TestApplyConvertsManagedFieldsVersion checks that the apply
// converts the API group-version in the field manager
func TestApplyConvertsManagedFieldsVersion(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	obj := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment",
			"labels": {"app": "nginx"},
			"managedFields": [
				{
					"manager": "sidecar_controller",
					"operation": "Apply",
					"apiVersion": "extensions/v1beta1",
					"fieldsV1": {
						"f:metadata": {
							"f:labels": {
								"f:sidecar_version": {}
							}
						},
						"f:spec": {
							"f:template": {
								"f: spec": {
									"f:containers": {
										"k:{\"name\":\"sidecar\"}": {
											".": {},
											"f:image": {}
										}
									}
								}
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest"
					}]
				}
			}
		}
	}`)

	_, err := client.CoreV1().RESTClient().Post().
		AbsPath("/apis/apps/v1").
		Namespace("default").
		Resource("deployments").
		Body(obj).Do(context.TODO()).Get()
	if err != nil {
		t.Fatalf("Failed to create object: %v", err)
	}

	obj = []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment",
			"labels": {"sidecar_version": "release"}
		},
		"spec": {
			"template": {
				"spec": {
					"containers": [{
						"name":  "sidecar",
						"image": "sidecar:latest"
					}]
				}
			}
		}
	}`)
	_, err = client.CoreV1().RESTClient().Patch(types.ApplyPatchType).
		AbsPath("/apis/apps/v1").
		Namespace("default").
		Resource("deployments").
		Name("deployment").
		Param("fieldManager", "sidecar_controller").
		Body([]byte(obj)).Do(context.TODO()).Get()
	if err != nil {
		t.Fatalf("Failed to apply object: %v", err)
	}

	object, err := client.AppsV1().Deployments("default").Get(context.TODO(), "deployment", metav1.GetOptions{})
	if err != nil {
		t.Fatalf("Failed to retrieve object: %v", err)
	}

	accessor, err := meta.Accessor(object)
	if err != nil {
		t.Fatalf("Failed to get meta accessor: %v", err)
	}

	managed := accessor.GetManagedFields()
	if len(managed) != 2 {
		t.Fatalf("Expected 2 field managers, but got managed fields: %v", managed)
	}

	var actual *metav1.ManagedFieldsEntry
	for i := range managed {
		entry := &managed[i]
		if entry.Manager == "sidecar_controller" && entry.APIVersion == "apps/v1" {
			actual = entry
		}
	}

	if actual == nil {
		t.Fatalf("Expected managed fields to contain entry with manager '%v' with converted api version '%v', but got managed fields:\n%v", "sidecar_controller", "apps/v1", managed)
	}

	expected := &metav1.ManagedFieldsEntry{
		Manager:    "sidecar_controller",
		Operation:  metav1.ManagedFieldsOperationApply,
		APIVersion: "apps/v1",
		Time:       actual.Time,
		FieldsType: "FieldsV1",
		FieldsV1: &metav1.FieldsV1{
			Raw: []byte(`{"f:metadata":{"f:labels":{"f:sidecar_version":{}}},"f:spec":{"f:template":{"f:spec":{"f:containers":{"k:{\"name\":\"sidecar\"}":{".":{},"f:image":{},"f:name":{}}}}}}}`),
		},
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected:\n%v\nbut got:\n%v", expected, actual)
	}
}

// TestClearManagedFieldsWithMergePatch verifies it's possible to clear the managedFields
func TestClearManagedFieldsWithMergePatch(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	_, err := client.CoreV1().RESTClient().Patch(types.ApplyPatchType).
		Namespace("default").
		Resource("configmaps").
		Name("test-cm").
		Param("fieldManager", "apply_test").
		Body([]byte(`{
			"apiVersion": "v1",
			"kind": "ConfigMap",
	}

	accessor, err := meta.Accessor(object)
	if err != nil {
		t.Fatalf("Failed to get meta accessor: %v", err)
	}

	if managedFields := accessor.GetManagedFields(); len(managedFields) != 0 {
		t.Fatalf("Failed to stop tracking managedFields, got: %v", managedFields)
	}

	if labels := accessor.GetLabels(); len(labels) < 1 {
		t.Fatalf("Expected other fields to stay untouched, got: %v", object)
	}
}

// TestApplyUnsetExclusivelyOwnedFields verifies that when owned fields are omitted from an applied
// configuration, and no other managers own the field, it is removed.
func TestApplyUnsetExclusivelyOwnedFields(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	// spec.replicas is a optional, defaulted field
	// spec.template.spec.hostname is an optional, non-defaulted field
	apply := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment-exclusive-unset",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"replicas": 3,
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"hostname": "test-hostname",
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest"
					}]
				}
			}
		}
	}`)

	_, err := client.CoreV1().RESTClient().Patch(types.ApplyPatchType).
		AbsPath("/apis/apps/v1").
		Namespace("default").
		Resource("deployments").
		Name("deployment-exclusive-unset").
		Param("fieldManager", "apply_test").
		Body(apply).
		Do(context.TODO()).
		Get()
	if err != nil {
		t.Fatalf("Failed to create object using Apply patch: %v", err)
	}

	// unset spec.replicas and spec.template.spec.hostname
	apply = []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment-exclusive-unset",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest"
					}]
				}
			}
		}
	if err != nil {
		t.Fatalf("Failed to create object using Apply patch: %v", err)
	}

	deployment, ok := patched.(*appsv1.Deployment)
	if !ok {
		t.Fatalf("Failed to convert response object to Deployment")
	}
	if *deployment.Spec.Replicas != 1 {
		t.Errorf("Expected deployment.spec.replicas to be 1 (default value), but got %d", deployment.Spec.Replicas)
	}
	if len(deployment.Spec.Template.Spec.Hostname) != 0 {
		t.Errorf("Expected deployment.spec.template.spec.hostname to be unset, but got %s", deployment.Spec.Template.Spec.Hostname)
	}
}

// TestApplyUnsetSharedFields verifies that when owned fields are omitted from an applied
// configuration, but other managers also own the field, is it not removed.
func TestApplyUnsetSharedFields(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	// spec.replicas is a optional, defaulted field
	// spec.template.spec.hostname is an optional, non-defaulted field
	apply := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment-shared-unset",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"replicas": 3,
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"hostname": "test-hostname",
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest"
					}]
				}
			}
		}
	}`)

	for _, fieldManager := range []string{"shared_owner_1", "shared_owner_2"} {
		_, err := client.CoreV1().RESTClient().Patch(types.ApplyPatchType).
			AbsPath("/apis/apps/v1").
			Namespace("default").
			Resource("deployments").
			Name("deployment-shared-unset").
			Param("fieldManager", fieldManager).
			Body(apply).
			Do(context.TODO()).
			Get()
		if err != nil {
			t.Fatalf("Failed to create object using Apply patch: %v", err)
		}
	}

	// unset spec.replicas and spec.template.spec.hostname
	apply = []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment-shared-unset",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest"
					}]
				}
			}
		}
	if err != nil {
		t.Fatalf("Failed to create object using Apply patch: %v", err)
	}

	deployment, ok := patched.(*appsv1.Deployment)
	if !ok {
		t.Fatalf("Failed to convert response object to Deployment")
	}
	if *deployment.Spec.Replicas != 3 {
		t.Errorf("Expected deployment.spec.replicas to be 3, but got %d", deployment.Spec.Replicas)
	}
	if deployment.Spec.Template.Spec.Hostname != "test-hostname" {
		t.Errorf("Expected deployment.spec.template.spec.hostname to be \"test-hostname\", but got %s", deployment.Spec.Template.Spec.Hostname)
	}
}

// TestApplyCanTransferFieldOwnershipToController verifies that when an applier creates an
// object, a controller takes ownership of a field, and the applier
// then omits the field from its applied configuration, that the field value persists.
func TestApplyCanTransferFieldOwnershipToController(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	// Applier creates a deployment with replicas set to 3
	apply := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment-shared-map-item-removal",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"replicas": 3,
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest",
					}]
				}
			}
		Resource("deployments").
		Name("deployment-shared-map-item-removal").
		Param("fieldManager", "test_applier").
		Body(apply).
		Do(context.TODO()).
		Get()
	if err != nil {
		t.Fatalf("Failed to create object using Apply patch: %v", err)
	}

	// a controller takes over the replicas field
	applied, ok := appliedObj.(*appsv1.Deployment)
	if !ok {
		t.Fatalf("Failed to convert response object to Deployment")
	}
	replicas := int32(4)
	applied.Spec.Replicas = &replicas
	_, err = client.AppsV1().Deployments("default").
		Update(context.TODO(), applied, metav1.UpdateOptions{FieldManager: "test_updater"})
	if err != nil {
		t.Fatalf("Failed to create object using Apply patch: %v", err)
	}

	// applier omits replicas
	apply = []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment-shared-map-item-removal",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest",
					}]
				}
			}
		}
		Do(context.TODO()).
		Get()
	if err != nil {
		t.Fatalf("Failed to create object using Apply patch: %v", err)
	}

	// ensure the container is deleted even though a controller updated a field of the container
	deployment, ok := patched.(*appsv1.Deployment)
	if !ok {
		t.Fatalf("Failed to convert response object to Deployment")
	}
	if *deployment.Spec.Replicas != 4 {
		t.Errorf("Expected deployment.spec.replicas to be 4, but got %d", deployment.Spec.Replicas)
	}
}

// TestApplyCanRemoveMapItemsContributedToByControllers verifies that when an applier creates an
// object, a controller modifies the contents of the map item via update, and the applier
// then omits the item from its applied configuration, that the item is removed.
func TestApplyCanRemoveMapItemsContributedToByControllers(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	// Applier creates a deployment with a name=nginx container
	apply := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment-shared-map-item-removal",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest",
					}]
				}
			}
		}
		Namespace("default").
		Resource("deployments").
		Name("deployment-shared-map-item-removal").
		Param("fieldManager", "test_applier").
		Body(apply).
		Do(context.TODO()).
		Get()
	if err != nil {
		t.Fatalf("Failed to create object using Apply patch: %v", err)
	}

	// a controller sets container.workingDir of the name=nginx container via an update
	applied, ok := appliedObj.(*appsv1.Deployment)
	if !ok {
		t.Fatalf("Failed to convert response object to Deployment")
	}
	applied.Spec.Template.Spec.Containers[0].WorkingDir = "/home/replacement"
	_, err = client.AppsV1().Deployments("default").
		Update(context.TODO(), applied, metav1.UpdateOptions{FieldManager: "test_updater"})
	if err != nil {
		t.Fatalf("Failed to create object using Apply patch: %v", err)
	}

	// applier removes name=nginx the container
	apply = []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment-shared-map-item-removal",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"replicas": 3,
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"hostname": "test-hostname",
					"containers": [{
						"name":  "other-container",
						"image": "nginx:latest",
					}]
				}
		Get()
	if err != nil {
		t.Fatalf("Failed to create object using Apply patch: %v", err)
	}

	// ensure the container is deleted even though a controller updated a field of the container
	deployment, ok := patched.(*appsv1.Deployment)
	if !ok {
		t.Fatalf("Failed to convert response object to Deployment")
	}
	if len(deployment.Spec.Template.Spec.Containers) != 1 {
		t.Fatalf("Expected 1 container after apply, got %d", len(deployment.Spec.Template.Spec.Containers))
	}
	if deployment.Spec.Template.Spec.Containers[0].Name != "other-container" {
		t.Fatalf("Expected container to be named \"other-container\" but got %s", deployment.Spec.Template.Spec.Containers[0].Name)
	}
}

// TestDefaultMissingKeys makes sure that the missing keys default is used when merging.
func TestDefaultMissingKeys(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	// Applier creates a deployment with containerPort but no protocol
	apply := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment-shared-map-item-removal",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest",
						"ports": [{
							"name": "foo",
							"containerPort": 80
						}]
					}]
				}
			}
		}
	}`)

	_, err := client.CoreV1().RESTClient().Patch(types.ApplyPatchType).
		AbsPath("/apis/apps/v1").
		Namespace("default").
		Resource("deployments").
		Name("deployment-shared-map-item-removal").
		Param("fieldManager", "test_applier").
		Body(apply).
		Do(context.TODO()).
		Get()
	if err != nil {
		t.Fatalf("Failed to create object using Apply patch: %v", err)
	}

	// Applier updates the name, and uses the protocol, we should get a conflict.
	apply = []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment-shared-map-item-removal",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"selector": {
				"matchLabels": {
					"app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest",
						"ports": [{
							"name": "bar",
							"containerPort": 80,
							"protocol": "TCP"
	if err == nil {
		t.Fatalf("Expecting to get conflicts when a different applier updates existing list item, got no error: %s", patched)
	}
	status, ok := err.(*apierrors.StatusError)
	if !ok {
		t.Fatalf("Expecting to get conflicts as API error")
	}
	if len(status.Status().Details.Causes) != 1 {
		t.Fatalf("Expecting to get one conflict when a different applier updates existing list item, got: %v", status.Status().Details.Causes)
	}
}

var podBytes = []byte(`
apiVersion: v1
kind: Pod
metadata:
  labels:
    app: some-app
    plugin1: some-value
    plugin2: some-value
    plugin3: some-value
    plugin4: some-value
  name: some-name
  namespace: default
  ownerReferences:
  - apiVersion: apps/v1
    blockOwnerDeletion: true
    controller: true
    kind: ReplicaSet
    name: some-name
    uid: 2d4a2e7c-448d-22f3-a753-79276f5741cc
spec:
  containers:
  - args:
    - one
    - two
    - three
    - four
    - five
    - six
    - seven
    - eight
    - nine
    env:
    - name: VAR_3
      valueFrom:
        secretKeyRef:
          key: some-other-key
          name: some-oher-name
    - name: VAR_2
      valueFrom:
	}
}

func benchRepeatedUpdate(client clientset.Interface, podName string) func(*testing.B) {
	return func(b *testing.B) {
		b.ResetTimer()
		b.ReportAllocs()
		for i := 0; i < b.N; i++ {
			_, err := client.CoreV1().RESTClient().Patch(types.JSONPatchType).
				Namespace("default").
				Resource("pods").
				Name(podName).
				Body([]byte(fmt.Sprintf(`[{"op": "replace", "path": "/spec/containers/0/image", "value": "image%d"}]`, i))).Do(context.TODO()).Get()
			if err != nil {
				b.Fatalf("Failed to patch object: %v", err)
			}
		}
	}
}

func TestUpgradeClientSideToServerSideApply(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	obj := []byte(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-deployment
  annotations:
    "kubectl.kubernetes.io/last-applied-configuration": |
      {"kind":"Deployment","apiVersion":"apps/v1","metadata":{"name":"my-deployment","labels":{"app":"my-app"}},"spec":{"replicas": 3,"template":{"metadata":{"labels":{"app":"my-app"}},"spec":{"containers":[{"name":"my-c","image":"my-image"}]}}}}
  labels:
    app: my-app
spec:
  replicas: 100000
  selector:
    matchLabels:
      app: my-app
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - name: my-c
        image: my-image
`)

	deployment, err := yamlutil.ToJSON(obj)
	if err != nil {
		t.Fatalf("Failed marshal yaml: %v", err)
	}

	_, err = client.CoreV1().RESTClient().Post().
		AbsPath("/apis/apps/v1").
		Namespace("default").
		Resource("deployments").
		Body(deployment).Do(context.TODO()).Get()
	if err != nil {
		t.Fatalf("Failed to create object: %v", err)
	}

	obj = []byte(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-deployment
  labels:
    app: my-new-label
spec:
  replicas: 3 # expect conflict
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - name: my-c
        image: my-image
`)

	deployment, err = yamlutil.ToJSON(obj)
	if err != nil {
		t.Fatalf("Failed marshal yaml: %v", err)
	}

	_, err = client.CoreV1().RESTClient().Patch(types.ApplyPatchType).
		AbsPath("/apis/apps/v1").
		Namespace("default").
		Resource("deployments").
		Name("my-deployment").
		Param("fieldManager", "kubectl").
		Body(deployment).
		Do(context.TODO()).
		Get()
	if !apierrors.IsConflict(err) {
		t.Fatalf("Expected conflict error but got: %v", err)
	}

	obj = []byte(`
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-deployment
  labels:
    app: my-new-label
spec:
  template:
    metadata:
      labels:
        app: my-app
    spec:
      containers:
      - name: my-c
        image: my-image-new
`)

	deployment, err = yamlutil.ToJSON(obj)
	if err != nil {
		t.Fatalf("Failed marshal yaml: %v", err)
	}

	_, err = client.CoreV1().RESTClient().Patch(types.ApplyPatchType).
		AbsPath("/apis/apps/v1").
		Namespace("default").
		Resource("deployments").
		Get()
	if err != nil {
		t.Fatalf("Failed to update: %v", err)
	}

	accessor, err := meta.Accessor(obj)
	if err != nil {
		t.Fatalf("Failed to get meta accessor for object: %v", err)
	}
	managedFields = accessor.GetManagedFields()
	if len(managedFields) != 2 {
		t.Fatalf("Expected object to have 2 managed fields entries, got: %d", len(managedFields))
	}
	entry := managedFields[1]
	if entry.Manager != "manager" || entry.Operation != "Update" || string(entry.FieldsV1.Raw) != `{"f:status":{"f:phase":{}}}` {
		t.Fatalf(`Unexpected entry, got: %v`, entry)
	}
}

func TestSubresourceField(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	// Creating a deployment
	deploymentBytes := []byte(`{
		"apiVersion": "apps/v1",
		"kind": "Deployment",
		"metadata": {
			"name": "deployment",
			"labels": {"app": "nginx"}
		},
		"spec": {
			"replicas": 3,
			"selector": {
				"matchLabels": {
					 "app": "nginx"
				}
			},
			"template": {
				"metadata": {
					"labels": {
						"app": "nginx"
					}
				},
				"spec": {
					"containers": [{
						"name":  "nginx",
						"image": "nginx:latest"
					}]
				}
			}

	err = json.Unmarshal(expectedManagedFieldsString, &expectedManagedFields)
	if err != nil {
		t.Fatalf("unexpectly failed to decode expected managed fields")
	}

	// Wipe timestamps before comparison
	for i := range expectedManagedFields {
		expectedManagedFields[i].Time = nil
	}

	for i := range managedFields {
		managedFields[i].Time = nil
	}

	if !reflect.DeepEqual(expectedManagedFields, managedFields) {
		t.Fatalf("unexpected managed fields: %v", cmp.Diff(expectedManagedFields, managedFields))
	}
}

func TestDuplicatesInAssociativeLists(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()

	ds := []byte(`{
  "apiVersion": "apps/v1",
  "kind": "DaemonSet",
  "metadata": {
    "name": "example-daemonset",
    "labels": {
      "app": "example"
    }
  },
  "spec": {
    "selector": {
      "matchLabels": {
        "app": "example"
      }
    },
    "template": {
      "metadata": {
        "labels": {
          "app": "example"
        }
      },
      "spec": {
        "containers": [
          {
            "name": "nginx",
            "image": "nginx",
            "ports": [
	// Create the object
	obj, err := client.AppsV1().RESTClient().
		Post().
		Namespace("default").
		Param("fieldManager", "create").
		Resource("daemonsets").
		Body(ds).
		Do(context.TODO()).
		Get()
	if err != nil {
		t.Fatalf("Failed to create the object: %v", err)
	}
	daemon := obj.(*appsv1.DaemonSet)
	if want, got := 3, len(daemon.Spec.Template.Spec.Containers[0].Env); want != got {
		t.Fatalf("Expected %v EnvVars, got %v", want, got)
	}
	if want, got := 3, len(daemon.Spec.Template.Spec.Containers[0].Ports); want != got {
		t.Fatalf("Expected %v Ports, got %v", want, got)
	}

	expectManagedFields(t, daemon.ManagedFields, `
[
	{
		"manager": "create",
		"operation": "Update",
		"apiVersion": "apps/v1",
		"time": null,
		"fieldsType": "FieldsV1",
		"fieldsV1": {
			"f:metadata": {
				"f:annotations": {
					".": {},
					"f:deprecated.daemonset.template.generation": {}
				},
				"f:labels": {
					".": {},
					"f:app": {}
				}
			},
			"f:spec": {
			"f:revisionHistoryLimit": {},
			"f:selector": {},
			"f:template": {
				"f:metadata": {
					"f:labels": {
						".": {},
						"f:app": {}
					}
				},
				"f:spec": {
					"f:containers": {
						"f:terminationMessagePolicy": {}
						}
					},
					"f:dnsPolicy": {},
					"f:restartPolicy": {},
					"f:schedulerName": {},
					"f:securityContext": {},
					"f:terminationGracePeriodSeconds": {}
					}
				},
				"f:updateStrategy": {
					"f:rollingUpdate": {
					".": {},
					"f:maxSurge": {},
					"f:maxUnavailable": {}
					},
					"f:type": {}
				}
			}
		}
	}
]`)

	// Apply unrelated fields, fieldmanager should be strictly additive.
	ds = []byte(`
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: example-daemonset
  labels:
    app: example
spec:
  selector:
    matchLabels:
      app: example
  template:
    spec:
      containers:
        - name: nginx
          image: nginx
          ports:
            - name: port3
              containerPort: 443
          env:
            - name: HOME
              value: "/usr/home"
`)
	obj, err = client.AppsV1().RESTClient().
		Patch(types.ApplyPatchType).
		Namespace("default").
		Name("example-daemonset").
		Param("fieldManager", "apply").
		Resource("daemonsets").
		Body(ds).
		Do(context.TODO()).
		Get()
	if err != nil {
		t.Fatalf("Failed to aply the first object: %v", err)
	}
	daemon = obj.(*appsv1.DaemonSet)
	if want, got := 4, len(daemon.Spec.Template.Spec.Containers[0].Env); want != got {
		t.Fatalf("Expected %v EnvVars, got %v", want, got)
	}
	if want, got := 4, len(daemon.Spec.Template.Spec.Containers[0].Ports); want != got {
		t.Fatalf("Expected %v Ports, got %v", want, got)
	}

	expectManagedFields(t, daemon.ManagedFields, `
[
	{
		"manager": "apply",
		"operation": "Apply",
		"apiVersion": "apps/v1",
		"time": null,
		"fieldsType": "FieldsV1",
		"fieldsV1": {
			"f:metadata": {
				"f:labels": {
					"f:app": {}
				}
			},
			"f:spec": {
				"f:selector": {},
				"f:template": {
					"f:spec": {
						"f:containers": {
							"k:{\"name\":\"nginx\"}": {
								".": {},
								"f:env": {
									"k:{\"name\":\"HOME\"}": {
										".": {},
										"f:name": {},
										"f:value": {}
									}
								},
								"f:image": {},
								"f:name": {},
								"f:ports": {
									"k:{\"containerPort\":443,\"protocol\":\"TCP\"}": {
										".": {},
										"f:containerPort": {},
										"f:name": {}
									}
								}
							}
						}
					}
				}
			}
		}
	},
	{
		"manager": "create",
		"operation": "Update",
		"apiVersion": "apps/v1",
		"time": null,
		"fieldsType": "FieldsV1",
		"fieldsV1": {
			"f:metadata": {
				"f:annotations": {
					".": {},
					"f:deprecated.daemonset.template.generation": {}
				},
				"f:labels": {
					".": {},
					"f:app": {}
				}
			},
			"f:spec": {
				"f:revisionHistoryLimit": {},
				"f:selector": {},
				"f:template": {
					"f:metadata": {
						"f:labels": {
							".": {},
							"f:app": {}
						}
					},
					"f:spec": {
						"f:containers": {
							}
						},
						"f:dnsPolicy": {},
						"f:restartPolicy": {},
						"f:schedulerName": {},
						"f:securityContext": {},
						"f:terminationGracePeriodSeconds": {}
					}
				},
				"f:updateStrategy": {
					"f:rollingUpdate": {
						".": {},
						"f:maxSurge": {},
						"f:maxUnavailable": {}
					},
					"f:type": {}
				}
			}
		}
	}
]
`)

	// Change name of some ports.
	ds = []byte(`{
  "apiVersion": "apps/v1",
  "kind": "DaemonSet",
  "metadata": {
    "name": "example-daemonset",
    "labels": {
      "app": "example"
    }
  },
  "spec": {
    "selector": {
      "matchLabels": {
        "app": "example"
      }
    },
    "template": {
      "metadata": {
        "labels": {
          "app": "example"
        }
      },
      "spec": {
        "containers": [
          {
            "name": "nginx",
            "image": "nginx",
            "ports": [
	obj, err = client.AppsV1().RESTClient().
		Put().
		Namespace("default").
		Name("example-daemonset").
		Param("fieldManager", "update").
		Resource("daemonsets").
		Body(ds).
		Do(context.TODO()).
		Get()
	if err != nil {
		t.Fatalf("Failed to update the object: %v", err)
	}
	daemon = obj.(*appsv1.DaemonSet)
	if want, got := 4, len(daemon.Spec.Template.Spec.Containers[0].Env); want != got {
		t.Fatalf("Expected %v EnvVars, got %v", want, got)
	}
	if want, got := 4, len(daemon.Spec.Template.Spec.Containers[0].Ports); want != got {
		t.Fatalf("Expected %v Ports, got %v", want, got)
	}

	expectManagedFields(t, daemon.ManagedFields, `
[
	{
		"manager": "apply",
		"operation": "Apply",
		"apiVersion": "apps/v1",
		"time": null,
		"fieldsType": "FieldsV1",
		"fieldsV1": {
			"f:metadata": {
				"f:labels": {
					"f:app": {}
				}
			},
			"f:spec": {
				"f:selector": {},
				"f:template": {
					"f:spec": {
						"f:containers": {
							"k:{\"name\":\"nginx\"}": {
								".": {},
								"f:env": {
									"k:{\"name\":\"HOME\"}": {
										".": {},
										"f:name": {},
										"f:value": {}
									}
								},
								"f:image": {},
								"f:name": {},
								"f:ports": {
									"k:{\"containerPort\":443,\"protocol\":\"TCP\"}": {
										".": {},
										"f:containerPort": {},
										"f:name": {}
									}
								}
							}
						}
					}
				}
			}
		}
	},
	{
		"manager": "create",
		"operation": "Update",
		"apiVersion": "apps/v1",
		"time": null,
		"fieldsType": "FieldsV1",
		"fieldsV1": {
			"f:metadata": {
				"f:annotations": {},
				"f:labels": {
					".": {},
					"f:app": {}
				}
			},
			"f:spec": {
				"f:revisionHistoryLimit": {},
				"f:selector": {},
				"f:template": {
					"f:metadata": {
						"f:labels": {
							".": {},
							"f:app": {}
						}
					},
					"f:spec": {
						"f:containers": {
							"k:{\"name\":\"nginx\"}": {
								".": {},
								"f:env": {
								"f:terminationMessagePath": {},
								"f:terminationMessagePolicy": {}
							}
						},
						"f:dnsPolicy": {},
						"f:restartPolicy": {},
						"f:schedulerName": {},
						"f:securityContext": {},
						"f:terminationGracePeriodSeconds": {}
					}
				},
				"f:updateStrategy": {
					"f:rollingUpdate": {
						".": {},
						"f:maxSurge": {},
						"f:maxUnavailable": {}
					},
					"f:type": {}
				}
			}
		}
	},
	{
		"manager": "update",
		"operation": "Update",
		"apiVersion": "apps/v1",
		"time": null,
		"fieldsType": "FieldsV1",
		"fieldsV1": {
			"f:metadata": {
				"f:annotations": {
				"f:deprecated.daemonset.template.generation": {}
				}
			},
			"f:spec": {
				"f:template": {
					"f:spec": {
						"f:containers": {
							"k:{\"name\":\"nginx\"}": {
								"f:env": {
									"k:{\"name\":\"PATH\"}": {}
								},
								"f:ports": {
									"k:{\"containerPort\":80,\"protocol\":\"TCP\"}": {}
								}
							}
						}
					}
				}
			}
		}
	}
]
`)

	// Replaces envvars and paths.
	ds = []byte(`
apiVersion: apps/v1
kind: DaemonSet
metadata:
  name: example-daemonset
  labels:
    app: example
spec:
  selector:
    matchLabels:
      app: example
  template:
    spec:
      containers:
        - name: nginx
          image: nginx
          ports:
            - name: port80
              containerPort: 80
          env:
            - name: PATH
              value: "/bin:/usr/bin:/usr/local/bin"
`)
	obj, err = client.AppsV1().RESTClient().
		Patch(types.ApplyPatchType).
		Namespace("default").
		Name("example-daemonset").
		Param("fieldManager", "apply").
		Param("force", "true").
		Resource("daemonsets").
		Body(ds).
		Do(context.TODO()).
		Get()
	if err != nil {
		t.Fatalf("Failed to apply the second object: %v", err)
	}

	daemon = obj.(*appsv1.DaemonSet)
	// HOME is removed, PATH is replaced with 1.
	if want, got := 2, len(daemon.Spec.Template.Spec.Containers[0].Env); want != got {
		t.Fatalf("Expected %v EnvVars, got %v", want, got)
	}
	if want, got := 2, len(daemon.Spec.Template.Spec.Containers[0].Ports); want != got {
		t.Fatalf("Expected %v Ports, got %v", want, got)
	}

	expectManagedFields(t, daemon.ManagedFields, `
[
	{
		"manager": "apply",
		"operation": "Apply",
		"apiVersion": "apps/v1",
		"time": null,
		"fieldsType": "FieldsV1",
		"fieldsV1": {
			"f:metadata": {
				"f:labels": {
					"f:app": {}
				}
			},
			"f:spec": {
				"f:selector": {},
				"f:template": {
					"f:spec": {
						"f:containers": {
							"k:{\"name\":\"nginx\"}": {
								".": {},
								"f:env": {
									"k:{\"name\":\"PATH\"}": {
										".": {},
										"f:name": {},
										"f:value": {}
									}
								},
								"f:image": {},
								"f:name": {},
								"f:ports": {
									"k:{\"containerPort\":80,\"protocol\":\"TCP\"}": {
										".": {},
										"f:containerPort": {},
										"f:name": {}
									}
								}
							}
						}
					}
				}
			}
		}
	},
	{
		"manager": "create",
		"operation": "Update",
		"apiVersion": "apps/v1",
		"time": null,
		"fieldsType": "FieldsV1",
		"fieldsV1": {
			"f:metadata": {
				"f:annotations": {},
				"f:labels": {
					".": {},
					"f:app": {}
				}
			},
			"f:spec": {
				"f:revisionHistoryLimit": {},
				"f:selector": {},
				"f:template": {
					"f:metadata": {
						"f:labels": {
							".": {},
							"f:app": {}
						}
					},
					"f:spec": {
						"f:containers": {
							"k:{\"name\":\"nginx\"}": {
								".": {},
								"f:env": {
								"f:terminationMessagePath": {},
								"f:terminationMessagePolicy": {}
							}
						},
						"f:dnsPolicy": {},
						"f:restartPolicy": {},
						"f:schedulerName": {},
						"f:securityContext": {},
						"f:terminationGracePeriodSeconds": {}
					}
				},
				"f:updateStrategy": {
					"f:rollingUpdate": {
						".": {},
						"f:maxSurge": {},
						"f:maxUnavailable": {}
					},
					"f:type": {}
				}
			}
		}
	},
	{
		"manager": "update",
		"operation": "Update",
		"apiVersion": "apps/v1",
		"time": null,
		"fieldsType": "FieldsV1",
		"fieldsV1": {
			"f:metadata": {
				"f:annotations": {
					"f:deprecated.daemonset.template.generation": {}
				}
			}
		}
	}
]
`)
}

func TestApplyMatchesFakeClientsetApply(t *testing.T) {
	client, closeFn := setup(t)
	defer closeFn()
	fakeClient := fake.NewClientset()

	// The fake client does not default fields, so we set all defaulted fields directly.
	deployment := appsv1ac.Deployment("deployment", "default").
		WithLabels(map[string]string{"app": "nginx"}).
		WithSpec(appsv1ac.DeploymentSpec().
			WithReplicas(3).
			WithStrategy(appsv1ac.DeploymentStrategy().
func expectManagedFields(t *testing.T, managedFields []metav1.ManagedFieldsEntry, expect string) {
	t.Helper()
	for i := range managedFields {
		managedFields[i].Time = &metav1.Time{}
	}
	got, err := json.MarshalIndent(managedFields, "", "  ")
	if err != nil {
		t.Fatalf("Failed to marshal managed fields: %v", err)
	}
	b := &bytes.Buffer{}
	err = json.Indent(b, []byte(expect), "", "  ")
	if err != nil {
		t.Fatalf("Failed to indent json: %v", err)
	}
	want := b.String()
	diff := cmp.Diff(strings.Split(strings.TrimSpace(string(got)), "\n"), strings.Split(strings.TrimSpace(want), "\n"))
	if len(diff) > 0 {
		t.Fatalf("Want:\n%s\nGot:\n%s\nDiff:\n%s", string(want), string(got), diff)
	}
}

// TestCreateOnApplyFailsWithForbidden makes sure that PATCH requests with the apply content type
// will not create the object if the user does not have both patch and create permissions.
func TestCreateOnApplyFailsWithForbidden(t *testing.T) {
	// Enable RBAC so we can exercise authorization errors.
	server := kubeapiservertesting.StartTestServerOrDie(t, nil, append([]string{"--authorization-mode=RBAC"}, framework.DefaultTestServerFlags()...), framework.SharedEtcd())
	t.Cleanup(server.TearDownFn)

	adminClient := clientset.NewForConfigOrDie(server.ClientConfig)

	pandaConfig := restclient.CopyConfig(server.ClientConfig)
	pandaConfig.Impersonate.UserName = "panda"
	pandaClient := clientset.NewForConfigOrDie(pandaConfig)

	errPatch := ssaPod(pandaClient)

	requireForbiddenPodErr(t, errPatch, `pods "test-pod" is forbidden: User "panda" cannot patch resource "pods" in API group "" in the namespace "default"`)

	createPodRBACAndWait(t, adminClient, "patch")

	errCreate := ssaPod(pandaClient)

	requireForbiddenPodErr(t, errCreate, `pods "test-pod" is forbidden: User "panda" cannot create resource "pods" in API group "" in the namespace "default"`)

	createPodRBACAndWait(t, adminClient, "create")

	errNone := ssaPod(pandaClient)
	require.NoError(t, errNone, "pod create via SSA should succeed now that RBAC is correct")
}

func requireForbiddenPodErr(t *testing.T, err error, message string) {