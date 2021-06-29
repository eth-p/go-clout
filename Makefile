example:
	for example in example/*; do (cd "$$example" && make); done

test:
	go test github.com/eth-p/clout

all: test
.PHONY: example test all
