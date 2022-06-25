package spacetrack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"spacetrack/utils"
	"strings"
)

type Auth struct {
	Username string         `json:"identity" xml:"identity"`
	Password string         `json:"password" xml:"password"`
	Cookies  []*http.Cookie `json:"cookies" xml:"cookies"`
}

func (s *SpaceTrack) jsonAuthInfo() ([]byte, error) {
	sta := &Auth{
		Username: s.Auth.Username,
		Password: s.Auth.Password,
	}

	return json.Marshal(sta)
}

func (s *SpaceTrack) Login() (*SpaceTrack, error) {

	authInfo, err := s.jsonAuthInfo()
	if err != nil {
		return nil, err
	}

	authInfo = bytes.TrimSpace(authInfo)

	loginUrl := strings.TrimSpace(DefaultLoginURL)
	if loginUrl == "" || len(authInfo) == 0 {
		return nil, fmt.Errorf("%s", utils.ErrorInvalidParameter)
	}

	loginUrl = utils.JoinURLPath(DefaultBaseURL, DefaultLoginURL)

	headers := map[string]interface{}{"Content-Type": "application/json"}

	resp, err := utils.NewHttpClient(
		http.MethodPost, loginUrl,
		bytes.NewReader(authInfo),
		&utils.HttpClientOpts{
			DisableCompression: true,
			DisableKeepAlives:  true,
			Headers:            headers,
		},
	)
	if err != nil {
		return nil, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf(
			utils.HttpResultFormat,
			resp.StatusCode,
			string(body),
		)
	}

	s.Auth.Cookies = resp.Cookies()

	return s, nil
}

func (s *SpaceTrack) Logout() (err error) {

	logoutUrl := strings.TrimSpace(DefaultLogoutURL)

	if logoutUrl == "" {
		return fmt.Errorf("%s", utils.ErrorInvalidParameter)
	}

	logoutUrl = utils.JoinURLPath(DefaultBaseURL, DefaultLogoutURL)

	headers := map[string]interface{}{"Content-Type": "application/json"}

	resp, err := utils.NewHttpClient(
		http.MethodPost,
		logoutUrl,
		nil,
		&utils.HttpClientOpts{
			DisableCompression: true,
			DisableKeepAlives:  true,
			Headers:            headers,
		},
	)

	if err != nil {
		return err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		return fmt.Errorf(
			utils.HttpResultFormat,
			resp.StatusCode,
			string(body),
		)
	}

	return nil
}
