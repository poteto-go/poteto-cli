# Poteto-Cli

- hot-reload run for golang
- support creating poteto app

## Create new app

We support cli tool. But if you doesn't like it, you can create poteto-app w/o cli of course.

```sh
go install github.com/poteto-go/poteto/cmd/poteto-cli@v0.26.5
```

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