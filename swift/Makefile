flags/prepare: Package.swift
	@swift package resolve
	@mkdir -p $(@D)
	@touch $@

flags/build: flags/prepare Sources/*.swift
	@swift build -c release
	@mkdir -p $(@D)
	@touch $@

run: flags/build .build/release/judge
	@.build/release/judge

clean:
	@rm -rf .build flags

.PHONY: run clean
