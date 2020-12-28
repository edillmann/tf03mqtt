package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	"strings"
)

type sensorReg struct {
	Name              string `json:"name"`
	UnitOfMeasurement string `json:"unit_of_measurement,omit_empty"`
	Icon              string `json:"icon,omit_empty"`
	StateTopic        string `json:"state_topic,omit_empty"`
	ValueTemplate     string `json:"value_template,omit_empty"`
	ExpireAfter       uint32 `json:"expire_after"`
	UniqueId          string `json:"unique_id"`
	Platform          string `json:"platform"`
	Device            struct {
		Name         string   `json:"name"`
		Manufacturer string   `json:"manufacturer"`
		Model        string   `json:"model"`
		Identifiers  []string `json:"identifiers"`
	} `json:"device"`
}

func (cm *Coulombmeter) register(c MQTT.Client) {
	for _, s := range cfg.Sensors {
		var r sensorReg
		r.Name = s.Name
		r.UnitOfMeasurement = s.Unit
		r.Icon = s.Icon
		r.StateTopic = s.StateTopic
		r.ValueTemplate = s.ValueTemplate
		r.ExpireAfter = uint32(60)
		r.UniqueId = strings.ToLower(cfg.Name + "_" + s.Name)
		r.Platform = "mqtt"
		r.Device.Name = cfg.Name
		r.Device.Manufacturer = cfg.Manufacturer
		r.Device.Model = cfg.Model
		r.Device.Identifiers = append(r.Device.Identifiers, strings.ToLower(cfg.Name))

		_ = publishAsJson(strings.ToLower("homeassistant/sensor/"+r.Name+"/config"), r, c)
	}
}
