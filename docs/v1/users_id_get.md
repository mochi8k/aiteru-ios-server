# GET  /v1/users/:user_id

ユーザー情報を取得

## Headers

| Key           | Value            |
|---------------|------------------|
| Content-Type  | application/json |
| Authorization | {access_token}   |

## Body

なし

## Response

```
{
  "user": {
    "id": "2",
    "name": "user1",
    "createdAt": "2017-01-01T00:00:00Z",
    "createdUserId": "1",
    "updatedAt": "2017-01-02T00:00:00Z,
    "updatedUserId": "1"
  }
}
```