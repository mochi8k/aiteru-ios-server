# GET  /v1/places

位置情報の一覧を取得

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
  "places": [
    {
      "id": "1",
      "name": "place1",
      "ownerIds": [
        "1",
        "2"
      ],
      "collaboratorIds": [
        "1",
        "2"
      ],
      "createdAt": "2017-01-03T00:00:00Z",
      "createdUserId": "2",
      "updatedAt": "2016-12-25T04:47:40Z",
      "updatedUserId": "1",
      "status": {
        "placeId": "1",
        "isOpen": true,
        "updatedAt": "2017-01-03T07:02:48Z",
        "updatedUserId": "5"
      }
    },
    {
      "id": "2",
      "name": "place2",
      "ownerIds": [
        "3"
      ],
      "collaboratorIds": [
        "3"
      ],
      "createdAt": "2017-01-03T00:00:00Z",
      "createdUserId": "3",
      "updatedAt": "2016-12-25T04:47:40Z",
      "updatedUserId": "1",
      "status": {
        "placeId": "2",
        "isOpen": true,
        "updatedAt": "2016-12-25T04:47:40Z",
        "updatedUserId": "5"
      }
    },
    {
      "id": "3",
      "name": "place3",
      "ownerIds": [
        "5"
      ],
      "collaboratorIds": [
        "5"
      ],
      "createdAt": "2017-01-02T02:46:03Z",
      "createdUserId": "5",
      "updatedAt": "2016-12-25T04:47:40Z",
      "updatedUserId": "1",
      "status": {
        "placeId": "3",
        "isOpen": false,
        "updatedAt": "",
        "updatedUserId": ""
      }
    }
  ]
}
```