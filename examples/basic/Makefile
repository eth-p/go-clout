MODULE = $(shell cat "../../go.mod" | grep "^module " | sed 's/^module \(.*\)/\1/')
EXAMPLE = $(notdir $(shell pwd))
EXAMPLE_PACKAGE = $(MODULE)/examples/$(EXAMPLE)

# Target: build
# Builds the example.
.PHONY: build
build:
	@echo "$(EXAMPLE)"
	go build -o "/dev/null" "$(EXAMPLE_PACKAGE)"

# Target: run
# Runs the example.
.PHONY: run
run:
	go run "$(EXAMPLE_PACKAGE)"
