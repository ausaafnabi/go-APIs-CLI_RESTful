# METRO API
---
### OVERVIEW
This RESTful API is designed to interact with the user to manage the metro CRUD transactions :

**This API consists of three services:**
|---|
|Driver|
|Station|
|Trains|

#### Endpoints
`www.metroAPIurl:8000/v1/trains`

`www.metroAPIurl:8000/v1/station`
 
`www.metroAPIurl:8000/v1/driver`

#### API USAGE:

##### POST :

`curl -X POST http://metroAPIurl:8000/v1/trains -H 'cache-control: no-cache' -H 'content-type: application/json' -d '{"driverName":"DriverName","operatingStatus":true}'
`
##### GET :

`curl -X GET "http://metroAPIurl:8000/v1/trains/1"`

##### DELETE :

`curl -X DELETE "http://metroAPIurl:8000/v1/trains/1"
`
