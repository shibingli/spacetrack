# SpaceTrack SDK for golang

## Description
Go SDK with RESTFul API package based on space-track.org

Reference: [https://www.space-track.org/documentation#/api](https://www.space-track.org/documentation#/api)

## Installation

```go
go get -v github.com/shibingli/spacetrack
```

## Example

```go

package main

import (
	"log"

	"github.com/shibingli/spacetrack"
	"github.com/shibingli/spacetrack/entity"
)

func main() {

	st := &spacetrack.SpaceTrack{
		Auth: spacetrack.Auth{
			Username: "shibingli@yeah.net",
			Password: "******************",
		},
	}

	_, err := st.Login()
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	defer func() {
		if err := st.Logout(); err != nil {
			log.Fatalf("%+v\n", err)
		}
	}()

	log.Printf("%+v\n", st.Auth.Cookies)

	var satcat []entity.Satcat

	err = st.BasicSpaceDataController().
		QuerySatcat().
		Predicates("satname").
		Like("starlink").
		OrderBy("norad_cat_id", "desc").
		Limit(10).
		FormatJSON().
		Build(&satcat)
	if err != nil {
		log.Fatalf("%+v", err)
	}

	log.Printf("%+v", satcat)
}

```