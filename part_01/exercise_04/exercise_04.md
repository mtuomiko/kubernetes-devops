## Part 1 exercise 4

Todo-app with deployment.yaml located in subfolder, published through Docker Hub at [mtuomiko/todo-app](https://hub.docker.com/r/mtuomiko/todo-app)

1. Delete previous deployment with `kubectl delete deployment todo-app-dep`

2. Realise you could have tried updating the deployment. We'll do it later.

3. Apply new deployment file `kubectl apply -f manifests/deployment.yaml` with tag 0.0.2

4. Get pods with `kubectl get pods`

```
NAME                            READY   STATUS    RESTARTS   AGE
ticker-dep-69d6f84486-klgb2    1/1     Running   0          10m
todo-app-dep-8bfbf886b-244cs   1/1     Running   0          6s
```

5. See logs with `kubectl logs -f todo-app-dep-8bfbf886b-244cs`

```
2021/03/22 19:07:43 Server started in port 5678
```

6. Create and push a new image with tag 0.0.3 and change message to *started **ON** port* instead of *started **IN** port*

7. Apply deployment file `kubectl apply -f manifests/deployment.yaml` this time with the 0.0.3 tag

8. Get pods with `kubectl get pods` and see that there is something going on (!)

```
NAME                            READY   STATUS        RESTARTS   AGE
ticker-dep-69d6f84486-klgb2     1/1     Running       0          15m
todo-app-dep-58cbbfcb97-v88gh   1/1     Running       0          6s
todo-app-dep-8bfbf886b-244cs    1/1     Terminating   0          4m57s
```

9. Check the logs from the new pod `kubectl logs -f todo-app-dep-58cbbfcb97-v88gh`

```
2021/03/22 19:12:34 Server started on port 5678
```
