package ctx

import (
	conf "github.com/nixys/nxs-go-conf"
)

type confOpts struct {
	LogFile  string  `conf:"logfile" conf_extraopts:"default=stdout"`
	LogLevel string  `conf:"loglevel" conf_extraopts:"default=info"`
	PidFile  string  `conf:"pidfile"`
	Src      srcConf `conf:"src" conf_extraopts:"required"`
	Dst      dstConf `conf:"dst" conf_extraopts:"required"`
}

type srcConf struct {
	MySQL mysqlConf `conf:"mysql" conf_extraopts:"required"`
	Redis redisConf `conf:"redis" conf_extraopts:"required"`
}

type dstConf struct {
	MySQL mysqlConf `conf:"mysql" conf_extraopts:"required"`
}

type mysqlConf struct {
	Host     string `conf:"host" conf_extraopts:"required"`
	Port     int    `conf:"port" conf_extraopts:"required"`
	DB       string `conf:"db" conf_extraopts:"required"`
	User     string `conf:"user" conf_extraopts:"required"`
	Password string `conf:"password" conf_extraopts:"required"`
}

type redisConf struct {
	Host string `conf:"host" conf_extraopts:"required"`
	Port int    `conf:"port" conf_extraopts:"required"`
}

func confRead(confPath string) (confOpts, error) {

	var c confOpts

	err := conf.Load(&c, conf.Settings{
		ConfPath:    confPath,
		ConfType:    conf.ConfigTypeYAML,
		UnknownDeny: true,
	})
	if err != nil {
		return c, err
	}

	return c, err
}
