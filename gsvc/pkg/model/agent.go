package model

import (
	"github.com/go-openapi/strfmt"
)

type CreateAgentRequest struct {
	Name            string `json:"name"`
	SecurityProfile string `json:"securityProfile"`
	EntityId        string `json:"entityId"`
}

type Agent struct {
	Name            string `json:"name"`
	SecurityProfile string `json:"securityProfile"`
	Id              string `json:"id"`
	ETag            string `json:"eTag"`
	EntityId        string `json:"entityId"`
}

type PagedAgent struct {
	Content          []Agent `json:"content"`
	Last             bool    `json:"last"`
	TotalPages       int     `json:"totalPages"`
	TotalElements    int     `json:"totalElements"`
	NumberOfElements int     `json:"numberOfElements"`
	First            bool    `json:"first"`
	Sort             string  `json:"sort"`
	Size             int     `json:"size"`
	Number           int     `json:"number"`
}

type Order struct {
	Direction    string `json:"direction"`
	Property     string `json:"property"`
	IgnoreCase   bool   `json:"ignoreCase"`
	NullHandling string `json:"nullHandling"`
	Descending   bool   `json:"descending"`
	Ascending    bool   `json:"ascending"`
}

type UpdateAgentRequest struct {
	Name            string `json:"name"`
	SecurityProfile string `json:"securityProfile"`
}

type OnlineStatus struct {
	Status string           `json:"status"`
	Since  *strfmt.DateTime `json:"timestamp"`
}

type UpdateDataSourceConfigurationRequest struct {
	ConfigurationId string       `json:"configurationId"`
	DataSources     []DataSource `json:"dataSources"`
}

type DataSourceConfigution struct {
	Id              string       `json:"id"`
	ETag            string       `json:"eTag"`
	ConfigurationId string       `json:"configurationId"`
	DataSources     []DataSource `json:"dataSources"`
}

type DataSource struct {
	Name        string      `json:"name"`
	Description string      `json:"description"`
	DataPoints  []DataPoint `json:"dataPoints"`
	CustomData  string      `json:"customData"`
}

type DataPoint struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Type        string `json:"type"`
	Unit        string `json:"unit"`
	CustomData  string `json:"customData"`
}

type Configuration struct {
	Content    OnboardingConfigurationContent `json:"id"`
	Expiration *strfmt.DateTime               `json:"timestamp"`
}

type OnboardingConfigurationContent struct {
	BaseUrl                 string `json:"baseUrl"`
	Iat                     string `json:"iat"`
	ClientCredentialProfile string `json:"clientCredentialProfile"`
	ClientId                string `json:"clientId"`
	Tenant                  string `json:"tenant"`
}

type OnboardingStatus struct {
	Status string `json:"status"`
}

type RegisterAgentRequest struct {
	Jwks Jwkeys `json:"jwks"`
}

type Jwkeys struct {
	Keys []Key `json:"keys"`
}

type Key struct {
	E   string `json:"e"`
	N   string `json:"n"`
	Kty string `json:"kty"`
	Kid string `json:"kid"`
}

type RotationKeys struct {
	ClientId string `json:"client_id"`
	Jwks     Jwkeys `json:"jwks"`
}

type ClientIdentifier struct {
	ClientId                string `json:"client_id"`
	ClientSecret            string `json:"client_secret"`
	ClientSecretExpiresAt   string `json:"client_secret_expires_at"`
	GrantTypes              string `json:"grant_types"`
	RegistrationAccessToken string `json:"registration_access_token"`
	RegistrationClientUri   string `json:"registration_client_uri"`
	TokenEndpointAuthMethod string `json:"token_endpoint_auth_method"`
}

type AccessTokenRequest struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   string `json:"expires_in"`
	Jti         string `json:"jti"`
	Scope       string `json:"scope"`
}

type TokenKey struct {
	Alg   string `json:"alg"`
	Use   string `json:"use"`
	Value string `json:"value"`
	E     string `json:"e"`
	N     string `json:"n"`
	Kty   string `json:"kty"`
	Kid   string `json:"kid"`
}

type TokenKeys struct {
	Keys []TokenKey `json:"keys"`
}

type UnAuthorized struct {
	Id      string `json:"id"`
	Message string `json:"message"`
}
