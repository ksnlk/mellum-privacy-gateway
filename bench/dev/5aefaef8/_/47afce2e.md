
* 9 lgtm, but not approved test PRs: [https://github.com/kubernetes/kubernetes/pulls?q=is%3Apr+is%3Aopen+label%3Asig%2Fnode++label%3Algtm+-label%3Aapproved+label%3Aarea%2Ftest](https://github.com/kubernetes/kubernetes/pulls?q=is%3Apr+is%3Aopen+label%3Asig%2Fnode++label%3Algtm+-label%3Aapproved+label%3Aarea%2Ftest)

**AI**: Sergey - can we exclude all PRs that are not sig/node specific. Release manager, API review blocked, etc.

* [derekwaynecarr] should we revisit [swap](https://github.com/kubernetes/kubernetes/issues/53533#issuecomment-334947625) in 1.21, volunteers to work together?
  * **AI**: Derek to start google doc
  * Alexander @kad interested to join
  * Karan (@karan) also interested
  * Peter (@haircommander) is interested
  * [added after meeting] @ehashman is also interested
* [dashpole/egernst] naming for pod metrics in [https://github.com/kubernetes/kubernetes/pull/95839](https://github.com/kubernetes/kubernetes/pull/95839)
  * kube_pod_resource_cpu_seconds_total (in-line with [clayton’s KEP](https://github.com/kubernetes/enhancements/pull/1916))
  * pod_cpu_usage_seconds_total (in-line with other metrics on the endpoint; [KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/kubelet-resource-metrics-endpoint.md))
  * **AI**: David Ashpole to start slack thread
* [mrunalp] CRI beta next step
  * [https://github.com/kubernetes/kubernetes/pull/96387](https://github.com/kubernetes/kubernetes/pull/96387)
* [SergeyKanzhelev] RuntimeClass GA: [https://github.com/kubernetes/kubernetes/pull/95718](https://github.com/kubernetes/kubernetes/pull/95718)
  * Should v1alpa1 be removed?
    * no objections - will send a separate PR
* [@jpepin] [cadvisor metrics deprecation](https://github.com/kubernetes/kubernetes/issues/68522) in the kubelet and [timeline for proposed KEP](https://github.com/kubernetes/enhancements/pull/2130/files?short_path=58871be#diff-58871be1cadaf7855efac07908b9fa9b589e4aac74ce6badc2c0f5cd05af8faa)
* ~~[fromani] quick question: podresources API PRs: suitable for code freeze exception? Any concern? (already answered by the following bullet point)~~
* PRs for 1.20 to be landed this week!!!
  * [dims] [Deprecate Dockershim](https://github.com/kubernetes/kubernetes/pull/94624)
  * ~~[AlexeyPerevalov]: [Implement TopologyInfo and cpu_ids in podresources interface](https://github.com/kubernetes/kubernetes/pull/93243)~~
  * [fromani]: [podresources APIs: concrete resources apis: implement GetAllocatableResources](https://github.com/kubernetes/kubernetes/pull/95734)
  * [alukiano, cez, krzwiatrzyk] [Memory manager](https://github.com/kubernetes/kubernetes/pull/95479)
  * ~~[andrewsykim] kubelet credential provider plugin [https://github.com/kubernetes/kubernetes/pull/94196](https://github.com/kubernetes/kubernetes/pull/94196)~~
  * ~~[SergeyKanzhelev] RuntimeClass [https://github.com/kubernetes/kubernetes/pull/95718](https://github.com/kubernetes/kubernetes/pull/95718)~~
  * [David Porter] Graceful Node shutdown -- [https://github.com/kubernetes/kubernetes/pull/96129](https://github.com/kubernetes/kubernetes/pull/96129)
  * ~~[Mrunal Patel] Add CRI v1beta1 proto[https://github.com/kubernetes/kubernetes/pull/96387](https://github.com/kubernetes/kubernetes/pull/96387)~~
* [karan] NPD dev is blocked due to failing e2e presubmit test: [https://github.com/kubernetes/kubernetes/issues/95955](https://github.com/kubernetes/kubernetes/issues/95955)
  * ~~PR 1: [https://github.com/kubernetes/kubernetes/pull/96381](https://github.com/kubernetes/kubernetes/pull/96381) (need approval)~~
  * ~~PR 2: [https://github.com/kubernetes/kubernetes/pull/96262](https://github.com/kubernetes/kubernetes/pull/96262) (need rebase + approval)~~
* [David Porter] Call out for graceful node shutdown PR - [https://github.com/kubernetes/kubernetes/pull/96129](https://github.com/kubernetes/kubernetes/pull/96129)

## Nov 3rd, 2020

Today’s SIG Node was cancelled due to the election.

* [@egernst] still need help with review:  adding pod usage to kubelet metrics/resource endpoint: <https://github.com/kubernetes/kubernetes/pull/95839>

## October 27th, 2020

Total active pull requests: [192](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+is%3Aopen+) (-8 from the last meeting)

**Incoming** | | **Completed** | |
---|---|---|---
|Created:    |[**24**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++created%3A2020-10-20T17%3A00%3A00%2B0000..2020-10-27T16%3A46%3A57%2B0000)                                            |Closed:      |[16](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++is%3Aunmerged+closed%3A2020-10-20T17%3A00%3A00%2B0000..2020-10-27T16%3A46%3A57%2B0000)|
|Updated:    |[**100**](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode+updated%3A2020-10-20T17%3A00%3A00%2B0000..2020-10-27T16%3A46%3A57%2B0000+created%3A%3C2020-10-20T17%3A00%3A00%2B0000)|Merged:      |[16](https://github.com/kubernetes/kubernetes/pulls?q=repo%3Akubernetes%2Fkubernetes+type%3Apr+label%3Asig%2Fnode++merged%3A2020-10-20T17%3A00%3A00%2B0000..2020-10-27T16%3A46%3A57%2B0000)              |

    * How do we handle runtimes that *may* sometimes require container restart in order to resize? How does that work with restartPolicy=Never
    * Should Restart be the default resize policy?
    * Should RuntimeClass [should be allowed to disable in-place update](https://github.com/kubernetes/enhancements/pull/1883/files#r448754705) so users get synchronous error. sounds reasonable to me if runtime doesn’t support this - thoughts?
    * Are RestartOnGrow, RestartOnShrink resize policies a good idea? (Doesn’t have to be added today, but quite easy to add if it feels useful)
* [mrunalp, nalin] /dev/fuse in pods [https://github.com/kubernetes/kubernetes/pull/79925](https://github.com/kubernetes/kubernetes/pull/79925)
  * Nalin and Mrunal to come up with use cases and design
* [mrunalp, zvonkok] Rlimits support - [https://github.com/kubernetes/kubernetes/issues/3595](https://github.com/kubernetes/kubernetes/issues/3595)
* [@rata/Kinvolk, @jirving] sidecar ordering ([KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/sidecarcontainers.md))
  * Switch KEP to provisionable state: [https://github.com/kubernetes/enhancements/pull/1874](https://github.com/kubernetes/enhancements/pull/1874)
* [@AlexeyPerevalov @swatisehgal] Topology Aware Scheduling
  * Topology Exporter Daemon in kubernetes-sig
* [harche/mrunalp] Fedora images for node testing
* [@yash97]Informal Kep:Sending Events to user directly to kubelet in distributed ways instead to receive it from api server.
* [alejandrox1] [https://github.com/kubernetes/kubernetes/pull/80917](https://github.com/kubernetes/kubernetes/pull/80917) needs review from SIG Node approver

## June 30, 2020

* [Balaji] Add support for disabling /logs endpoint
  * [https://github.com/kubernetes/kubernetes/pull/87273#issuecomment-649652855](https://github.com/kubernetes/kubernetes/pull/87273#issuecomment-649652855) would like to get some opinion on which option to pick. Handling backward compatibility complicated some of the things, so I’m thinking of going with enableSystemLogHandler requiring enableDebuggingHandler to be set to true as well. Looking for some suggestions and feedback here.
* [@vinaykul] [In-Place Pod Vertical Scaling](https://github.com/kubernetes/enhancements/blob/master/keps/sig-node/20181106-in-place-update-of-pod-resources.md) for v1.19 - Update &amp; next step
  * Feature code complete with unit tests, reviewed by @dashpole.
    * PR: [https://github.com/kubernetes/kubernetes/pull/92127](https://github.com/kubernetes/kubernetes/pull/92127)
  * Basic e2e test framework done, adding test-cases promised for alpha.
  * Need reviews @liggitt, @thockin, @ahg-g, and Derek, Dawn, sig-testing.
  * Identify process for informing CRI changes to runtime folks.
* [@mukesh-dua/@qerorsjoqld0237] CPU Reservation for Infrastructure Services Deployed in Kubernete
  * Would like to introduce the enhancement and its requirements.
* [@rata/Kinvolk, @alban/Kinvolk, @jirving], sidecar ordering ([KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/sidecarcontainers.md))
  * PR to update to provisional state and depend on kubelet node shutdown [in place](https://github.com/kubernetes/enhancements/pull/1874)
  * Reviewers/approvers?
    * SergeyKanzhelev
  * PR coming \~1 or 2 weeks (hopefully sooner) with design “callout” discussed in the previous meeting, alternatives and history of all decisions made in the past (WIP)
    * Plan to update summary, motivation, etc. too.
* [@hasheddan] seccomp-operator: [https://github.com/kubernetes/org/issues/1873](https://github.com/kubernetes/org/issues/1873)

## June 23, 2020

* [@harche]: GCP accounts for hosting custom images for node test
* [dawnchen] Introducing PoC from google for node images and critool/critest
  * Karan [GKE] and Sergey Kanzhelev
* [mrunal]: CRI missing detach for exec. See [https://github.com/kubernetes/kubernetes/issues/](https://github.com/kubernetes/kubernetes/issues/92057)[92057](https://github.com/kubernetes/kubernetes/issues/92057).
* [@rata, @alban](Kinvolk) sidecar ordering ([KEP](https://github.com/kubernetes/enhancements/blob/master/keps/sig-apps/sidecarcontainers.md))
  * How to raise concerns about KEP and have review/feedback from sig-node? To address all concerns at the KEP level
    * Yes, PR to KEP calling them out
  * Add KEP use cases?
    * Yes, PR to KEP
  * Should we make a PR to the KEP to depend on the kubelet graceful shutdown?
    * Yes, PR to KEP to set clear expectations to everyone
* [@alban, @mauriciovasquezbernal, @rata](Kinvolk) User namespaces [[design-proposal](https://github.com/kubernetes/community/blob/master/contributors/design-proposals/node/node-usernamespace-remapping.md)] [[issue](https://github.com/kubernetes/enhancements/issues/127)] [[PR](https://github.com/kubernetes/kubernetes/pull/64005)]
  * Reusing the upstream PR above, we’re working on a PoC based on Kubernetes 1.17 + containerd/cri 1.3
    * [Kubernetes PoC branch](https://github.com/kinvolk/kubernetes/pull/3)
* [k.wiatrzyk (krzwiatrzyk), c.zukowski, bg.chun] Enhancements proposals for 5G packet processing in Kubernetes.
  * Towards high-performance 5G packet processing, [slide](https://docs.google.com/presentation/d/1DueZ65OEPHvEMBf_nck7YOxWIXhOpYTIiPNix7rGIik/edit#slide=id.p)
  * Overall Proposals in detail, [docs](https://docs.google.com/document/d/1ASbIJNfeWmYrBIEMQl1MxHBPgvoztb_Fx5GvFhoPXF0/edit)

## March 31, 2020

* [tedyu] Protect log rotation against concurrent symlink removal
    [https://github.com/kubernetes/kubernetes/pull/89160](https://github.com/kubernetes/kubernetes/pull/89160)
* [harche] Add CRIO tests for Kubernes Prow (CI)
  * [https://github.com/kubernetes/test-infra/pull/17014](https://github.com/kubernetes/test-infra/pull/17014)
* [c.zukowski, cynepco3hahue, bg.chun] a warm invitation to leave your comments for a new component proposed in: [https://docs.google.com/document/d/1XvOUwI1DiZf0vTlIa2vPOT_QJYh-g3vpOa1D1pajqEs/edit?usp=sharing](https://docs.google.com/document/d/1XvOUwI1DiZf0vTlIa2vPOT_QJYh-g3vpOa1D1pajqEs/edit?usp=sharing)
    Memory Manager opens a series of further enablers of NUMA-aware resource management in Kubernetes. Do not hesitate to ask &quot;Why?&quot; questions, and point out whenever the document should expand on or stress something else (we will elaborate).
* [mrunalp] KEP for custom cri errors - [https://github.com/kubernetes/enhancements/pull/1654](https://github.com/kubernetes/enhancements/pull/1654)
* [jberkus] Discuss [year-of-support KEP](https://github.com/kubernetes/enhancements/pull/1497), particularly what the effect would be on required kubelet version skew
  * will just keep skew tests running for one more version
  * don't see this as serious problem
* [mattjmcnaughton/dims] KEP- Build Kubelet without Docker
  * [https://github.com/kubernetes/enhancements/pull/1546](https://github.com/kubernetes/enhancements/pull/1546)
    * KEP adds a “dockerless” tag to avoid docker/docker imports, builds kubelet without dockershim
  * PR [https://github.com/kubernetes/enhancements/pull/1546](https://github.com/kubernetes/enhancements/pull/1546)
* Need a way to mock network related testing for example to test issue:
  * [https://github.com/kubernetes/kubernetes/issues/88543](https://github.com/kubernetes/kubernetes/issues/88543)

## March 24, 2020

* [mrunalp] CRI custom errors update[https://docs.google.com/presentation/d/1DdmlEmK4TkfJWW19wXPReAw8MA0oRdzPwDDhjtDKJo8/edit#slide=id.g8190ad4d27_0_38](https://docs.google.com/presentation/d/1DdmlEmK4TkfJWW19wXPReAw8MA0oRdzPwDDhjtDKJo8/edit#slide=id.g8190ad4d27_0_38)
* [mikebrow] Ensure Secret Pulled Images (without pull always)
  * [https://github.com/kubernetes/enhancements/pull/1608](https://github.com/kubernetes/enhancements/pull/1608)
* [decarr] Identifying areas to improve kubelet reliability across contributor
  * A number of subtle prs have been opened recently around edge cases on pod status reporting, wanted us to have a plan of record as a group on what we tackle so we can have full context when reviewing.  They are non-trivial to review given their subtle updates, so just having a dialog on what the rationale will help all.
  * Examples:
    * [https://github.com/kubernetes/kubernetes/pull/89155/files](https://github.com/kubernetes/kubernetes/pull/89155/files) which targets issue #80968
    * Other
* [decarr] flaky node conformance [help desired]
  * [https://github.com/kubernetes/kubernetes/issues/85762#issuecomment-601273307](https://github.com/kubernetes/kubernetes/issues/85762#issuecomment-601273307)
  * [https://github.com/kubernetes/kubernetes/pull/89379/files](https://github.com/kubernetes/kubernetes/pull/89379/files)
* [keerthan] Issue: [https://github.com/kubernetes/kubernetes/issues/88543](https://github.com/kubernetes/kubernetes/issues/88543)
  * [dashpole] test do create delete pod through kubelet and interact with cri to ensure its gone (similar to checking inside cgroups for kubelet settings)
* [dashpole] Out-of-Pid eviction update [https://github.com/kubernetes/kubernetes/pull/89359](https://github.com/kubernetes/kubernetes/pull/89359).  Pod-level PID stats in the summary API?  KEP update: [https://github.com/kubernetes/enhancements/pull/1631](https://github.com/kubernetes/enhancements/pull/1631)
* [tallclair] RuntimeClass to GA
* [smarterclayton] Doc pending on Kubelet status suggestions (ran out of time last week), will email sig-node with a shared google doc soon
* [smarterclayton] Discuss exposing calculated resource limits as metrics briefly - too hard today to emulate the complex calculations scheduler and node do for resource limit in external monitoring tool

## March 17, 2020

* [vpickard] Please upload latest video recording
* [liorokman] Enable defining Pod-level resource limit
  * [https://github.com/kubernetes/enhancements/pull/1592](https://github.com/kubernetes/enhancements/pull/1592)
  * Way forward is to implement this using a daemonset that manipulates the cgroups externally to Kubernetes, and report back once there is some real-world experience around this feature.

## March 10, 2020