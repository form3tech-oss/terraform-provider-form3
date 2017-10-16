package form3

import (
	"encoding/base64"
	"net/http"
	"bytes"
	"mime/multipart"
	"encoding/json"
	"io/ioutil"
	"path"
	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
	"fmt"
	"github.com/go-openapi/runtime"
	"github.com/ewilde/go-form3/client"
)

type AuthenticatedClient struct {
	Config 			*client.TransportConfig
	AccessToken		string
	ApiClients		*client.Form3CorelibDataStructures
	HttpClient		*http.Client
}

type AuthenticatedClientCheckRedirect struct {

}

func (r *AuthenticatedClientCheckRedirect) CheckRedirect(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", via[0].Header.Get("Authorization"))
	return nil
}

func NewAuthenticatedClient(config *client.TransportConfig) *AuthenticatedClient {
	a := &AuthenticatedClientCheckRedirect{}
	h := &http.Client{
		Transport: http.DefaultTransport,
		CheckRedirect: a.CheckRedirect,
	}

	rt := rc.NewWithClient(config.Host, config.BasePath, config.Schemes, h)
	authClient := &AuthenticatedClient{
		ApiClients: client.New(rt, strfmt.Default),
		HttpClient: h,
		Config: config,
	}

	rt.Consumers["application/vnd.api+json;charset=UTF-8"] = runtime.JSONConsumer()
	rt.Consumers["application/vnd.api+json"] = runtime.JSONConsumer()
	rt.Do = authClient.Do

	return authClient
}

func (r *AuthenticatedClient) Authenticate(clientId string, clientSecret string) error {

	mpbody := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(mpbody)
	_ = writer.WriteField("grant_type", "client_credentials")
	writer.Close()
	req, _ := http.NewRequest("POST", "/oauth2/token", mpbody)
	req.URL.Host = r.Config.Host
	req.URL.Path = path.Join(r.Config.BasePath, req.URL.Path)
	req.URL.Scheme = r.Config.Schemes[0]

	req.Header.Set("Content-Type", writer.FormDataContentType())
	encoded := base64.StdEncoding.EncodeToString([]byte(clientId + ":" + clientSecret))
	req.Header.Set("Authorization", "Basic "+encoded)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	loginResponse, err := getLoginResponse([]byte(body))

	if err != nil {
		return err
	}

	r.AccessToken = loginResponse.AccessToken
	return nil
}



func (r *AuthenticatedClient) Do(ctx context.Context, client *http.Client, req *http.Request) (*http.Response, error) {
	if client == nil {
		client = r.HttpClient
	}

	if len(r.AccessToken) > 0 {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", r.AccessToken))
	}

	resp, err := client.Do(req.WithContext(ctx))
	// If we got an error, and the context has been canceled,
	// the context's error is probably more useful.
	if err != nil {
		select {
		case <-ctx.Done():
			err = ctx.Err()
		default:
		}
	}
	return resp, err
}

func getLoginResponse(body []byte) (*LoginResponse, error) {
	var s = new(LoginResponse)
	err := json.Unmarshal(body, &s)
	return s, err
}

type LoginResponse struct {
	TokenType string `json:"token_type,omitempty"`
	AccessToken string `json:"access_token"`
	ExpiresIn int `json:"expires_in"`
}

