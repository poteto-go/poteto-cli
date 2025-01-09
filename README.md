# Poteto-Cli

We support cli tool. But if you doesn't like it, you can create poteto-app w/o cli of course.

```sh
go install github.com/poteto-go/poteto-cli/cmd/poteto-cli@latest
```

- hot-reload run for golang
- support creating poteto app

## Build From Docker

You can also use from docker image

https://hub.docker.com/repository/docker/poteto17/poteto-go/general

```sh
docker pull poteto17/poteto-go
docker -it --rm poteto17/poteto-go:1.23 bash
```

## Create new app

Create file.

```sh
poteto-cli new
```

fast mode.

```sh
poteto-cli new --fast
```

## run app with hot-reload

- create `poteto.yaml`

```yaml
version: "0.27"
build_script_path: "main.go"
debug_mode: true
```

- command

```sh
poteto-cli run
```
