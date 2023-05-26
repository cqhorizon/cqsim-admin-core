package config

import (
	"fmt"
	"github.com/cqhorizon/cqsim-admin-core/config"
	"github.com/cqhorizon/cqsim-admin-core/config/source"
	"log"
)

var (
	_cfg *Settings
)

// Settings 兼容原先的配置结构
type Settings struct {
	Settings  Config `yaml:"settings"`
	callbacks []func()
}

func (e *Settings) runCallback() {
	for i := range e.callbacks {
		e.callbacks[i]()
	}
}

func (e *Settings) OnChange() {
	e.init()
	log.Println("!!! config change and reload")
}

func (e *Settings) Init() {
	e.init()
	log.Println("!!! config init")
}

func (e *Settings) init() {
	e.Settings.Logger.Setup()
	e.Settings.multiDatabase()
	e.runCallback()
}

// Config 配置集合
type Config struct {
	Application *Application          `yaml:"application"`
	Logger      *Logger               `yaml:"logger"`
	Jwt         *Jwt                  `yaml:"jwt"`
	Resource    *Resource             `yaml:"mongo"`
	Database    *Database             `yaml:"database"`
	Databases   *map[string]*Database `yaml:"databases"`
	Cache       *Cache                `yaml:"cache"`
}

// 多db改造
func (e *Config) multiDatabase() {
	if len(*e.Databases) == 0 {
		*e.Databases = map[string]*Database{
			"*": e.Database,
		}

	}
}

// Setup 载入配置文件
func Setup(s source.Source,
	fs ...func()) {
	_cfg = &Settings{
		Settings: Config{
			Application: ApplicationConfig,
			Logger:      LoggerConfig,
			Jwt:         JwtConfig,
			Resource:    ResourceConfig,
			Database:    DatabaseConfig,
			Databases:   &DatabasesConfig,
			Cache:       CacheConfig,
		},
		callbacks: fs,
	}
	var err error
	_, err = config.NewConfig(
		config.WithSource(s),
		config.WithEntity(_cfg),
	)
	if err != nil {
		log.Fatal(fmt.Sprintf("New config object fail: %s", err.Error()))
	}
	_cfg.Init()
}
