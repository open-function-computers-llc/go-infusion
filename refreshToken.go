package infusion

import (
	"encoding/json"
	"io/ioutil"
	"net/url"
)

func GetRefreshToken() error {
	data := url.Values{}
	data.Add("grant_type", "refresh_token")
	data.Add("refresh_token", refreshToken)

	r, err := postFormRequest("https://api.infusionsoft.com/token", data)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	type newTokenResponse struct {
		AccessToken  string `json:"access_token"`
		TokenType    string `json:"token_type"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		Scope        string `json:"scope"`
	}

	ntr := newTokenResponse{}
	err = json.Unmarshal(body, &ntr)
	if err != nil {
		return err
	}

	authToken = ntr.AccessToken
	refreshToken = ntr.RefreshToken
	log.Info("Updated tokens. New access token: " + authToken + " refresh token: " + refreshToken)
	return nil
}
