# PUT  /v1/users/:user_id

ユーザー情報を更新

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
    "createdAt": "2017-01-01T00:00:00Z",
    "createdUserId": "1",
    "updatedAt": "2017-01-02T00:00:00Z",
    "updatedUserId": "1"
  }
}
```