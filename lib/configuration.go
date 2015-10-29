package lib

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Duration struct{ time.Duration }

// MarshalJSON transforms a duration into JSON.
func (d *Duration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarshalJSON transform JSON into a Duration.
func (d *Duration) UnmarshalJSON(b []byte) error {
	var str string
	err := json.Unmarshal(b, &str)
	d.Duration, err = time.ParseDuration(str)
	return err
}

type Configuration struct {
	Carbon   CarbonConfiguration
	Router   RouterConfiguration
	Interval Duration
	Prefix   string
}

func (c *Configuration) Load(filename string) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(bytes, &c); err != nil {
		panic(err)
	}
}

type CarbonConfiguration struct {
	Host string
	Port int
}

func (c *CarbonConfiguration) GetAddress() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

type RouterConfiguration struct {
	Host string
	Port int
}

func (r *RouterConfiguration) GetAddress() string {
	return fmt.Sprintf("http://%s:%d", r.Host, r.Port)
}
