PostgreSQL Backend for straumur
===============================

[![Build Status](https://secure.travis-ci.org/straumur/postgres_backend.png)](http://travis-ci.org/straumur/postgres_backend)

Databackend for [straumur](http://www.github.com/straumur/straumur)

```go

package main

import (
	"github.com/straumur/straumur"
	"github.com/straumur/postgres_backend"
)

func main() {

	connString := "dbname=activityfeed host=localhost sslmode=disable"
	
	d, err := postgres_backend.NewPostgresDataSource(connString)
	if err != nil {
		panic(err)
	}	
	h := eventhub.NewHub("Application", d)
	h.Run()

}
```