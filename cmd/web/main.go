package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/lucasvavon/slipx-api/internal/adapters/repositories/mysql"
	"github.com/lucasvavon/slipx-api/internal/core/services"
)

type App struct {
	userService  *services.UserService
	videoService *services.VideoService
}

func main() {
	// Initialisation de la DB et des services
	db := mysql.InitDB()
	// userStore := mysql.NewUserGORMRepository(db)
	videoStore := mysql.NewVideoGORMRepository(db)

	// Initialisation des services avec les repositories
	// us := services.NewUserService(userStore)
	vs := services.NewVideoService(videoStore)

	// Récupérer toutes les vidéos
	videos, err := vs.GetVideos()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des vidéos:", err)
		return
	}

	// Afficher les vidéos récupérées
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		tmpl.Execute(w, videos)
	})
	http.ListenAndServe(":8080", nil)
}
