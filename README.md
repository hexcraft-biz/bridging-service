# Bridging Service

```sh
# Prepare Mysql & Redis service
# Set up environment variable refers to .example.env
# Execute main.go
$ go run main.go
# Or Docker
$ docker-compose -f dev.yml up --build -d
```

## Endpoint
### healthcheck
GET /healthcheck/v1/ping
- Params
  - None
- Resonse
  - 200

### bridging
#### endpoints
POST /bridging/v1/endpoints
- Params
  - Headers
    - Content-Type : application/json
  - Body
    - path
      - Required : True
- Response
  - 201
  	```json
	{
	  "id": "a66a9a1b-f4d9-44b5-ae48-46e4dd3077c0",
	  "path": "/aaa/bbc",
	  "createdAt": "2022-08-30 05:22:42",
	  "updatedAt": "2022-08-30 05:22:42"
	}
	```
  - 400
  - 409
  - 500
  
GET /bridging/v1/endpoints
- Params
  - QueryString
    - limit
      - Required : False
      - Default : 20
    - offset
      - Required : False
      - Default : 0
- Response
  - 200
  	```json
	[
	  {
	    "id": "a66a9a1b-f4d9-44b5-ae48-46e4dd3077c0",
	    "path": "/aaa/bbc",
	    "createdAt": "2022-08-30 05:22:42",
	    "updatedAt": "2022-08-30 05:22:42"
	  }
	]
	```
  - 400
  - 500
  
GET /bridging/v1/endpoints/:id
- Params
  - Path
    - id
      - Required : True
- Response
  - 200
  	```json
	{
	  "id": "a66a9a1b-f4d9-44b5-ae48-46e4dd3077c0",
	  "path": "/aaa/bbc",
	  "createdAt": "2022-08-30 05:22:42",
	  "updatedAt": "2022-08-30 05:22:42"
	}
	```
  - 400
  - 404
  - 500

DELETE /bridging/v1/endpoints/:id
- Params
  - Path
    - id
      - Required : True
- Response
  - 204
  - 400
  - 404
  - 500

#### topics
POST /bridging/v1/topics
- Params
  - Headers
    - Content-Type : application/json
  - Body
    - name
      - Required : True
- Response
  - 201
  	```json
	{
	  "id": "a66a9a1b-f4d9-44b5-ae48-46e4dd3077c0",
	  "name": "pubsub-demo1",
	  "createdAt": "2022-08-30 05:22:42",
	  "updatedAt": "2022-08-30 05:22:42"
	}
	```
  - 400
  - 409
  - 500

GET /bridging/v1/topics
- Params
  - QueryString
    - limit
      - Required : False
      - Default : 20
    - offset
      - Required : False
      - Default : 0
- Response
  - 200
  	```json
	[
	  {
	    "id": "a66a9a1b-f4d9-44b5-ae48-46e4dd3077c0",
	    "name": "pubsub-demo1",
	    "createdAt": "2022-08-30 05:22:42",
	    "updatedAt": "2022-08-30 05:22:42"
	  }
	]
	```
  - 400
  - 500

GET /bridging/v1/topics/:id
- Params
  - Path
    - id
      - Required : True
- Response
  - 200
  	```json
	{
	  "id": "a66a9a1b-f4d9-44b5-ae48-46e4dd3077c0",
	  "name": "pubsub-demo1",
	  "createdAt": "2022-08-30 05:22:42",
	  "updatedAt": "2022-08-30 05:22:42"
	}
	```
  - 400
  - 404
  - 500

DELETE /bridging/v1/topics/:id
- Params
  - Path
    - id
      - Required : True
- Response
  - 204
  - 400
  - 404
  - 500

#### endpoint-topic-rels
POST /bridging/v1/endpoint-topic-rels
- Params
  - Headers
    - Content-Type : application/json
  - Body
    - endpointId
      - Required : True
    - topicId
      - Required : True
- Response
  - 201
  	```json
	{
	  "id": "0302f81c-ed0f-4ebe-b216-1e3248cdb8a1",
	  "endpointId": "a66a9a1b-f4d9-44b5-ae48-46e4dd3077c0",
	  "endpointPath": "/aaa/bbc",
	  "topicId": "dbf25409-27af-4866-bfaa-7920d724bb04",
	  "topicName": "pubsub-demo10",
	  "createdAt": "2022-08-30T05:42:57Z",
	  "updatedAt": "2022-08-30T05:42:57Z"
	}
	```
  - 400
  - 409
  - 500

GET /bridging/v1/endpoint-topic-rels
- Params
  - QueryString
    - limit
      - Required : False
      - Default : 20
    - offset
      - Required : False
      - Default : 0
    - endpointId
      - Required : False
    - endpointPath
      - Required : False
    - topicId
      - Required : False
    - topicName
      - Required : False
- Response
  - 200
  	```json
	[
	  {
	    "id": "0302f81c-ed0f-4ebe-b216-1e3248cdb8a1",
	    "endpointId": "a66a9a1b-f4d9-44b5-ae48-46e4dd3077c0",
	    "endpointPath": "/aaa/bbc",
	    "topicId": "dbf25409-27af-4866-bfaa-7920d724bb04",
	    "topicName": "pubsub-demo10",
	    "createdAt": "2022-08-30T05:42:57Z",
	    "updatedAt": "2022-08-30T05:42:57Z"
	  }
	]
	```
  - 400
  - 500

GET /bridging/v1/endpoint-topic-rels/:id
- Params
  - Path
    - id
      - Required : True
- Response
  - 200
  	```json
	{
	  "id": "0302f81c-ed0f-4ebe-b216-1e3248cdb8a1",
	  "endpointId": "a66a9a1b-f4d9-44b5-ae48-46e4dd3077c0",
	  "endpointPath": "/aaa/bbc",
	  "topicId": "dbf25409-27af-4866-bfaa-7920d724bb04",
	  "topicName": "pubsub-demo10",
	  "createdAt": "2022-08-30T05:42:57Z",
	  "updatedAt": "2022-08-30T05:42:57Z"
	}
	```
  - 400
  - 404
  - 500

DELETE /bridging/v1/endpoint-topic-rels/:id
- Params
  - Path
    - id
      - Required : True
- Response
  - 204
  - 400
  - 404
  - 500
