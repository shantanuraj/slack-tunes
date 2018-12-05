all:
	go run release/main.go

nix:
	go run release/main.go --arch amd64 --os linux

win:
	go run release/main.go --arch amd64 --os windows

mac:
	go run release/main.go --arch amd64 --os darwin

install:
	go run release/main.go --base-dir $(GOPATH)

install-nix:
	go run release/main.go --base-dir $(GOPATH) --arch amd64 --os linux

install-win:
	go run release/main.go --base-dir $(GOPATH) --arch amd64 --os windows

install-mac:
	go run release/main.go --base-dir $(GOPATH) --arch amd64 --os darwin
