
.PHONY: build clean test help default day-template



BIN_NAME=AoC2024

VERSION := $(shell grep "const Version " version/version.go | sed -E 's/.*"(.+)"$$/\1/')
GIT_COMMIT=$(shell git rev-parse HEAD)
GIT_DIRTY=$(shell test -n "`git status --porcelain`" && echo "+CHANGES" || true)
BUILD_DATE=$(shell date '+%Y-%m-%d-%H:%M:%S')

default: test

help:
	@echo 'Management commands for AoC2024:'
	@echo
	@echo 'Usage:'
	@echo '    make build           Compile the project.'
	@echo '    make get-deps        runs dep ensure, mostly used for ci.'
	@echo '    make clean           Clean the directory tree.'
	@echo '    make test            Run tests.'
	@echo '    make day-template    Create a new day command from the template.'
	@echo

build:
	@echo "building ${BIN_NAME} ${VERSION}"
	@echo "GOPATH=${GOPATH}"
	go build -ldflags "-X github.com/JEHoctor/AoC2024/version.GitCommit=${GIT_COMMIT}${GIT_DIRTY} -X github.com/JEHoctor/AoC2024/version.BuildDate=${BUILD_DATE}" -o bin/${BIN_NAME}

get-deps:
	dep ensure

clean:
	@test ! -e bin/${BIN_NAME} || rm bin/${BIN_NAME}

test:
	go test ./...

.inflect-venv:
	rm -rf .inflect-venv
	python3 -m venv .inflect-venv
	.inflect-venv/bin/python3 -m pip install inflect

day-template: .inflect-venv
	./scripts/day_template.sh templates/day.go.template cmd/ .inflect-venv/bin/python3
