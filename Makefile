

dev-account: install-refresh
	@refresh run -c ".dev/refresh-account.yml"

dev-gql: install-refresh
	@refresh run -c ".dev/refresh-graphql.yml"

install-refresh:
	@which refresh > /dev/null 2>&1  || GO111MODULE=off go get -v github.com/markbates/refresh