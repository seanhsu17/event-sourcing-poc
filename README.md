# event-sourcing-poc

## Purpose
In event-driven system, we would like to know the publishers/subscribers sending completion rate to estimate how stable of our system.
Therefore, we designed a concept to approach this goal.
## Design Concept

![plot](./docs/images/event-sourcing.jpg)

### Documentation
[doc](https://17media.atlassian.net/wiki/spaces/ST/pages/712442097/Event+Sourcing+POC)

## Participants
- An API service with publishing event function
- one or many Subscribers that subscribe the same topic
- mongoDB with PublishRecord and ReceivedRecord collections
- gcp pub/sub
  - One topic
  - one or multiple subscriptions that subscribe the same topic
- A Scheduler

## Message Format
Our message should have 2 parts `payload` and `attributes`.

**payload** is any json format or struct in Go. It is defined by publisher.
```json
{
  "userId": "d4c6d768-7a5b-4fa4-86b7-f39a944d63d1",
  "name": "Zakk Wylde",
  "age": 18,
  "gender": "male"
}
```
**attributes** is a kind of metadata. It should be included traceID, eventID at least.
In our case, we have defined two keys for developers to use. We will follow goapi style if there are no other concerns.
```json
{
  "Cloud-Trace-Context": "<trace-id>/0;o=0", 
  "Options": "<json-string>"
}
```

## Launch
To run and stop:
```shell
make stop && make run
```

## Test
- start test
```shell
curl --location --request POST 'http://localhost:3000/api/v1/users'
```
- response for debugging
```
{
  "payload": {
    "userId": "ef0df7a5-4c94-49df-926d-ac22380f8f91",
    "name": "Zakk Wylde",
    "age": 18,
    "gender": "male"
  },
  "traceID": "7ac5a053-ce16-4f10-8cac-519a381fc5f9",
  "metadata": {
    "Cloud-Trace-Context": "7ac5a053-ce16-4f10-8cac-519a381fc5f9",
    "Options": "{}",
    "eventID": "8d68024a-fad3-489b-8a23-d49a3190c27a" // generate from Publisher, only for PoC
  }
}
```

- Publish record
```
{
  _id: ObjectId('62a849745b2fd0614ba18ab6'),
  traceID: '7ac5a053-ce16-4f10-8cac-519a381fc5f9',
  eventID: '8d68024a-fad3-489b-8a23-d49a3190c27a',
  topic: 'create-user',
  payload: '{"userId":"ef0df7a5-4c94-49df-926d-ac22380f8f91","name":"Zakk Wylde","age":18,"gender":"male"}',
  publishTime: 1655196020,
  createdTime: 1655196020
}
```

- Receive record
```
{
  _id: ObjectId('62a84975a2e9605831529f98'),
  topic: 'create-user',
  traceID: '7ac5a053-ce16-4f10-8cac-519a381fc5f9',
  eventID: '8d68024a-fad3-489b-8a23-d49a3190c27a',
  publishTime: 0,
  receiveTime: 1655196021,
  createdTime: 1655196021
}
```

- Scheduler calculate sending completion rate

The logs will like as below:
```log
scheduler_1      | rev record: create-user, traceID: 7ac5a053-ce16-4f10-8cac-519a381fc5f9, eventID: 8d68024a-fad3-489b-8a23-d49a3190c27a
scheduler_1      | total: 1 
scheduler_1      | missed: 0 
scheduler_1      | completion rate of last 10 sec. : 100.000000 
```