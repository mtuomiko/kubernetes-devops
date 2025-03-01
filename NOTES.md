### #k3d traefik dashboard access

Dashboard enabled automatically but not open.

Manual access

kubectl -n kube-system get pods --selector "app.kubernetes.io/name=traefik" --output=name

get the name, and do a direct forward on port 9000

kubectl -n kube-system port-forward pod/traefik-57b79cf995-h5v64 9000:9000

http://localhost:9000/dashboard/ (note last slash)

### build shorthand functions

```
db() {
    app_name=${PWD##*/}
    version=$(cat version | tr -d '\r\n')

    set -x
    docker build -t mtuomiko/${app_name}:${version} -f Containerfile .
    set +x
}

dbt() {
    app_name=${PWD##*/}
    version=$(cat version | tr -d '\r\n')

    set -x
    docker push mtuomiko/${app_name}:${version}
    set +x
}
```
