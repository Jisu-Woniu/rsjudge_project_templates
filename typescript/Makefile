flags/prepare: package.json
	@pnpm install
	@mkdir -p $(@D)
	@touch $@

flags/build: flags/prepare index.ts tsconfig.json
	@pnpm run build >/dev/null
	@touch $@

run: flags/build build/index.js
	@node --enable-source-maps build/index.js

clean:
	@rm -rf build node_modules flags/ pnpm-lock.yaml

.PHONY: run clean
