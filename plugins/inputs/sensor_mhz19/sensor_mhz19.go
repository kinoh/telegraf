package sensor_mhz19

import (
	_ "embed"
	"fmt"

	z19 "github.com/eternal-flame-AD/mh-z19"
	"github.com/tarm/serial"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

//go:embed sample.conf
var sampleConfig string

type SensorMhz19 struct{}

func (*SensorMhz19) SampleConfig() string {
	return sampleConfig
}

func (*SensorMhz19) Init() error {
	return nil
}

func (*SensorMhz19) Gather(acc telegraf.Accumulator) error {
	connConfig := z19.CreateSerialConfig()
	connConfig.Name = "/dev/serial0"

	port, err := serial.OpenPort(connConfig)
	if err != nil {
		return fmt.Errorf("OpenPort error: %w", err)
	}
	defer port.Close()

	concentration, err := z19.TakeReading(port)
	if err != nil {
		return fmt.Errorf("reading error: %w", err)
	}

	acc.AddFields("sensor_mhz19", map[string]interface{}{"co2": concentration}, map[string]string{})

	return nil
}

func init() {
	inputs.Add("sensor_mhz19", func() telegraf.Input { return &SensorMhz19{} })
}
