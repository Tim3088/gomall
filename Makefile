.PHONY: build-svc
build-svc:
	docker build -f ./deploy/Dockerfile.svc -t ${svc}:${v} --build-arg SVC=${svc} .
.PHONY: build-client
build-client:
	docker build -f ./deploy/Dockerfile.client -t client:${v} .

##@ Open Browser

.PHONY: open.consul
open-consul: ## open `consul ui` in the default browser
	@open "http://localhost:8500/ui/"

.PHONY: open.jaeger
open-jaeger: ## open `jaeger ui` in the default browser
	@open "http://localhost:16686/search"

.PHONY: open.grafana
open-grafana: ## open `grafana ui` in the default browser
	@open "http://localhost:3000"