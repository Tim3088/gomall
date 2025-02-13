.PHONY: build-svc
build-svc:
	docker build -f ./deploy/Dockerfile.svc -t ${svc}:${v} --build-arg SVC=${svc} .
.PHONY: build-client
build-client:
	docker build -f ./deploy/Dockerfile.client -t client:${v} .