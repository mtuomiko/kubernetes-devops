## Part 2 exercises 3 & 4

Apps located located in [/apps](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps) at commit **placeholder**. Published through Docker Hub yada yada etc.

1. Namespaces created: `project` for the project and `exercises` for everything else. Everything is now defined in the corresponding namespace: persistent volume and claim for the ticker, deployments, services etc. 

2. These exercises included only one change in the actual applications: `ticker-server` which listens at external path `/status/status` was changed so that the host from which the pingpong count is retrieved was `pingpong-svc.exercises` instead of just `pingpong-svc`. This might have worked either way?

3. Had a bit of trouble with Ingress since I first tried to use a single Ingress instance in `default` namespace which would route to other namespaces with services defined by `externalName`. This did not succeed in a reasonable amount of time and rather than trying to force it I tried just using one Ingress inside each namespace which actually worked even though I do not completely understand why. My motivation for originally using only a single Ingress was so that I could keep the single external endpoint :8081.