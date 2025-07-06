// services/auth/config/config.go
package config

import "github.com/spf13/viper"

type OIDCConfig struct {
    Issuer       string
    ClientID     string
    ClientSecret string
    RedirectURL  string
}

type ServerConfig struct {
    Port string
}

type Config struct {
    OIDC   OIDCConfig
    Server ServerConfig
}

func Load() Config {
    // No establecemos ningún EnvPrefix, para leer las vars tal cual las defines
    viper.AutomaticEnv()
    viper.SetDefault("SERVER_PORT", "8080")

    return Config{
        OIDC: OIDCConfig{
            Issuer:       viper.GetString("OIDC_ISSUER_URL"),   // antes era OIDC_ISSUER
            ClientID:     viper.GetString("OIDC_CLIENT_ID"),
            ClientSecret: viper.GetString("OIDC_CLIENT_SECRET"),
            RedirectURL:  viper.GetString("OIDC_REDIRECT_URL"),
        },
        Server: ServerConfig{
            Port: viper.GetString("SERVER_PORT"),
        },
    }
}
