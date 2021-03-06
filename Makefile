dev:
	docker-compose up

build:
	docker buildx build --platform=linux/amd64 -t go-ci-cd-prod . --target production -f Dockerfile.production --no-cache

start:
	docker run -p 8000:8000 --name go-ci-cd-prod go-ci-cd-prod 