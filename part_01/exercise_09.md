## Part 1 exercise 9

Todo-app located in [/apps/pingpong](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps/pingpong) at commit **placeholder**. Published through Docker Hub at [mtuomiko/pingpong](https://hub.docker.com/r/mtuomiko/pingpong) with tag `0.0.1`

1. Create deployment for pingpong `kubectl apply -f manifests/deployment.yaml`

2. Create service for pingpong `kubectl apply -f manifests/service.yaml`

3. Update Ingress (note: Ingress manifest now located at apps root since affects multiple apps) `kubectl apply -f ../manifests/ingress.yaml` 

4. Check that todo-app still works `curl localhost:8081`

```
<!DOCTYPE html>
<html lang="en">
...
```

5. Check that pingpong works `curl localhost:8081/pingpong`

```
pong 0
```

6. Again `curl localhost:8081/pingpong`

```
pong 1
```
