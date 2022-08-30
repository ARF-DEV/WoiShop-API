package helpers

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"

	"golang.org/x/oauth2"
)

func GetUserInfo(State string, Code string, OAuthStateString string, GoogleOAuthConfig *oauth2.Config) ([]byte, error) {
	if State != OAuthStateString {
		return nil, fmt.Errorf("invalid oauth state")
	}

	token, err := GoogleOAuthConfig.Exchange(context.Background(), Code)
	if err != nil {
		return nil, fmt.Errorf("code exchange failed: %s", err.Error())
	}

	response, err := http.Get("https://www.googleapis.com/oauth2/v2/userinfo?access_token=" + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed reading response body: %s", err.Error())
	}
	return contents, nil
}
