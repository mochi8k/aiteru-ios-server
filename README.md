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

# Progress

### /auth
POST: /auth
```
{
  "name": "user-name"
}
```
### /users
POST: /v1/users
```
{
  "name": "user-name"
}
```
GET: /v1/users  
GET: /v1/users/{user-id}  
### /places
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