package config

type Config struct {
	OAuth2Google struct {
		ClientID     string `toml:"client_id"`
		ClientSecret string `toml:"client_secret"`
		RedirectURL  string `toml:"redirect_url"`
	} `toml:"oauth2_google"`
}
