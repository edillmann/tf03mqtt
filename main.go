package main

import (
	"flag"
	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"os"
)

var cfg config

//define a function for the default message handler
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	log.Warn("TOPIC: %s\n", msg.Topic())
	log.Warn("MSG: %s\n", msg.Payload())
}

func main() {

	var loglevel = flag.String("loglevel", "info", "log level [trace,debug,info,warn,error,fatal,panic]")
	var configFile = flag.String("config", "config.yaml", "config file")
	flag.Parse()

	// configure log formatter
	setLogLevel(*loglevel)
	log.SetFormatter(&formatter{})

	err := parseYaml(*configFile, &cfg)

	if err != nil {
		log.Error(err.Error())
		os.Exit(-1)
	}

	opts := MQTT.NewClientOptions().AddBroker(cfg.MqttServer)
	opts.SetClientID(cfg.Name)
	opts.SetDefaultPublishHandler(f)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// allocate coulometer
	cm := Coulombmeter{
		port: cfg.SerialPort,
	}
	// register sensors to Home Assistant
	if cfg.HaRegister {
		cm.register(c)
	}

	// init serial port
	err = cm.init()
	if err != nil {
		log.Errorln(err)
		os.Exit(-1)
	}

	// allocate channel to store frame
	ch := make(chan uint8, 64)
	// start frame parser goroutine
	go cm.parseFrame(ch, c)
	// infinite loop reading on serial port
	err = cm.poll(ch)
	if err != nil {
		log.Errorln(err)
		os.Exit(-2)
	}

}
