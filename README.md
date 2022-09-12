# Bridging Service

```sh
# Prepare Mysql & Redis service
# Set up environment variable refers to .example.env
# Execute main.go
$ go run main.go
# Or Docker
$ docker-compose -f dev.yml up --build -d
```

## Basic Endpoint
### healthcheck
GET /healthcheck/v1/ping
- Params
  - None
- Resonse
  - 200

### Bridging
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
  - 401
  - 403
  - 404
  - 500

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
  - 401
  - 403
  - 404
  - 500

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
	  "path": "/aaa/bbc",
	  "topicId": "dbf25409-27af-4866-bfaa-7920d724bb04",
	  "name": "pubsub-demo10",
	  "createdAt": "2022-08-30T05:42:57Z",
	  "updatedAt": "2022-08-30T05:42:57Z"
	}
	```
  - 400
  - 401
  - 403
  - 404
  - 500
