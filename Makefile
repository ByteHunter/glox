BINAY_NAME=glox

build:
	go build -o glox .

run:
	@./${BINAY_NAME} main.lox

repl:
	@./${BINAY_NAME}

test:
	go test -v