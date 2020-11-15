export NOW=$(shell date +"%Y/%m/%d")
export PKGS=$(shell go list ./... | grep -vE '(vendor|cmd|entity|docs|cmd/test|model|pkg|generated)')
export TEST_OPTS=-cover -race

dev-easy:
	@echo "${NOW} == RUNNING ALL SERVICE INSIDE DOCKER"
	@docker-compose -f .dev/docker-compose.dev.yaml up

dev-account: install-air
	@echo "${NOW} == RUNNING ACCOUNT SERVICE"
	@air -c ".dev/account.air.toml"

dev-shop: install-air
	@echo "${NOW} == RUNNING SHOP SERVICE"
	@air -c ".dev/shop.air.toml"

dev-gql: install-air
	@echo "${NOW} == RUNNING GRAPHQL SERVICE"
	@air -c ".dev/graphql.air.toml"

install-air:
	@which air > /dev/null 2>&1 || GO111MODULE=off go get -u github.com/cosmtrek/air

install-refresh:
	@which refresh > /dev/null 2>&1  || GO111MODULE=off go get -v github.com/markbates/refresh

generate-mock:
	@echo "${NOW} == GENERATING MOCK FILES"
	@which mockgen > /dev/null 2>&1 || GO111MODULE=off go get -u github.com/golang/mock/mockgen
	@go generate .dev/generate.go

test:
	@echo "${NOW} == TESTING..."
	@go test ${TEST_OPTS} ${PKGS} -short | tee test.out
	@.dev/test.sh test.out 15