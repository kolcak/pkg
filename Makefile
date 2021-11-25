#!/usr/bin/env bash

PROJECT_ENV=./assets/.env.project

include ${PROJECT_ENV}

# Setup project name
PROJECT=$(or ${CI_PROJECT_NAME}, $(notdir ${DOCKERC_PROJECT_PATH}))

# These are the generic values
RELEASE=`date -u +%s`
VERSION=`date +"v%Y.%m%d.%H%M%S"`
REVISION=$(or ${CI_COMMIT_SHA}, ${VERSION})

CMD_DOCKER_COMPOSE=env $$(cat ${PROJECT_ENV} | sed 's/\#.*//g' | xargs) docker-compose

# If the first argument is "run"...
ifeq (run,$(firstword $(MAKECMDGOALS)))
  # use the rest as arguments for "run"
  RUN_ARGS := $(wordlist 2,$(words $(MAKECMDGOALS)),$(MAKECMDGOALS))
  # ...and turn them into do-nothing targets
  $(eval $(RUN_ARGS):;@:)
endif

# Build GO devel container
# ===================================================================================================================================================================
dc.build:
	${CMD_DOCKER_COMPOSE} build go

# Start container bash (external terminal)
# ===================================================================================================================================================================
dc.go:
	${CMD_DOCKER_COMPOSE} run --rm go make do.bash
do.bash:
	bash
dc.go.ssh:
	${CMD_DOCKER_COMPOSE} run --rm go make do.bash.ssh
do.bash.ssh:
	cp -R /tmp/ssh/* /root/.ssh
	chmod 700 /root/.ssh
	chmod 600 /root/.ssh/* -R
	eval `ssh-agent -s` && ssh-add /root/.ssh/${SSH_PRIVATE_KEY_FILENAME} && bash

# GO Prepare dependencies
# ===================================================================================================================================================================
dc.prepare:
	${CMD_DOCKER_COMPOSE} run --rm go make prepare
prepare:
	go mod tidy && go mod vendor || true

# GO Run section
# ===================================================================================================================================================================
dc.run:
	${CMD_DOCKER_COMPOSE} run --rm go make run
run:
	go run ./src ${cmd}

# GO test security
# ===================================================================================================================================================================
dc.security:
	${CMD_DOCKER_COMPOSE} run --rm go make security
security:
	gosec -exclude=G307 ./src/...

# GO test
# ===================================================================================================================================================================
dc.test:
	${CMD_DOCKER_COMPOSE} run --rm go make test
test:
	mkdir -p cover
	go test -v -timeout 30s -coverprofile=cover/cover.out -cover ./src/...
	go tool cover -html=cover/cover.out -o cover/cover.html

# Clean
# ===================================================================================================================================================================
dc.clean:
	${CMD_DOCKER_COMPOSE} rm -fv

%:
	@:
