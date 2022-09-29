### About Project
ping-pong is an api with one post handler for route /ping
that accepts any request with a json body { times: x }, and 
responses with a json data { "pongs": ["pong",...] } that contains
array of pongs amount of x given in request

### How To Build

You can run following make command to build the application.
```shell
make build
```

If you want to build docker image you can use the make command
```shell
make dockerize
```

### How To Run
If you already built the application you can run with the following make command
```shell
make run
```
If you did not build, you can use the following command to build first, then run
```shell
make start
```

If you want to run ,and you were build the image with `make dockerize`,
you can use
```shell
make dockerun
```
to run the container. Otherwise, you should build your image with
```shell
docker build -t <image-name>:<image-tag> .
```
command ,and then run it with the following command:
```shell
docker run -p <host_port>:4200 <image-name>:<image-tag>
```
