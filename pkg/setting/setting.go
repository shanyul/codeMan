package setting

import (
	"github.com/go-ini/ini"
	"io/ioutil"
	"log"
	"time"
)

// App 应用基础信息
type App struct {
	AppName         string
	AppHost         string
	PageSize        int
	JwtSecret       string
	SignKey         string
	Environment     string
	ImageSavePath   string
	ImageMaxSize    int
	ImageAllowExt   []string
	RuntimeRootPath string
}

// Server 服务基础信息
type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// Log 日志配置
type Log struct {
	RootPath   string
	SavePath   string
	SaveName   string
	FileExt    string
	TimeFormat string
}

// Database 数据库配置
type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
	Timeout  time.Duration
}

// Redis 数据库配置
type Redis struct {
	Host        string
	Password    string
	MaxIdle     int
	MaxActive   int
	IdleTimeout time.Duration
}

// Storage 存储配置
type Tencent struct {
	CosKey    string
	CosSecret string
	AppId     string
	AppSecret string
}

// Rsa 公私钥配置
type Rsa struct {
	PublicKey  string
	PrivateKey string
}

var (
	cfg             *ini.File
	AppSetting      = &App{}
	ServerSetting   = &Server{}
	LogSetting      = &Log{}
	DatabaseSetting = &Database{}
	RedisSetting    = &Redis{}
	TencentSetting  = &Tencent{}
	RsaSetting      = &Rsa{}
)

// Setup 初始化配置
func Setup() {
	var err error
	cfg, err = ini.Load("config/app.ini")
	if err != nil {
		log.Fatalf("setting.Setup, fail to parse 'config/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("server", ServerSetting)
	mapTo("log", LogSetting)
	//mapTo("database", DatabaseSetting)
	//mapTo("redis", RedisSetting)
	//mapTo("tencent", TencentSetting)

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	//DatabaseSetting.Timeout = DatabaseSetting.Timeout * time.Second
	//RedisSetting.IdleTimeout = RedisSetting.IdleTimeout * time.Second

	//publicKey, privateKey := getRsaKey()
	//RsaSetting.PublicKey = publicKey
	//RsaSetting.PrivateKey = privateKey
}

// 参数赋值
func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo %s err: %v", section, err)
	}
}

func getRsaKey() (string, string) {
	publicKeyPath := "config/rsa/id_rsa_public.pem"
	privateKeyPath := "config/rsa/id_rsa_private.pem"

	publicKey, err := ioutil.ReadFile(publicKeyPath)

	if err != nil {
		log.Fatalf("read publicKey err: %v", err)
	}

	privateKey, err := ioutil.ReadFile(privateKeyPath)

	if err != nil {
		log.Fatalf("read privateKey err: %v", err)
	}

	return string(publicKey), string(privateKey)
}
