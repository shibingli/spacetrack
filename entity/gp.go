package entity

import (
	"github.com/SharkEzz/sgp4"
	"strings"
)

type GP struct {
	CCSDSOmmVers       string `json:"CCSDS_OMM_VERS,omitempty" xml:"CCSDS_OMM_VERS,omitempty" bson:"ccsds_omm_vers,omitempty"`
	Comment            string `json:"COMMENT,omitempty" xml:"COMMENT,omitempty" bson:"comment,omitempty"`
	CreationDate       string `json:"CREATION_DATE,omitempty" xml:"CREATION_DATE,omitempty" bson:"creation_date,omitempty"`
	Originator         string `json:"ORIGINATOR,omitempty" xml:"ORIGINATOR,omitempty" bson:"originator,omitempty"`
	ObjectName         string `json:"OBJECT_NAME,omitempty" xml:"OBJECT_NAME,omitempty" bson:"object_name,omitempty"`
	ObjectID           string `json:"OBJECT_ID,omitempty" xml:"OBJECT_ID,omitempty" bson:"object_id,omitempty"`
	CenterName         string `json:"CENTER_NAME,omitempty" xml:"CENTER_NAME,omitempty" bson:"center_name,omitempty"`
	RefFrame           string `json:"REF_FRAME,omitempty" xml:"REF_FRAME,omitempty" bson:"ref_frame,omitempty"`
	TimeSystem         string `json:"TIME_SYSTEM,omitempty" xml:"TIME_SYSTEM,omitempty" bson:"time_system,omitempty"`
	MeanElementTheory  string `json:"MEAN_ELEMENT_THEORY,omitempty" xml:"MEAN_ELEMENT_THEORY,omitempty" bson:"mean_element_theory,omitempty"`
	Epoch              string `json:"EPOCH,omitempty" xml:"EPOCH,omitempty" bson:"epoch,omitempty"`
	MeanMotion         string `json:"MEAN_MOTION,omitempty" xml:"MEAN_MOTION,omitempty" bson:"mean_motion,omitempty"`
	Eccentricity       string `json:"ECCENTRICITY,omitempty" xml:"ECCENTRICITY,omitempty" bson:"eccentricity,omitempty"`
	Inclination        string `json:"INCLINATION,omitempty" xml:"INCLINATION,omitempty" bson:"inclination,omitempty"`
	RaOfAscNode        string `json:"RA_OF_ASC_NODE,omitempty" xml:"RA_OF_ASC_NODE,omitempty" bson:"ra_of_asc_node,omitempty"`
	ArgOfPericenter    string `json:"ARG_OF_PERICENTER,omitempty" xml:"ARG_OF_PERICENTER,omitempty" bson:"arg_of_pericenter,omitempty"`
	MeanAnomaly        string `json:"MEAN_ANOMALY,omitempty" xml:"MEAN_ANOMALY,omitempty" bson:"mean_anomaly,omitempty"`
	EphemerisType      string `json:"EPHEMERIS_TYPE,omitempty" xml:"EPHEMERIS_TYPE,omitempty" bson:"ephemeris_type,omitempty"`
	ClassificationType string `json:"CLASSIFICATION_TYPE,omitempty" xml:"CLASSIFICATION_TYPE,omitempty" bson:"classification_type,omitempty"`
	NoradCatID         string `json:"NORAD_CAT_ID,omitempty" xml:"NORAD_CAT_ID,omitempty" bson:"norad_cat_id,omitempty"`
	ElementSetNO       string `json:"ELEMENT_SET_NO,omitempty" xml:"ELEMENT_SET_NO,omitempty" bson:"element_set_no,omitempty"`
	RevAtEpoch         string `json:"REV_AT_EPOCH,omitempty" xml:"REV_AT_EPOCH,omitempty" bson:"rev_at_epoch,omitempty"`
	Bstar              string `json:"BSTAR,omitempty" xml:"BSTAR,omitempty" bson:"bstar,omitempty"`
	MeanMotionDot      string `json:"MEAN_MOTION_DOT,omitempty" xml:"MEAN_MOTION_DOT,omitempty" bson:"mean_motion_dot,omitempty"`
	MeanMotionDDot     string `json:"MEAN_MOTION_DDOT,omitempty" xml:"MEAN_MOTION_DDOT,omitempty" bson:"mean_motion_d_dot,omitempty"`
	SemimajorAxis      string `json:"SEMIMAJOR_AXIS,omitempty" xml:"SEMIMAJOR_AXIS,omitempty" bson:"semimajor_axis,omitempty"`
	Period             string `json:"PERIOD,omitempty" xml:"PERIOD,omitempty" bson:"period,omitempty"`
	Apoapsis           string `json:"APOAPSIS,omitempty" xml:"APOAPSIS,omitempty" bson:"apoapsis,omitempty"`
	Periapsis          string `json:"PERIAPSIS,omitempty" xml:"PERIAPSIS,omitempty" bson:"periapsis,omitempty"`
	ObjectType         string `json:"OBJECT_TYPE,omitempty" xml:"OBJECT_TYPE,omitempty" bson:"object_type,omitempty"`
	RcsSize            string `json:"RCS_SIZE,omitempty" xml:"RCS_SIZE,omitempty" bson:"rcs_size,omitempty"`
	CountryCode        string `json:"COUNTRY_CODE,omitempty" xml:"COUNTRY_CODE,omitempty" bson:"country_code,omitempty"`
	LaunchDate         string `json:"LAUNCH_DATE,omitempty" xml:"LAUNCH_DATE,omitempty" bson:"launch_date,omitempty"`
	Site               string `json:"SITE,omitempty" xml:"SITE,omitempty" bson:"site,omitempty"`
	DecayDate          string `json:"DECAY_DATE,omitempty" xml:"DECAY_DATE,omitempty" bson:"decay_date,omitempty"`
	File               string `json:"FILE,omitempty" xml:"FILE,omitempty" bson:"file,omitempty"`
	GpID               string `json:"GP_ID,omitempty" xml:"GP_ID,omitempty" bson:"gp_id,omitempty"`
	TLELine0           string `json:"TLE_LINE0,omitempty" xml:"TLE_LINE0,omitempty" bson:"tle_line_0,omitempty"`
	TLELine1           string `json:"TLE_LINE1,omitempty" xml:"TLE_LINE1,omitempty" bson:"tle_line_1,omitempty"`
	TLELine2           string `json:"TLE_LINE2,omitempty" xml:"TLE_LINE2,omitempty" bson:"tle_line_2,omitempty"`
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
