## Part 1 exercise 12

App located in [/apps/todo-app](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps/todo-app) at commit **placeholder**. Published through Docker Hub at [mtuomiko/todo-app](https://hub.docker.com/r/mtuomiko/todo-app) with tag `0.0.9`.

1. Update deployment with emptyDir volume defined `kubectl apply -f manifests/deployment.yaml`

2. We get an static image on main page that stays the same even with refreshes. Let's try a restart with `kubectl rollout restart deployment todo-app-dep`

3. That does change the image since the volume is dependent on the pod. However I believe this solution satisfies the exercise requirement because the container crashing should cause only the container to restart.
