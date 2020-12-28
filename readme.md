# TF03K serial to Home Assistant 

This projet aim's to export TF03K coulomb meter information
to home assistant

![](images/tf03k.png)

## Getting Started

### Prerequisites

This software was written with docker in mind, the software itself
is written in go and should run on any platform where go is running

You can find instruction to install and configure docker here:
```
https://docs.docker.com/engine/install/debian
```

### Configuration

To configure you docker image you have to setup the following entries
in the supplied config.yaml

```
name: tf03k
manufacturer: Baiway
model: tf03k
mqtt_server: tcp://192.168.5.180:1883
serial_port: /dev/ttyUSB0
```

### Installing

```
make clean docker run
```

## Debugging

... to be written

## TODO

* handle MQTT user / password
* enhance documentation
* add lovelace card template

## Authors

* **Eric Dillmann** - *Initial work* - [edillmann](https://github.com/edillmann)

## License

This project is licensed under the Apache 2 License - see the [apache-license-2.0.md](apache-license-2.0.md) file for details

## Acknowledgments


