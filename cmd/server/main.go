package main

import (
	"log"
	"strings"

	"github.com/bwilczynski/homeapi"
	"github.com/bwilczynski/homeapi/lights"
	"github.com/bwilczynski/homeapi/lights/hue"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type config struct {
	Port    int
	HueHost string `mapstructure:"hue-host"`
	HueUser string `mapstructure:"hue-user"`
}

func main() {
	pflag.Int("port", 50051, "The server port")
	pflag.String("hue-host", "", "Address of hue bridge")
	pflag.String("hue-user", "", "Username to access Hue API")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)
	viper.SetEnvPrefix("hapi")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	var c config
	err := viper.Unmarshal(&c)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	s := homeapi.NewServer(c.Port)
	s.Run(func(s *grpc.Server) {
		lights.RegisterLightServiceServer(s, hue.NewServer(c.HueHost, c.HueUser))
	})
}
