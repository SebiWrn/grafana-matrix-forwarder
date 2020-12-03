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

generateChangelog:
	./tools/git-chglog_linux_amd64 --config tools/chglog/config.yml 0.1.0.. > CHANGELOG.md

build/snapshot:
	./tools/goreleaser_linux_amd64 --snapshot --rm-dist --skip-publish

build/release:
	./tools/goreleaser_linux_amd64 --rm-dist --skip-publish
