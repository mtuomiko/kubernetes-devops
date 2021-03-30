## Part 2 exercise 1

Updated `pingpong` and `ticker-server` apps located in [/apps](https://github.com/mtuomiko/kubernetes-devops/tree/main/apps) at commit https://github.com/mtuomiko/kubernetes-devops/commit/e96cceb2b1223671157b850442c712019dbe21eb. Published through Docker Hub at [mtuomiko/pingpong](https://hub.docker.com/r/mtuomiko/pingpong) with tag `0.0.3` and [mtuomiko/ticker-server](https://hub.docker.com/r/mtuomiko/ticker-server) with tag `0.0.3`.

1. I removed the persistent storage from pingpong since since it wasn't a clear requirement for this exercise. Pingpong count is still cached into a JSON file in emptyDir volume.
