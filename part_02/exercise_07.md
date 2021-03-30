## Part 2 exercise 7

Apps located located in [/apps](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps) at commit **placeholder**. `pingpong` is published through Docker Hub at [mtuomiko/pingpong](https://hub.docker.com/r/mtuomiko/pingpong) with tag `0.0.4`.

Secret `DB_PASSWORD` is inserted into `sealedsecret_exercises.yaml` which is, after decrypting inside the cluster, then used by `pingpong` and `pingpong-db` deployments. 

Database deployment uses `DB_PASSWORD` as `POSTGRES_PASSWORD` which is understood and used by the image.

I went with just the Deployment route for DB since that seemed like a sufficient solution (no need for multiple pods).