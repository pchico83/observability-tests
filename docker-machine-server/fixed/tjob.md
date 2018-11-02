* Configure a sut running the docker image `pchico83/docker-machine-server:fixed` on port 8080 and setting the environment variables `AWS_ACCESS_KEY`, `AWS_SECRET_KEY`.

* Call the sut api server on path `/create` to start a docker host on AWS.

* `/create` prints the host ip, which is received by the tjob using EMS.

* Check that the server is accessible calling the endpoint `/status`.

* Call the sut api server on path `/restart` to start a docker host on AWS.

* `/restart` prints the host ip, which is received by the tjob using EMS.

* Call `/restart` until the host ip is different than the one returned by `/create` (max 5 iterations).

* Check that the server is accessible calling the endpoint `/status`. It should success.
