# make hello
hello:
	echo "XXX"


# make build
build:
	go build -o bin/main main.go


# make run
run:
	go run main.go


# make compile
compile:
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go


# make all
all: compile hello
