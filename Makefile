BINAY_NAME=glox
TESTS_DIR=./tests
TEST_SET=. ./utils ./reporting ./token ./scanner ./expression ./astprinter ./parser ./interpreter ./cmd/ast

build:
	go build -o glox .

run:
	@./${BINAY_NAME} main.lox

repl:
	@./${BINAY_NAME}

ast:
	go run ./cmd/ast/main.go expression

test:
	go test ${TEST_SET}

test-verbose:
	go test -v ${TEST_SET}

coverage:
	go test -coverprofile=${TESTS_DIR}/coverage.out ${TEST_SET}

serve-coverage:
	go tool cover -html=${TESTS_DIR}/coverage.out -o=${TESTS_DIR}/coverage.html
	@echo "Check at: http://localhost:3000/coverage.html"
	php -t ${TESTS_DIR} -S localhost:3000
