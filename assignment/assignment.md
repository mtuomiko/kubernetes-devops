## Unity assignment

App and manifests located located in [/assignment](https://github.com/mtuomiko/kubernetes-devops/tree/main/assignment)

1. `kube-prometheus-stack` was already installed in part 2
2. Created a new namespace `assignment`
3. Created Deployment, Service and Ingress (for making requests outside the cluster) in the new namespace
4. Create PodMonitor inside the `prometheus` namespace (Prometheus by default discovers Monitors within its namespace, but only when labeled with the same release tag)
5. We need to change the `prometheus.prometheusSpec.podMonitorSelectorNilUsesHelmValues` setting to `false` to allow Prometheus the find our PodMonitor (since it's not labeled with the same release tag). `helm upgrade -n prometheus kube-prometheus-stack-1617221463 prometheus-community/kube-prometheus-stack --set prometheus.prometheusSpec.podMonitorSelectorNilUsesHelmValues=false`
6. We finally can query e.g. the `cat_view_counter` in Prometheus.
7. Add Grafana dashboard with the metrics using rate() over a 3 minute period (e.g. `rate(cat_view_counter[3m])`)
![Dashboard](/grafana.png)

#### Other stuff

- Port forward (for Prometheus) `kubectl -n <namespace> port-forward prometheus-kube-prometheus-stack-...-prometheus-0 9090`
- Port forward (for Grafana) `kubectl -n <namespace> port-forward kube-prometheus-stack-...-grafana-... 3000`

