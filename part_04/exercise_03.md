## Part 4 exercise 3

Prometheus query

`scalar(count(kube_pod_info{namespace="prometheus",created_by_kind="StatefulSet"}))`