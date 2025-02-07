package qgoconf

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"reflect"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

/**

**/

type AppConfig struct {
	Viper           *viper.Viper
	ConfigFile      string
	SavedConfigFile string
}

func NewYamlConfig(configFilePath any) (*AppConfig, error) {
	appConfig := &AppConfig{
		Viper:           viper.New(),
		SavedConfigFile: "config-json-example.json",
	}
	if configFilePath != nil && reflect.TypeOf(configFilePath).Kind() == reflect.String {
		appConfig.Viper.SetConfigType(DefaultConfigType)
		appConfig.Viper.SetConfigFile(configFilePath.(string))
		err := appConfig.Viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}

	return appConfig, nil
}

var DefaultAppConfig, _ = NewYamlConfig(DefaultConfigFile)

func (a *AppConfig) ToStruct(toStruct any) {
	err := a.Viper.Unmarshal(toStruct)
	if err != nil {
		panic(err)
	}
}

func (a *AppConfig) ToStructByKey(key string, toStruct any) {
	err := a.Viper.UnmarshalKey(key, toStruct)
	if err != nil {
		panic(err)
	}
}

func (a *AppConfig) Get(key string) any {
	return a.Viper.Get(key)
}

func (a *AppConfig) AddJsonConfig(config any) error {
	configJson, err := json.Marshal(config)
	if err != nil {
		return err
	}

	a.Viper.SetConfigType("json")
	err = a.Viper.ReadConfig(bytes.NewBuffer(configJson))
	return err
}

func (a *AppConfig) WriteConfig(filePath string) {
	err := a.Viper.SafeWriteConfigAs(filePath)
	if err != nil {
		slog.Error("write configuration failed", "error", err)
		_ = a.Viper.SafeWriteConfigAs(a.SavedConfigFile)
		return
	}
}

// WatchConfigChanges Watched Changed Overtime to update, it is dangerous operations
func (a *AppConfig) WatchConfigChanges() {
	a.Viper.WatchConfig()
	a.Viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := a.Viper.ReadInConfig(); err != nil {
			fmt.Printf("couldn't load config: %s", err)
		}
	})
}
