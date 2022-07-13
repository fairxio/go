

build-docker-base:
	docker build -f deployment/docker/Dockerfile-base -t fairxio/base-golang-service .

push-docker-base:
	docker push fairxio/base-golang-service

build-docker-did: build-docker-base
	docker build -f deployment/docker/Dockerfile-did -t fairxio/did .

push-docker-did:
	docker push fairxio/did

build-docker-auth: build-docker-base
	docker build -f deployment/docker/Dockerfile-auth -t fairxio/auth .

push-docker-auth:
	docker push fairxio/auth

build-docker-dwn: build-docker-base
	docker build -f deployment/docker/Dockerfile-dwn -t fairxio/dwn .

push-docker-dwn:
	docker push fairxio/dwn

build-docker-all: build-docker-base build-docker-did build-docker-auth build-docker-dwn

push-docker-all: push-docker-base push-docker-did push-docker-auth push-docker-dwn


generate-mocks:
	mockgen -source=comms/channel.go -destination=mock/channel_mock.go -package=mock
