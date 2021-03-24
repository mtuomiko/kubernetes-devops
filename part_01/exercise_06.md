## Part 1 exercise 6

Todo-app located in [/apps/todo-app](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps/todo-app) at commit https://github.com/mtuomiko/kubernetes-devops/commit/e94cd4d6fc714359c0da053aea52f44b77b27b84. Published through Docker Hub at [mtuomiko/todo-app](https://hub.docker.com/r/mtuomiko/todo-app) with tag `0.0.4`

1. Delete previous cluster `k3d cluster delete`

2. Create a new one `k3d cluster create --port '8082:32130@agent[0]' -p 8081:80@loadbalancer --agents 2`

3. Apply deployment file `kubectl apply -f manifests/deployment.yaml`

4. Apply service file `kubectl apply -f manifests/service.yaml`

5. Try it out with `curl localhost:8082`

```
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  ...
```
