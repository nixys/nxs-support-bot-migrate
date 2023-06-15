package migrate

import (
	"context"

	"github.com/nixys/nxs-support-bot-migrate/ctx"
	"github.com/nixys/nxs-support-bot-migrate/modules/dsmigrate"

	appctx "github.com/nixys/nxs-go-appctx/v2"
)

// Runtime executes the routine
func Runtime(cr context.Context, appCtx *appctx.AppContext, crc chan interface{}) {

	cc := appCtx.CustomCtx().(*ctx.Ctx)

	de := dsmigrate.Init(
		dsmigrate.Settings{
			Src: dsmigrate.SrcSettings{
				MySQL: cc.Src.MySQL,
				Redis: cc.Src.Redis,
			},
			Dst: dsmigrate.DstSettings{
				MySQL: cc.Dst.MySQL,
			},
		},
	)

	if err := de.Migrate(); err != nil {
		appCtx.Log().Errorf(err.Error())
		appCtx.RoutineDoneSend(appctx.ExitStatusFailure)
		return
	}

	appCtx.RoutineDoneSend(appctx.ExitStatusSuccess)
}
