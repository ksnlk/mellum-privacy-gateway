
        - Calendar to be finalized soon

        - Community feedback is that it’s better to have more time for
          development

**  
Open Discussion:**

- \[Sascha\] Possible image signing enhancements:

  - Supporting of recursive signing of multi-arch images in kpromo:  
    [<u>https://github.com/kubernetes-sigs/release-sdk/pull/193</u>](https://github.com/kubernetes-sigs/release-sdk/pull/193)  
    [<u>https://github.com/kubernetes-sigs/release-sdk/releases/tag/v0.10.1</u>](https://github.com/kubernetes-sigs/release-sdk/releases/tag/v0.10.1)  
    [<u>https://github.com/kubernetes-sigs/promo-tools/pull/868</u>](https://github.com/kubernetes-sigs/promo-tools/pull/868)

  - RFC: Being able to set the docker-reference to registry.k8s.io in
    image signatures:  
    [<u>https://github.com/sigstore/cosign/pull/2984</u>](https://github.com/sigstore/cosign/pull/2984)  
    Right now we still point them to the actual mirror, which makes
    signature validation hard:  
    \> cosign verify \\  
    --certificate-identity
    krel-trust@k8s-releng-prod.iam.gserviceaccount.com \\  
    --certificate-oidc-issuer https://accounts.google.com \\  
    registry.k8s.io/kube-apiserver:v1.27.1 \| jq
    .\[0\].critical.identity  
    …  
    {"docker-reference":"europe-southwest1-docker.pkg.dev/k8s-artifacts-prod/images/kubernetes/kube-apiserver"}

  <!-- -->

  - \[Sascha\] We don’t use recursive signing, so in case of multi-arch
    images, we sign the manifest, but don’t sign concrete images. Added
    support for recursive signing to fix this, will be introduced in
    promo-tools soon. Hopefully it’s not going to increase the promotion
    time, but we’ll properly sign and validate images

  - \[Sascha\] We want to be able to replace docker-reference with
    registry.k8s.io, but this requires changing cosign, but it’ll be
    better for users. We’ll see how it turns out, depends on if the
    change lands in cosign

  - \[Sascha\] Those are considered small changes, we don’t need to
    change our current enhancement

  - \[Jim\] What do we need to move with this?

    - \[Sascha\] Proposed change as a RFC, we need to wait for feedback.
      Without it, we don’t know how to make it better for users. If you
      make it really strict, it’s wrong, because the reference is not
  - [<u>SIG Release issues
    needs-triage</u>](https://github.com/search?q=org%3Akubernetes+is%3Aopen+is%3Aissue+label%3Asig%2Frelease+label%3Aneeds-triage)

  - [<u>SIG Release PRs
    needs-triage</u>](https://github.com/search?q=org%3Akubernetes+is%3Aopen+is%3Apr+label%3Asig%2Frelease+label%3Aneeds-triage)

  - [<u>Kubernetes
    issues</u>](https://github.com/kubernetes/kubernetes/issues?q=is%3Aopen+is%3Aissue+label%3Asig%2Frelease)

## Mar 28, 2023

**Host (pronouns):** Jeremy Rickard (he/him)

**Attendees (pronouns):**

- Marko Mudrinić (he/him)

- [<u>Adolfo García Veytia</u>](mailto:puercozon@gmail.com)(he/him)

- [<u>Grace Nguyen</u>](mailto:nng.grace@gmail.com) (she/her)

- Laura Lorenz (she/her)

- Benjamin Elder (he/him)

- [<u>Ashwin Kumar Uppala</u>](unlock_code:myrprwamjrx7169@gmail.com)(He/Him)

**Note Taker (pronouns):**

- Joseph Sandoval

**Recurring Topics \[timebox to 20 min\]:**

- Welcome any new members or attendees

  - Ashwin is checking out how things are working. Joining to follow
    along. Observe the meeting with the time he has.

  - Josh Berkus / Sig-Contribex co-chair.

  - Adolfo/Puerco - technical lead

  - Jeremy - does release team stuff

  - Marco - Release Mgr

  - Cici - 1.27 Branch Mgr

  - Mahamed - Contributor to Sig-release and Sig-K8s infra

  - Joseph - Release Mgr Assoc