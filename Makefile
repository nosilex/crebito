
dev: clean
	@echo "⚡️Starting docker dev..."
	@docker-compose -f docker/docker-compose-local.yml up -d

prod: build
	@echo "⚡️Starting docker prod..."
	@docker-compose -f docker/docker-compose.yml up -d

build: clean purge
	@echo "🚀 Building docker image..."
	@docker build -t nosilex/crebito .  2> /dev/null ||:

clean:
	@echo "🧹 Cleaning docker containers..."
	@docker stop rinha-api1 2> /dev/null ||:
	@docker rm rinha-api1 2> /dev/null ||:
	@docker stop rinha-api2 2> /dev/null ||:
	@docker rm rinha-api2 2> /dev/null ||:
	@docker stop rinha-alb 2> /dev/null ||:
	@docker rm rinha-alb 2> /dev/null ||:
	@docker stop rinha-db 2> /dev/null ||:
	@docker rm rinha-db 2> /dev/null ||:

purge:
	@echo "🧹 Purging docker image..."
	@docker rmi nosilex/crebito 2> /dev/null ||:

push:
	@echo "⚡️ Push docker image..."
	@docker push nosilex/crebito