package salesforceclient

import (
	"context"
	"fmt"
	"github.com/go-resty/resty/v2"
	"golang.org/x/oauth2"
	"net/http"
)

var cfg *Config

type Config struct {
	ClientID     string
	ClientSecret string
	RedirectUrl  string
	Scopes       []string
}

type Salesforce struct {
	Client *resty.Client
}

func (s *Salesforce) CreateCase(request CreateCaseRequest) error {
	resp, err := s.Client.R().
		SetBody(request).
		SetResult(&CreateCaseResponse{}).
		Post("/services/data/v56.0/sobjects/Case/")

	if err != nil {
		return fmt.Errorf("salesforce: %s", err.Error())
	}

	if resp.StatusCode() != http.StatusCreated {
		return fmt.Errorf(
			"salesforce: Expected: %d http response got %s. Body: %s",
			http.StatusCreated,
			resp.Status(),
			string(resp.Body()),
		)
	}

	caseResponse := resp.Result().(*CreateCaseResponse)
	if caseResponse.Success == false {
		return fmt.Errorf("%v", caseResponse.Errors)
	}

	return nil
}

func (s *Salesforce) Token() (*oauth2.Token, error) {
	httpClient := s.Client.GetClient()
	transport := httpClient.Transport.(*oauth2.Transport)

	token, err := transport.Source.Token()

	if err != nil {
		return nil, err
	}

	return token, nil
}

func CreateInstance(ctx context.Context, domain string, token *oauth2.Token) (*Salesforce, error) {
	conf := createConfig()
	if conf == nil {
		return nil, fmt.Errorf("salesforce: please inititate the Config")
	}

	httpClient := conf.Client(ctx, token)

	return &Salesforce{
		Client: resty.NewWithClient(httpClient).SetBaseURL(domain),
	}, nil
}

func CreateGrantLink(State string) (string, error) {
	conf := createConfig()
	if conf == nil {
		return "", fmt.Errorf("salesforce: please inititate the Config")
	}

	return conf.AuthCodeURL(State), nil
}

func ExchangeCode(ctx context.Context, code string) (*oauth2.Token, error) {
	conf := createConfig()
	if conf == nil {
		return nil, fmt.Errorf("salesforce: please inititate the Config")
	}

	token, err := conf.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func Initiate(config Config) {
	cfg = &config
}

func createConfig() *oauth2.Config {
	if cfg == nil {
		return nil
	}

	return &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		Scopes:       cfg.Scopes,
		Endpoint: oauth2.Endpoint{
			TokenURL:  "https://login.salesforce.com/services/oauth2/token",
			AuthURL:   "https://login.salesforce.com/services/oauth2/authorize",
		},
		RedirectURL: cfg.RedirectUrl,
	}
}
