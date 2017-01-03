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

# API

### Auth
* [POST /auth](docs/v1/auth_post.md) - 認証

### Users
* [POST /v1/users](docs/v1/users_post.md) - ユーザー情報を登録
* [GET /v1/users](docs/v1/users_get.md) - ユーザー情報の一覧を取得
* [GET /v1/users/:user_id](docs/v1/users_id_get.md) - ユーザー情報を取得


### Places
* [POST /v1/places](docs/v1/places_post.md) - 場所情報を登録
```
{
  "name": "place-name",
  "owners": ["user-id1", "user-id2"],
  "collaborators": ["user-id1", "user-id2"]
}
```
※owners, collaboratorsにはデフォルトで作成者が追加される.

* [GET /v1/places](docs/v1/places_get.md) - 場所情報の一覧を取得
* [GET /v1/places/:place_id](docs/v1/places_id_get.md) - 場所情報を取得
* [POST /v1/places/:place_id/status](docs/v1/places_id_status_post.md) - 場所情報に紐づく状態を登録
```
{
  "isOpen": true
}
```
* [GET /v1/places/:place_id/status](docs/v1/places_id_status_get.md) - 場所情報に紐づく状態を取得