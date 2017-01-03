# POST  /v1/users

ユーザー情報を登録

## Headers

| Key           | Value            |
|---------------|------------------|
| Content-Type  | application/json |
| Authorization | {access_token}   |

## Body

```
{
  "name": "user-name"
}
```

## Response

```
{
  "user": {
    "id": "2",
    "name": "user-name",
    "createdAt": "2017-01-01 00:00:00",
    "createdUserId": "1",
    "updatedAt": "",
    "updatedUserId": ""
  }
}
```