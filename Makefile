GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

treasure-backed:
	env GO111MODULE=on go build -o treasure-backend -v $(LDFLAGS) ./cmd

clean:
	rm treasure-backend

test:
	go test -v ./...

lint:
	golangci-lint run ./...

bindings:
	$(eval temp := $(shell mktemp))

	cat $(TM_ABI_ARTIFACT) \
    	| jq -r .bytecode.object > $(temp)

	cat $(TM_ABI_ARTIFACT) \
		| jq .abi \
		| abigen --pkg bindings \
		--abi - \
		--out bindings/treasure_manager.go \
		--type TreasureManager \
		--bin $(temp)

		rm $(temp)

.PHONY: \
	treasure-backed \
	clean \
	test \
	bindings \
	lint