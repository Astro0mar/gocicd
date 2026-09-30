# Go CI/CD Demo

A practical DevSecOps CI/CD project using:

- Go
- GitHub Actions
- Docker
- SonarQube
- Trivy
- Go unit tests
- Code coverage

## Pipeline

```text
Git Push / Pull Request
        |
        v
    Go Test
        |
        +--> gofmt
        +--> go vet
        +--> coverage
        +--> build
        |
        +--> SonarQube
        |
        +--> Trivy filesystem
        |
        v
   Docker Build
        |
        v
   Trivy image scan
        |
        v
 Docker image artifact
```

## Run locally

```bash
go test ./...
go test -coverprofile=coverage.out ./...
go run ./cmd/server
```

Then:

```bash
curl http://localhost:8080/
curl http://localhost:8080/health
```

## Build Docker image

```bash
docker build -t go-cicd-demo:local .
docker run --rm -p 8080:8080 go-cicd-demo:local
```

## Run SonarQube locally

```bash
docker compose up -d
```

Open:

```text
http://localhost:9000
```

Create a SonarQube token and configure these GitHub repository secrets:

```text
SONAR_HOST_URL
SONAR_TOKEN
```

For a self-hosted SonarQube server, `SONAR_HOST_URL` should point to a URL reachable by the GitHub Actions runner.

## Run Trivy locally

```bash
trivy fs --severity HIGH,CRITICAL .
```

After building the image:

```bash
trivy image --severity HIGH,CRITICAL go-cicd-demo:local
```

## GitHub setup

1. Create a GitHub repository.
2. Replace the module path in `go.mod` and imports if desired.
3. Push this project.
4. Add `SONAR_HOST_URL` and `SONAR_TOKEN` as repository secrets.
5. Push to `main` or open a pull request.

## Suggested next stage

After this pipeline is working, add:

```text
GitHub Actions
      |
      v
Docker Build
      |
      v
Trivy
      |
      v
Docker Hub / GHCR / AWS ECR
      |
      v
Kubernetes
      |
      v
Deployment
```

Then add Kubernetes manifests, Helm, GitOps/Argo CD, and environment promotion.
