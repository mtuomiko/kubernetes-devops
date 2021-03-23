## Part 1 exercise 2

Todo-app located in [/apps/todo-app](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps/todo-app) at commit https://github.com/mtuomiko/kubernetes-devops/commit/409693fdf52297423c4181d4cb5eebf03f635ae0. Published through Docker Hub at [mtuomiko/todo-app](https://hub.docker.com/r/mtuomiko/todo-app)

1. Create a deployment with `kubectl create deployment todo-app-dep --image=mtuomiko/todo-app`

2. Get pods with `kubectl get pods`

```
NAME                            READY   STATUS    RESTARTS   AGE
ticker-dep-55dbc9bb68-l7bvd     1/1     Running   0          173m
todo-app-dep-5f46796d5b-4nxx9   1/1     Running   0          7s
```

3. See logs with `kubectl logs -f todo-app-dep-5f46796d5b-4nxx9`

```
2021/03/22 18:11:32 Server started in port 5678
```