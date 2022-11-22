package main

import (
	"bytes"
	"fmt"
	"github.com/apolloconfig/agollo/v4"
	"github.com/apolloconfig/agollo/v4/env/config"
	"github.com/spf13/viper"
	"path/filepath"
	"text/template"
)

var apolloConfig *config.AppConfig
var agolloFuncMap template.FuncMap = template.FuncMap{
	"key": func(key string) string {
		// 远程获取key对应的配置值
		return key
	},
}

func main() {
	apolloConfig = &config.AppConfig{
		AppID:          "testApplication_yang",
		Cluster:        "dev",
		IP:             "http://106.54.227.205:8080",
		NamespaceName:  "dubbo",
		IsBackupConfig: true,
		Secret:         "6ce3ff7e96a24335a9634fe9abca6d51",
	}
	client, _ := agollo.StartWithConfig(func() (*config.AppConfig, error) {
		return apolloConfig, nil
	})
	cache := client.GetConfigCache(apolloConfig.NamespaceName)
	val, _ := cache.Get("key")
	fmt.Println(val)
	fileBase := filepath.Base("./pkg/agollo/conf.toml")
	templ := template.New(fileBase)
	tpl, err := templ.Funcs(agolloFuncMap).ParseFiles("./pkg/agollo/conf.toml")
	if err != nil {
		panic(err)
	}
	var b bytes.Buffer
	err = tpl.Execute(&b, "")
	if err != nil {
		panic(err)
	}
	viper.SetConfigType("toml")
	err = viper.ReadConfig(&b)
	fmt.Println(viper.GetString("redis.addr"))
}
