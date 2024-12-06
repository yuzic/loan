package app

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	appconfig "loan/internal/bootstrap/config"
	"loan/internal/definition"
	utilscfg "loan/internal/utils/config"

	httpmiddleware "loan/internal/bootstrap/http/middleware"

	"loan/internal/utils/logger"

	"github.com/brpaz/echozap"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/jmoiron/sqlx"
)

const shutdownTimeout = 10 * time.Second

type App struct {
	Logger *zap.Logger
	Config *appconfig.AppConfig
	DB     *sqlx.DB
}

func NewApp(configPath string) (*App, error) {
	utilscfg.InitViperByEnv(configPath)

	cfg := appconfig.NewAppConfig()
	log := logger.NewLogger()

	var (
		db *sqlx.DB
	)

	return &App{
		Logger: log,
		Config: cfg,
		DB:     db,
	}, nil
}

func (app *App) Start() error {
	fxApp := fx.New(
		fx.Provide(func() *appconfig.AppConfig { return app.Config }),
		fx.Provide(func() *zap.Logger { return app.Logger }),
		fx.Provide(func() *sqlx.DB { return app.DB }),
		fx.Provide(func() *echo.Echo {
			e := echo.New()
			e.Use(echozap.ZapLogger(app.Logger))
			e.Use(httpmiddleware.SetHeaders)
			e.Use(middleware.RequestID())
			e.Use(middleware.Recover())

			return e
		}),
		definition.NewOption(),
		fx.Invoke(app.startHTTPServer),
	)

	ctx := context.Background()
	if err := fxApp.Start(ctx); err != nil {
		return err
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	if err := fxApp.Stop(ctx); err != nil {
		return err
	}

	return nil
}

func (app *App) startHTTPServer(lifecycle fx.Lifecycle, e *echo.Echo) {
	lifecycle.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				err := e.Start(fmt.Sprintf(":%d", app.Config.HTTP.Port))
				if err != nil && errors.Is(err, http.ErrServerClosed) {
					e.Logger.Fatal("server shutdown", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			ctx, cancel := context.WithTimeout(ctx, shutdownTimeout)
			defer cancel()

			e.Logger.Info("shutting down server...")

			return e.Shutdown(ctx)
		},
	})
}
