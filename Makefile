.PHONY: init
init:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/golang/mock/mockgen@latest

.PHONY: mock
mock:
	mockgen -source=internal/service/user.go -destination test/mocks/service/user.go
	mockgen -source=internal/repository/user.go -destination test/mocks/repository/user.go

.PHONY: test
test:
	go test -coverpkg=./internal/handler,./internal/service,./internal/repository -coverprofile=./coverage.out ./test/server/...
	go tool cover -html=./coverage.out -o coverage.html

.PHONY: build
build:
	go build -ldflags="-s -w" -o ./bin/server ./cmd/server/...

.PHONY: docker
docker:
	docker build --no-cache -f deploy/build/Dockerfile --build-arg APP_RELATIVE_PATH=./cmd/server/... -t nunu-test/demo-api:v1 .

