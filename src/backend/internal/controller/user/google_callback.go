package user

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"github.com/Roongkun/software-eng-ii/internal/controller/util"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

func (r *Resolver) GoogleCallback(c *gin.Context) {
	config := util.GetGoogleLibConfig(c)
	code := c.Query("code")
	var pathURL string = "/"

	if c.Query("state") != "" {
		pathURL = c.Query("state")
	}

	if code == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  "fail",
			"message": "Authorization code not provided",
		})
	}

	tokenRes, err := getGoogleOAuth2Token(config, code)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  "fail",
			"message": err.Error(),
		})
		c.Abort()
		return
	}

	googleUser, err := getGoogleUser(tokenRes.AccessToken, tokenRes.IdToken)

}

type GoogleOAuth2Token struct {
	AccessToken string
	IdToken     string
}

func getGoogleOAuth2Token(config *oauth2.Config, code string) (*GoogleOAuth2Token, error) {
	const rootURL = "https://oath2.googleapis.com/token"

	values := url.Values{}
	values.Add("grant_type", "authorization_code")
	values.Add("code", code)
	values.Add("client_id", config.ClientID)
	values.Add("client_secret", config.ClientSecret)
	values.Add("redirect_url", config.RedirectURL)

	query := values.Encode()

	req, err := http.NewRequest("POST", rootURL, bytes.NewBufferString(query))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve token")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleOAuth2TokenRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleOAuth2TokenRes); err != nil {
		return nil, err
	}

	tokenBody := &GoogleOAuth2Token{
		AccessToken: GoogleOAuth2TokenRes["access_token"].(string),
		IdToken:     GoogleOAuth2TokenRes["id_token"].(string),
	}

	return tokenBody, nil
}

type GoogleUserResult struct {
	Id            string
	Email         string
	VerifiedEmail bool
	Name          string
	GivenName     string
}

func getGoogleUser(accessToken string, idToken string) (*GoogleUserResult, error) {
	rootURL := fmt.Sprintf("https://www.googleapis.com/oauth2/v1/userinfo?alt=json&access_token=%s", accessToken)

	req, err := http.NewRequest("GET", rootURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", idToken))

	client := http.Client{
		Timeout: time.Second * 30,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("could not retrieve user")
	}

	var resBody bytes.Buffer
	_, err = io.Copy(&resBody, res.Body)
	if err != nil {
		return nil, err
	}

	var GoogleUserRes map[string]interface{}

	if err := json.Unmarshal(resBody.Bytes(), &GoogleUserRes); err != nil {
		return nil, err
	}

	userBody := &GoogleUserResult{
		Id:            GoogleUserRes["id"].(string),
		Email:         GoogleUserRes["email"].(string),
		VerifiedEmail: GoogleUserRes["verified_email"].(bool),
		Name:          GoogleUserRes["name"].(string),
		GivenName:     GoogleUserRes["given_name"].(string),
	}

	return userBody, nil
}
