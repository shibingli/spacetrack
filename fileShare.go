package spacetrack

import (
	"net/http"
)

type FileShareController struct {
	*BuildCommand
}

func (s *SpaceTrack) FileShareController() *FileShareController {
	return &FileShareController{
		BuildCommand: &BuildCommand{
			Paths:   s.getControllerPath(RequestControllerFileShare),
			Cookies: s.Auth.Cookies,
			Method:  http.MethodGet,
		},
	}
}
