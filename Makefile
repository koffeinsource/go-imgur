DOCKER_COMPOSE_TEST := docker-compose -f dev/test.yml

ifdef TEST_RUN
	TESTRUN := -run ${TEST_RUN}
endif

GOPACKAGES := $(shell go list ./... | egrep -v 'github.com/mix/mix-api$$')
TEST_MODULES ?= $(GOPACKAGES)


test: # run unit tests
	${DOCKER_COMPOSE_TEST} rm --force || true
	${DOCKER_COMPOSE_TEST} run test_goimgur
	${DOCKER_COMPOSE_TEST} down

test-direct: # [INTERNAL]
	go test -p 1 -v -race  -coverprofile=$(COVERAGE_FILE) $(TESTRUN)

lint: # Run go lint
	${DOCKER_COMPOSE_TEST} run test_goimgur bash -c "GOGC=50 make -e lint-direct"

lint-direct: # [INTERNAL]
	@golangci-lint run
