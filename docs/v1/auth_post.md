# POST /auth

認証

## Headers

| Key           | Value            |
|---------------|------------------|
| Content-Type  | application/json |

## Body

```
{
  "name": "user-name"
}
```

## Response

```
{
  "session": {
    "AccessToken": "access-token", // キーを"accessToken"に変更予定
    "user": {
      "id": "2",
      "name": "user-name",
      "createdAt": "2017-01-01T00:00:00Z",
      "createdUserId": "1",
      "updatedAt": "2017-01-02T00:00:00Z",
      "updatedUserId": "1"
    }
  }
}
```