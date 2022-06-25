package spacetrack

import "net/http"

type BasicSpaceDataController struct {
	*BuildCommand
}

func (s *SpaceTrack) BasicSpaceDataController() *BasicSpaceDataController {
	return &BasicSpaceDataController{
		BuildCommand: &BuildCommand{
			Paths:      s.getControllerPath(RequestControllerBasicSpaceData),
			Cookies:    s.Auth.Cookies,
			Method:     http.MethodGet,
			StreamData: false,
		},
	}
}

func (b *BasicSpaceDataController) QueryAnnouncement() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesAnnouncement)
	return b
}

func (b *BasicSpaceDataController) QueryBoxScore() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesBoxScore)
	return b
}

func (b *BasicSpaceDataController) QueryCDMPublic() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesCdmPublic)
	return b
}

func (b *BasicSpaceDataController) QueryDecay() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesDecay)
	return b
}

func (b *BasicSpaceDataController) QueryGP() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesGp)
	return b
}

func (b *BasicSpaceDataController) QueryGPHistory() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesGpHistory)
	return b
}

func (b *BasicSpaceDataController) QueryLaunchSite() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesLaunchSite)
	return b
}

func (b *BasicSpaceDataController) QueryOMM() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesOMM)
	return b
}

func (b *BasicSpaceDataController) QuerySatcat() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesSatcat)
	return b
}

func (b *BasicSpaceDataController) QuerySatcatChange() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesSatcatChange)
	return b
}

func (b *BasicSpaceDataController) QuerySatcatDebut() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesSatcatDebut)
	return b
}

func (b *BasicSpaceDataController) QueryTIP() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesTIP)
	return b
}

func (b *BasicSpaceDataController) QueryTLE() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesTLE)
	return b
}

func (b *BasicSpaceDataController) QueryTLELatest() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesTLELatest)
	return b
}

func (b *BasicSpaceDataController) QueryTLEPublish() *BasicSpaceDataController {
	b.Paths = append(b.Paths, RequestClassesTLEPublish)
	return b
}
