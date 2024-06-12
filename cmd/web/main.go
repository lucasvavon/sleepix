package main

import (
	"fmt"
	"github.com/lucasvavon/slipx-api/internal/core/domain"
	"net/http"
	"text/template"

	"github.com/lucasvavon/slipx-api/internal/adapters/repositories/mysql"
	"github.com/lucasvavon/slipx-api/internal/core/services"
)

type App struct {
	userService  *services.UserService
	videoService *services.VideoService
}

type Data struct {
	Videos []domain.Video
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
		fmt.Println(err)
		return
	}

	// Afficher les vidéos récupérées
	tmpl := template.Must(template.ParseFiles("./templates/index.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Data{Videos: videos}

		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8081", nil)
}
