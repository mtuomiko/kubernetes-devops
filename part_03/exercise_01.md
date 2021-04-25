## Part 3 exercise 1

Apps and manifests located in [/apps/](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps/) and at branch **placeholder**.

We're going to need:
- a Deployment for the actual `pingpong` application and a StatefulSet for `pingpong-db` (this was was previously done with a Deployment but we'll switch over to a StatefulSet just for fun 😺)
- a hostname for the `pingpong-db` through a Service so `pingpong` can contact the database.
- a sealed secret for the database password
- a load balancer for connecting from outside the cluster

Steps

1. Recreate the sealed secret for our cluster `kubeseal -o yaml < secret_exercises.yaml > sealedsecret_exercises_gke.yaml`

2. Install Sealed Secrets to the cluster in `kube-system` namespace. Then apply our new sealed secret with the DB password (/apps/manifests/sealedsecret_exercises_gke.yaml)

3. Create `pingpong-db` StatefulSet and Service. We needed the instructed `subPath` property to get the DB running. Looks like 1Gi is the minimum size for storage since our 256Mi request became such. The actual VolumeClaim was created automatically as told in the GKE instructions.

```
NAME                                       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM                                   STORAGECLASS   REASON   AGE
pvc-6efd0afb-8a4a-427e-87a4-58da5a439c3e   1Gi        RWO            Delete           Bound    exercises/db-storage-pingpong-db-ss-0   standard                10m

NAME                          STATUS   VOLUME                                     CAPACITY   ACCESS MODES   STORAGECLASS   AGE
db-storage-pingpong-db-ss-0   Bound    pvc-6efd0afb-8a4a-427e-87a4-58da5a439c3e   1Gi        RWO            standard       11m
```

4. Create `pingpong` Deployment and LoadBalancer Service.

5. Once we have an external IP, let's try it out with multiple `curl http://35.228.253.162:5500`

```
pong 6
pong 7
pong 8
...
```


