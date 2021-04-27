## Part 3 exercise 2

Apps and manifests located in [/apps/](https://github.com/mtuomiko/kubernetes-devops/tree/ex_03-02/apps) at branch [ex_03-02](https://github.com/mtuomiko/kubernetes-devops/tree/ex_03-02).

Commands for recreating the cluster after deletion. Maybe should do a script for this if this becomes a regular thing. Zone assumed to be specified in gcloud configuration.
```
gcloud container clusters create dwk-cluster
gcloud container clusters get-credentials dwk-cluster
kubectl apply -n kube-system -f https://github.com/bitnami-labs/sealed-secrets/releases/download/v0.15.0/controller.yaml
kubectl create namespace exercises
kubens exercises
kubeseal -o yaml < secret_exercises.yaml > sealedsecret_exercises_gke.yaml
kubectl apply -f sealedsecret_exercises_gke.yaml
kubectl apply -f configmap_exercises.yaml
```

After that apply the needed configuration files.

I had some trouble with this exercise regarding the health checks and exact routes. I assumed that having a HTTP 200 OK root route for the Ingress would be enough but naturally the health of each of the services are monitored. I ended up creating a dummy health check response in the pingpong application that responds on the local `/` route to get a green status for the service which enabled Ingress to actually route traffic to `pingpong`. Additionally the pingpong app responds on `/pingpong/` and `/pingpong/pingpongs` for JSON. Main app (`ticker`) responds on Ingress root `/`.
