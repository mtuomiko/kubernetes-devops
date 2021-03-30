## Part 2 exercise 2

Apps `todo-app-backend` and `todo-app-frontend` located located in [/apps](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps) at commit https://github.com/mtuomiko/kubernetes-devops/commit/3fa729be16391cad0b14dc6f50bcd64e3026a752. Published through Docker Hub at [mtuomiko/todo-app-backend](https://hub.docker.com/r/mtuomiko/todo-app-backend) with tag `0.0.3` and [mtuomiko/todo-app-frontend](https://hub.docker.com/r/mtuomiko/todo-app-frontend) with tag `0.0.1`.

1. Todo app was split to a React frontend served from a nginx-alpine image and a Golang Echo framework backend. Image logic was moved to the backend. Backend is routed by Ingress on the path `/api`. Previous todo-app deployment and service were deleted.
