MOCKS_DESTINATION=server/internal/infrastructure/mocks/mocks.go
.PHONY: mocks

mocks: 
	@rm -rf $(MOCKS_DESTINATION)
	mockgen -source=server/internal/usecase/interfaces.go -destination=$(MOCKS_DESTINATION)
	

# build:
# 	go build -v ./server/cmd/main.go

swag:
	swag init -g ./server/cmd/main.go -o ./server/docs

docker-build:
	sudo docker build -t server ./server 
	sudo docker build -t client ./client

docker-run:
	sudo docker compose -f ./deploy/docker-compose.yml up

docker-run-background:
	sudo docker compose -f ./deploy/docker-compose.yml up -d

docker-stop:
	sudo docker compose -f ./deploy/docker-compose.yml down -v

docker-clean:
	sudo docker rmi server client
