package form3

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ewilde/go-form3/client"
	"github.com/go-openapi/runtime"
	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"golang.org/x/net/context"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"path"
)

type AuthenticatedClient struct {
	AccessToken        string
	SecurityClient     *client.Form3CorelibDataStructures
	NotificationClient *client.Form3CorelibDataStructures
	Config             *client.TransportConfig
	HttpClient         *http.Client
	OrganisationId     string
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
		Transport:     http.DefaultTransport,
		CheckRedirect: a.CheckRedirect,
	}

	config.WithBasePath("/v1/security")
	rt1 := rc.NewWithClient(config.Host, config.BasePath, config.Schemes, h)
	securityClient := client.New(rt1, strfmt.Default)

	config.WithBasePath("/v1/notification")
	rt2 := rc.NewWithClient(config.Host, config.BasePath, config.Schemes, h)
	notificationClient := client.New(rt2, strfmt.Default)

	authClient := &AuthenticatedClient{
		SecurityClient:     securityClient,
		NotificationClient: notificationClient,
		HttpClient:         h,
		Config:             config,
	}

	configureRuntime(rt1, authClient)
	configureRuntime(rt2, authClient)

	return authClient
}
func configureRuntime(rt *rc.Runtime, authClient *AuthenticatedClient) {
	rt.Consumers["application/vnd.api+json;charset=UTF-8"] = runtime.JSONConsumer()
	rt.Consumers["application/vnd.api+json"] = runtime.JSONConsumer()
	rt.Do = authClient.Do
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
	TokenType   string `json:"token_type,omitempty"`
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}
