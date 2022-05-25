package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"strings"
)

func Init(cfg string) error {
	// 初始化配置文件
	if err := initConfig(cfg); err != nil {
		return err
	}
	// 监控配置文件并热加载配置文件
	watchConfig()
	// 初始化日志包
	initLog()
	return nil
}

// 初始化配置文件
func initConfig(cfg string) error {
	//如果命令行通过flag的方式传来了配置文件路径，就通过第一种方式解析
	//比如viper.SetConfigFile("./config/config.yaml")
	if cfg != "" {
		viper.SetConfigFile(cfg)
	} else {
		//如果命令行通过flag的方式传来了配置文件路径，就通过第二种方式解析
		//设置配置文件名称(无扩展名)
		viper.AddConfigPath("conf")
		// 设置配置文件的名称，此处不包括配置文件的拓展名
		viper.SetConfigName("config")
	}
	// 设置配置文件格式为YAML格式
	viper.SetConfigType("yaml")
	// 读取匹配的环境变量
	viper.AutomaticEnv()

	// 读取环境变量的前缀为APISERVER，以下配置可以使程序读取环境变量
	viper.SetEnvPrefix("APISERVER")
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// 前边已经配置好一些参数了，现在找到并且读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}
	return nil
}

func watchConfig() {
	// 监控配置文件变化并热加载程序
	viper.WatchConfig()
	//配置文件发生变更之后会调用的回调函数,通过该函数的viper设置，可以使viper监控配置文件变更，如有变更则热更新程序
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Infof("配置文件已更改: %s", in.Name)
	})
}

func initLog() {
	passLagerConfig := log.PassLagerCfg{
		// 输出位置，有两个可选项 —— file 和 stdout:
		//   选择 file 会将日志记录到 logger_file 指定的日志文件中
		//   选择 stdout 会将日志输出到标准输出，当然也可以两者同时选择
		Writers: viper.GetString("log.writers"),
		// 配置日志级别(类型)
		LoggerLevel: viper.GetString("log.logger_level"),
		// 日志文件
		LoggerFile: viper.GetString("log.logger_file"),
		// 日志的输出格式，JSON或者plaintext，如果是true会输出非json格式
		LogFormatText: viper.GetBool("log.log_format_text"),
		// rotate 依据，可选的有 daily 和 size。如果选 daily 则根据天进行转存，如果是 size 则根据大小进行转存
		RollingPolicy: viper.GetString("log.rollingPolicy"),
		// 配合RollingPolicy：daily  转存时间
		LogRotateDate: viper.GetInt("log.log_rotate_date"),
		// 配合RollingPolicy：size  转存大小
		LogRotateSize: viper.GetInt("log.log_rotate_size"),
		// 当日志文件达到转存标准时，log 系统会将该日志文件进行压缩备份，这里指定了备份文件的最大个数
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}
	_ = log.InitWithConfig(&passLagerConfig)

}
