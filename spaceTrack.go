package spacetrack

import (
	"spacetrack/utils"
)

type SpaceTrack struct {
	Auth Auth
}

func (s *SpaceTrack) getControllerPath(path string) []string {
	path = utils.PathCleanSep(path)

	paths := make([]string, 0, 0)
	paths = append(paths, DefaultBaseURL, path, RequestActionQuery, RESTPredicatesClass)
	return paths
}
