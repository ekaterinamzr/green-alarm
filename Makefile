MOCKS_DESTINATION=internal/infrastructure/mocks/mocks.go
.PHONY: mocks

mocks: 
	@rm -rf $(MOCKS_DESTINATION)
	mockgen -source=internal/usecase/interfaces.go -destination=$(MOCKS_DESTINATION)
	

build:
	go build -v ./cmd/main.go

swag:
	swag init -g ./cmd/main.go

# sudo docker build .

# sudo docker compose up -d
# sudo docker compose down -v