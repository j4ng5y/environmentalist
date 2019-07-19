package environmentalistrestful

import "net/http"

type NewAWSSSMSecretRequest struct {
	SecretName  string `json:"secretName"`
	SecretValue string `json:"secretValue"`
}

type NewAWSSSMSecretResponse struct {
	Success bool `json:"success"`
}

type UpdateAWSSSMSecretRequest struct {
	SecretName     string `json:"secretName"`
	SecretNewValue string `json:"secretNewValue"`
}

type UpdateAWSSSMSecretResponse struct {
	Success        bool   `json:"success"`
	SecretOldvalue string `json:"secretOldValue"`
	SecretNewValue string `json:"secretNewValue"`
}

type DeleteAWSSSMSecretRequest struct {
	SecretName string `json:"secretName"`
}

type DeleteAWSSSMSecretResponse struct {
	Success bool `json:"success"`
}

type ViewAWSSSMSecretRequest struct {
	SecretName string `json:"secretName"`
}

type ViewAWSSSMSecretResponse struct {
	SecretValue string `json:"secretValue"`
}

func NewAWSSSMSecret(w http.ResponseWriter, r *http.Request) {

}

func UpdateAWSSSMSecret(w http.ResponseWriter, r *http.Request) {

}

func DeleteAWSSSMSecret(w http.ResponseWriter, r *http.Request) {

}

func ViewAWSSSMSecret(w http.ResponseWriter, r *http.Request) {

}
