		Signal:   evictionapi.SignalMemoryAvailable,
		Operator: evictionapi.OpLessThan,
		Value: evictionapi.ThresholdValue{
			Quantity: quantityMustParse("1Gi"),
		},
		MinReclaim: &evictionapi.ThresholdValue{
			Quantity: quantityMustParse("500Mi"),
		},
	}
	testCases := map[string]struct {
		enforceMinReclaim bool
		thresholds        []evictionapi.Threshold
		observations      signalObservations
		result            []evictionapi.Threshold
	}{
		"empty": {
			enforceMinReclaim: false,
			thresholds:        []evictionapi.Threshold{},
			observations:      signalObservations{},
			result:            []evictionapi.Threshold{},
		},
		"threshold-met-memory": {
			enforceMinReclaim: false,
			thresholds:        []evictionapi.Threshold{hardThreshold},
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					available: quantityMustParse("500Mi"),
				},
			},
			result: []evictionapi.Threshold{hardThreshold},
		},
		"threshold-not-met": {
			enforceMinReclaim: false,
			thresholds:        []evictionapi.Threshold{hardThreshold},
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					available: quantityMustParse("2Gi"),
				},
			},
			result: []evictionapi.Threshold{},
		},
		"threshold-met-with-min-reclaim": {
			enforceMinReclaim: true,
			thresholds:        []evictionapi.Threshold{hardThreshold},
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					available: quantityMustParse("1.05Gi"),
				},
			},
			result: []evictionapi.Threshold{hardThreshold},
		},
		"threshold-not-met-with-min-reclaim": {
			enforceMinReclaim: true,
			thresholds:        []evictionapi.Threshold{hardThreshold},
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					available: quantityMustParse("2Gi"),
				},
			},
			result: []evictionapi.Threshold{},
		},
	}
	for testName, testCase := range testCases {
		actual := thresholdsMet(testCase.thresholds, testCase.observations, testCase.enforceMinReclaim)
		if !thresholdList(actual).Equal(thresholdList(testCase.result)) {
			t.Errorf("Test case: %s, expected: %v, actual: %v", testName, testCase.result, actual)
		}
	}
}

func TestThresholdsUpdatedStats(t *testing.T) {
	updatedThreshold := evictionapi.Threshold{
		Signal: evictionapi.SignalMemoryAvailable,
	}
	locationUTC, err := time.LoadLocation("UTC")
	if err != nil {
		t.Error(err)
		return
	}
	testCases := map[string]struct {
		thresholds   []evictionapi.Threshold
		observations signalObservations
		last         signalObservations
		result       []evictionapi.Threshold
	}{
		"empty": {
			thresholds:   []evictionapi.Threshold{},
			observations: signalObservations{},
			last:         signalObservations{},
			result:       []evictionapi.Threshold{},
		},
		"no-time": {
			thresholds: []evictionapi.Threshold{updatedThreshold},
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{},
			},
			last:   signalObservations{},
			result: []evictionapi.Threshold{updatedThreshold},
		},
		"no-last-observation": {
			thresholds: []evictionapi.Threshold{updatedThreshold},
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					time: metav1.Date(2016, 1, 1, 0, 0, 0, 0, locationUTC),
				},
			},
			last:   signalObservations{},
			result: []evictionapi.Threshold{updatedThreshold},
		},
		"time-machine": {
			thresholds: []evictionapi.Threshold{updatedThreshold},
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					time: metav1.Date(2016, 1, 1, 0, 0, 0, 0, locationUTC),
				},
			},
			last: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					time: metav1.Date(2016, 1, 1, 0, 1, 0, 0, locationUTC),
				},
			},
			result: []evictionapi.Threshold{},
		},
		"same-observation": {
			thresholds: []evictionapi.Threshold{updatedThreshold},
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					time: metav1.Date(2016, 1, 1, 0, 0, 0, 0, locationUTC),
				},
			},
			last: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					time: metav1.Date(2016, 1, 1, 0, 0, 0, 0, locationUTC),
				},
			},
			result: []evictionapi.Threshold{},
		},
		"new-observation": {
			thresholds: []evictionapi.Threshold{updatedThreshold},
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					time: metav1.Date(2016, 1, 1, 0, 1, 0, 0, locationUTC),
				},
			},
			last: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					time: metav1.Date(2016, 1, 1, 0, 0, 0, 0, locationUTC),
				},
			},
			result: []evictionapi.Threshold{updatedThreshold},
		},
	}
	for testName, testCase := range testCases {
		actual := thresholdsUpdatedStats(testCase.thresholds, testCase.observations, testCase.last)
		if !thresholdList(actual).Equal(thresholdList(testCase.result)) {
			t.Errorf("Test case: %s, expected: %v, actual: %v", testName, testCase.result, actual)
		}
	}
}

func TestPercentageThresholdsMet(t *testing.T) {
	specificThresholds := []evictionapi.Threshold{
		{
			Signal:   evictionapi.SignalMemoryAvailable,
			Operator: evictionapi.OpLessThan,
			Value: evictionapi.ThresholdValue{
				Percentage: 0.2,
			},
			MinReclaim: &evictionapi.ThresholdValue{
				Percentage: 0.05,
			},
		},
		{
			Signal:   evictionapi.SignalNodeFsAvailable,
			Operator: evictionapi.OpLessThan,
			Value: evictionapi.ThresholdValue{
				Percentage: 0.3,
			},
		},
	}

	testCases := map[string]struct {
		enforceMinRelaim bool
		thresholds       []evictionapi.Threshold
		observations     signalObservations
		result           []evictionapi.Threshold
	}{
		"BothMet": {
			enforceMinRelaim: false,
			thresholds:       specificThresholds,
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					available: quantityMustParse("100Mi"),
					capacity:  quantityMustParse("1000Mi"),
				},
				evictionapi.SignalNodeFsAvailable: signalObservation{
					available: quantityMustParse("100Gi"),
					capacity:  quantityMustParse("1000Gi"),
				},
			},
			result: specificThresholds,
		},
		"NoneMet": {
			enforceMinRelaim: false,
			thresholds:       specificThresholds,
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					available: quantityMustParse("300Mi"),
					capacity:  quantityMustParse("1000Mi"),
				},
				evictionapi.SignalNodeFsAvailable: signalObservation{
					available: quantityMustParse("400Gi"),
					capacity:  quantityMustParse("1000Gi"),
				},
			},
			result: []evictionapi.Threshold{},
		},
		"DiskMet": {
			enforceMinRelaim: false,
			thresholds:       specificThresholds,
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					available: quantityMustParse("300Mi"),
					capacity:  quantityMustParse("1000Mi"),
				},
				evictionapi.SignalNodeFsAvailable: signalObservation{
					available: quantityMustParse("100Gi"),
					capacity:  quantityMustParse("1000Gi"),
				},
			},
			result: []evictionapi.Threshold{specificThresholds[1]},
		},
		"MemoryMet": {
			enforceMinRelaim: false,
			thresholds:       specificThresholds,
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					available: quantityMustParse("100Mi"),
					capacity:  quantityMustParse("1000Mi"),
				},
				evictionapi.SignalNodeFsAvailable: signalObservation{
					available: quantityMustParse("400Gi"),
					capacity:  quantityMustParse("1000Gi"),
				},
			},
			result: []evictionapi.Threshold{specificThresholds[0]},
		},
		"MemoryMetWithMinReclaim": {
			enforceMinRelaim: true,
			thresholds:       specificThresholds,
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					available: quantityMustParse("225Mi"),
					capacity:  quantityMustParse("1000Mi"),
				},
			},
			result: []evictionapi.Threshold{specificThresholds[0]},
		},
		"MemoryNotMetWithMinReclaim": {
			enforceMinRelaim: true,
			thresholds:       specificThresholds,
			observations: signalObservations{
				evictionapi.SignalMemoryAvailable: signalObservation{
					available: quantityMustParse("300Mi"),
					capacity:  quantityMustParse("1000Mi"),
				},
			},
			result: []evictionapi.Threshold{},
		},
	}
	for testName, testCase := range testCases {
		actual := thresholdsMet(testCase.thresholds, testCase.observations, testCase.enforceMinRelaim)
		if !thresholdList(actual).Equal(thresholdList(testCase.result)) {
			t.Errorf("Test case: %s, expected: %v, actual: %v", testName, testCase.result, actual)
		}
	}
}

func TestThresholdsFirstObservedAt(t *testing.T) {
	hardThreshold := evictionapi.Threshold{
		Signal:   evictionapi.SignalMemoryAvailable,
		Operator: evictionapi.OpLessThan,
		Value: evictionapi.ThresholdValue{
			Quantity: quantityMustParse("1Gi"),
		},
	}
	now := metav1.Now()
	oldTime := metav1.NewTime(now.Time.Add(-1 * time.Minute))
			// Terminated init containers aren't included in the containerMemoryStats
			containerMemoryStats: []containerMemoryStat{
				{name: restartableInit1, usage: "150Mi"},
				{name: regular1, usage: "100Mi"},
			},
			expectedContainerMessage: memoryExceededEvictionMessage(nil),
		},
		{
			name: "Init container exceeds memory usage",
			initContainers: []v1.Container{
				newContainer(init1, newResourceList("", "100Mi", ""), newResourceList("", "", "")),
				newRestartableInitContainer(restartableInit1, newResourceList("", "200Mi", ""), newResourceList("", "", "")),
				newContainer(init2, newResourceList("", "150Mi", ""), newResourceList("", "", "")),
			},
			containers: []v1.Container{
				newContainer(regular1, newResourceList("", "200Mi", ""), newResourceList("", "", "")),
			},
			// An eviction occurred while the init container was running
			containerMemoryStats: []containerMemoryStat{
				{name: init1, usage: "150Mi"},
			},
			expectedContainerMessage: memoryExceededEvictionMessage([]memoryExceededContainer{
				{name: init1, request: "100Mi", usage: "150Mi"},
			}),
			expectedAnnotations: map[string]string{
				OffendingContainersKey:      init1,
				OffendingContainersUsageKey: "150Mi",
			},
		},
		{
			name: "Restartable init container exceeds memory usage",
			initContainers: []v1.Container{
				newContainer(init1, newResourceList("", "100Mi", ""), newResourceList("", "", "")),
				newRestartableInitContainer(restartableInit1, newResourceList("", "200Mi", ""), newResourceList("", "", "")),
				newRestartableInitContainer(restartableInit2, newResourceList("", "200Mi", ""), newResourceList("", "", "")),
			},
			containers: []v1.Container{
				newContainer(regular1, newResourceList("", "200Mi", ""), newResourceList("", "", "")),
			},
			// Terminated init containers aren't included in the containerMemoryStats
			containerMemoryStats: []containerMemoryStat{
				{name: restartableInit1, usage: "250Mi"},
				{name: restartableInit2, usage: "150Mi"},
				{name: regular1, usage: "200Mi"},
			},
			expectedContainerMessage: memoryExceededEvictionMessage([]memoryExceededContainer{
				{name: restartableInit1, request: "200Mi", usage: "250Mi"},
			}),
			expectedAnnotations: map[string]string{
				OffendingContainersKey:      restartableInit1,
				OffendingContainersUsageKey: "250Mi",
			},
		},
		{
			name: "Regular container exceeds memory usage",
			initContainers: []v1.Container{
				newContainer(init1, newResourceList("", "100Mi", ""), newResourceList("", "", "")),
				newRestartableInitContainer(restartableInit1, newResourceList("", "200Mi", ""), newResourceList("", "", "")),
			},
			containers: []v1.Container{
				newContainer(regular1, newResourceList("", "200Mi", ""), newResourceList("", "", "")),
				newContainer(regular2, newResourceList("", "300Mi", ""), newResourceList("", "", "")),
			},
			// Terminated init containers aren't included in the containerMemoryStats
			containerMemoryStats: []containerMemoryStat{
				{name: restartableInit1, usage: "200Mi"},
				{name: regular1, usage: "250Mi"},
				{name: regular2, usage: "250Mi"},
			},
			expectedContainerMessage: memoryExceededEvictionMessage([]memoryExceededContainer{
				{name: regular1, request: "200Mi", usage: "250Mi"},
			}),
			expectedAnnotations: map[string]string{
				OffendingContainersKey:      regular1,
				OffendingContainersUsageKey: "250Mi",
			},
		},
		{
			name: "Ephemeral container exceeds memory usage",
			initContainers: []v1.Container{
				newContainer(init1, newResourceList("", "100Mi", ""), newResourceList("", "", "")),
				newRestartableInitContainer(restartableInit1, newResourceList("", "200Mi", ""), newResourceList("", "", "")),
			},
			containers: []v1.Container{
				newContainer(regular1, newResourceList("", "200Mi", ""), newResourceList("", "", "")),
			},
			ephemeralContainers: []v1.EphemeralContainer{
				{TargetContainerName: regular1, EphemeralContainerCommon: v1.EphemeralContainerCommon{Name: ephemeral1}},
			},
			// Terminated init containers aren't included in the containerMemoryStats
			containerMemoryStats: []containerMemoryStat{
				{name: restartableInit1, usage: "200Mi"},
				{name: regular1, usage: "150Mi"},
				{name: ephemeral1, usage: "250Mi"},
			},
			expectedContainerMessage: memoryExceededEvictionMessage(nil),
		},
		{
			name: "Both regular and restartable init containers exceed memory usage due to missing memory requests",
			initContainers: []v1.Container{
				newContainer(init1, newResourceList("", "100Mi", ""), newResourceList("", "", "")),
				newRestartableInitContainer(restartableInit1, newResourceList("", "", ""), newResourceList("", "", "")),
				newRestartableInitContainer(restartableInit2, newResourceList("", "300Mi", ""), newResourceList("", "", "")),
			},
			containers: []v1.Container{
				newContainer("regular1", newResourceList("", "", ""), newResourceList("", "", "")),
			},
			// Terminated init containers aren't included in the containerMemoryStats
			containerMemoryStats: []containerMemoryStat{
				{name: restartableInit1, usage: "250Mi"},
				{name: restartableInit2, usage: "250Mi"},
				{name: regular1, usage: "200Mi"},
			},
			expectedContainerMessage: memoryExceededEvictionMessage([]memoryExceededContainer{
				{name: restartableInit1, request: "0", usage: "250Mi"},
				{name: regular1, request: "0", usage: "200Mi"},
			}),
			expectedAnnotations: map[string]string{
				OffendingContainersKey:      strings.Join([]string{restartableInit1, regular1}, ","),
				OffendingContainersUsageKey: "250Mi,200Mi",
			},
		},
		{
			name: "Multiple regular and restartable init containers exceed memory usage",
			initContainers: []v1.Container{
				newContainer(init1, newResourceList("", "100Mi", ""), newResourceList("", "", "")),
				newContainer(init2, newResourceList("", "100Mi", ""), newResourceList("", "", "")),
				newContainer(init3, newResourceList("", "100Mi", ""), newResourceList("", "", "")),
				newRestartableInitContainer(restartableInit1, newResourceList("", "100Mi", ""), newResourceList("", "", "")),
				newRestartableInitContainer(restartableInit2, newResourceList("", "200Mi", ""), newResourceList("", "", "")),
				newRestartableInitContainer(restartableInit3, newResourceList("", "300Mi", ""), newResourceList("", "", "")),
			},
			containers: []v1.Container{
				newContainer(regular1, newResourceList("", "300Mi", ""), newResourceList("", "", "")),
				newContainer(regular2, newResourceList("", "200Mi", ""), newResourceList("", "", "")),
				newContainer(regular3, newResourceList("", "100Mi", ""), newResourceList("", "", "")),
			},
			// Terminated init containers aren't included in the containerMemoryStats
			containerMemoryStats: []containerMemoryStat{
				{name: restartableInit1, usage: "150Mi"},
				{name: restartableInit2, usage: "250Mi"},
				{name: restartableInit3, usage: "250Mi"},
				{name: regular1, usage: "50Mi"},
				{name: regular2, usage: "250Mi"},
				{name: regular3, usage: "400Mi"},
			},
			expectedContainerMessage: memoryExceededEvictionMessage([]memoryExceededContainer{
				{name: restartableInit1, request: "100Mi", usage: "150Mi"},
				{name: restartableInit2, request: "200Mi", usage: "250Mi"},
				{name: regular2, request: "200Mi", usage: "250Mi"},
				{name: regular3, request: "100Mi", usage: "400Mi"},
			}),
			expectedAnnotations: map[string]string{
				OffendingContainersKey:      strings.Join([]string{restartableInit1, restartableInit2, regular2, regular3}, ","),
				OffendingContainersUsageKey: "150Mi,250Mi,250Mi,400Mi",
			},
		},
	}

	threshold := []evictionapi.Threshold{}
	observations := signalObservations{}
	for _, tc := range testcase {
		pod := newPodWithInitContainers("pod", 1, tc.initContainers, tc.containers, nil)
		if len(tc.ephemeralContainers) > 0 {
			pod.Spec.EphemeralContainers = tc.ephemeralContainers
		}
		// Create PodMemoryStats with dummy container memory
		// and override container stats with the provided values.
		dummyContainerMemory := resource.MustParse("1Mi")
		podStats := newPodMemoryStats(pod, dummyContainerMemory)
		podStats.Containers = make([]statsapi.ContainerStats, len(tc.containerMemoryStats))
		for _, stat := range tc.containerMemoryStats {
			memoryStat := resource.MustParse(stat.usage)
			memoryBytes := uint64(memoryStat.Value())
			podStats.Containers = append(podStats.Containers, statsapi.ContainerStats{
				Name: stat.name,
				Memory: &statsapi.MemoryStats{
					WorkingSetBytes: &memoryBytes,
				},
			})