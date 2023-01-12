BINARY=bagoScan
OUTPUT_DIR=output

run:
	go run main.go

build:
	go build -o $(BINARY) main.go

clean:
	if [ -f $(BINARY) ]; then rm $(BINARY); fi
	rm -rf $(OUTPUT_DIR)

