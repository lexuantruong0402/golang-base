# BE SOSC App



## Getting started

To make it easy for you to get started with GitLab, here's a list of recommended next steps.

Already a pro? Just edit this README.md and make it your own. Want to make it easy? [Use the template at the bottom](#editing-this-readme)!

## Setup Environment

- [ ] [Install Golang](https://go.dev/doc/install)
- [ ] [Install Docker and Docker Compose]()
- [ ] [Install Swagger for Golang](https://goswagger.io/install.html)

## API service
## Setup local development
1. copy ``.env.example`` to ``.env``

### Install tools
- [Docker](https://docs.docker.com/engine/install/)
  After install docker and docker-compose. We should add your user in docker group.
  And then, you can use docker command without 'sudo'
    ```bash
  sudo usermod -aG docker ${USER}
  sudo services docker restart
    ```
- [TablePlus](https://tableplus.com/)
- [Golang](https://golang.org/doc/install)
- [Migrate](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)

    ```bash
    curl -s https://packagecloud.io/install/repositories/golang-migrate/migrate/script.deb.sh | sudo bash
    apt-get install -y migrate
    ```

- [Sqlc](https://github.com/kyleconroy/sqlc#installation)

    ```bash
    sudo snap install sqlc
    ```

- [Ngrok](https://github.com/kyleconroy/sqlc#installation)

    ```bash
    sudo snap install ngrok
    ngrok authtoken 1k7wxqj7ayAu0Hxd1ZN9qtmabRl_4fFwqoYj22YeJXJckUdGe
    ```
  Run Ngrok in background with:
    ```bash
    ngrok http --bind-tls=true --log=stdout 8080 > /dev/null &
    ```

- [Protoc + Proto-gen]()

    ```bash
    apt install -y protobuf-compiler
    sudo snap install protobuf --classic
    ```

- [Swagger CMD]()

    ```bash
    go get -u github.com/swaggo/swag/cmd/swag
    export PATH=$PATH:/usr/local/go/bin
    export PATH=$PATH:$HOME/go/bin
    ```

### Setup infrastructure
- Start postgres container:
    ```bash
    make postgres
    ```
- Create database:
    ```bash
    make createdb
    ```



- Run db migration:
    ```bash
    make migrateup
    ```
### How to run
- Setup env to build without cgo, before running all.bash or make.bash:
    ```bash
    go env -w CGO_ENABLED=0
    ```
- Download package dependence:
    ```bash
    go mod tidy
    ```