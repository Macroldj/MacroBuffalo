package models

import (
	"MacroBuffalo/conf"
	"MacroBuffalo/models/item"
	"MacroBuffalo/models/user"
	"fmt"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

func MysqlConfig() *conf.Yaml {
	mysqlConf := new(conf.Yaml)
	yamlFile, err := ioutil.ReadFile("conf/resources.yaml")
	if err != nil {
		log.Printf("resources file get err #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, mysqlConf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
	return mysqlConf
}

func init() {
	config := MysqlConfig()
	fmt.Println("[init database]......")
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterModel(new(user.User))
	orm.RegisterModel(new(user.Post))
	orm.RegisterModel(new(item.Item))
	mysqlConfig := fmt.Sprintf("%s:%s@%s(%s:%d)/%s", config.Mysql.User, config.Mysql.Password, config.Mysql.Server,
		config.Mysql.Host, config.Mysql.Port, config.Mysql.Name)
	orm.RegisterDataBase("default", "mysql", mysqlConfig)
	orm.SetMaxIdleConns("default", config.Mysql.MaxIdleConns)
	orm.SetMaxOpenConns("default", config.Mysql.MaxOpenConns)
	fmt.Println("[end init database]......")
}
