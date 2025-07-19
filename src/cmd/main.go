package main

import (
	"github.com/tdutanton/Tictactoe_go/internal/di"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

// NewApp создает и конфигурирует новое FX-приложение с зависимостями
// Возвращает:
//   - *fx.App - экземпляр приложения с настроенным DI-графом и HTTP-сервером
//
// ------
//   - Автоматически включает модуль зависимостей (di.Module)
//   - Устанавливает Gin в режим Release (без debug-логов)
//   - Регистрирует конечную точку запуска HTTP-сервера на порту 8080
//   - Для запуска приложения требуется вызвать NewApp().Run()
//
// Можно отправлять запросы через extension для VSCode - REST Client
// Пример запросов в internal/app/example_requests.http
// Через REST Client использовать путем выделения запроса и нажатия
// Ctrl+Alt+R или в контекстном меню через ПКМ
func NewApp() *fx.App {
	gin.SetMode(gin.ReleaseMode)
	return fx.New(di.Module, fx.Invoke(func(router *gin.Engine) { router.Run(":8080") }))
}

func main() {
	NewApp().Run()
}
