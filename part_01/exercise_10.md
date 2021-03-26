## Part 1 exercise 10

Ticker apps located in [/apps/ticker-generator](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps/ticker-generator) and [/apps/ticker-server](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps/ticker-server) at commit **placeholder**. Published through Docker Hub at [mtuomiko/ticker-generator](https://hub.docker.com/r/mtuomiko/ticker-generator) with tag `0.0.1` and [mtuomiko/ticker-server](https://hub.docker.com/r/mtuomiko/ticker-server) with tag `0.0.1`.

1. Update the deployment with new apps `kubectl apply -f manifests/ticker/deployment.yaml`. This should stop the previous ticker app and start the generator and the server.

2. I forgot to push the images to Docker Hub so a new pod ended up in `ImagePullBackOff` state. After pushing the images I ran `kubectl rollout restart deployment ticker-dep` since nothing seemed to happen. Although the situation would probably correct itself after some time?

3. Apply service file `kubectl apply -f manifests/ticker/service.yaml`

```
service/ticker-svc unchanged
```

4. Okay so no changes. Let's update Ingress `kubectl apply -f manifests/ingress.yaml` which should route `/status/` to the ticker-server.

5. Let's try it out with `curl localhost:8081/status/status`. Second status needed since Ingress will strip the status prefix and our server is still listening to `/status/` at its end.

```
2021-03-26T11:42:23Z: 1e1d0ca7-8e26-11eb-8cde-ea82e332dc3e
```

6. Cool. We'll wait a bit and run it again.

```
2021-03-26T11:44:03Z: 1e1d0ca7-8e26-11eb-8cde-ea82e332dc3e
```

7. We get a new timestamp so everything seems to work.