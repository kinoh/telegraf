package sensor_mhz19

import (
	"fmt"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	z19 "github.com/eternal-flame-AD/mh-z19"
	"github.com/tarm/serial"
)

type SensorMhz19 struct {
}

func (s *SensorMhz19) Description() string {
	return "Environment sensor MH-Z19 series"
}

func (s *SensorMhz19) SampleConfig() string {
	return ""
}

func (s *SensorMhz19) Init() error {
	return nil
}

func (s *SensorMhz19) Gather(acc telegraf.Accumulator) error {
	connConfig := z19.CreateSerialConfig()
	connConfig.Name = "/dev/serial0"

	port, err := serial.OpenPort(connConfig)
	if err != nil {
		return fmt.Errorf("OpenPort error: %s", err)
	}
	defer port.Close()
	concentration, err := z19.TakeReading(port)
	if err != nil {
		return fmt.Errorf("Reading error: %s", err)
	}

	fields := make(map[string]interface{})
	fields["co2"] = concentration

	tags := make(map[string]string)

	acc.AddFields("sensor_mhz19", fields, tags)

	return nil
}

func init() {
	inputs.Add("sensor_mhz19", func() telegraf.Input { return &SensorMhz19{} })
}
