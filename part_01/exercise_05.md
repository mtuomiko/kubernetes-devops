## Part 1 exercise 5

Todo-app located in [/apps/todo-app](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps/todo-app) at commit https://github.com/mtuomiko/kubernetes-devops/commit/8d2442cf9d8781a92c0e322d3ffba512adefa98e. Published through Docker Hub at [mtuomiko/todo-app](https://hub.docker.com/r/mtuomiko/todo-app) with tag `0.0.4`

1. Apply new deployment file `kubectl apply -f manifests/deployment.yaml` with tag 0.0.4

2. Get pods with `kubectl get pods`

```
NAME                            READY   STATUS    RESTARTS   AGE
ticker-dep-69d6f84486-klgb2     1/1     Running   1          23h
todo-app-dep-6b595cdf8d-rlwxm   1/1     Running   0          12s
```

3. Set up port forwarding `kubectl port-forward todo-app-dep-6b595cdf8d-rlwxm 5678:5678`

```
Forwarding from 127.0.0.1:5678 -> 5678
Forwarding from [::1]:5678 -> 5678
```

4. Check `curl localhost:5678`

```
<!DOCTYPE html>
<html lang="en">

<head>
  <meta charset="UTF-8">
  <meta http-equiv="X-UA-Compatible" content="IE=edge">
  <meta name="viewport" content="width=device-width, initial-scale=1.0">
  <title>Main page</title>
</head>

<body>
  <h1>Main page</h1>
</body>

</html>
```
