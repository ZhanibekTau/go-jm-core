package structures

type RedisConfig struct {
	RedisHost       string `mapstructure:"REDIS_HOST" json:"redis_host"`
	RedisPassword   string `mapstructure:"REDIS_PASSWORD" json:"redis_password"`
	RedisDb         int    `mapstructure:"REDIS_DB" json:"redis_db"`
	PoolSize        int    `mapstructure:"REDIS_POOL_SIZE" json:"redis_pool_size"`
	IdleTimeout     int    `mapstructure:"REDIS_IDLE_TIMEOUT" json:"redis_idle_timeout"`
	MaxConnLifetime int    `mapstructure:"REDIS_MAX_CONN_LIFETIME" json:"redis_max_conn_lifetime"`
}
