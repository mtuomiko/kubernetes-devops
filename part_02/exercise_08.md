## Part 2 exercise 8

Apps and manifests located located in [/apps](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps) at commit **placeholder**. `todo-app-backend` is published through Docker Hub at [mtuomiko/todo-app-backend](https://hub.docker.com/r/mtuomiko/todo-app-backend) with tag `0.0.4`.

Database `todo-app-db` is based on `postgres:alpine` image and is created as a StatefulSet. Environment variables for database creation and utilization are passed from a ConfigMap `configmap_project` and from a SealedSecret `sealedsecret_project.yaml` to both the DB and the backend.