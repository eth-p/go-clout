example:
	for example in example/*; do (cd "$$example" && make); done

test:
	go test go.eth-p.dev/clout

all: test
.PHONY: example test all
