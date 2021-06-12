## Part 3 exercise 8

I had some issues with not being able to fit all the pods on the cluster since at first I didn't realize that without explicit requests the resource limits will become the requests. Additionally I started with a 2000 milliCPU request (actually limit, but anyway) for the database container but even though the nodes had 2 vCPUs all the extra Kubernetes workloads mean that there isn't a completely free node to run our 2 CPU request pod.
