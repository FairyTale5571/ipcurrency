# IPCurrency

## Description:

Test task for Privat24 Golang Developer position.

## Tests:

```bash 
make test
```

## Configuration:

Use `config.yml` in `config` folder.

# How to start:

## Install:

```bash
make init-go
```

## Build:

```bash
make build
```

## Run:

```bash
make run
```


## API:

`/ip-info` - POST request with body:

```json
{
  "ip": ["92.62.121.6","2a01:4f8:120:210d::"]
}
```

### Swagger:

```bash
http://localhost:3000/docs/index.html
```