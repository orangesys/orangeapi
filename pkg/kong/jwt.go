package kong

import (
	"net/http"

	"github.com/dghubble/sling"
	"github.com/orangesys/orangeapi/pkg/config"
)

type JWTCredentials struct {
	Data  []JWTCredential `json:"data,omitempty"`
	Total int             `json:"total,omitempty"`
}

//            "algorithm": "HS256",
//            "consumer_id": "e22e6550-d00b-478f-9b40-f71c3b9a1614",
//            "created_at": 1482679688000,
//            "id": "c3378e15-80fd-4eb3-a9aa-445179c52a18",
//            "key": "a36c3049b36249a3c9f8891cb127243c",
//            "secret": "e71829c351aa4242c2719cbfbe671c09"

type JWTCredential struct {
	Algorithm  string `json:"algorithm,omitempty"`
	Secret     string `json:"secret,omitempty"`
	ConsumerID string `json:"consumer_id,omitempty"`
	ID         string `json:"id,omitempty"`
	Key        string `json:"key,omitempty"`
	CreatedAt  int    `json:"created_at,omitempty"`
}

type JWTService struct {
	sling  *sling.Sling
	config *config.KongConfiguration
}

func NewJWTService(httpClient *http.Client, config *config.KongConfiguration) *JWTService {
	return &JWTService{
		sling:  sling.New().Client(httpClient).Base(config.KongAdminURL + "consumers/"),
		config: config,
	}
}

func (s *JWTService) Create(consumerID string, params *JWTCredential) (*JWTCredential, *http.Response, error) {
	credential := new(JWTCredential)
	resp, err := s.sling.New().Post(s.config.KongAdminURL + "consumers/" + consumerID + "/jwt").BodyJSON(params).ReceiveSuccess(credential)
	return credential, resp, err
}

func (s *JWTService) Get(consumerID, credentialID string) (*JWTCredential, *http.Response, error) {
	credential := new(JWTCredential)
	resp, err := s.sling.New().Path(consumerID + "/jwt/" + credentialID).ReceiveSuccess(credential)
	return credential, resp, err
}

func (s *JWTService) List(consumerID string) (*JWTCredentials, *http.Response, error) {
	credentials := new(JWTCredentials)
	resp, err := s.sling.New().Path(consumerID + "/jwt").ReceiveSuccess(credentials)
	return credentials, resp, err
}

func (s *JWTService) Update(consumerID string, params *JWTCredential) (*JWTCredential, *http.Response, error) {
	credential := new(JWTCredential)
	resp, err := s.sling.New().Patch(s.config.KongAdminURL + "consumers/" + consumerID + "/jwt/" + params.ID).BodyJSON(params).ReceiveSuccess(credential)
	return credential, resp, err
}

func (s *JWTService) Delete(consumerID, credentialID string) (string, *http.Response, error) {
	var message string
	resp, err := s.sling.New().Delete(s.config.KongAdminURL + "consumers/" + consumerID + "/jwt/" + credentialID).ReceiveSuccess(message)
	return message, resp, err
}
