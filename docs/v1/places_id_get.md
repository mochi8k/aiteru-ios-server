# GET  /v1/places/:place_id

位置情報を取得

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
    "createdAt": "2017-01-01T00:00:00Z",
    "createdUserId": "1",
    "updatedAt": "2017-01-02T00:00:00Z",
    "updatedUserId": "2",
    "status": {
      "placeId": "1", // ここ怪しい
      "isOpen": false,
      "updatedAt": "",
      "updatedUserId": ""
    }
  }
}
```
