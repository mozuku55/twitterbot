package entity

// keys
type apiKey string
type accessToken string
type bearerToken string
type apiKeySecret string
type accessTokenSecret string
type keys = struct {
	apikey            apiKey
	accessToken       accessToken
	bearerToken       bearerToken
	apiSecretkey      apiKeySecret
	accessTokenSecret accessTokenSecret
}
