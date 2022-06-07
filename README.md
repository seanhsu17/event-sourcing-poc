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
- A Subscriber
- mongoDB with PublishRecord and ReceivedRecord collections
- gcp pub/sub
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
In our case, we have defined two keys for developers to use.
```json lines
// example
// attributes (metadata)
// Follow goapi style
{
  "Cloud-Trace-Context": "<<trace-id>>/0;o=0", 
  "Options": "{\"source\":\"UserService\",\"eventType\":\"createUser\",\"key\":\"<<event_id>>\",\"timestamp\":1654568594}"
}
```


## Launch
To run and stop:
```shell
make stop && make run
```
