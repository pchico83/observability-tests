* Configure a kubernetes cluster with two nodes (could this be automated).

* Configure a sut running the docker image `pchico83/ksync` on port 8080 and setting the environment variables to access the external kubernetes cluster.

* Call the sut api server on path `/create` to start a ksync folder.

* `ksync` prints the node where the sync is happening. Obtain this value via EMS.

* Check that sync is working.

* Remove the node wjere the sync is not happening.

* Check that sync is not working anymore.

