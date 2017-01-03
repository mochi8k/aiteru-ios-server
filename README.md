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
POST: /auth
```
{
  "name": "user-name"
}
```
### Users
* [POST /v1/users](docs/v1/users_post.md) - ユーザー情報を登録
```
{
  "name": "user-name"
}
```
* [GET /v1/users](docs/v1/users_get.md) - ユーザー情報の一覧を取得
* [GET /v1/users/:user_id](docs/v1/users_id_get.md) - ユーザー情報を取得



### Places
POST: /v1/places
```
{
  "name": "place-name",
  "owners": ["user-id1", "user-id2"],
  "collaborators": ["user-id1", "user-id2"]
}
```
※owners, collaboratorsにはデフォルトで作成者が追加される.

GET: /v1/places  
GET: /v1/places/{place-id}  
POST: /v1/places/{place-id}/status
```
{
  "isOpen": true
}
```
GET: /v1/places/{place-id}/status  