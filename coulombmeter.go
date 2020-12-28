package main

import (
	MQTT "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
	"github.com/tarm/serial"
	"io"
	"os"
	"time"
)

// parse state
const (
	// wait start of frame
	sof = iota
	// wait soc
	soc
	// wait voltage
	voltage
	// wait capacity
	capacity
	// wait current
	current
	// wait time remaining
	remaining
	// wait crc
	crc
)

// tf03 frame content
type frame struct {
	Soc uint8 `json:"soc"`
	DeciVolt uint16 `json:"decivolt"`
	CapaMAh uint32 `json:"capamah"`
	CurrentMA int32 `json:"currentma"`
	RemainSec uint32 `json:"remainsec"`
	Sum uint8 `json:"-"`
}

type Coulombmeter struct {
	dev               *serial.Port
	port              string
}

// init serial port
func (cm *Coulombmeter) init() error {
	var err error

	c := &serial.Config{Name: cm.port, Baud: 9600, ReadTimeout: 2 * time.Second}
	cm.dev, err = serial.OpenPort(c)
	if err != nil {
		return err
	}
	// flush device buffers
	rs := make([]byte, 300)
	_, _ = cm.dev.Read(rs)
	return nil
}

// frame parser goroutine
func (cm *Coulombmeter) parseFrame(ch chan uint8, ct MQTT.Client) {
	var f frame
	state := sof
	var shift = 0

	for c := range ch {
		switch state {
		case sof:
			if c == 0xA5 {
				state = soc
				f.Sum = 0xA5
			}
			// ignore byte
			continue
		case soc:
			f.Soc = c
			f.Sum += c
			state = voltage
			shift = 8
			continue
		case voltage:
			f.DeciVolt = f.DeciVolt | (uint16(c) << shift)
			f.Sum += c
			shift -= 8
			if shift < 0 {
				state = capacity
				shift = 24
			}
			continue
		case capacity:
			f.CapaMAh = f.CapaMAh | (uint32(c) << shift)
			f.Sum += c
			shift -= 8
			if shift < 0 {
				state = current
				shift = 24
			}
			continue
		case current:
			f.CurrentMA = f.CurrentMA | (int32(c) << shift)
			f.Sum += c
			shift -= 8
			if shift < 0 {
				state = remaining
				shift = 16
			}
			continue
		case remaining:
			f.RemainSec = f.RemainSec | (uint32(c) << shift)
			f.Sum += c
			shift -= 8
			if shift < 0 {
				state = crc
			}
			continue
		case crc:
			if f.Sum == c {
				err := publishAsJson(cfg.Topic, f, ct)
				if err != nil {
					log.Errorln(err)
					os.Exit(-3)
				}
			}
			state = sof
			shift = 0
			f.reset()
			continue
		}
	}
}

// reset frame
func (f *frame) reset() {
	f.RemainSec = 0
	f.CurrentMA = 0
	f.CapaMAh = 0
	f.DeciVolt = 0
	f.Soc = 0
	f.Sum = 0
}

// serial port infinite reader
func (cm *Coulombmeter) poll(ch chan uint8) error {
	rs := make([]byte, 64)
	defer close(ch)
	for true {
		n, err := cm.dev.Read(rs)
		if err == io.EOF || n == 0 {
			log.Warnln("timeout")
			continue
		}
		if err != nil {
			log.Errorln(err)
			return err
		}
		for i := 0; i < n; i++ {
			ch <- rs[i]
		}
	}
	return nil
}

