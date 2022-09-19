NAME=greenalarm

build:
	go build -v ./cmd/main.go

build-docker: 
	docker build --rm -t $(NAME) .

run-docker:
	docker run $(NAME)
