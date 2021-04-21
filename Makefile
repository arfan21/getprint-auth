build-dev:
	docker build -f dev.Dockerfile -t getprint-service-auth-dev .

build-prod:
	docker build -f prod.Dockerfile -t getprint-service-auth-prod .