// services/auth/cmd/server/main.go
package main

import (
    "context"
    "net/http"
    "time"

    "github.com/coreos/go-oidc/v3/oidc"
    "github.com/gorilla/mux"
    "github.com/sirupsen/logrus"

    "github.com/tu-org/actiora/services/auth/config"
    "github.com/tu-org/actiora/services/auth/handlers"
    "github.com/tu-org/actiora/services/auth/middleware"
)

func main() {
    // Carga config desde env / config file
    cfg := config.Load()
    log := logrus.New()

    // DEBUG: imprimimos lo que tenemos
    log.Infof("🔑 OIDC issuer URL: %q", cfg.OIDC.Issuer)
    log.Infof("🔑 OIDC client ID:    %q", cfg.OIDC.ClientID)

    if cfg.OIDC.Issuer == "" {
        log.Fatalf("🛑 OIDC issuer URL está vacío, asegúrate de que OIDC_ISSUER_URL esté definido en tu Deployment")
    }

    // Inicializa OIDC verifier
    ctx := context.Background()
    provider, err := oidc.NewProvider(ctx, cfg.OIDC.Issuer)
    if err != nil {
        log.Fatalf("failed to init OIDC provider: %v", err)
    }
    verifier := provider.Verifier(&oidc.Config{ClientID: cfg.OIDC.ClientID})

    r := mux.NewRouter()
    // Health check
    r.HandleFunc("/health", handlers.Health).Methods("GET")

    // Login (redirección a OIDC)
    r.HandleFunc("/login", handlers.LoginHandler(cfg)).Methods("GET")

    // Callback OIDC
    r.HandleFunc("/callback", handlers.CallbackHandler(cfg, verifier)).Methods("GET")

    // Protected endpoint
    api := r.PathPrefix("/api").Subrouter()
    api.Use(middleware.AuthMiddleware(verifier, cfg))
    api.HandleFunc("/userinfo", handlers.UserInfo).Methods("GET")

    srv := &http.Server{
        Handler:      r,
        Addr:         ":" + cfg.Server.Port,
        WriteTimeout: 15 * time.Second,
        ReadTimeout:  15 * time.Second,
    }
    log.Infof("auth service listening on %s", srv.Addr)
    log.Fatal(srv.ListenAndServe())
}
