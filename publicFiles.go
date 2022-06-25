package spacetrack

import "net/http"

type PublicFilesController struct {
	*BuildCommand
}

func (s *SpaceTrack) PublicFilesController() *PublicFilesController {
	return &PublicFilesController{
		BuildCommand: &BuildCommand{
			Paths:   s.getControllerPath(RequestControllerPublicFiles),
			Cookies: s.Auth.Cookies,
			Method:  http.MethodGet,
		},
	}
}

func (p *PublicFilesController) QueryDirs() *PublicFilesController {
	p.Paths = append(p.Paths, RequestClassesDirs)
	return p
}

func (p *PublicFilesController) QueryGetPublicDataFile() *PublicFilesController {
	p.Paths = append(p.Paths, RequestClassesGetPublicDataFile)
	p.Method = http.MethodPost
	return p
}

func (p *PublicFilesController) QueryLoadPublicData() *PublicFilesController {
	p.Paths = append(p.Paths, RequestClassesLoadPublicData)
	return p
}
