package environmentalistrestful

import "net/http"

type NewHashicorpVaultSecretRequest struct {
	SecretName  string `json:"secretName"`
	SecretValue string `json:"secretValue"`
}

type NewHashicorpVaultSecretResponse struct {
	Success bool `json:"success"`
}

type UpdateHashicorpVaultSecretRequest struct {
	SecretName     string `json:"secretName"`
	SecretNewValue string `json:"secretNewValue"`
}

type UpdateHashicorpVaultSecretResponse struct {
	Success        bool   `json:"success"`
	SecretOldvalue string `json:"secretOldValue"`
	SecretNewValue string `json:"secretNewValue"`
}

type DeleteHashicorpVaultSecretRequest struct {
	SecretName string `json:"secretName"`
}

type DeleteHashicorpVaultSecretResponse struct {
	Success bool `json:"success"`
}

type ViewHashicorpVaultSecretRequest struct {
	SecretName string `json:"secretName"`
}

type ViewHashicorpVaultSecretResponse struct {
	SecretValue string `json:"secretValue"`
}

func NewHashicorpVaultSecret(w http.ResponseWriter, r *http.Request) {

}

func UpdateHashicorpVaultSecret(w http.ResponseWriter, r *http.Request) {

}

func DeleteHashicorpVaultSecret(w http.ResponseWriter, r *http.Request) {

}

func ViewHashicorpVaultSecret(w http.ResponseWriter, r *http.Request) {

}
