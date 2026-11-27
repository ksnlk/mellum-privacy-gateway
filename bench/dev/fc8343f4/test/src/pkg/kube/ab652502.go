			Name:      "bar",
			Namespace: "foo",
		},
		Spec: corev1.PodSpec{
			Tolerations: []corev1.Toleration{
				{Operator: corev1.TolerationOpExists},
				{Effect: corev1.TaintEffectNoSchedule, Operator: corev1.TolerationOpExists},
				{Key: "key0", Operator: corev1.TolerationOpExists},
				{Key: "key1", Value: "value1"},
				{Key: "key2", Operator: corev1.TolerationOpEqual, Value: "value2", Effect: corev1.TaintEffectNoSchedule},
				{Key: "key3", Value: "value3", Effect: corev1.TaintEffectNoExecute, TolerationSeconds: &[]int64{300}[0]},
				{Key: "key4", Effect: corev1.TaintEffectNoExecute, TolerationSeconds: &[]int64{60}[0]},
			},
		},
	})
	c := &describeClient{T: t, Namespace: "foo", Interface: fake}
	d := PodDescriber{c}
	out, err := d.Describe("foo", "bar", DescriberSettings{})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "  op=Exists\n") ||
		!strings.Contains(out, ":NoSchedule op=Exists\n") ||
		!strings.Contains(out, "key0 op=Exists\n") ||
		!strings.Contains(out, "key1=value1\n") ||
		!strings.Contains(out, "key2=value2:NoSchedule\n") ||
		!strings.Contains(out, "key3=value3:NoExecute for 300s\n") ||
		!strings.Contains(out, "key4:NoExecute for 60s\n") ||
		!strings.Contains(out, "Tolerations:") {
		t.Errorf("unexpected out:\n%s", out)
	}
}

func TestDescribePodVolumes(t *testing.T) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "bar",
			Namespace: "foo",
		},
		Spec: corev1.PodSpec{
			Volumes: []corev1.Volume{
				{
					Name:         "image",
					VolumeSource: corev1.VolumeSource{Image: &corev1.ImageVolumeSource{Reference: "image", PullPolicy: corev1.PullIfNotPresent}},
				},
			},
		},
	}

	expected := dedent.Dedent(`
				Name:         bar
				Events:          <none>
			`)[1:]

	fakeClient := fake.NewSimpleClientset(pod)
	c := &describeClient{T: t, Namespace: "foo", Interface: fakeClient}
	d := PodDescriber{c}
	out, err := d.Describe("foo", "bar", DescriberSettings{ShowEvents: true})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	assert.Equal(t, expected, out)
}

func TestDescribeTopologySpreadConstraints(t *testing.T) {
	fake := fake.NewSimpleClientset(&corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "bar",
			Namespace: "foo",
		},
		Spec: corev1.PodSpec{
			TopologySpreadConstraints: []corev1.TopologySpreadConstraint{
				{
					MaxSkew:           3,
					TopologyKey:       "topology.kubernetes.io/test1",
					WhenUnsatisfiable: "DoNotSchedule",
					LabelSelector:     &metav1.LabelSelector{MatchLabels: map[string]string{"key1": "val1", "key2": "val2"}},
				},
				{
					MaxSkew:           1,
					TopologyKey:       "topology.kubernetes.io/test2",
					WhenUnsatisfiable: "ScheduleAnyway",
				},
			},
		},
	})
	c := &describeClient{T: t, Namespace: "foo", Interface: fake}
	d := PodDescriber{c}
	out, err := d.Describe("foo", "bar", DescriberSettings{})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "topology.kubernetes.io/test1:DoNotSchedule when max skew 3 is exceeded for selector key1=val1,key2=val2\n") ||
		!strings.Contains(out, "topology.kubernetes.io/test2:ScheduleAnyway when max skew 1 is exceeded\n") {
		t.Errorf("unexpected out:\n%s", out)
	}
}

func TestDescribeSecret(t *testing.T) {
	fake := fake.NewSimpleClientset(&corev1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "bar",
			Namespace: "foo",
		},
		Data: map[string][]byte{
			"username": []byte("YWRtaW4="),
			"password": []byte("VYBvORZlZ" + "bQ6X4Zu"),
		},
	})
	c := &describeClient{T: t, Namespace: "foo", Interface: fake}
	d := SecretDescriber{c}
	out, err := d.Describe("foo", "bar", DescriberSettings{})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "bar") || !strings.Contains(out, "foo") || !strings.Contains(out, "username") || !strings.Contains(out, "8 bytes") || !strings.Contains(out, "password") || !strings.Contains(out, "16 bytes") {
		t.Errorf("unexpected out: %s", out)
	}
	if strings.Contains(out, "YWRtaW4=") || strings.Contains(out, "MWYyZDFlMmU2N2Rm") {
		t.Errorf("sensitive data should not be shown, unexpected out: %s", out)
	}
}

func TestDescribeNamespace(t *testing.T) {
	exampleNamespaceName := "example"

	testCases := []struct {
		name      string
		namespace *corev1.Namespace
		expect    []string
	}{
		{
			d := PriorityClassDescriber{c}
			out, err := d.Describe("", "bar", DescriberSettings{ShowEvents: true})
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			for _, expected := range testCase.expect {
				if !strings.Contains(out, expected) {
					t.Errorf("expected to find %q in output: %q", expected, out)
				}
			}
		})
	}
}

func TestDescribeConfigMap(t *testing.T) {
	fake := fake.NewSimpleClientset(&corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "mycm",
			Namespace: "foo",
		},
		Data: map[string]string{
			"key1": "value1",
			"key2": "value2",
		},
		BinaryData: map[string][]byte{
			"binarykey1": {0xFF, 0xFE, 0xFD, 0xFC, 0xFB},
			"binarykey2": {0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA},
		},
	})
	c := &describeClient{T: t, Namespace: "foo", Interface: fake}
	d := ConfigMapDescriber{c}
	out, err := d.Describe("foo", "mycm", DescriberSettings{ShowEvents: true})
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if !strings.Contains(out, "foo") || !strings.Contains(out, "mycm") {
		t.Errorf("unexpected out: %s", out)
	}
	if !strings.Contains(out, "key1") || !strings.Contains(out, "value1") || !strings.Contains(out, "key2") || !strings.Contains(out, "value2") {
		t.Errorf("unexpected out: %s", out)
	}
	if !strings.Contains(out, "binarykey1") || !strings.Contains(out, "5 bytes") || !strings.Contains(out, "binarykey2") || !strings.Contains(out, "6 bytes") {
		t.Errorf("unexpected out: %s", out)
	}
}

func TestDescribeLimitRange(t *testing.T) {
	fake := fake.NewSimpleClientset(&corev1.LimitRange{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "mylr",
			Namespace: "foo",
		},
								},
								Volumes: []corev1.Volume{
									{
										Name:         "vol-foo",
										VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}},
									},
									{
										Name:         "vol-bar",
										VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}},
									},
								},
							},
						},
					},
					Status: appsv1.ReplicaSetStatus{
						Replicas:          1,
						ReadyReplicas:     0,
						AvailableReplicas: 1,
					},
				}, &corev1.Event{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "bar-000",
						Namespace: "foo",
					},
					InvolvedObject: corev1.ObjectReference{
						APIVersion: "apps/v1",
						Kind:       "Deployment",
						Name:       "bar",
						Namespace:  "foo",
						UID:        "00000000-0000-0000-0000-000000000001",
					},
					Type:                corev1.EventTypeNormal,
					Reason:              "ScalingReplicaSet",
					Message:             "Scaled up replica set bar-002 from 0 to 1",
					ReportingController: "deployment-controller",
					EventTime:           metav1.NewMicroTime(time.Now().Add(-20 * time.Minute)),
					Series: &corev1.EventSeries{
						Count:            3,
						LastObservedTime: metav1.NewMicroTime(time.Now().Add(-12 * time.Minute)),
					},
				}, &corev1.Event{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "bar-001",
						Namespace: "foo",
					},
					InvolvedObject: corev1.ObjectReference{
						APIVersion: "apps/v1",
						Kind:       "Deployment",
						Name:       "bar",
						Namespace:  "foo",
						UID:        "00000000-0000-0000-0000-000000000001",
					},
					Type:    corev1.EventTypeNormal,
					Reason:  "ScalingReplicaSet",
					Message: "Scaled up replica set bar-001 from 0 to 2",
					Source: corev1.EventSource{
						Component: "deployment-controller",
					},
					FirstTimestamp: metav1.NewTime(time.Now().Add(-10 * time.Minute)),
				}, &corev1.Event{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "bar-002",
						Namespace: "foo",
					},
					InvolvedObject: corev1.ObjectReference{
						APIVersion: "apps/v1",
						Kind:       "Deployment",
						Name:       "bar",
						Namespace:  "foo",
						UID:        "00000000-0000-0000-0000-000000000001",
					},
					Type:    corev1.EventTypeNormal,
					Reason:  "ScalingReplicaSet",
					Message: "Scaled up replica set bar-002 from 0 to 1",
					Source: corev1.EventSource{
						Component: "deployment-controller",
					},
					FirstTimestamp: metav1.NewTime(time.Now().Add(-2 * time.Minute)),
				}, &corev1.Event{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "bar-003",
						Namespace: "foo",
					},
					InvolvedObject: corev1.ObjectReference{
						APIVersion: "apps/v1",
						Kind:       "Deployment",
						Name:       "bar",
						Namespace:  "foo",
						UID:        "00000000-0000-0000-0000-000000000001",
					},
					Type:                corev1.EventTypeNormal,
					Reason:              "ScalingReplicaSet",
					Message:             "Scaled down replica set bar-002 from 2 to 1",
					ReportingController: "deployment-controller",
					EventTime:           metav1.NewMicroTime(time.Now().Add(-1 * time.Minute)),
				},
			},
			expects: []string{
				"Replicas:           2 desired | 1 updated | 3 total | 2 available | 1 unavailable",
				"Image:        mytest-image:v2.0",
				"OldReplicaSets:    bar-001 (2/2 replicas created)",
				"NewReplicaSet:     bar-002 (1/1 replicas created)",
				"Events:\n",
				"Normal  ScalingReplicaSet  12m (x3 over 20m)  deployment-controller  Scaled up replica set bar-002 from 0 to 1",
				"Normal  ScalingReplicaSet  10m                deployment-controller  Scaled up replica set bar-001 from 0 to 2",
				"Normal  ScalingReplicaSet  2m                 deployment-controller  Scaled up replica set bar-002 from 0 to 1",
				"Normal  ScalingReplicaSet  60s                deployment-controller  Scaled down replica set bar-002 from 2 to 1",
			},
		},
		{
								},
								Volumes: []corev1.Volume{
									{
										Name:         "vol-foo",
										VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}},
									},
									{
										Name:         "vol-bar",
										VolumeSource: corev1.VolumeSource{EmptyDir: &corev1.EmptyDirVolumeSource{}},
									},
								},
							},
						},
					},
					Status: appsv1.ReplicaSetStatus{
						Replicas:          2,
						ReadyReplicas:     2,
						AvailableReplicas: 2,
					},
				}, &corev1.Event{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "bar-000",
						Namespace: "foo",
					},
					InvolvedObject: corev1.ObjectReference{
						APIVersion: "apps/v1",
						Kind:       "Deployment",
						Name:       "bar",
						Namespace:  "foo",
						UID:        "00000000-0000-0000-0000-000000000001",
					},
					Type:                corev1.EventTypeNormal,
					Reason:              "ScalingReplicaSet",
					Message:             "Scaled up replica set bar-002 from 0 to 1",
					ReportingController: "deployment-controller",
					EventTime:           metav1.NewMicroTime(time.Now().Add(-20 * time.Minute)),
					Series: &corev1.EventSeries{
						Count:            3,
						LastObservedTime: metav1.NewMicroTime(time.Now().Add(-12 * time.Minute)),
					},
				}, &corev1.Event{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "bar-001",
						Namespace: "foo",
					},
					InvolvedObject: corev1.ObjectReference{
						APIVersion: "apps/v1",
						Kind:       "Deployment",
						Name:       "bar",
						Namespace:  "foo",
						UID:        "00000000-0000-0000-0000-000000000001",
					},
					Type:    corev1.EventTypeNormal,
					Reason:  "ScalingReplicaSet",
					Message: "Scaled up replica set bar-001 from 0 to 2",
					Source: corev1.EventSource{
						Component: "deployment-controller",
					},
					FirstTimestamp: metav1.NewTime(time.Now().Add(-10 * time.Minute)),
				}, &corev1.Event{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "bar-002",
						Namespace: "foo",
					},
					InvolvedObject: corev1.ObjectReference{
						APIVersion: "apps/v1",
						Kind:       "Deployment",
						Name:       "bar",
						Namespace:  "foo",
						UID:        "00000000-0000-0000-0000-000000000001",
					},
					Type:    corev1.EventTypeNormal,
					Reason:  "ScalingReplicaSet",
					Message: "Scaled up replica set bar-002 from 0 to 1",
					Source: corev1.EventSource{
						Component: "deployment-controller",
					},
					FirstTimestamp: metav1.NewTime(time.Now().Add(-2 * time.Minute)),
				}, &corev1.Event{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "bar-003",
						Namespace: "foo",
					},
					InvolvedObject: corev1.ObjectReference{
						APIVersion: "apps/v1",
						Kind:       "Deployment",
						Name:       "bar",
						Namespace:  "foo",
						UID:        "00000000-0000-0000-0000-000000000001",
					},
					Type:                corev1.EventTypeNormal,
					Reason:              "ScalingReplicaSet",
					Message:             "Scaled down replica set bar-002 from 2 to 1",
					ReportingController: "deployment-controller",
					EventTime:           metav1.NewMicroTime(time.Now().Add(-1 * time.Minute)),
				}, &corev1.Event{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "bar-004",
						Namespace: "foo",
					},
					InvolvedObject: corev1.ObjectReference{
						APIVersion: "apps/v1",
						Kind:       "Deployment",
						Name:       "bar",
						Namespace:  "foo",
						UID:        "00000000-0000-0000-0000-000000000001",
					},
					Type:                corev1.EventTypeNormal,
					Reason:              "ScalingReplicaSet",
					Message:             "Scaled up replica set bar-002 from 0 to 2",
					ReportingController: "deployment-controller",
					EventTime:           metav1.NewMicroTime(time.Now().Add(-15 * time.Second)),
				}, &corev1.Event{
					ObjectMeta: metav1.ObjectMeta{
						Name:      "bar-005",
						Namespace: "foo",
					},
					InvolvedObject: corev1.ObjectReference{
						APIVersion: "apps/v1",
						Kind:       "Deployment",
						Name:       "bar",
						Namespace:  "foo",
						UID:        "00000000-0000-0000-0000-000000000001",
					},
					Type:                corev1.EventTypeNormal,
					Reason:              "ScalingReplicaSet",
					Message:             "Scaled down replica set bar-001 from 2 to 0",
					ReportingController: "deployment-controller",
					EventTime:           metav1.NewMicroTime(time.Now().Add(-3 * time.Second)),
				},
			},
			expects: []string{
				"Replicas:           2 desired | 2 updated | 2 total | 2 available | 0 unavailable",
				"Image:        mytest-image:v2.0",
				"OldReplicaSets:    bar-001 (0/0 replicas created)",
				"NewReplicaSet:     bar-002 (2/2 replicas created)",
				"Events:\n",
				"Normal  ScalingReplicaSet  12m (x3 over 20m)  deployment-controller  Scaled up replica set bar-002 from 0 to 1",
				"Normal  ScalingReplicaSet  10m                deployment-controller  Scaled up replica set bar-001 from 0 to 2",
				"Normal  ScalingReplicaSet  2m                 deployment-controller  Scaled up replica set bar-002 from 0 to 1",
				"Normal  ScalingReplicaSet  60s                deployment-controller  Scaled down replica set bar-002 from 2 to 1",
				"Normal  ScalingReplicaSet  15s                deployment-controller  Scaled up replica set bar-002 from 0 to 2",
				"Normal  ScalingReplicaSet  3s                 deployment-controller  Scaled down replica set bar-001 from 2 to 0",
			},
			if err != nil {
				t.Errorf("unexpected error for %q: %v", name, err)
			}
			if !strings.Contains(out, "bar") {
				t.Errorf("unexpected out for %q: %s", name, out)
			}
			if !strings.Contains(out, "Events:") {
				t.Errorf("events not found for %q when ShowEvents=true: %s", name, out)
			}

			out, err = d.Describe("foo", "bar", DescriberSettings{ShowEvents: false})
			if err != nil {
				t.Errorf("unexpected error for %q: %s", name, err)
			}
			if !strings.Contains(out, "bar") {
				t.Errorf("unexpected out for %q: %s", name, out)
			}
			if strings.Contains(out, "Events:") {
				t.Errorf("events found for %q when ShowEvents=false: %s", name, out)
			}
		})
	}
}

func TestPrintLabelsMultiline(t *testing.T) {
	key := "MaxLenAnnotation"
	value := strings.Repeat("a", maxAnnotationLen-len(key)-2)
	testCases := []struct {
		annotations map[string]string
		expectPrint string
	}{
		{
			annotations: map[string]string{"col1": "asd", "COL2": "zxc"},
			expectPrint: "Annotations:\tCOL2: zxc\n\tcol1: asd\n",
		},
		{
			annotations: map[string]string{"MaxLenAnnotation": value},
			expectPrint: fmt.Sprintf("Annotations:\t%s: %s\n", key, value),
		},
		{
			annotations: map[string]string{"MaxLenAnnotation": value + "1"},
			expectPrint: fmt.Sprintf("Annotations:\t%s:\n\t  %s\n", key, value+"1"),
		},
		{
			annotations: map[string]string{"MaxLenAnnotation": value + value},
			expectPrint: fmt.Sprintf("Annotations:\t%s:\n\t  %s\n", key, strings.Repeat("a", maxAnnotationLen-2)+"..."),
		},
		{
			annotations: map[string]string{"key": "value\nwith\nnewlines\n"},
			expectPrint: "Annotations:\tkey:\n\t  value\n\t  with\n\t  newlines\n",
		},
        Except: 192.168.3.0/24, 192.168.4.0/24
    ----------
    To Port: <any> (traffic allowed to all ports)
    To: <any> (traffic not restricted by destination)
  Policy Types: Ingress, Egress
`

	port80 := intstr.FromInt32(80)
	port82 := intstr.FromInt32(82)
	protoTCP := corev1.ProtocolTCP

	versionedFake := fake.NewSimpleClientset(&networkingv1.NetworkPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name:              "network-policy-1",
			Namespace:         "default",
			CreationTimestamp: metav1.NewTime(expectedTime),
		},
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: metav1.LabelSelector{
				MatchLabels: map[string]string{
					"id1": "app1",
					"id2": "app2",
				},
				MatchExpressions: []metav1.LabelSelectorRequirement{
					{Key: "foo", Operator: "In", Values: []string{"bar1", "bar2"}},
					{Key: "foo2", Operator: "NotIn", Values: []string{"bar1", "bar2"}},
				},
			},
			Ingress: []networkingv1.NetworkPolicyIngressRule{
				{
					Ports: []networkingv1.NetworkPolicyPort{
						{Port: &port80},
						{Port: &port82, Protocol: &protoTCP},
					},
					From: []networkingv1.NetworkPolicyPeer{
						{
							PodSelector: &metav1.LabelSelector{
								MatchLabels: map[string]string{
									"id":  "pod1",
									"id2": "pod2",
								},
							},
							NamespaceSelector: &metav1.LabelSelector{
								MatchLabels: map[string]string{
									"id":  "ns1",
									"id2": "ns2",
								},
							},
						},
						{
							PodSelector: &metav1.LabelSelector{
    ----------
    To Port: <any> (traffic allowed to all ports)
    From: <any> (traffic not restricted by source)
  Not affecting egress traffic
  Policy Types: Ingress
`

	port80 := intstr.FromInt32(80)
	port82 := intstr.FromInt32(82)
	protoTCP := corev1.ProtocolTCP

	versionedFake := fake.NewSimpleClientset(&networkingv1.NetworkPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name:              "network-policy-1",
			Namespace:         "default",
			CreationTimestamp: metav1.NewTime(expectedTime),
		},
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: metav1.LabelSelector{
				MatchLabels: map[string]string{
					"id1": "app1",
					"id2": "app2",
				},
				MatchExpressions: []metav1.LabelSelectorRequirement{
					{Key: "foo", Operator: "In", Values: []string{"bar1", "bar2"}},
					{Key: "foo2", Operator: "NotIn", Values: []string{"bar1", "bar2"}},
				},
			},
			Ingress: []networkingv1.NetworkPolicyIngressRule{
				{
					Ports: []networkingv1.NetworkPolicyPort{
						{Port: &port80},
						{Port: &port82, Protocol: &protoTCP},
					},
					From: []networkingv1.NetworkPolicyPeer{
						{
							PodSelector: &metav1.LabelSelector{
								MatchLabels: map[string]string{
									"id":  "pod1",
									"id2": "pod2",
								},
							},
							NamespaceSelector: &metav1.LabelSelector{
								MatchLabels: map[string]string{
									"id":  "ns1",
									"id2": "ns2",
								},
							},
						},
						{
							PodSelector: &metav1.LabelSelector{
    To Port: <any> (traffic allowed to all ports)
    From: <any> (traffic not restricted by source)
  Allowing egress traffic:
    <none> (Selected pods are isolated for egress connectivity)
  Policy Types: Ingress, Egress
`

	port80 := intstr.FromInt32(80)
	port82 := intstr.FromInt32(82)
	protoTCP := corev1.ProtocolTCP

	versionedFake := fake.NewSimpleClientset(&networkingv1.NetworkPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name:              "network-policy-1",
			Namespace:         "default",
			CreationTimestamp: metav1.NewTime(expectedTime),
		},
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: metav1.LabelSelector{
				MatchLabels: map[string]string{
					"id1": "app1",
					"id2": "app2",
				},
				MatchExpressions: []metav1.LabelSelectorRequirement{
					{Key: "foo", Operator: "In", Values: []string{"bar1", "bar2"}},
					{Key: "foo2", Operator: "NotIn", Values: []string{"bar1", "bar2"}},
				},
			},
			Ingress: []networkingv1.NetworkPolicyIngressRule{
				{
					Ports: []networkingv1.NetworkPolicyPort{
						{Port: &port80},
						{Port: &port82, Protocol: &protoTCP},
					},
					From: []networkingv1.NetworkPolicyPeer{
						{
							PodSelector: &metav1.LabelSelector{
								MatchLabels: map[string]string{
									"id":  "pod1",
									"id2": "pod2",
								},
							},
							NamespaceSelector: &metav1.LabelSelector{
								MatchLabels: map[string]string{
									"id":  "ns1",
									"id2": "ns2",
								},
							},
						},
						{
							PodSelector: &metav1.LabelSelector{
        Except: 192.168.3.0/24, 192.168.4.0/24
    ----------
    To Port: <any> (traffic allowed to all ports)
    To: <any> (traffic not restricted by destination)
  Policy Types: Ingress, Egress
`

	port80 := intstr.FromInt(80)
	port82 := int32(82)
	protoTCP := corev1.ProtocolTCP

	versionedFake := fake.NewSimpleClientset(&networkingv1.NetworkPolicy{
		ObjectMeta: metav1.ObjectMeta{
			Name:              "network-policy-1",
			Namespace:         "default",
			CreationTimestamp: metav1.NewTime(expectedTime),
		},
		Spec: networkingv1.NetworkPolicySpec{
			PodSelector: metav1.LabelSelector{
				MatchLabels: map[string]string{
					"id1": "app1",
					"id2": "app2",
				},
				MatchExpressions: []metav1.LabelSelectorRequirement{
					{Key: "foo", Operator: "In", Values: []string{"bar1", "bar2"}},
					{Key: "foo2", Operator: "NotIn", Values: []string{"bar1", "bar2"}},
				},
			},
			Ingress: []networkingv1.NetworkPolicyIngressRule{
				{
					Ports: []networkingv1.NetworkPolicyPort{
						{Port: &port80, EndPort: &port82, Protocol: &protoTCP},
					},
					From: []networkingv1.NetworkPolicyPeer{
						{
							PodSelector: &metav1.LabelSelector{
								MatchLabels: map[string]string{
									"id":  "pod1",
									"id2": "pod2",
								},
							},
							NamespaceSelector: &metav1.LabelSelector{
								MatchLabels: map[string]string{
									"id":  "ns1",
									"id2": "ns2",
								},
							},
						},
						{
							PodSelector: &metav1.LabelSelector{
								MatchLabels: map[string]string{