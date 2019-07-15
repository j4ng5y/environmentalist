package environmentalist

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

// Using the Hashicorp Vault product with envrionmentalist assumes a few things:
//
// 1.) That you have Hashicorp's Vault's HTTP RESTful API running and that it is available to Envrionmentalist
// 2.) That you have credentials to access your Vault Instance
//
// If any of the above is not the case, you may not use the Hashicorp Vault option.
//
// If you require assistance on setting any of these things up, please see the hashicorp vault documentation at https://learn.hashicorp.com/vault/getting-started/apis

// HashicorpVault is a structure to hold information about an external Hashicorp Vault instance
type HashicorpVault struct {
	Hostname    string
	Port        string
	ClientToken string
}

// NewHCV returns a new, blank instance of HashicorpVault
//
// Arguments:
//     None
//
// Returns:
//     (*HashicorpVault): A pointer to a new instance of the HashicorpVault structure
func NewHCV() *HashicorpVault {
	var H HashicorpVault
	return &H
}

// GetToken will retrieve the secure API token from the running Hashicorp Vault server instance
func (H *HashicorpVault) GetToken(authResponse *http.Response) *HashicorpVault {
	type resp struct {
		Auth struct {
			ClientToken   string   `json:"client_token"`
			Accessor      string   `json:"accessor"`
			Policies      []string `json:"policies"`
			TokenPolicies []string `json:"token_policies"`
			MetaData      struct {
				RoleName string `json:"role_name"`
			} `json:"metadata"`
			LeaseDuration int    `json:"lease_duration"`
			Renewable     bool   `json:"renewable"`
			EntityID      string `json:"entity_id"`
			TokenType     string `json:"token_type"`
		} `json:"auth"`
	}
	var r resp
	b, err := ioutil.ReadAll(authResponse.Body)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	err = json.Unmarshal(b, &r)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	H.ClientToken = r.Auth.ClientToken
	return H
}

// AppRoleAuth uses the Hashicorp Vault AppRole authentication method to try to gather the client token and returns the http response
//
// Arguments:
//     roleID (string): your Role ID
//     secretID (string): your Secret ID
//
// Returns:
//     (*http.Response): the http response object
func (H *HashicorpVault) AppRoleAuth(roleID string, secretID string) *http.Response {
	type role struct {
		RoleID   string `json:"role_id"`
		SecretID string `json:"secret_id"`
	}

	j, err := json.Marshal(role{
		RoleID:   roleID,
		SecretID: secretID,
	})
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	resp, err := http.Post(fmt.Sprintf("http://%s:%s/v1/auth/approle/login", H.Hostname, H.Port), "application/json", bytes.NewBuffer(j))
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return resp
}

// AWSAuth TODO: Build This
func (H *HashicorpVault) AWSAuth() *http.Response {
	var resp *http.Response
	return resp
}

// AzureAuth TODO: Build This
func (H *HashicorpVault) AzureAuth() *http.Response {
	var resp *http.Response
	return resp
}

// GitHubAuth TODO: Build This
func (H *HashicorpVault) GitHubAuth() *http.Response {
	var resp *http.Response
	return resp
}

// GoogleCloudAuth TODO: Build This
func (H *HashicorpVault) GoogleCloudAuth() *http.Response {
	var resp *http.Response
	return resp
}

// JWTOIDCAuth TODO: Build This
func (H *HashicorpVault) JWTOIDCAuth() *http.Response {
	var resp *http.Response
	return resp
}

// K8SAuth TODO: Build This
func (H *HashicorpVault) K8SAuth() *http.Response {
	var resp *http.Response
	return resp
}

// LDAPAuth TODO: Build This
func (H *HashicorpVault) LDAPAuth() *http.Response {
	var resp *http.Response
	return resp
}

// OktaAuth TODO: Build This
func (H *HashicorpVault) OktaAuth() *http.Response {
	var resp *http.Response
	return resp
}

// RADIUSAuth TODO: Build This
func (H *HashicorpVault) RADIUSAuth() *http.Response {
	var resp *http.Response
	return resp
}

// TLSCertAuth TODO: Build This
func (H *HashicorpVault) TLSCertAuth() *http.Response {
	var resp *http.Response
	return resp
}

// TokenAuth TODO: Build This
func (H *HashicorpVault) TokenAuth() *http.Response {
	var resp *http.Response
	return resp
}

// UserPassAuth TODO: Build This
func (H *HashicorpVault) UserPassAuth() *http.Response {
	var resp *http.Response
	return resp
}
