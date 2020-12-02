install-deps:
	go mod download

# Standard go test
test:
	go test ./... -v -race

# Make sure no unnecessary dependencies are present
go-mod-tidy:
	go mod tidy -v
	git diff-index --quiet HEAD

format:
	go fmt $(go list ./... | grep -v /vendor/)
	go vet $(go list ./... | grep -v /vendor/)

build/snapshot:
	./tools/goreleaser --snapshot --rm-dist --skip-publish

build/release:
	./tools/goreleaser --rm-dist --skip-publish
