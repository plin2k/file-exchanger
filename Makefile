BUILD_ARCH=arm64
BUILD_OS=darwin


OS_SETTINGS=env GOOS=$(BUILD_OS) GOARCH=$(BUILD_ARCH)
# Build Downloader ...
.PHONY: d-build
d-build:
	$(OS_SETTINGS) go build -o bin/downloader/app cmd/downloader/*.go
# Build Uploader ...
.PHONY: u-build
u-build:
	$(OS_SETTINGS) go build -o bin/uploader/app cmd/uploader/*.go



# Run Downloader ...
.PHONY: d-run
d-run:
	go run cmd/downloader/*.go
# Run Uploader ...
.PHONY: u-run
u-run:
	go run cmd/uploader/*.go
