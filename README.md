# go-restful

restful API in go, a learning project

## up and running

- copy and update .envrc file

```sh
cp .envrc-sample .envrc
```

- source .envrc or use `direnv`

### start api

```sh
docker compose up api
```

### create user

```sh
curl -sSL \
  -H "Content-Type: application/json" \
  -X POST \
  -d '{"username":"test", "password":"test"}' \
  http://localhost:${SERVER_PORT}/auth/register
```

### login

```sh
jwt=$(curl -sSL \
  -H "Content-Type: application/json" \
  -X POST \
  -d '{"username":"test", "password":"test"}' \
  http://localhost:${SERVER_PORT}/auth/login \
  | jq -r .jwt)
```

### create entry

```sh
curl -sSL \
  -d '{"content":"Sample content"}' \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $jwt" \
  -X POST \
  http://localhost:${SERVER_PORT}/api/entry
```

### get entires

```sh
curl -sSL \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer $jwt" \
  -X GET \
  http://localhost:${SERVER_PORT}/api/entry \
  | jq
```
