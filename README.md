# FairX Golang Monorepo
Golang Monorepo for the FairX Project, projects and code supporting TBD's Web5 Protocol with some WASM extensions :)

* [TBD Web5 Specification](https://developer.tbd.website/projects/web5/)


# Applications

## Building the `fairx` executable

The `fairx` executable from this monorepo is following the [commander pattern](https://github.com/spf13/viper) for golang apps. 

To build:

```shell

git clone https://github.com/fairxio/go
cd go
go mod download
cd go/applications/fairx
go build

```

## Decentralized Web Node

The Decentralized Web Node attempts is an opinionated implementation of [DIF Decentralized Web Nodes](https://identity.foundation/decentralized-web-node/spec/), an
interesting approach to decentralized application development.

**To run (after building the `fairx` executable):**

```shell

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
make build-docker-dwn
```


## FairX Authentication Service

Generally speaking, any authentication service can be plugged in here, however one wants to issue a valid JWT, using a shared key between services requiring authorization.
This service automatically issues a JWT for any valid DID authentication flow where the DID presented was not on a blacklist, or is on a whitelist.  

**To run (after building the `fairx` executable):**

```shell

./fairx auth

```

**Use Prebuilt Docker Image:**

```shell
docker pull fairxio/auth:latest
docker run -it -v "$(PWD):/etc/fairx" -p "7080:8080" fairxio/auth:latest
```

**Build Docker Image:**

```shell
git clone https://github.com/fairxio/go
cd go
make build-docker-auth
```

## FairX DID Service

The FairX DID Service is a loose implementation of a [DID Registration and Resolution](https://github.com/decentralized-identity/did-registration/blob/main/spec/spec.md) spec 
led by the identity.foundation, but using REST over TLS explicitly as the base underlying protocol.

**To run (after building the `fairx` executable):**

```shell

./fairx did

```

**Use Prebuilt Docker Image:**

```shell
docker pull fairxio/auth:latest
docker run -it -v "$(PWD):/etc/fairx" -p "6080:8080" fairxio/did:latest
```

**Build Docker Image:**

```shell
git clone https://github.com/fairxio/go
cd go
make build-docker-did
```