## Part 1 exercise 8

Todo-app located in [/apps/todo-app](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps/todo-app) at commit https://github.com/mtuomiko/kubernetes-devops/commit/b20e3f1d8941f12c833ce2f58cb2fa3cb372440f. Published through Docker Hub at [mtuomiko/todo-app](https://hub.docker.com/r/mtuomiko/todo-app) with tag `0.0.4`

1. Get services `kubectl get service`

```
NAME           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
kubernetes     ClusterIP   10.43.0.1       <none>        443/TCP          12h
todo-app-svc   NodePort    10.43.198.211   <none>        5000:32130/TCP   12h
ticker-svc     ClusterIP   10.43.130.211   <none>        2345/TCP         17m
```

2. Delete previous NodePort service for todo-app `kubectl delete service todo-app-svc`

3. Apply new ClusterIp service for todo-app `kubectl apply -f manifests/service.yaml`

4. Update ingress with `kubectl apply -f manifests/ingress.yaml` where the service name and port are updated

5. Try it out with `curl localhost:8081`

```
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  ...
```
