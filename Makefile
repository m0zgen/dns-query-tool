# Binary name, arch, ver
BINARY_NAME := dns-query-tool

PLATFORMS := linux/amd64 darwin/amd64 # windows/amd64
EXTENSIONS := '' '' '.exe'

VERSION := $(shell cat version.txt)

# Build
build: update_version
	@mkdir -p builds
	@for idx in $$(seq 0 $$(($(words $(PLATFORMS)) - 1))); do \
		platform=$$(echo $(PLATFORMS) | cut -d' ' -f$$(($$idx + 1))); \
		os=$$(echo $$platform | cut -d'/' -f1); \
		arch=$$(echo $$platform | cut -d'/' -f2); \
		ext=$$(echo $(EXTENSIONS) | cut -d' ' -f$$(($$idx + 1))); \
		if [ $$os = "linux" ]; then \
			build_dir="tools/builds"; \
		elif [ $$os = "darwin" ]; then \
			build_dir="tools/builds/darwin"; \
		elif [ $$os = "windows" ]; then \
			build_dir="tools/builds/windows"; \
		else \
			echo "Unknown platform: $$os"; \
			exit 1; \
		fi; \
		mkdir -p $$build_dir; \
		echo "Build for $$os/$$arch..."; \
		if [ $$os = "linux" ]; then \
			CGO_ENABLED=0 GOOS=$$os GOARCH=$$arch go build -ldflags "-s -w -X main.version=$(VERSION)" -o $$build_dir/$(BINARY_NAME)$$ext main.go; \
		else \
			GOOS=$$os GOARCH=$$arch go build -ldflags "-X main.version=$(VERSION)" -o $$build_dir/$(BINARY_NAME)$$ext main.go; \
		fi; \
	done

# Update version
update_version:
	../go-upapp-version-tool/update-version-tool*
	echo $(shell cat version.txt) > VERSION

# Clean
clean:
	rm -rf builds

# Usage
help:
	@echo "Use 'make build' to build the application for all platforms."
	@echo "Use 'make clean' to remove generated files."
