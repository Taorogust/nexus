// services/auth/config/config.go
package config

import (
    "fmt"
    "github.com/spf13/viper"
)

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

// Load reads environment variables and returns a validated Config.
// Panics if any required OIDC variables are missing.
func Load() Config {
    // Read environment variables as-is
    viper.AutomaticEnv()
    viper.SetDefault("SERVER_PORT", "8080")

    cfg := Config{
        OIDC: OIDCConfig{
            Issuer:       viper.GetString("OIDC_ISSUER_URL"),
            ClientID:     viper.GetString("OIDC_CLIENT_ID"),
            ClientSecret: viper.GetString("OIDC_CLIENT_SECRET"),
            RedirectURL:  viper.GetString("OIDC_REDIRECT_URL"),
        },
        Server: ServerConfig{
            Port: viper.GetString("SERVER_PORT"),
        },
    }

    // Validate required OIDC fields
    missing := []string{}
    if cfg.OIDC.Issuer == "" {
        missing = append(missing, "OIDC_ISSUER_URL")
    }
    if cfg.OIDC.ClientID == "" {
        missing = append(missing, "OIDC_CLIENT_ID")
    }
    if cfg.OIDC.ClientSecret == "" {
        missing = append(missing, "OIDC_CLIENT_SECRET")
    }
    if cfg.OIDC.RedirectURL == "" {
        missing = append(missing, "OIDC_REDIRECT_URL")
    }
    if len(missing) > 0 {
        panic(fmt.Errorf("missing required OIDC environment variables: %v", missing))
    }

    return cfg
}
