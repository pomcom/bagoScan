BINARY=bagoScan
OUTPUT_DIR=output

run:
	go run cmd/bagoScan/main.go

build:
	go build -o $(BINARY) cmd/bagoScan/main.go

clean:
	rm -r $(BINARY)
	rm -rf $(OUTPUT_DIR)

