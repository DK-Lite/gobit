package config

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type BybitConfig struct {
	AccessKey string `yaml:"ACCESS_KEY"`
	SecretKey string `yaml:"SECRET_KEY"`
}

type Config struct {
	Bybit BybitConfig `yaml:"Bybit"`
}

// var envLocal config

// func init() {
// 	envLocal = config{
// 		AccessKey: getenvStr("AccessKey", ""),
// 		SecretKey: getenvStr("SecretKey", ""),
// 	}
// }

func getenvStr(key string, defaultVal string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		log.Printf("not exists %s from env\n", value)
		return defaultVal
	}
	return value
}

func LoadConfig(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	v := Config{}
	if err := yaml.Unmarshal(buf, &v); err != nil {
		return nil, err
	}

	return &v, nil
}

// func GetBybit(path string) (*BybitConfig, error) {
// 	buf, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		return nil, err
// 	}

// 	p := &BybitConfig{}
// 	err = yaml.Unmarshal(buf, p)
// 	if err != nil {
// 		log.Fatalf("Unmarshal: %v", err)
// 	}
// 	os.Setenv("AccessKey", p.AccessKey)
// 	os.Setenv("SecretKey", p.SecretKey)
// 	log.Println(os.Getenv("AccessKey"), os.Getenv("SecretKey"))

// 	return p, nil
// }
