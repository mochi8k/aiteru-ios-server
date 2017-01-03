# POST /v1/places

場所情報を登録

## Headers

| Key           | Value            |
|---------------|------------------|
| Content-Type  | application/json |
| Authorization | {access_token}   |

## Body

```
{
  "name": "place-name",
  "owners": ["user-id1", "user-id2"],
  "collaborators": ["user-id1"]
}
```
※owners, collaboratorsにはデフォルトで作成者が追加される.

## Response

```
{
  "place": {
    "id": "1",
    "name": "place-name",
    "ownerIds": [
      "user-id1",
      "user-id2"
    ],
    "collaboratorIds": [
      "user-id1"
    ],
    "createdAt": "2017-01-01 00:00:00",
    "createdUserId": "1",
    "updatedAt": "",
    "updatedUserId": "",
    "status": {
      "placeId": "1", // ここ怪しい
      "isOpen": false,
      "updatedAt": "",
      "updatedUserId": ""
    }
  }
}
```
