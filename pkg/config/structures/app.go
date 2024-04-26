package structures

type AppConfig struct {
	Name           string `mapstructure:"APP_NAME" json:"app_name"`
	HostName       string `mapstructure:"HOST_NAME" json:"host_name"`
	Version        string `mapstructure:"APP_VERSION" json:"app_version"`
	ServerAddress  string `mapstructure:"SERVER_ADDRESS" json:"server_address"`
	SentryDsn      string `mapstructure:"SENTRY_DSN"    json:"sentry_dsn"`
	AppEnv         string `mapstructure:"APP_ENV"    json:"app_env"`
	PinbaHost      string `mapstructure:"PINBA_HOST"    json:"pinba_host"`
	HandlerTimeout int    `mapstructure:"HANDLER_TIMEOUT"    json:"handler_timeout"`
}
