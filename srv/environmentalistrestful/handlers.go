package environmentalistrestful

import "net/http"

type NewSecretRequest struct {
	SecretName  string `json:"secretName"`
	SecretValue string `json:"secretValue"`
}

type NewSecretResponse struct {
	Success bool `json:"success"`
}

type UpdateSecretRequest struct {
	SecretName     string `json:"secretName"`
	SecretNewValue string `json:"secretNewValue"`
}

type UpdateSecretResponse struct {
	Success        bool   `json:"success"`
	SecretOldvalue string `json:"secretOldValue"`
	SecretNewValue string `json:"secretNewValue"`
}

type DeleteSecretRequest struct {
	SecretName string `json:"secretName"`
}

type DeleteSecretResponse struct {
	Success bool `json:"success"`
}

type ViewSecretRequest struct {
	SecretName string `json:"secretName"`
}

type ViewSecretResponse struct {
	SecretValue string `json:"secretValue"`
}

func NewSecret(w http.ResponseWriter, r *http.Request) {

}

func UpdateSecret(w http.ResponseWriter, r *http.Request) {

}

func DeleteSecret(w http.ResponseWriter, r *http.Request) {

}

func ViewSecret(w http.ResponseWriter, r *http.Request) {

}

// identifyBackingVault identifies the current active backing secrets vault
//
// Arguments:
//     None
//
// Returns:
//     I'm not sure yet
func identifyBackingVault() {

}
