# Getting Started

## MySQL
```
$ brew install mysql
$ mysql.server start
$ mysql -uroot

create database aiteru;
```

# Progress

POST: /auth
```
{
  "name": "user-name"
}
```

POST: /v1/users
```
{
  "name": "user-name"
}
```
GET: /v1/users
GET: /v1/users/{user-id}
GET: /v1/places  
GET: /v1/places/{place-id}