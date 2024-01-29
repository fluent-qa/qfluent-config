# GO-Lib: Configurations

## How to Use

- Read Configuration from file

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

