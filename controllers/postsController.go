package controllers

import (
	"fmt"
	"migration/initializer"
	"migration/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	////////////C///////////////////////////////////////

	image := c.Request.FormValue("imageEmty")
	fmt.Println(image)
	if image == "" {
		// Получаем файл с изображением
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to retrieve image from form data",
			})
			return
		}

		// Создаем временный файл для сохранения изображения
		tempFile, err := os.CreateTemp("images", "upload-*.jpg")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create temporary file for image",
			})
			return
		}
		defer tempFile.Close()

		// Сохраняем файл с изображением на диск
		err = c.SaveUploadedFile(file, tempFile.Name())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save image file",
			})
			return
		}
		//Получаем отсавльные поля и заполняем в структуру
		title := c.Request.FormValue("title")
		body := c.Request.FormValue("body")
		post := models.Post{
			Title:     title,
			Body:      body,
			ImagePath: tempFile.Name(),
		}
		fmt.Println(post.ImagePath)
		//Create POST
		result := initializer.DB.Create(&post) // pass pointer of data to Create
		if result.Error != nil {
			c.Status(400)
			return
		}

		//Возврат
		c.JSON(200, gin.H{
			"post": post,
		})

	} else {
		//Получаем отсавльные поля и заполняем в структуру
		title := c.Request.FormValue("title")
		body := c.Request.FormValue("body")
		post := models.Post{
			Title:     title,
			Body:      body,
			ImagePath: "",
		}
		fmt.Println(post.ImagePath)
		//Create POST
		result := initializer.DB.Create(&post) // pass pointer of data to Create
		if result.Error != nil {
			c.Status(400)
			return
		}

		//Возврат
		c.JSON(200, gin.H{
			"post": post,
		})
	}

}

// Получить все посты
func PostsIndex(c *gin.Context) {
	// Get all records
	var posts []models.Post
	initializer.DB.Find(&posts)
	// c.Header("Access-Control-Allow-Origin", "*")

	//respons
	c.JSON(200, gin.H{
		"post": posts,
	})
}

func PostsShow(c *gin.Context) {
	//get id off url
	id := c.Param("id")
	//get the posts
	var post models.Post
	initializer.DB.First(&post, &id)
	// c.Header("Access-Control-Allow-Origin", "*")

	//respons
	c.JSON(200, gin.H{
		"post": post,
	})

}

// Изменить пост
func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	image := c.Request.FormValue("imageEmty")
	if image == "" {
		// Получаем файл с изображением
		file, err := c.FormFile("image")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to retrieve image from form data",
			})
			return
		}

		// Создаем временный файл для сохранения изображения
		tempFile, err := os.CreateTemp("images", "upload-*.jpg")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create temporary file for image",
			})
			return
		}
		defer tempFile.Close()

		// Сохраняем файл с изображением на диск
		err = c.SaveUploadedFile(file, tempFile.Name())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save image file",
			})
			return
		}
		//Получаем отсавльные поля и заполняем в структуру
		title := c.Request.FormValue("title")
		body := c.Request.FormValue("body")

		//find the posts
		var post models.Post
		initializer.DB.First(&post, &id)
		//Update it
		initializer.DB.Model(&post).Updates(models.Post{
			Title:     title,
			Body:      body,
			ImagePath: tempFile.Name(),
		})

		result := initializer.DB.Create(&post) // pass pointer of data to Create
		if result.Error != nil {
			c.Status(400)
			return
		}

		//Возврат
		c.JSON(200, gin.H{
			"post": post,
		})

	} else {

		//Получаем отсавльные поля и заполняем в структуру
		title := c.Request.FormValue("title")
		body := c.Request.FormValue("body")

		//find the posts
		var post models.Post
		initializer.DB.First(&post, &id)
		//Update it
		initializer.DB.Model(&post).Updates(models.Post{
			Title:     title,
			Body:      body,
			ImagePath: "",
		})

		initializer.DB.Create(&post) // pass pointer of data to Create
		// if result.Error != nil {
		// 	c.Status(400)
		// 	return
		// }
		//Возврат
		c.JSON(200, gin.H{
			"post": post,
		})
	}
}

// Удаление Поста
func PostsDelete(c *gin.Context) {
	// Получить id из URL
	id := c.Param("id")

	// Получить информацию о посте из базы данных
	post := models.Post{}
	if err := initializer.DB.First(&post, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}
	//проверяем если картинки нет то не удаляем  картинку
	if post.ImagePath == "" {
	} else {
		// Удалить файл изображения
		if err := os.Remove(post.ImagePath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete image"})
			return
		}
	}

	// Удалить пост из базы данных
	if err := initializer.DB.Delete(&post, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete post"})
		return
	}

	// Отправить успешный JSON-ответ
	c.JSON(http.StatusOK, gin.H{"message": "Post deleted successfully"})
}
