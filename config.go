package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type sensor struct {
	Name          string `yaml:"name"`
	Unit          string `yaml:"unit"`
	Icon          string `yaml:"icon"`
	StateTopic    string `yaml:"state_topic"`
	ValueTemplate string `yaml:"value_template"`
}

type config struct {
	Name         string   `yaml:"name"`
	Manufacturer string   `yaml:"manufacturer"`
	Model        string   `yaml:"model"`
	MqttServer   string   `yaml:"mqtt_server"`
	SerialPort   string   `yaml:"serial_port"`
	Topic        string   `yaml:"topic"`
	Sensors      []sensor `yaml:"sensors"`
}

func parseYaml(path string, c interface{}) error {
	cfg, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(cfg, c)
	if err != nil {
		return err
	}

	return nil
}
