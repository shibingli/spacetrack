package spacetrack

import (
	"net/http"
)

type ExpandedSpaceDataController struct {
	*BuildCommand
}

func (s *SpaceTrack) ExpandedSpaceDataController() *ExpandedSpaceDataController {
	return &ExpandedSpaceDataController{
		BuildCommand: &BuildCommand{
			Paths:   s.getControllerPath(RequestControllerExpandedSpaceData),
			Cookies: s.Auth.Cookies,
			Method:  http.MethodGet,
		},
	}
}
