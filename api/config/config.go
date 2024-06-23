package config

import (
  "github.com/charmbracelet/log"
  "github.com/Salz-Team/salz/api/db"
  "github.com/Salz-Team/salz/api/objectstore"
  "golang.org/x/oauth2"
  "golang.org/x/oauth2/github"
  "os"
)

// Config is a struct that holds the configuration for the application
type Config struct {
  ApiDBHandler db.ApiDBHandler
  AuthDBHandler db.AuthDBHandler
  ObjectStoreHandler objectstore.ObjectStoreHandler
  LogLevel log.Level
  OAuth2Config *oauth2.Config
}

func getEnvOrDie(key string, defaultValue string) string {
  value, ok := os.LookupEnv(key)
  if !ok && defaultValue == "" {
    log.Fatal("Environment variable not set. No default value configured.", "key", key)
  }
  if !ok {
    return defaultValue
  }
  return value
}

func NewLocalConfig() *Config {
  oauthConfig := &oauth2.Config{
      ClientID:     getEnvOrDie("OAUTH_CLIENT_KEY", ""),
      ClientSecret: getEnvOrDie("OAUTH_CLIENT_SECRET", ""),
      RedirectURL:  "", // Dynamically constructed in /login controller
      Scopes:       []string{"read-user", "user-email"},
      Endpoint:     github.Endpoint,
  }
  return &Config{
    ApiDBHandler: db.NewPostgresHandler(),
    AuthDBHandler: db.NewPostgresHandler(),
    ObjectStoreHandler: objectstore.NewMinIOHandler(),
    LogLevel: log.DebugLevel,
    OAuth2Config: oauthConfig,
  }
}

func NewDevelopmentConfig() *Config {
  // Development uses SQLite for the database (although maybe it shouldn't? lol)
  // Actually, both databases
  oauthConfig := &oauth2.Config{
      ClientID:     getEnvOrDie("OAUTH_CLIENT_KEY", ""),
      ClientSecret: getEnvOrDie("OAUTH_CLIENT_SECRET", ""),
      RedirectURL:  "", // Dynamically constructed in /login controller
      Scopes:       []string{"read-user", "user-email"},
      Endpoint:     github.Endpoint,
  }
  return &Config{
    ApiDBHandler: db.NewPostgresHandler(),
    AuthDBHandler: db.NewPostgresHandler(),
    ObjectStoreHandler: objectstore.NewMinIOHandler(),
    LogLevel: log.DebugLevel,
    OAuth2Config: oauthConfig,
  }
}

func NewProductionConfig() *Config {
  var logLevel log.Level
  logLevelStr := getEnvOrDie("LOG_LEVEL", "Info")
  switch logLevelStr {
  case "Debug":
    logLevel = log.DebugLevel
  case "Info":
    logLevel = log.InfoLevel
  case "Warn":
    logLevel = log.WarnLevel
  case "Error":
    logLevel = log.ErrorLevel
  case "Fatal":
    logLevel = log.FatalLevel
  default:
    log.Fatal("Invalid log level", "level", logLevelStr)
  }
  oauthConfig := &oauth2.Config{
      ClientID:     getEnvOrDie("OAUTH_CLIENT_KEY", ""),
      ClientSecret: getEnvOrDie("OAUTH_CLIENT_SECRET", ""),
      RedirectURL:  "", // Dynamically constructed in /login controller
      Scopes:       []string{"read-user", "user-email"},
      Endpoint:     github.Endpoint,
  }
  return &Config{
    ApiDBHandler: db.NewPostgresHandler(),
    AuthDBHandler: db.NewPostgresHandler(),
    ObjectStoreHandler: objectstore.NewMinIOHandler(),
    LogLevel: logLevel,
    OAuth2Config: oauthConfig,
  }
}

func NewConfig() *Config {
  env := getEnvOrDie("ENV", "development")
  log.Info("Environment configuration", "env", env)
  switch env {
  case "local":
    return NewLocalConfig()
  case "development":
    return NewDevelopmentConfig()
  case "production":
    return NewProductionConfig()
  default:
    log.Fatal("Invalid environment", "env", env)
    return nil
  }
}
