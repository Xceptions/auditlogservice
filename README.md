## Audit Log Service

**System for receiving events of any structure, and querying them**
![alt text](./imageforreadme/design.png)


### Overview

Based on the instruction, there are two main components

1. Service for accepting events
2. Service for querying events

Then there are features:

1. Authentication
2. Exposable via a http endpoint

### My approach

If you look at the root of the application, you will be met with 6 files:

1. auditlogeventservice
2. auditlogcustomerservice
3. curlrequests.sh
4. docker-compose
5. imageforreadme
6. readme
7. .env

The assignment is in `auditlogcustomerservice`, can be deployed using the `docker-compose` file, and tested using the `curlrequests.sh`.

The reason there is an `auditlogeventservice` file
*From my understanding of the service, the events will be received from different systems which can send events at any time. This seems to me like a http endpoint will not be the best way to approach this, rather a TCP connection to these systems will be better. This is why I added `auditlogeventservice`. It contains a tcp listener. A TCP listener is faster than a http one, and ensures correct order of events, this will help us to treat the event reception as a stream and handle appropriately*

Despite my opinion, I have followed the instruction in the email. It is contained in `auditlogcustomerservice`, and will now explain my approach.

Looking at the `docker-compose.yml` file, you will see various services
1. auditlogcustomerservice - where the application is contained
2. mongodb - choice database for storing events. The reason for this choice is the flexibility of each entry. It can help create flexible events
3. mongoexpress - a GUI for viewing mongodb
4. postgres - choice database for storing user detail. The reason for this choice is the ability to make aggregations it brings to our data. For example, we may choose to store more user detail like location data, then we need to ask how many users are in that location?
5. pgadmin - a GUI for viewing postgres
6. redis - choice cache system for caching get queries

In `auditlogcustomerservice`,
the `main.go` file

Contains four endpoints
- createuser: for creating the user account that can send and query the database, returns a jwt token for authentication
- loginuser: returns a jwt token for authentication
- submitevent: for submitting the event to the database
- getevent: for querying the event

*All files with _test are test files for the appropriate methods*

the submitevent method is optimized using a buffered channel for temporarily storing event data until a threshold set in the .env file, then an `insertMany` operation is done on this number of events. I do this to ensure that I do not keep making calls to the database for each event that comes in. This method is also rate-limited using a custom rate-limiter

The getevent method is optimized by storing queries in a cache after it has been made. This is to ensure that I do not keep hitting the database for the same query. An expiration time of 1 hour is set for expiring the cached data. This method is also rate-limited using a custom rate-limiter

The overall application has versioning for the APIs


To deploy the application:
```
docker-compose up -d
```

To test the endpoint:
```
./curlrequests.sh
```