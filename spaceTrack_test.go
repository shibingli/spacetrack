package spacetrack

import (
	"encoding/json"
	"log"
	"testing"
	"time"

	"github.com/shibingli/spacetrack/entity"
)

var (
	st *SpaceTrack
)

func init() {
	st = &SpaceTrack{
		Auth: Auth{
			Username: "shibingli@yeah.net",
			Password: "******************",
		},
	}
}

type SpaceTrackCodec struct {
}

func (s *SpaceTrackCodec) Unmarshal(data []byte, v any) error {
	return json.Unmarshal(data, v)
}

func (s *SpaceTrackCodec) Marshal(v any) ([]byte, error) {
	return json.Marshal(v)
}

func (s *SpaceTrackCodec) Write(p []byte) (n int, err error) {
	log.Printf("SpaceTrackCodecRet: %s\n", string(p))
	return 0, nil
}

func TestBasicSpaceDataController_QuerySatcat(t *testing.T) {
	_, err := st.Login()
	if err != nil {
		t.Fatalf("%+v", err.Error())
	}

	defer func() {
		if err := st.Logout(); err != nil {
			t.Fatalf("%+v", err.Error())
		}
	}()

	t.Logf("%+v", st.Auth.Cookies)

	var satcat []entity.Satcat

	stc := &SpaceTrackCodec{}

	err = st.BasicSpaceDataController().
		QuerySatcat().
		Predicate("satname").
		Like("starlink").
		OrderBy("norad_cat_id", "desc").
		Limit(10).
		FormatJSON().
		Build(&satcat, stc)
	if err != nil {
		t.Fatalf("%+v", err.Error())
	}

	b, _ := json.Marshal(satcat)

	t.Logf("%+v", string(b))
}

func TestBasicSpaceDataController_QueryGP(t *testing.T) {
	_, err := st.Login()
	if err != nil {
		t.Fatalf("%+v", err.Error())
	}

	defer func() {
		if err := st.Logout(); err != nil {
			t.Fatalf("%+v", err.Error())
		}
	}()

	var gp []entity.GP

	stc := &SpaceTrackCodec{}

	err = st.BasicSpaceDataController().
		QueryGP().
		Predicate("NORAD_CAT_ID").
		OR("52882").
		Limit(1).
		FormatJSON().
		Build(&gp, stc)
	if err != nil {
		t.Fatalf("%+v", err.Error())
	}

	gpData := gp[0]

	b, _ := json.Marshal(gpData)

	t.Logf("Gp Data: %+v", string(b))

	tle, err := gpData.TLE()
	if err != nil {
		t.Fatalf("%+v", err.Error())
	}

	t.Logf("TLE Data: %+v", tle)

	sgp, err := gpData.SGP4()
	if err != nil {
		t.Fatalf("%+v", err.Error())
	}

	lat, lng, alt, err := sgp.Position(time.Now().Local().UTC())
	if err != nil {
		t.Fatalf("%+v", err.Error())
	}

	t.Logf("lat: %f, lng: %f, alt: %f\n", lat, lng, alt)
}
