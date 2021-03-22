## Part 1 exercise 3

Ticker/hashgenerator app with deployment.yaml located in subfolder, image published through Docker Hub at [mtuomiko/todo-app](https://hub.docker.com/r/mtuomiko/ticker)

1. Delete previous deployment with `kubectl delete deployment ticker-dep`

2. Apply new deployment file `kubectl apply -f manifests/deployment.yaml` with new tag 0.0.2

3. Get pods with `kubectl get pods`

```
NAME                            READY   STATUS    RESTARTS   AGE
todo-app-dep-5f46796d5b-4nxx9   1/1     Running   0          45m
ticker-dep-69d6f84486-klgb2     1/1     Running   0          3s
```

3. See logs with `kubectl logs -f ticker-dep-69d6f84486-klgb2`

```
2021-03-22T18:57:17Z: 60689fff-8b40-11eb-86ed-1a35c9acf31f
2021-03-22T18:57:22Z: 60689fff-8b40-11eb-86ed-1a35c9acf31f
2021-03-22T18:57:27Z: 60689fff-8b40-11eb-86ed-1a35c9acf31f
2021-03-22T18:57:32Z: 60689fff-8b40-11eb-86ed-1a35c9acf31f
...
```