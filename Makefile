build-linux:
	GOOS=linux GOARCH=amd64 go build -o ./build/optiod-linux-amd64 ./cmd/optiod/main.go
	GOOS=linux GOARCH=arm64 go build -o ./build/optiod-linux-arm64 ./cmd/optiod/main.go

do-checksum-linux:
	cd build && sha256sum \
		optiod-linux-amd64 optiod-linux-arm64 \
		> optiod-checksum-linux

build-linux-with-checksum: build-linux do-checksum-linux

build-linux-amd64:
	GOOS=linux GOARCH=amd64 go build -o ./build/optiod-linux-amd64 ./cmd/optiod/main.go

build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o ./build/optiod-linux-arm64 ./cmd/optiod/main.go

build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o ./build/optiod-darwin-amd64 ./cmd/optiod/main.go
	GOOS=darwin GOARCH=arm64 go build -o ./build/optiod-darwin-arm64 ./cmd/optiod/main.go

build-all: build-linux build-darwin

do-checksum-darwin:
	cd build && sha256sum \
		optiod-darwin-amd64 optiod-darwin-arm64 \
		> optiod-checksum-darwin

build-darwin-with-checksum: build-darwin do-checksum-darwin

build-with-checksum: build-linux-with-checksum build-darwin-with-checksum

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 go build -o ./build/optiod-darwin-amd64 ./cmd/optiod/main.go

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o ./build/optiod-darwin-arm64 ./cmd/optiod/main.go
