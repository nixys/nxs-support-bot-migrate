package ctx

import (
	"fmt"

	"github.com/nixys/nxs-support-bot-migrate/ds/mysql"
	"github.com/nixys/nxs-support-bot-migrate/ds/redis"

	appctx "github.com/nixys/nxs-go-appctx/v2"
)

// Ctx defines application custom context
type Ctx struct {
	Conf confOpts
	Src  SrcCtx
	Dst  DstCtx
}

type SrcCtx struct {
	MySQL mysql.MySQL
	Redis redis.Redis
}

type DstCtx struct {
	MySQL mysql.MySQL
}

// Init initiates application custom context
func (c *Ctx) Init(opts appctx.CustomContextFuncOpts) (appctx.CfgData, error) {

	//a := opts.Args.(*Args)

	// Read config file
	conf, err := confRead(opts.Config)
	if err != nil {
		return appctx.CfgData{}, err
	}

	// Set application context
	c.Conf = conf

	redisHost := fmt.Sprintf("%s:%d", c.Conf.Src.Redis.Host, c.Conf.Src.Redis.Port)

	// Connect to source MySQL
	c.Src.MySQL, err = mysql.Connect(mysql.Settings{
		Host:     c.Conf.Src.MySQL.Host,
		Port:     c.Conf.Src.MySQL.Port,
		Database: c.Conf.Src.MySQL.DB,
		User:     c.Conf.Src.MySQL.User,
		Password: c.Conf.Src.MySQL.Password,
	})
	if err != nil {
		return appctx.CfgData{}, err
	}

	// Connect to source Redis
	c.Src.Redis, err = redis.Connect(redisHost)
	if err != nil {
		return appctx.CfgData{}, err
	}

	// Connect to destination MySQL
	c.Dst.MySQL, err = mysql.Connect(mysql.Settings{
		Host:     c.Conf.Dst.MySQL.Host,
		Port:     c.Conf.Dst.MySQL.Port,
		Database: c.Conf.Dst.MySQL.DB,
		User:     c.Conf.Dst.MySQL.User,
		Password: c.Conf.Dst.MySQL.Password,
	})
	if err != nil {
		return appctx.CfgData{}, err
	}

	return appctx.CfgData{
		LogFile:  c.Conf.LogFile,
		LogLevel: c.Conf.LogLevel,
		PidFile:  c.Conf.PidFile,
	}, nil
}

// Reload reloads application custom context
func (c *Ctx) Reload(opts appctx.CustomContextFuncOpts) (appctx.CfgData, error) {

	opts.Log.Debug("reloading context")

	c.Src.MySQL.Close()
	c.Src.Redis.Close()

	c.Dst.MySQL.Close()

	return c.Init(opts)
}

// Free frees application custom context
func (c *Ctx) Free(opts appctx.CustomContextFuncOpts) int {

	opts.Log.Debug("freeing context")

	c.Src.MySQL.Close()
	c.Src.Redis.Close()

	c.Dst.MySQL.Close()

	return 0
}
