install-deps:
	cd src/ && go mod download

# Standard go test
test:
	cd src/ && go test ./... -v -race

# Make sure no unnecessary dependencies are present
go-mod-tidy:
	cd src/ && go mod tidy -v
	git diff-index HEAD

format:
	cd src/ && go fmt $(go list ./... | grep -v /vendor/)
	cd src/ && go vet $(go list ./... | grep -v /vendor/)

generateChangelog:
	./tools/git-chglog_linux_amd64 --config tools/chglog/config.yml 0.1.0.. > CHANGELOG.md

build/snapshot:
	./tools/goreleaser_linux_amd64 --snapshot --rm-dist --skip-publish

build/release:
	./tools/goreleaser_linux_amd64 --rm-dist --skip-publish

build/docker:
	docker build -t registry.gitlab.com/hectorjsmith/grafana-matrix-forwarder:latest .

build/docker-tag:
	docker build -t registry.gitlab.com/hectorjsmith/grafana-matrix-forwarder:$(shell git describe --tags) .
