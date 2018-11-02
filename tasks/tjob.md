* Configure a kubernetes cluster with two nodes (could this be automated).

* Configure a sut running the docker compose in this repo.

* Call the sut api server on path `/create` to start an operation.

* `tasks` prints the container exxecuting the add operation. Obtain this value via EMS.

* Kill the container.

* Wait for the operations to finish (60s).

* Check the value in the database. It should be 21.

