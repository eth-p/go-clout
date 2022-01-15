MODULE = $(shell cat "go.mod" | grep "^module " | sed 's/^module \(.*\)/\1/')

# Target: test
# Run the package tests.
.PHONY: test
test: dependencies
	go test -race $(MODULE)/...

# Target: bench
# Benchmark the package tests.
.PHONY: bench
bench: dependencies
	go test -bench=. $(MODULE)/...

# Target: dependencies
# Download all dependencies.
dependencies:
	@go mod download

# Target: run-example/%
# Run an example.
.PHONY: example/%
run-example/%: dependencies examples/%/Makefile
	@cd "examples/$*" && $(MAKE) -s run

# Target: example/%
# Build an example.
.PHONY: example/%
example/%: dependencies examples/%/Makefile
	@cd "examples/$*" && $(MAKE) -s

# Target: example
# Run all examples.
.PHONY: example
example: $(addprefix example/,$(notdir $(wildcard examples/*)))
