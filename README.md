# Toll calculator

POC of a simple toll calculator made to learn about microservices in go.


## Tech stack
 - golang
 - docker (and docker compose for basic orchestration)
 - kafka
 - http
 - protobuf
 - grpc
 - prometheus
 - grafana

## Architecture

### OBU send
Random data generator used to test the application. It will send data via websocket continuously in JSON format.

### OBU recv
Receives OBU data via websocket and produces kafka messages.

### Distance calculator
Consumes kafka messages containing OBU data and sends to the aggregator; it can use http or grpc transport layer to send this data.  
In the POC, we aggregate data in memory to keep things simple, but the interface allows for more persistent implementations.

### Aggregator
Aggregates the distance received by the distance calculator and exposes the invoice, given the OBU Id.  
The aggregation is as simple as it gets, no maps or any analytics on the path taken by the truck, that's business logic outside the scope of
this repository.
The invoice is as simple as it gets as well, i just need to prove a point, there's no real business logic involved.

### Gateway
User-facing backend interface to interact qith our API.

### Monitoring
TBD

## How to use
Run `make containers` to spawn the cluster of containers  
TBD

