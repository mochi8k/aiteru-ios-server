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
    "AccessToken": "access-token",
    "User": {
      "id": "2",
      "name": "user-name",
      "createdAt": "2017-01-01 00:00:00",
      "createdUserId": "1",
      "updatedAt": "2017-01-02 00:00:00",
      "updatedUserId": "1"
    }
  }
}
```