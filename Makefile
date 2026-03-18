BINAY_NAME=glox
TESTS_DIR=./tests

build:
	go build -o glox .

run:
	@./${BINAY_NAME} main.lox

repl:
	@./${BINAY_NAME}

test:
	go test -v . ./utils ./reporting ./token ./scanner

coverage:
	go test -coverprofile=${TESTS_DIR}/coverage.out . ./utils ./reporting ./token ./scanner

serve-coverage:
	go tool cover -html=${TESTS_DIR}/coverage.out -o=${TESTS_DIR}/coverage.html
	@echo "Check at: http://localhost:3000/coverage.html"
	php -t ${TESTS_DIR} -S localhost:3000
