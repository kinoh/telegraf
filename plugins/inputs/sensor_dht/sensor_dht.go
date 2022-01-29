package sensor_dht

import (
	"fmt"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
	"github.com/MichaelS11/go-dht"
)

type SensorDht struct {
	PinName string `toml:"pin_name"`
	Model   string `toml:"model"`
	Retry   int    `toml:"retry"`
}

func (s *SensorDht) Description() string {
	return "Environment sensor DHT-11 / DHT-22"
}

func (s *SensorDht) SampleConfig() string {
	return `
  pin_name = "GPIO4"
  model = "dht22"
  retry = 11
`
}

func (s *SensorDht) Init() error {
	return nil
}

func (s *SensorDht) Gather(acc telegraf.Accumulator) error {
	err := dht.HostInit()
	if err != nil {
		return fmt.Errorf("HostInit error: %s", err)
	}

	dht, err := dht.NewDHT(s.PinName, dht.Celsius, s.Model)
	if err != nil {
		return fmt.Errorf("NewDHT error: %s", err)
	}

	humidity, temperature, err := dht.ReadRetry(s.Retry)
	if err != nil {
		return fmt.Errorf("Read error: %s", err)
	}

	fields := make(map[string]interface{})
	fields["temperature"] = temperature
	fields["humidity"] = humidity

	tags := make(map[string]string)

	acc.AddFields("sensor_dht", fields, tags)

	return nil
}

func init() {
	inputs.Add("sensor_dht", func() telegraf.Input { return &SensorDht{} })
}
