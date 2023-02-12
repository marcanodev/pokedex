bin=./pokedex

build:
	go build

run: build
	@$(bin)
