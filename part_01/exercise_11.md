## Part 1 exercise 11

Apps located in [/apps/ticker-server](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps/ticker-server) and [/apps/pingpong](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps/pingpong) at commit https://github.com/mtuomiko/kubernetes-devops/commit/80513372d791ebca57c2248546c460c49e7d7b85. Published through Docker Hub at [mtuomiko/ticker-server](https://hub.docker.com/r/mtuomiko/ticker-server) with tag `0.0.2` and [mtuomiko/pingpong](https://hub.docker.com/r/mtuomiko/pingpong) with tag `0.0.2`.

1. The persistent volume directory was created on agent-0 node where our apps seem to run, but that's not a safe assumption. We'll go with it for these exercises. Let's create the volume and the claim with `kubectl apply -f manifests/persistentvolume.yaml` and `kubectl apply -f manifests/persistentvolumeclaim.yaml`

2. We'll update the pingpong and ticker deployments `kubectl apply -f pingpong/manifests/deployment.yaml` and `kubectl apply -f manifests/ticker/deployment.yaml`

3. Pingpong will create the count file upon starting. After this the ticker-server can read the pingpong count and use it on response.

4. Does it work? `curl localhost:8081/status/status`

```
2021-03-26T17:08:33Z: e7fcf650-8e54-11eb-899c-46cf550a329a
Ping / Pongs: 0
```

5. Looking good. Let's ping some pongs with `for ((i=1;i<=10;i++)); do curl localhost:8081/pingpong; done`

```
...
pong 8
pong 9
pong 10
```

6. `curl localhost:8081/status/status`

```
2021-03-26T17:10:13Z: e7fcf650-8e54-11eb-899c-46cf550a329a
Ping / Pongs: 11
```

7. Let's test the persistence by deleting both deployments and creating them again with `kubectl delete deployment ticker-dep` & `kubectl delete deployment pingpong-dep` & `kubectl apply -f pingpong/manifests/deployment.yaml` & `kubectl apply -f manifests/persistentvolumeclaim.yaml`. After all that we'll check the ping pong count with `curl localhost:8081/status/status`

```
2021-03-26T17:10:13Z: e7fcf650-8e54-11eb-899c-46cf550a329a
Ping / Pongs: 11
```
