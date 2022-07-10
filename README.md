# FairX Golang Monorepo
Golang Monorepo for the FairX Project, projects and code supporting TBD's Web5 Protocol with some WASM extensions :)

* [TBD Web5 Specification](https://developer.tbd.website/projects/web5/)


# Applications

## Decentralized Web Node

**To install, build, and run:**

```shell

git clone https://github.com/fairxio/go
cd go
go mod download
cd go/applications/fairx
go build

./fairx dwn

```

**Use Prebuilt Docker Image:**

```shell
docker pull fairxio/dwn:latest
docker run -it -v "$(PWD):/etc/fairx" -p "8080:8080" fairxio/dwn
```

**Build Docker Image:**

```shell
git clone https://github.com/fairxio/go
cd go
docker build -f deployment/docker/Dockerfile -t fairxio/dwn .
```