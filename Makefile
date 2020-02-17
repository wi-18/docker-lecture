.PHONY: compile

all: compile

compile:
	@scripts/compile.sh

docker-compile:
	@scripts/docker-compile.sh
