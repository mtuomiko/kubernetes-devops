## Part 1 exercise 1

Ticker app located in [/apps/ticker](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps/ticker) at commit https://github.com/mtuomiko/kubernetes-devops/commit/8c41f8f7d13dd933c71b68e7f695d06f4b734b76. Published through Docker Hub at [mtuomiko/ticker](https://hub.docker.com/r/mtuomiko/ticker)

1. Create a cluster with `k3d cluster create -a 1` (one agent node)

2. Create a deployment with `kubectl create deployment ticker-dep --image=mtuomiko/ticker`

3. Get pods with `kubectl get pods`

```
NAME                          READY   STATUS    RESTARTS   AGE
ticker-dep-55dbc9bb68-l7bvd   1/1     Running   0          11s
```

4. See ticker logs with `kubectl logs -f ticker-dep-55dbc9bb68-l7bvd`

```
2021-03-22T15:17:56Z: c4be0931-8b21-11eb-b72f-4e161d9536ee
2021-03-22T15:18:01Z: c4be0931-8b21-11eb-b72f-4e161d9536ee
2021-03-22T15:18:06Z: c4be0931-8b21-11eb-b72f-4e161d9536ee
...
```