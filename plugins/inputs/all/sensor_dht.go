//go:build !custom || inputs || inputs.sensor_dht

package all

import _ "github.com/influxdata/telegraf/plugins/inputs/sensor_dht" // register plugin
