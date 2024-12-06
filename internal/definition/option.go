package definition

import (
	"loan/internal/api"
	"loan/internal/handler/public"
	"loan/internal/service"

	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func runRouter(e *echo.Echo, logger *zap.Logger) {
	loanService := service.NewLoanService()
	loanHandler := public.NewLoanHandler(loanService, logger)

	api.RegisterHandlers(e, loanHandler)
}

func NewOption() fx.Option {
	return fx.Options(
		fx.Provide(),
		fx.Invoke(runRouter),
	)
}
