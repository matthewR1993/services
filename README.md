# Auth api with golang

Dev:
```sh
$ make rundev
```

Build:
```sh
$ make build
```

Run:
```sh
$ ./auth-service
```

Run as docker container:
```sh
$ sudo docker build -t auth-service .
$ sudo docker run --name auth-service -p 8080:8080 -it auth-service:latest
```
