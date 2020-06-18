package main

import (
	"github.com/profzone/eden-framework/pkg/application"
	"github.com/profzone/eden-framework/pkg/context"
	"longhorn/proxy/internal/global"
	"longhorn/proxy/internal/routers"
	"longhorn/proxy/internal/storage"
	"longhorn/proxy/pkg"
)

func main() {
	app := application.NewApplication(runner, &global.Config)
	go app.Start()
	app.WaitStop(func(ctx *context.WaitStopContext) error {
		ctx.Cancel()
		return nil
	})
}

func runner(app *application.Application) error {
	// init database
	pkg.Generator = pkg.NewSnowflake(global.Config.SnowflakeConfig)
	storage.Database.Init(global.Config.DBConfig, app.Context())

	// start administrator server
	go global.Config.GRPCServer.Serve(routers.RootRouter)
	go global.Config.HTTPServer.Serve(routers.RootRouter)
	return nil
}
