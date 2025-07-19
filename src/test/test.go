package test

/* func main() {
	game := model.NewGame()
	game.Board[1][1] = enum.CROSS
	clirender.RenderBoard(game.Board)
	gs := datasource.NewGameStore()
	r := datasource.NewGameRepository(gs)
	s := service.NewGameService(r)
	game, _ = s.NextMove(game)
	fmt.Println()
	clirender.RenderBoard(game.Board)
	game.Board[0][2] = enum.CROSS
	clirender.RenderBoard(game.Board)
	game, _ = s.NextMove(game)
	fmt.Println()
	clirender.RenderBoard(game.Board)
} */

/* func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "HELLO"})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Run() // listen and serve on 0.0.0.0:8080
} */

/* func main() {
	// Инициализация хранилища
	gameStore := datasource.NewGameStore()
	gameRepo := datasource.NewGameRepository(gameStore)
	gameService := service.NewGameService(gameRepo)

	// Настройка маршрутов
	router := web.SetupRouter(gameService)

	// Запуск сервера
	router.Run(":8080")
} */

/* func main() {
	gin.SetMode(gin.ReleaseMode)
	fx.New(
		di.Module,
		fx.Invoke(func(router *gin.Engine) {
			router.Run(":8080")
		}),
	).Run()
} */
