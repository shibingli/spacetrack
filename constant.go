package spacetrack

const (
	DefaultBaseURL   = "https://www.space-track.org"
	DefaultLoginURL  = "/ajaxauth/login"
	DefaultLogoutURL = "/ajaxauth/logout"

	RequestControllerBasicSpaceData    = "basicspacedata"
	RequestControllerExpandedSpaceData = "expandedspacedata"
	RequestControllerFileShare         = "fileshare"
	RequestControllerCombinedOpsData   = "combinedopsdata"
	RequestControllerPublicFiles       = "publicfiles"

	RequestActionQuery    = "query"
	RequestActionModelDef = "modeldef"

	RequestClassesAnnouncement      = "announcement"
	RequestClassesBoxScore          = "boxscore"
	RequestClassesCdmPublic         = "cdm_public"
	RequestClassesDecay             = "decay"
	RequestClassesDirs              = "dirs"              //GET
	RequestClassesGetPublicDataFile = "getpublicdatafile" //POST
	RequestClassesGp                = "gp"
	RequestClassesGpHistory         = "gp_history"
	RequestClassesLaunchSite        = "launch_site"
	RequestClassesLoadPublicData    = "loadpublicdata" //GET
	RequestClassesOMM               = "omm"
	RequestClassesSatcat            = "satcat"
	RequestClassesSatcatChange      = "satcat_change"
	RequestClassesSatcatDebut       = "satcat_debut"
	RequestClassesTIP               = "tip"
	RequestClassesTLE               = "tle"
	RequestClassesTLELatest         = "tle_latest"
	RequestClassesTLEPublish        = "tle_publish"

	RESTOperatorsGT                    = ">"
	RESTOperatorsLT                    = "<"
	RESTOperatorsNE                    = "<>"
	RESTOperatorsOR                    = ","
	RESTOperatorsRange                 = "--"
	RESTOperatorsNullValue             = "null-val"
	RESTOperatorsLike                  = "~~"
	RESTOperatorsWildcard              = "^"
	RESTOperatorsCurrentSystemDateTime = "now"

	RESTPredicatesClass           = "class"
	RESTPredicatesPredicates      = "predicates"
	RESTPredicatesMetaData        = "metadata"
	RESTPredicatesLimit           = "limit"
	RESTPredicatesDistinct        = "distinct"
	RESTPredicatesOrderBy         = "orderby"
	RESTPredicatesFormat          = "format"
	RESTPredicatesEmptyResultShow = "emptyresult"
	RESTPredicatesFavorites       = "favorites"
	RESTPredicatesRecursive       = "recursive"
)
