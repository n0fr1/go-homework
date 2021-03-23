package yaml

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/url"

	"gopkg.in/yaml.v2"
)

//Config is ...
type Config struct {
	Port        int    `yaml:"port"`
	DbURL       string `yaml:"db_url"`
	JaegerURL   string `yaml:"jaeger_url"`
	SentryURL   string `yaml:"sentry_url"`
	KafkaBroker string `yaml:"kafka_broker"`
	SomeAppID   string `yaml:"some_app_id"`
	SomeAppKey  string `yaml:"some_app_key"`
}

//ReadYaml is ...
func ReadYaml() {

	file, err := ioutil.ReadFile("./conf.yaml")

	if err != nil {
		log.Fatal(err)
	}

	conf := Config{}
	err = yaml.Unmarshal(file, &conf)

	//проверяем url
	jaegerURL, err := url.Parse(conf.JaegerURL)
	if err != nil || (jaegerURL.Scheme != "http" && jaegerURL.Scheme != "https") {
		log.Fatalf("Не верный формат jaeger_url: %v", err)
	}

	SentryURL, err := url.Parse(conf.SentryURL)
	if err != nil || (SentryURL.Scheme != "http" && SentryURL.Scheme != "https") {
		log.Fatalf("Не верный формат sentry_url: %v", err)
	}

	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	fmt.Printf("%v\n", conf)
}
