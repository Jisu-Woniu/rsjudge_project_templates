flags/prepare: fsharp.fsproj
	@dotnet restore
	@mkdir -p flags
	@touch flags/prepare

flags/build: flags/prepare Program.fs
	@dotnet build -c Release --no-restore
	@touch flags/build

run: flags/build
	@dotnet run -c Release --no-build

clean:
	@rm -rf bin/ obj/ flags

.PHONY: run clean
