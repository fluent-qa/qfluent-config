# Go Configurations

## How to Use

Steps:
- Define Configuration Struct
```go
type testConfig struct {
	BaseConfig `mapstructure:"core"`
	Name       string
}

type AnotherConfig struct {
	Desc string
	Misc string
	Name string
}
```

- Read Configuration from file

```yaml
desc: desc
misc: config.yaml
name: FLUENT
nested:
  kevin: string
  smith: smith
```

- Configuration to Struct
- Configuration to Struct by Key:nested configuration


```shell
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
```

