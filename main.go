package main

import (
	"migration/controllers"
	"migration/initializer"
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

func init() {
	initializer.LoadEnvVariables()
	initializer.ConnectToDB()
}
func MwCors() gin.HandlerFunc {
	return cors.New(cors.Options{
		AllowOriginFunc: func(origin string) bool { return true },
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodConnect,
			http.MethodOptions,
			http.MethodTrace,
		},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           604800,
	})
}
func main() {
	r := gin.Default()
	// Загрузка ключей
	certFile := "/etc/letsencrypt/live/quantico.kz/fullchain.pem"
	keyFile := "/etc/letsencrypt/live/quantico.kz/privkey.pem"

	// Создание сервера HTTP
	server := &http.Server{
		Addr:    ":3300",
		Handler: r,
	}

	r.Use(MwCors())
	r.Static("/images", "./images")
	//Create
	r.POST("/posts", controllers.PostsCreate)
	//Update
	r.PUT("/posts/:id", controllers.PostsUpdate)
	//GET all
	r.GET("/posts", controllers.PostsIndex)
	//GET  id
	r.GET("/posts/:id", controllers.PostsShow)
	//Delete
	r.DELETE("/posts/:id", controllers.PostsDelete)

	// Запуск сервера с использованием SSL
	err := server.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		panic(err)
	}
}
