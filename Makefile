GITCOMMIT := $(shell git rev-parse HEAD)
GITDATE := $(shell git show -s --format='%ct')

LDFLAGSSTRING +=-X main.GitCommit=$(GITCOMMIT)
LDFLAGSSTRING +=-X main.GitDate=$(GITDATE)
LDFLAGS := -ldflags "$(LDFLAGSSTRING)"

TM_ABI_ARTIFACT := /Users/kit/code/go/com.web3/src/treasure-backend/abis/TreasureManager.sol/TreasureManager.json
ABIGEN := /Users/kit/code/go/com.web3/bin/abigen

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
		| $(ABIGEN) --pkg bindings \
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