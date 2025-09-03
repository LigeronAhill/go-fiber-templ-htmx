# https://just.systems

default:
    just --list

gen:
    @templ generate

run: gen
    @go run cmd/main.go
