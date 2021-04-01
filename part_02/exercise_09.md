## Part 2 exercise 9

Apps and manifests located located in [/apps](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps) at commit https://github.com/mtuomiko/kubernetes-devops/commit/dabd5da4cdbfbc9842ba1fb3cccd9632cdebd0c5. `todo-generator` is published through Docker Hub at [mtuomiko/todo-generator](https://hub.docker.com/r/mtuomiko/todo-generator) with tag `0.0.2`.

Todo generator is a Go program which is set to run at 6:15 (UTC) every day so that would be 9:15 local time (at least as I'm writing this right now on 31.3.2021). Gotta love timezones. Todo backend url is provided via `configmap_project.yaml`.
