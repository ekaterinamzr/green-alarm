build:
	go build -v ./cmd/main.go 

bench:
	go build -v ./benchmarking/main.go

clean:
	rm main
