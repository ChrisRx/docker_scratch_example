# docker_scratch_example
example using static compiled Golang binary and scratch docker image to produce small docker containers

Take a look at the `Makefile` to see how it works.  Included in the `build-cgo` target is the go build options necessary to ensure that even when dealing with cgo code that the go executable is statically compiled. This tiny [echo](https://github.com/labstack/echo) server doesn't include any C, but figured I would put it in there in as a reference.

Running `docker images | grep tiny-server` then shows that the resulting docker image is only ~9mb (uncompressed).
