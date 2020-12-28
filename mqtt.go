package main

import (
	"encoding/json"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

func publishAsJson(topic string, s interface{}, c MQTT.Client) error {
	buf, err := json.Marshal(s)
	json := string(buf)
	log.Trace(json)
	c.Publish(topic, 0, false, json)
	return err
}
