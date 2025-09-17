# https://just.systems

# list available commands
default:
    just --list

# install tools
install-tools:
    go install github.com/a-h/templ/cmd/templ@latest
    go install github.com/air-verse/air@latest

# generate template code
gen:
    @templ generate

# initialize configuration for air
init-air:
    air init

# run application
run: gen
    @go run cmd/main.go

# air application
air:
    @air

# watch app live reload
watch:
    templ generate --watch --proxy="http://localhost:3000" --cmd="go run cmd/main.go"

# run docker compose file
up:
    @docker compose up -d
# stop docker container
stop:
    @docker compose down

# build docker image
docker:
    docker build -t go-app .
