run:
	go run ./cmd/server

test:
	go test ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out

build:
	go build -o bin/server ./cmd/server

docker-build:
	docker build -t go-cicd-demo:local .

trivy-fs:
	trivy fs --severity HIGH,CRITICAL .

trivy-image:
	trivy image --severity HIGH,CRITICAL go-cicd-demo:local
