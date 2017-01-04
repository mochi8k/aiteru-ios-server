# GET  /v1/users

ユーザー情報の一覧を取得

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
  "users": [
    {
      "id": "1",
      "name": "admin",
      "createdAt": "2017-01-03T00:00:00Z",
      "createdUserId": "",
      "updatedAt": "2017-01-04T00:00:00Z",
      "updatedUserId": "1"
    },
    {
      "id": "2",
      "name": "user1",
      "createdAt": "2017-01-03T00:00:00Z",
      "createdUserId": "1",
      "updatedAt": "2017-01-04T00:00:00Z",
      "updatedUserId": "1"
    },
    {
      "id": "3",
      "name": "user2",
      "createdAt": "2017-01-03T00:00:00Z",
      "createdUserId": "1",
      "updatedAt": "2017-01-04T00:00:00Z",
      "updatedUserId": "1"
    }
  ]
}
```