bin=./pokedex

build:
	go build

run: build
	@$(bin)

clean:
	go mod tidy