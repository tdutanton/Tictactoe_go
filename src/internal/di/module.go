package di

import (
	"github.com/tdutanton/Tictactoe_go/internal/datasource"
	"github.com/tdutanton/Tictactoe_go/internal/domain/service"
	"github.com/tdutanton/Tictactoe_go/internal/web"

	"go.uber.org/fx"
)

// Module - di модуль для управления зависимостями
var Module = fx.Module("app",
	// NewGameStore Инициализация хранилища (singleton)
	fx.Provide(
		datasource.NewGameStore,
	),
	// NewGameRepository Репозиторий
	fx.Provide(
		datasource.NewGameRepository,
	),
	// NewGameService Сервис
	fx.Provide(
		service.NewGameService,
	),
	// NewGameHandler Хендлеры
	fx.Provide(
		web.NewGameHandler,
	),
	// SetupRouter Роутер
	fx.Provide(
		web.SetupRouter,
	),
)
