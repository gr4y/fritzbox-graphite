package lib

type Configuration struct {
	Carbon CarbonConfiguration
}

type CarbonConfiguration struct {
	Host string
	Port int
}
