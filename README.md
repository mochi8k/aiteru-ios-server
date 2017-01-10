# Getting Started
## MySQL
```
$ brew install mysql
$ mysql.server start
$ mysql -uroot

create database aiteru;
```

## Go
```
$ go run server.go
```

## Config
```
export APP_PORT=8000
export APP_MYSQL_USER=user-name
export APP_MYSQL_PASSWORD=password
export APP_MYSQL_DB=db-name
```

# API

### Auth
* [POST /auth](docs/v1/auth_post.md) - 認証

### Users
* [POST /v1/users](docs/v1/users_post.md) - ユーザー情報を登録
* [GET /v1/users](docs/v1/users_get.md) - ユーザー情報の一覧を取得
* [GET /v1/users/:user_id](docs/v1/users_id_get.md) - ユーザー情報を取得
* [PUT /v1/users/:user_id](docs/v1/users_id_put.md) - ユーザー情報を更新


### Places
* [POST /v1/places](docs/v1/places_post.md) - 場所情報を登録
* [GET /v1/places](docs/v1/places_get.md) - 場所情報の一覧を取得
* [GET /v1/places/:place_id](docs/v1/places_id_get.md) - 場所情報を取得
* [POST /v1/places/:place_id/status](docs/v1/places_id_status_post.md) - 場所情報に紐づく状態を登録
* [GET /v1/places/:place_id/status](docs/v1/places_id_status_get.md) - 場所情報に紐づく状態を取得