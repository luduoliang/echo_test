package config
import (
	"encoding/json"
	"io/ioutil"
	"os"
)
type DatabaseConf struct {
	Appname string `json:"appname"`
	Httpport string `json:"httpport"`
	SessionCookieName string `json:"session_cookie_name"`
	SessionPath string `json:"session_path"`
	AesKey string `json:"aes_key"`
	AesIv string `json:"aes_iv"`
	RedisAddr string `json:"redis_addr"`
	RedisPassword string `json:"redis_password"`
	RedisDb int `json:"redis_db"`
	DatabaseDriver string `json:"database_driver"`
	DatabaseConnectString string `json:"database_connect_string"`
}
var (
	Config = LoadDatabaseConf()
)
//读取配置文件，并解析json , 有错误抛异常，中止，不允许带病上岗
func LoadDatabaseConf() *DatabaseConf {
	dcobj := DatabaseConf{}
	file, err := os.Open("config/config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	fd, err := ioutil.ReadAll(file)
	conf := string(fd)
	//json解析到结构体里面
	err = json.Unmarshal([]byte(conf), &dcobj)
	if err != nil {
		panic(err)
	}
	return &dcobj
}
