flags/prepare: csharp.csproj
	@dotnet restore
	@mkdir -p flags
	@touch flags/prepare

flags/build: flags/prepare Program.cs
	@dotnet build -c Release --no-restore
	@touch flags/build

run: flags/build
	@dotnet run -c Release --no-build

clean:
	@rm -rf bin/ obj/ flags

.PHONY: run clean
