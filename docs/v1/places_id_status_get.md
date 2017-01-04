# GET /v1/places/:place_id/status

場所情報に紐づく状態を取得


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
  "status": {
    "placeId": "1",
    "isOpen": true,
    "updatedAt": "2017-01-01T00:00:00Z",
    "updatedUserId": "1"
  }
}
```
