dev:
	docker-compose up

build:
	docker build -t go-ci-cd-prod . --target production -f Dockerfile.production

start:
	docker run -p 80:8000 --name go-ci-cd-prod go-ci-cd-prod 