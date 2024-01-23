package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-resty/resty/v2"
)

const (
	identityServerURL = "https://tu-identity-server"
	clientID          = "tu-client-id"
	clientSecret      = "tu-client-secret"
	username          = "tu-usuario"
	password          = "tu-contraseña"
)

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
	Scope       string `json:"scope"`
}

func main() {
	// Paso 1: Obtener el token de acceso mediante el flujo de contraseña de OAuth2
	token, err := getAccessToken()
	if err != nil {
		log.Fatal(err)
	}

	// Utiliza el token de acceso para realizar otras operaciones, como llamar a APIs protegidas.

	fmt.Printf("Token de acceso: %s\n", token.AccessToken)
}

func getAccessToken() (*TokenResponse, error) {
	client := resty.New()

	// Paso 1: Realizar la solicitud para obtener el token de acceso
	resp, err := client.R().
		SetBasicAuth(clientID, clientSecret).
		SetFormData(map[string]string{
			"grant_type": "password",
			"username":   username,
			"password":   password,
			"scope":      "openid",
		}).
		Post(fmt.Sprintf("%s/oauth2/token", identityServerURL))

	if err != nil {
		return nil, err
	}

	// Verificar si la solicitud fue exitosa (código de estado 200)
	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("Error al obtener el token de acceso. Código de estado: %d", resp.StatusCode())
	}

	// Analizar la respuesta JSON para obtener el token de acceso
	var tokenResponse TokenResponse
	err = json.Unmarshal(resp.Body(), &tokenResponse)
	if err != nil {
		return nil, err
	}

	return &tokenResponse, nil
}
