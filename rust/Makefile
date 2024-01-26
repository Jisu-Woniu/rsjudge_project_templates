flags/prepare: Cargo.toml
	@cargo fetch
	@mkdir -p $(@D)
	@touch $@

flags/build: flags/prepare src/main.rs
	@cargo build --release
	@touch $@

run: flags/build target/release/judge
	@target/release/judge

clean:
	@rm -rf target flags

.PHONY: run clean
