## Part 3 exercise 4

GKE project files located in [/gke_project](https://github.com/mtuomiko/kubernetes-devops/tree/feat-new/gke_project) at branch [feat-new](https://github.com/mtuomiko/kubernetes-devops/tree/feat-new). Workflow located in [gke_project.yaml](https://github.com/mtuomiko/kubernetes-devops/blob/feat-new/.github/workflows/gke_project.yaml).

EDIT: `feat-new` has since been deleted in exercise 5. Relevant changes should be present in the `main` branch.

I got stuck on SealedSecrets for a while until vicious googling told what the actual problem was: by default a SealedSecret has its namespace and name locked as in they are baked as part of the encrypted data and cannot be changed afterwards. Providing a cluster-wide scope flag for kubeseal (`kubeseal -o yaml --scope cluster-wide < secret.yaml > sealedsecret.yaml`) creates SealedSecrets that can be accessed in other namespaces. I'm sure there are other solutions for this issue but this was an easy fix for this particular situation since allowing cluster-wide access is not a security risk because there are no other (non-admin) users.
