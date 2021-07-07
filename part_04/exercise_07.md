## Part 4 exercise 7

Manual operations needed for clean startup:

- Create cluster: `k3d cluster create -p 8081:80@loadbalancer --agents 2`
- Create storage folder for PersistentVolume: `docker exec k3d-k3s-default-agent-0 mkdir -p /tmp/storage`
- Setup GITHUB_TOKEN for flux: `export GITHUB_TOKEN=foo`
- Bootstrap to GitHub: `flux bootstrap github --owner=mtuomiko --repository=kubernetes-devops --personal --private=false --path=cluster --branch=main`
- Get Sealed Secrets public key: `kubeseal --fetch-cert --controller-name=sealed-secrets --controller-namespace=flux-system > pub-sealed-secrets.pem`
- Use public key to create SealedSecrets, for example: `kubeseal -o yaml --cert=pub-sealed-secrets.pem < secret.yaml > sealedsecret.yaml`
- Add SealedSecrets to Git and push
- Add applications and services