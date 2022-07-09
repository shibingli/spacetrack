package entity

type Satcat struct {
	Intldes      string `json:"INTLDES,omitempty" xml:"INTLDES,omitempty" bson:"intldes,omitempty"`
	NoradCatID   string `json:"NORAD_CAT_ID,omitempty" xml:"NORAD_CAT_ID,omitempty" bson:"norad_cat_id,omitempty"`
	ObjectType   string `json:"OBJECT_TYPE,omitempty" xml:"OBJECT_TYPE,omitempty" bson:"object_type,omitempty"`
	SatName      string `json:"SATNAME,omitempty" xml:"SATNAME,omitempty" bson:"sat_name,omitempty"`
	Country      string `json:"COUNTRY,omitempty" xml:"COUNTRY,omitempty" bson:"country,omitempty"`
	Site         string `json:"SITE,omitempty" xml:"SITE,omitempty" bson:"site,omitempty"`
	Decay        string `json:"DECAY,omitempty" xml:"DECAY,omitempty" bson:"decay,omitempty"`
	Period       string `json:"PERIOD,omitempty" xml:"PERIOD,omitempty" bson:"period,omitempty"`
	Inclination  string `json:"INCLINATION,omitempty" xml:"INCLINATION,omitempty" bson:"inclination,omitempty"`
	Apogee       string `json:"APOGEE,omitempty" xml:"APOGEE,omitempty" bson:"apogee,omitempty"`
	Perigee      string `json:"PERIGEE,omitempty" xml:"PERIGEE,omitempty" bson:"perigee,omitempty"`
	Comment      string `json:"COMMENT,omitempty" xml:"COMMENT,omitempty" bson:"comment,omitempty"`
	CommentCode  string `json:"COMMENTCODE,omitempty" xml:"COMMENTCODE,omitempty" bson:"comment_code,omitempty"`
	RcsValue     string `json:"RCSVALUE,omitempty" xml:"RCSVALUE,omitempty" bson:"rcs_value,omitempty"`
	RcsSize      string `json:"RCS_SIZE,omitempty" xml:"RCS_SIZE,omitempty" bson:"rcs_size,omitempty"`
	File         string `json:"FILE,omitempty" xml:"FILE,omitempty" bson:"file,omitempty"`
	Launch       string `json:"LAUNCH,omitempty" xml:"LAUNCH,omitempty" bson:"launch,omitempty"`
	LaunchYear   string `json:"LAUNCH_YEAR,omitempty" xml:"LAUNCH_YEAR,omitempty" bson:"launch_year,omitempty"`
	LaunchNum    string `json:"LAUNCH_NUM,omitempty" xml:"LAUNCH_NUM,omitempty" bson:"launch_num,omitempty"`
	LaunchPiece  string `json:"LAUNCH_PIECE,omitempty" xml:"LAUNCH_PIECE,omitempty" bson:"launch_piece,omitempty"`
	Current      string `json:"CURRENT,omitempty" xml:"CURRENT,omitempty" bson:"current,omitempty"`
	ObjectName   string `json:"OBJECT_NAME,omitempty" xml:"OBJECT_NAME,omitempty" bson:"object_name,omitempty"`
	ObjectID     string `json:"OBJECT_ID,omitempty" xml:"OBJECT_ID,omitempty" bson:"object_id,omitempty"`
	ObjectNumber string `json:"OBJECT_NUMBER,omitempty" xml:"OBJECT_NUMBER,omitempty" bson:"object_number,omitempty"`
}
