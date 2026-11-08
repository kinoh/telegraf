package sensor_dht

import (
	_ "embed"
	"fmt"

	"github.com/MichaelS11/go-dht"

	"github.com/influxdata/telegraf"
	"github.com/influxdata/telegraf/plugins/inputs"
)

//go:embed sample.conf
var sampleConfig string

type SensorDht struct {
	PinName string `toml:"pin_name"`
	Model   string `toml:"model"`
	Retry   int    `toml:"retry"`
}

func (*SensorDht) SampleConfig() string {
	return sampleConfig
}

func (*SensorDht) Init() error {
	return nil
}

func (s *SensorDht) Gather(acc telegraf.Accumulator) error {
	if err := dht.HostInit(); err != nil {
		return fmt.Errorf("HostInit error: %w", err)
	}

	sensor, err := dht.NewDHT(s.PinName, dht.Celsius, s.Model)
	if err != nil {
		return fmt.Errorf("NewDHT error: %w", err)
	}

	humidity, temperature, err := sensor.ReadRetry(s.Retry)
	if err != nil {
		return fmt.Errorf("read error: %w", err)
	}

	fields := map[string]interface{}{
		"temperature": temperature,
		"humidity":    humidity,
	}
	acc.AddFields("sensor_dht", fields, map[string]string{})

	return nil
}

func init() {
	inputs.Add("sensor_dht", func() telegraf.Input { return &SensorDht{} })
}
