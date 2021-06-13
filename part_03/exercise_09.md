## Part 3 exercise 9

Added some resource requests and limits to the pod descriptions in ticker Deployment, pingpong Deployment and pingpong-db StatefulSet. I also noticed that I had an issue with applying a new `deployment.yaml` files for the ticker app as I was trying to set the resources limits. The ReadWriteOnce PersistentVolumeClaim used by the deployment caused the new pod in a rolling update to halt if it was deployed in another node. I made a patchy solution for this with a PodAffinity setting.
