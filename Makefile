compile:
	go build -o ./build/ ./...

dev:
	go run ./...

exec:
	./build/app

run: compile exec

	