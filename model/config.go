package model

// Config ...
type Config struct {
	// APP
	App AppConfig `json:"app"`

	// VENDOR
	Kafka KafkaConfig `json:"kafka"`
	Nats  NatsConfig  `json:"rabbitmq"`

	// DATABASE
	Mysql MySQLConfig `json:"mysql"`

	// CRON
	Cron CronConfig `json:"cron"`

	// New Relic
	NewRelic NewRelicConfig `json:"new_relic"`
}

// AppConfig ...
type AppConfig struct {
	Env  string `json:"env" mapstructure:"env"`
	Port string `json:"port" mapstructure:"port"`
}

// KafkaConfig ...
type KafkaConfig struct {
	BrokerURL     string `json:"broker_url" mapstructure:"broker_url"`
	ConsumerGroup string `json:"consumer_group" mapstructure:"consumer_group"`
}

type NatsConfig struct {
	BrokerURL     string   `json:"broker_url" mapstructure:"broker_url"`
	ConsumerGroup string   `json:"consumer_group" mapstructure:"consumer_group"`
	UsersKey      KeyValue `json:"users_key" mapstructure:"users_key"`
	RunningTrade  KeyValue `json:"running_trade" mapstructure:"running_trade"`
	MaxConnection int      `json:"max_connection" mapstructure:"max_connection"`
}

type KeyValue struct {
	Bucket  string `json:"bucket" mapstructure:"bucket"`
	TTL     int    `json:"ttl" mapstructure:"ttl"`
	History int    `json:"history" mapstructure:"history"`
}

// New Relic
type NewRelicConfig struct {
	ServiceName string `json:"service_name"`
	LicenseKey  string `json:"license_key"`
}

type MySQLConfig struct {
	Host            string `json:"host"`
	Port            int    `json:"port"`
	Username        string `json:"username"`
	Password        string `json:"password"`
	DBname          string `json:"dbname"`
	MaxOpenConns    int    `json:"max_open_conns"`
	MaxIdleConns    int    `json:"max_idle_conns"`
	ConnMaxLifetime int    `json:"conn_max_lifetime"`
}

type CronConfig struct {
	HistoricalPrevious string `json:"historical_previous"`
}
