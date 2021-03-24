## Part 1 exercise 7

Ticker app located in [/apps/ticker](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps/ticker) at commit **placeholder**. Published through Docker Hub at [mtuomiko/todo-app](https://hub.docker.com/r/mtuomiko/todo-app) with tag `0.0.3`

Our previous cluster should work fine since we already port routing to port 80 at the load balancer.

1. Update the deployment `kubectl apply -f manifests/deployment.yaml` with tag 0.0.3

2. Apply service file `kubectl apply -f manifests/service.yaml`

3. Apply ingress file `kubectl apply -f manifests/ingress.yaml`

```
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
ingress.extensions/dwk-ingress created
```

4. We'll fix that deprecation warning but lets first try this out with `curl localhost:8081/status`

```
2021-03-24T08:57:22Z: e64acbde-8c7e-11eb-9893-ce74434f189e
```

5. Seems to work OK. Switching to `networking.k8s.io/v1` required some minor changes to the ingress file.