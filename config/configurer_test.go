package config

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// `mapstructure:",squash"` is required for embedded structs
type testConfig struct {
	BaseConfig `mapstructure:"core"`
	Name       string
}

type AnotherConfig struct {
	Desc string
	Misc string
	Name string
}

var testConfigInstance = testConfig{
	BaseConfig: BaseConfig{
		DB: DBConfig{
			Driver: "sqlite",
			DSN:    "./test.db",
		},
		HTTP: HTTPConfig{
			Address: ":8080",
		},
		LogLevel: "info",
	},
	Name: "Test Config",
}

var AnotherConfigInstance = &AnotherConfig{
	Desc: "desc",
	Misc: DefaultConfigFile,
	Name: EvnPrefix,
}

// Test Default Configuration
func TestReadConfigurationFromModelAndWriteToFile(t *testing.T) {
	appconfig := DefaultAppConfig
	_ = appconfig.AddJsonConfig(testConfigInstance)
	//_ = appconfig.AddJsonConfig(AnotherConfigInstance)
	appconfig.WriteConfig("config.yaml")
}

type NamedMan struct {
	Kevin string
	Smith string
}

func TestReadConfigurationFromFileAndWriteToStruct(t *testing.T) {
	appconfig, _ := NewYamlConfig("config.yaml")
	result := appconfig.Viper.Get("name")
	assert.Equal(t, "FLUENT", result)
	named := &NamedMan{}

	appconfig.ToStruct(AnotherConfigInstance)
	appconfig.ToStructByKey("nested", named)
	fmt.Println(named.Kevin)
	fmt.Println(named.Smith)

	assert.Equal(t, "FLUENT", AnotherConfigInstance.Name)
	assert.Equal(t, DefaultConfigFile, AnotherConfigInstance.Misc)

}

func TestReadConfigurationFromFileAndWriteToFile(t *testing.T) {
	appconfig, _ := NewYamlConfig("config.yaml")
	config := &testConfig{}
	fmt.Println(appconfig.Viper.Get("db"))
	appconfig.ToStruct(config)
	fmt.Println(config)
}
