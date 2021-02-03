package config

import "github.com/google/wire"

func SetupInject(config *Config){
	wire.NewSet(config, NewMongoClient)

}
