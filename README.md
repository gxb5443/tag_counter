#Namely Tag Counter
HTTP Microservice to accept POST requests with a tag string and counts the number of tag occurances.  Stores data in Postgres database.

##System Requirements
* Golang
* Postgres database
* Endpoints:
  * POST /tags
  * GET /tags/[tag]

##Usage
###Configuration
The tag counter comes with a configuration file.  This must be configured before the server is run.  The configuration file is in the project root directory.  It's called the 'app.config' file.  It contains the configuration information for connecting to Postgres.
```
user=admin
password=admin
dbname=tag_counting
sslmode=disable
```

###Dependencies
The system requires some external libraries to make things easier and cleaner.  The libraries include:
* [Gin-Gonic](http://gin-gonic.github.io/gin/) - This handles all the http requests and routing
* [PQ](https://github.com/lib/pq) - Golang Postgres driver
* [Gorp](https://github.com/coopernurse/gorp) - ORM wrapper library for easier DB access

If needed, use go get:
```
go get github.com/gin-gonic/gin
go get github.com/lib/pq
go get github.com/coopernurse/gorp
```

###Running
Golang runs like C.  It is compiled and runs off an executable.
```
./tag_counter
```

###Building
While the repo should come with the executable, if you need to build it yourself it's as simple as:
```
go run tag_counter.go
```

##API Usage
 Method |Endpoint | Description |
 --- | --- | --- |
GET | /tags/[tag] | Fetches count statistics for specified tag.  If tag is not in database, it returns a 0 in the count attribute.```{"tag"}```
POST | /tags | Accepts a JSON object ```{"tag":"desired_tag"}``` and increments desired tag by one.
Both endpoints return a tag JSON object.  If the requested tag is not in the database, the count attribute is 0.
```
{
  "tag": "desired_tag",
  "count": 1
}
```


##Author
* **Gian Biondi** <gianbiondijr@gmail.com>
