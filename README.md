# FlightTracker

## Overview

This repository contains the implementation of a flight tracker service. 
Given a set of flights, the service calculate for each request, the source
and the destination airports. The service also perform basic input validation. 
The service is implemented using the gin web framework.

# API documentation.

The service consist of one end point used to calculate the source and destination airports for 
a flight.

## End Points

### Calculate 

```code 
GET /v1/calculate
```

The endpoint is used to calculate the start and destination airport for a flight.

#### Attributes

```code
flights  array of flights
```

List of flights in the form of [Source,Destination]. 

For example:

```code 
GET http://localhost:8080/v1/calculate
Accept: application/json

{ "flights": [["IND", "EWR"], ["SFO", "ATL"], ["GSO", "IND"], ["ATL", "GSO"]] }

```


#### Response

```code
route  pair of start and source airport.
```

The service answer a route, which is a pair of the source and destination airport.

For example:

```code

{
  "route": [
    "SFO",
    "EWR"
  ]
}

```

#### Return Codes

The service perform basic input validation. The following errors are returned in response


| Code | Description | Summary                                                                   |
|------|-------------|---------------------------------------------------------------------------|
| 200  | OK          | Everything worked as expected                                             |
| 400  | Bad Request | The request was unacceptable, often due to missing a required parameter. |
 | 402 | Request Failed | The parameters were valid but the request failed. |
 | 500 | Internal Error | Something went wrong |


## Internal Design

The design follow the canonical Json based Microservice design. At its core the design is
based on the Gin web framework for request routing, as well as Json parsing.

The following are the important go packages


| Package           | Description                                                   |
|-------------------|---------------------------------------------------------------|
| cmd/flighttracker | Contain the main program. Start the microservice on port 8080 |
| api/v1            | The service API and data structures |
| server            | The actual service implementation |


## Running the service

In order to run the service type the following at the command line

```code
flighttracker serve
```

## API Design

The API is composed of a single method 

```code 
	Calculate(context.Context, *CalculateRequest) (*CalculateReply, error)
```

The method is bound to the path

```code 
 /v1/calculate
```

## Service Implementation

The service code reside in the server package. The service is composed of
two components:


* Server struct that implement the actual server
* The implementation of the calculate method, which contain the algorithm.

## The Server

The server is based on the GIN web framework. The framework handles all
the HTTP method routing. The GIN framework also have the concept of a
middleware, which are go function handlers that help during request
processing. 

The server implementation contain the three main methods:

| Package  | Description                                                  |
|----------|--------------------------------------------------------------|
| New      | Initialize the Server and setup the middleware and the routes |
| Serve    | Start the service and listen to HTTP requests                |
| Shutdown | Gracfully Shutdown the service                               |




## Calculate Algorithm 

To calculate() algorithm is based on graph theory. The output of the calculation
should be the first and last airport of the trip. Hence, the following insights
were used as part of the algorithm:

* The first node, is the graph node which does not have any incoming edges 
* The last node, is the graph node which does not have any outgoing edges.

The algorithm first build the graph data structure, and then uses 
the above insights to find the first and last node in O(N).


# Important insights and Ideas

* Time it took to finish the assignment : 5 hours
* At the beginning, I was thinking about implementing a topological sort, 
  however since we need only the first and last airport, Using the graph data structure is enough.
* The runtime is O(N), and the storage is also O(N)
* I added versioning to the api. I.e. instead of /calculate the API is /v1/calculate. This would
  assure API versioning. 
* I used GIN (including GIN middleware)










