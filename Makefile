BINARY=bagoScan
OUTPUT_DIR=output

run:
	go run cmd/bagoScan/main.go

build:
	go build -o $(BINARY) cmd/bagoScan/main.go

clean:
	if [ -f $(BINARY) ]; then rm $(BINARY); fi
	rm -rf $(OUTPUT_DIR)

