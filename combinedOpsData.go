package spacetrack

import (
	"net/http"
)

type CombinedOpsDataController struct {
	*BuildCommand
}

func (s *SpaceTrack) CombinedOpsDataController() *CombinedOpsDataController {
	return &CombinedOpsDataController{
		BuildCommand: &BuildCommand{
			Paths:   s.getControllerPath(RequestControllerCombinedOpsData),
			Cookies: s.Auth.Cookies,
			Method:  http.MethodGet,
		},
	}
}
