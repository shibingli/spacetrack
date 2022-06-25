# SpaceTrack SDK for golang

## Description
Go SDK with RESTFul API package based on space-track.org

Reference: [https://www.space-track.org/documentation#/api](https://www.space-track.org/documentation#/api)

## Example

```go

st := &SpaceTrack{
    Auth: Auth{
        Username: "shibingli@yeah.net",
        Password: "***********",
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
    Do(&satcat)
if err != nil {
    log.Fatalf("%+v", err)
}

log.Printf("%+v", satcat)

```