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

new-migration:
	@if [ -z "$(FILENAME)" ]; then \
		echo "Usage: make new-migration FILENAME=<your_filename>"; \
		exit 1; \
	fi; \
	TIMESTAMP=$$(date +"%s"); \
	NEW_FILENAME=$$(printf "%s_%s%s" $$TIMESTAMP $$FILENAME ".up.sql"); \
	touch "migrations/$$NEW_FILENAME"; \
	NEW_FILENAME=$$(printf "%s_%s%s" $$TIMESTAMP $$FILENAME ".down.sql"); \
    touch "migrations/$$NEW_FILENAME"; \
