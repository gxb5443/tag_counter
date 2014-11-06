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
1. [Gin-Gonic]<http://gin-gonic.github.io/gin/> - This handles all the http requests and routing
2. [PQ]<"https://github.com/lib/pq"> - Golang Postgres driver
3. [Gorp]<https://github.com/coopernurse/gorp> - ORM wrapper library for easier DB access

###Running the server
Golang runs like C.  It is compiled

##Author
* **Gian Biondi** <gianbiondijr@gmail.com>
