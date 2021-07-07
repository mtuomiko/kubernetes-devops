## Part 4 exercise 7

Manual operations needed:

- Create cluster: `k3d cluster create -a 2`
- Create storage folder for PersistentVolume: `docker exec k3d-k3s-default-agent-0 mkdir -p /tmp/storage`
- Setup GITHUB_TOKEN for flux: `export GITHUB_TOKEN=bla`
- Bootstrap to GitHub `flux bootstrap github --owner=mtuomiko --repository=kubernetes-devops --personal --private=false --path=cluster --branch=main`
- Add applications