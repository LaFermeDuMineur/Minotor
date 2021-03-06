package config

import (
	"2miner-monitoring/log2miner"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var Cfg *Config

var Wtw *WalletToWatch

type WalletToWatch struct {
	Adress []string `yaml:"adress"`
}

type Config struct {
	ElasticsearchUser     string   `yaml:"elasticsearch_user"`
	ElasticsearchPassword string   `yaml:"elasticsearch_password"`
	APITokenEtherscan     string   `yaml:"api_token_etherscan"`
	LogLevel              string   `yaml:"log_level"`
	CaPath                string   `yaml:"ca_path"`
	TwoMinersURL          string   `yaml:"2miners_url"`
	RedisHost             string   `yaml:"redis_host"`
	RedisPassword         string   `yaml:"redis_password"`
	MinerListing          string   `yaml:"miner_listing"`
	APILogFile            string   `yaml:"api_log_file"`
	APIUsername           string   `yaml:"api_username"`
	APIPassword           string   `yaml:"api_password"`
	APIAdress             string   `yaml:"api_adress"`
	AdressFilePath        string   `yaml:"adress_file_path"`
	LockPath              string   `yaml:"lock_path"`
	CardsConfigFile       string   `yaml:"cards_config_file"`
	HiveosUrl             string   `yaml:"hiveos_api_url"`
	MinotorHiveOsUser     string   `yaml:"hiveos_minotor_user"`
	MinotorHiveOsPass     string   `yaml:"hiveos_minotor_password"`
	MinotorHiveosToken    string   `yaml:"hiveos_minotor_token"`
	CoinList              []string `yaml:"coin_list"`
	ElasticsearchHosts    []string `yaml:"elasticsearch_hosts"`
	ElasticsearchPort     int      `yaml:"elasticsearch_port"`
	RedisPort             int      `yaml:"redis_port"`
	RedisShortLifetime    int      `yaml:"redis_short_lifetime"`
	RedisMidLifetime      int      `yaml:"redis_mid_lifetime"`
	RedisLongLifetime     int      `yaml:"redis_long_lifetime"`
	APIPort               int      `yaml:"api_port"`
	APIFrontPort          int      `yaml:"api_front_port"`
	Factor                float64  `yaml:"factor"`
	EthFactor             float64  `yaml:"ether_factor"`
	GazFactor             float64  `yaml:"gaz_factor"`
}

func LoadYamlConfig(ConfigFilePath string) {
	t := Config{}
	data, err := ioutil.ReadFile(ConfigFilePath)
	log2miner.Error(err)
	err = yaml.Unmarshal(data, &t)
	log2miner.Error(err)
	Cfg = &t
	data, err = ioutil.ReadFile(Cfg.AdressFilePath)
	log2miner.Error(err)
	a := WalletToWatch{}
	err = yaml.Unmarshal(data, &a)
	log2miner.Error(err)
	Wtw = &a
	Wtw.Adress = a.Adress
}
