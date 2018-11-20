# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
OUTPUT_DIR=../../bin
BINARY_NAME=stanford_algos_week1

all: test build
build: 
	$(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(OUTPUT_DIR)/$(BINARY_NAME)
run:
	$(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME) -v 
	$(OUTPUT_DIR)/$(BINARY_NAME)
#deps:
#    $(GOGET) github.com/markbates/goth
#    $(GOGET) github.com/markbates/pop
