APP := echo-realworld
set-goproxy:
	go env -w GOPROXY=https://proxy.golang.org

############################################################
# Build & Run
############################################################
build: set-goproxy
	go build -v -race .

serve: build
	./echo-realworld serve

build-static-vendor-linux: vendor
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -mod vendor -v -o $(APP) -installsuffix cgo

vendor: set-goproxy
	go mod vendor -v