package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/form3tech-oss/terraform-provider-form3/api/httputil"
	"github.com/form3tech-oss/terraform-provider-form3/client"
	"github.com/giantswarm/retry-go"
	"github.com/go-openapi/runtime"
	rc "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/logging"
)

var tokenCache = sync.Map{}

type CachedToken struct {
	Token   string
	Expires time.Time
}

type AuthenticatedClient struct {
	AccessToken           string
	SecurityClient        *client.Form3CorelibDataStructures
	NotificationClient    *client.Form3CorelibDataStructures
	Config                *client.TransportConfig
	HttpClient            *http.Client
	OrganisationId        string
	OrganisationClient    *client.Form3CorelibDataStructures
	AssociationClient     *client.Form3CorelibDataStructures
	AccountClient         *client.Form3CorelibDataStructures
	LimitsClient          *client.Form3CorelibDataStructures
	PaymentdefaultsClient *client.Form3CorelibDataStructures
	TransactionClient     *client.Form3CorelibDataStructures
	SystemClient          *client.Form3CorelibDataStructures
}

type AuthenticatedClientCheckRedirect struct {
}

type roundTripperFunc func(*http.Request) (*http.Response, error)

type Request struct {
	Body io.ReadSeeker
	*http.Request
}

type ReadSeekerCloser interface {
	io.ReadSeeker
	io.Closer
}

var _ ReadSeekerCloser = (*ReadSeekerCloserImpl)(nil)

type ReadSeekerCloserImpl struct {
	ReadSeeker io.ReadSeeker
	Closer     io.Closer
}

type debugReqResp struct {
	req string
	res string
}

func (r *ReadSeekerCloserImpl) Read(p []byte) (n int, err error) {
	return r.ReadSeeker.Read(p)
}

func (r *ReadSeekerCloserImpl) Seek(offset int64, whence int) (int64, error) {
	return r.ReadSeeker.Seek(offset, whence)
}

func (r *ReadSeekerCloserImpl) Close() error {
	return nil
}

func NewReaderSeekerCloser(request *http.Request) (ReadSeekerCloser, error) {
	readerSeekerCloser := ReadSeekerCloserImpl{}
	if request.Body != nil {
		bodyBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			return nil, err
		}
		readerSeekerCloser.ReadSeeker = bytes.NewReader(bodyBytes)
	}
	return &readerSeekerCloser, nil
}

func (fn roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return fn(req)
}

func (r *AuthenticatedClientCheckRedirect) CheckRedirect(req *http.Request, via []*http.Request) error {
	req.Header.Add("Authorization", via[0].Header.Get("Authorization"))
	return nil
}

func getEnvIntDefault(name string, defaultValue int) int {

	strValue, ok := os.LookupEnv(name)
	if !ok {
		return defaultValue
	}
	value, err := strconv.Atoi(strValue)
	if err != nil {
		panic(fmt.Sprintf("expected an int value for environment variable %s, got %q", name, strValue))
	}
	return value
}

func NewAuthenticatedClient(config *client.TransportConfig) *AuthenticatedClient {
	a := &AuthenticatedClientCheckRedirect{}
	var authClient *AuthenticatedClient

	retrySleep := 500 * time.Millisecond
	retryMax := getEnvIntDefault("MAX_API_RETRIES", 10)
	retryTimeout := time.Duration(retryMax) * retrySleep * 2

	h := &http.Client{
		Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			if req.Body != nil && req.Body != http.NoBody {
				rewindableBody, err := NewReaderSeekerCloser(req)
				if err != nil {
					return nil, err
				}
				req.Body = rewindableBody
			}

			if len(authClient.AccessToken) > 0 {
				req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", authClient.AccessToken))
			}

			debugReqResp := debugReqResp{}

			if logging.IsDebugOrHigher() {
				dump, errDump := httputil.SecureDumpRequest(req)
				if errDump != nil {
					log.Fatal(errDump)
				}

				if req.Body != nil {
					rewindableBody, err := NewReaderSeekerCloser(req)
					if err != nil {
						return nil, err
					}
					req.Body = rewindableBody
				}

				debugReqResp.req = string(dump)
			}

			// In case some API initially responds with 403, retry the request until permissions propagate.
			var resp *http.Response
			retryableFunc := func() error {
				var err error
				log.Printf("[DEBUG] retrying %s %s", req.Method, req.URL)
				if req.Body != nil && req.Body != http.NoBody {
					if body, ok := req.Body.(ReadSeekerCloser); ok {
						if _, err := body.Seek(0, 0); err != nil {
							return fmt.Errorf("failed to seek request body: %s", err)
						}
					}
				}

				resp, err = http.DefaultTransport.RoundTrip(req)
				if err != nil {
					return err
				}
				if resp.StatusCode == 403 {
					return fmt.Errorf("status code: %d", resp.StatusCode)
				}

				return nil
			}
			err := retry.Do(retryableFunc, retry.MaxTries(retryMax), retry.Sleep(retrySleep), retry.Timeout(retryTimeout))

			if logging.IsDebugOrHigher() {
				if resp != nil {
					dump, errDump := httputil.SecureDumpResponse(resp)
					if errDump != nil {
						log.Fatal(errDump)
					}

					debugReqResp.res = string(dump)
				}

				log.Printf("%s%s", debugReqResp.req, debugReqResp.res)
			}

			if err != nil {
				return resp, err
			}

			return resp, nil
		}),
		CheckRedirect: a.CheckRedirect,
	}

	config.WithBasePath("/v1/security")
	rt1 := rc.NewWithClient(config.Host, config.BasePath, config.Schemes, h)
	securityClient := client.New(rt1, strfmt.Default)

	config.WithBasePath("/v1/notification")
	rt2 := rc.NewWithClient(config.Host, config.BasePath, config.Schemes, h)
	notificationClient := client.New(rt2, strfmt.Default)

	config.WithBasePath("/v1/organisation")
	rt3 := rc.NewWithClient(config.Host, config.BasePath, config.Schemes, h)
	organisationClient := client.New(rt3, strfmt.Default)

	config.WithBasePath("/v1/organisation/units/associations")
	rt4 := rc.NewWithClient(config.Host, config.BasePath, config.Schemes, h)
	associationsClient := client.New(rt4, strfmt.Default)

	config.WithBasePath("/v1/organisation")
	rt5 := rc.NewWithClient(config.Host, config.BasePath, config.Schemes, h)
	accountClient := client.New(rt5, strfmt.Default)

	config.WithBasePath("/v1/organisation/units/")
	rt6 := rc.NewWithClient(config.Host, config.BasePath, config.Schemes, h)
	limitsClient := client.New(rt6, strfmt.Default)

	config.WithBasePath("/v1/transaction")
	rt7 := rc.NewWithClient(config.Host, config.BasePath, config.Schemes, h)
	transactionClient := client.New(rt7, strfmt.Default)

	config.WithBasePath("/v1/organisation/units/")
	rt8 := rc.NewWithClient(config.Host, config.BasePath, config.Schemes, h)
	paymentdefaultsClient := client.New(rt8, strfmt.Default)

	config.WithBasePath("/v1/system/")
	rt9 := rc.NewWithClient(config.Host, config.BasePath, config.Schemes, h)
	systemClient := client.New(rt9, strfmt.Default)

	authClient = &AuthenticatedClient{
		AssociationClient:     associationsClient,
		SecurityClient:        securityClient,
		NotificationClient:    notificationClient,
		OrganisationClient:    organisationClient,
		AccountClient:         accountClient,
		LimitsClient:          limitsClient,
		PaymentdefaultsClient: paymentdefaultsClient,
		TransactionClient:     transactionClient,
		SystemClient:          systemClient,
		HttpClient:            h,
		Config:                config,
	}

	configureRuntime(rt1, authClient)
	configureRuntime(rt2, authClient)
	configureRuntime(rt3, authClient)
	configureRuntime(rt4, authClient)
	configureRuntime(rt5, authClient)
	configureRuntime(rt6, authClient)
	configureRuntime(rt7, authClient)
	configureRuntime(rt8, authClient)
	configureRuntime(rt9, authClient)

	return authClient
}

func configureRuntime(rt *rc.Runtime, authClient *AuthenticatedClient) {
	consumer := runtime.ConsumerFunc(func(reader io.Reader, data interface{}) error {
		b, _ := ioutil.ReadAll(reader)

		return fmt.Errorf("[ERROR] non json error: %s", string(b))
	})

	rt.Consumers[runtime.HTMLMime] = consumer
	rt.Consumers[runtime.TextMime] = consumer
	rt.Consumers[runtime.DefaultMime] = consumer
	rt.Consumers["application/vnd.api+json;charset=UTF-8"] = runtime.JSONConsumer()
	rt.Consumers["application/vnd.api+json"] = runtime.JSONConsumer()
}

func (r *AuthenticatedClient) Authenticate(clientId string, clientSecret string) error {
	token, cached := tokenCache.Load(clientId)
	if cached && time.Now().Before(token.(CachedToken).Expires) {
		r.AccessToken = token.(CachedToken).Token
		return nil
	}

	req, _ := http.NewRequest("POST", "/v1/oauth2/token", bytes.NewBufferString("grant_type=client_credentials"))
	req.URL.Host = r.Config.Host
	req.URL.Scheme = r.Config.Schemes[0]

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientId+":"+clientSecret)))

	if logging.IsDebugOrHigher() {
		dump, errDump := httputil.SecureDumpRequest(req)
		if errDump != nil {
			log.Fatal(errDump)
		}

		log.Printf("[DEBUG] %s %s", req.URL.String(), string(dump))
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	if resp.StatusCode != 200 {
		err = fmt.Errorf("Error returned while authenticating, response code was %v", resp.StatusCode)
		return err
	}

	defer resp.Body.Close()

	if logging.IsDebugOrHigher() {
		dump, errDump := httputil.SecureDumpResponse(resp)
		if errDump != nil {
			log.Fatal(errDump)
		}

		log.Printf("[DEBUG] %s %s", req.URL.String(), string(dump))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}

	loginResponse, err := getLoginResponse([]byte(body))

	if err != nil {
		return err
	}

	r.AccessToken = loginResponse.AccessToken
	tokenCache.Store(clientId, CachedToken{Token: r.AccessToken, Expires: time.Now().Add(time.Duration(loginResponse.ExpiresIn/2) * time.Second)})

	return nil
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
