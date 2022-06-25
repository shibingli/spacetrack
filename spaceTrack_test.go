package spacetrack

import (
	"encoding/json"
	"log"
	"testing"

	"github.com/shibingli/spacetrac/entity"
)

var (
	st *SpaceTrack
)

func init() {
	st = &SpaceTrack{
		Auth: Auth{
			Username: "shibingli@yeah.net",
			Password: "****************",
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
		Predicates("satname").
		Like("starlink").
		OrderBy("norad_cat_id", "desc").
		Limit(10).
		FormatJSON().
		Do(&satcat, stc)
	if err != nil {
		t.Fatalf("%+v", err.Error())
	}

	t.Logf("%+v", satcat)
}
