package entity

import (
	"github.com/SharkEzz/sgp4"
	"strings"
)

type GP struct {
	CCSDSOmmVers       string `json:"CCSDS_OMM_VERS" xml:"CCSDS_OMM_VERS"`
	Comment            string `json:"COMMENT" xml:"COMMENT"`
	CreationDate       string `json:"CREATION_DATE" xml:"CREATION_DATE"`
	Originator         string `json:"ORIGINATOR" xml:"ORIGINATOR"`
	ObjectName         string `json:"OBJECT_NAME" xml:"OBJECT_NAME"`
	ObjectID           string `json:"OBJECT_ID" xml:"OBJECT_ID"`
	CenterName         string `json:"CENTER_NAME" xml:"CENTER_NAME"`
	RefFrame           string `json:"REF_FRAME" xml:"REF_FRAME"`
	TimeSystem         string `json:"TIME_SYSTEM" xml:"TIME_SYSTEM"`
	MeanElementTheory  string `json:"MEAN_ELEMENT_THEORY" xml:"MEAN_ELEMENT_THEORY"`
	Epoch              string `json:"EPOCH" xml:"EPOCH"`
	MeanMotion         string `json:"MEAN_MOTION" xml:"MEAN_MOTION"`
	Eccentricity       string `json:"ECCENTRICITY" xml:"ECCENTRICITY"`
	Inclination        string `json:"INCLINATION" xml:"INCLINATION"`
	RaOfAscNode        string `json:"RA_OF_ASC_NODE" xml:"RA_OF_ASC_NODE"`
	ArgOfPericenter    string `json:"ARG_OF_PERICENTER" xml:"ARG_OF_PERICENTER"`
	MeanAnomaly        string `json:"MEAN_ANOMALY" xml:"MEAN_ANOMALY"`
	EphemerisType      string `json:"EPHEMERIS_TYPE" xml:"EPHEMERIS_TYPE"`
	ClassificationType string `json:"CLASSIFICATION_TYPE" xml:"CLASSIFICATION_TYPE"`
	NoradCatID         string `json:"NORAD_CAT_ID" xml:"NORAD_CAT_ID"`
	ElementSetNO       string `json:"ELEMENT_SET_NO" xml:"ELEMENT_SET_NO"`
	RevAtEpoch         string `json:"REV_AT_EPOCH" xml:"REV_AT_EPOCH"`
	Bstar              string `json:"BSTAR" xml:"BSTAR"`
	MeanMotionDot      string `json:"MEAN_MOTION_DOT" xml:"MEAN_MOTION_DOT"`
	MeanMotionDDot     string `json:"MEAN_MOTION_DDOT" xml:"MEAN_MOTION_DDOT"`
	SemimajorAxis      string `json:"SEMIMAJOR_AXIS" xml:"SEMIMAJOR_AXIS"`
	Period             string `json:"PERIOD" xml:"PERIOD"`
	Apoapsis           string `json:"APOAPSIS" xml:"APOAPSIS"`
	Periapsis          string `json:"PERIAPSIS" xml:"PERIAPSIS"`
	ObjectType         string `json:"OBJECT_TYPE" xml:"OBJECT_TYPE"`
	RcsSize            string `json:"RCS_SIZE" xml:"RCS_SIZE"`
	CountryCode        string `json:"COUNTRY_CODE" xml:"COUNTRY_CODE"`
	LaunchDate         string `json:"LAUNCH_DATE" xml:"LAUNCH_DATE"`
	Site               string `json:"SITE" xml:"SITE"`
	DecayDate          string `json:"DECAY_DATE" xml:"DECAY_DATE"`
	File               string `json:"FILE" xml:"FILE"`
	GpID               string `json:"GP_ID" xml:"GP_ID"`
	TLELine0           string `json:"TLE_LINE0" xml:"TLE_LINE0"`
	TLELine1           string `json:"TLE_LINE1" xml:"TLE_LINE1"`
	TLELine2           string `json:"TLE_LINE2" xml:"TLE_LINE2"`
}

func (g *GP) TLE() (*sgp4.TLE, error) {
	name := g.TLELine0
	names := strings.Fields(name)

	if len(names) > 1 {
		name = strings.Join(names[1:], " ")
	}

	tle, err := sgp4.NewTLE(name, g.TLELine1, g.TLELine2)
	if err != nil {
		return nil, err
	}

	return tle, nil
}

func (g *GP) SGP4() (*sgp4.SGP4, error) {
	tle, err := g.TLE()
	if err != nil {
		return nil, err
	}

	sgp, err := sgp4.NewSGP4(tle)
	if err != nil {
		return nil, err
	}

	return sgp, nil
}
