package utils

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"note/config"
)

/*
拓展类读取yaml文件
**/
type DatabaseConf config.PostgresYaml

func (c *DatabaseConf) getDbYamlConf() *DatabaseConf {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return c
}

func GetDbConf() DatabaseConf {
	var c DatabaseConf
	c.getDbYamlConf()
	return c
}

// 函数读取yaml文件
//func getDbYamlConf() config.PostgresYaml{
//	c :=config.PostgresYaml{}
//	yamlFile, err := ioutil.ReadFile("config.yaml")
//	if err != nil {
//		log.Printf("yamlFile.Get err   #%v ", err)
//	}
//
//	err = yaml.Unmarshal(yamlFile, &c)
//	if err != nil {
//		log.Fatalf("Unmarshal: %v", err)
//	}
//	return c
//}
//
