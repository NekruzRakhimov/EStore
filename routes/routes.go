package routes

import (
	"EStore/pkg/controller"
	"EStore/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Настройка router'a
func RunAllRoutes() {
	// Создание роутера
	r := gin.Default()

	// Исползование CORS
	r.Use(controller.CORSMiddleware())

	// Установка Logger-а
	utils.SetLogger()

	// Форматирование логов
	utils.FormatLogs(r)

	// Статус код 500, при любых panic()
	r.Use(gin.Recovery())

	// Запуск роутов
	initAllRoutes(r)

	// Запуск сервера
	_ = r.Run(utils.AppSettings.AppParams.PortRun)
}

// Инициализация роутов
func initAllRoutes(router *gin.Engine) {

	router.GET("/ping", PingPong) // Проверка связи

	/****** Открытые роуты*********************************************************************************************/
	router.GET("/goods", controller.GetAllGoods) // Получение списка всех продуктов

	auth := router.Group("/")
	auth.POST("/sign-up", controller.SignUp) // Регистрация
	auth.POST("/sign-in", controller.SignIn) // Аутентификация

	/****** Закрытые роуты*********************************************************************************************/
	api := router.Group("/api")

	user := api.Group("/user", controller.UserIdentity)
	user.GET("/cart", controller.GetUserCart)                    // для просмотра корзины
	user.POST("/cart", controller.AddToCart)                     // для добавления товара в корзину
	user.DELETE("/cart/:good_id", controller.DeleteGoodFromCart) // удаление товара с корзины

	manager := api.Group("/manager", controller.ManagerIdentity)
	manager.GET("/carts", controller.GetAllCarts)                        // получение списка корзин, всех ползователей
	manager.POST("/good", controller.CreateGood)                         // создание нового продукта
	manager.PUT("/good/:type_of_operation", controller.СrementGoodCount) // увеличение или уменьшение количаства товара
	manager.DELETE("/good/:id", controller.DeleteGood)                   // удаление товара
}

// PingPong Проверка связи
func PingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}
