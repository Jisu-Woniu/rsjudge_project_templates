flags/prepare: go.mod
	@go mod tidy
	@mkdir -p $(@D)
	@touch $@

flags/build: flags/prepare main.go
	@go build
	@touch $@

run: flags/build judge
	@./judge

clean:
	@rm -rf flags judge

.PHONY: run clean
