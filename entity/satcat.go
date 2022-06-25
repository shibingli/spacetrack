package entity

type Satcat struct {
	Intldes      string `json:"INTLDES" xml:"INTLDES"`
	NoradCatID   string `json:"NORAD_CAT_ID" xml:"NORAD_CAT_ID"`
	ObjectType   string `json:"OBJECT_TYPE" xml:"OBJECT_TYPE"`
	SatName      string `json:"SATNAME" xml:"SATNAME"`
	Country      string `json:"COUNTRY" xml:"COUNTRY"`
	Site         string `json:"SITE" xml:"SITE"`
	Decay        string `json:"DECAY" xml:"DECAY"`
	Period       string `json:"PERIOD" xml:"PERIOD"`
	Inclination  string `json:"INCLINATION" xml:"INCLINATION"`
	Apogee       string `json:"APOGEE" xml:"APOGEE"`
	Perigee      string `json:"PERIGEE" xml:"PERIGEE"`
	Comment      string `json:"COMMENT" xml:"COMMENT"`
	CommentCode  string `json:"COMMENTCODE" xml:"COMMENTCODE"`
	RcsValue     string `json:"RCSVALUE" xml:"RCSVALUE"`
	RcsSize      string `json:"RCS_SIZE" xml:"RCS_SIZE"`
	File         string `json:"FILE" xml:"FILE"`
	Launch       string `json:"LAUNCH" xml:"LAUNCH"`
	LaunchYear   string `json:"LAUNCH_YEAR" xml:"LAUNCH_YEAR"`
	LaunchNum    string `json:"LAUNCH_NUM" xml:"LAUNCH_NUM"`
	LaunchPiece  string `json:"LAUNCH_PIECE" xml:"LAUNCH_PIECE"`
	Current      string `json:"CURRENT" xml:"CURRENT"`
	ObjectName   string `json:"OBJECT_NAME" xml:"OBJECT_NAME"`
	ObjectID     string `json:"OBJECT_ID" xml:"OBJECT_ID"`
	ObjectNumber string `json:"OBJECT_NUMBER" xml:"OBJECT_NUMBER"`
}
