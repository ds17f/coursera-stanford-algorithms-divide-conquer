# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
OUTPUT_DIR=../../bin
BINARY_NAME=stanford_algos_week4

## Params for external tests
# update the next 2 vars to change the tests
COURSE=course1
ASSIGNMENT=assignment4MinCut

TEST_DIR=../../stanford-algs
TESTER=$(TEST_DIR)/tester/shell/tester.sh
TEST_CASES=$(TEST_DIR)/testCases/$(COURSE)/$(ASSIGNMENT)/input*

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
testall:
	$(GOBUILD) -o $(OUTPUT_DIR)/$(BINARY_NAME) -v
	$(TESTER) -c $(OUTPUT_DIR)/$(BINARY_NAME) $(TEST_CASES)
