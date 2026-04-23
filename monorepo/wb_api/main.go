package main

import (
	"fmt"
	"log"
	"net/http"
	commuRepo "v1/community/repository"
	commuServ "v1/community/service"

	postRepo "v1/post/repository"
	postServ "v1/post/service"
	"v1/users/repository"
	"v1/users/service"
	util "v1/util/cors"
	dbconfig "v1/util/db_config"
	config "v1/util/route_config"
	"v1/v1/handlers"
	"v1/wb_router/routes"

	"github.com/gorilla/mux"
)

// USED TO GENERATE THE KEY - very simple
// func init() {
// 	key := make([]byte, 64)

// 	if _, err := rand.Read(key); err != nil {
// 		log.Fatal(err)
// 	}

// 	stringBase64 := base64.StdEncoding.EncodeToString(key)
// 	fmt.Println(stringBase64)
// }

func main() {
	config.Load()

	db, err := dbconfig.Connect()
	if err != nil {
		log.Fatal("DB connection error:", err)
	}
	defer db.Close()

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	postRepository := postRepo.NewPostRepository(db)
	postService := postServ.NewPostService(postRepository)
	postHandler := handlers.NewPostHandler(postService)

	communityRepo := commuRepo.NewCommunityRepository(db)
	communityService := commuServ.NewCommunityService(communityRepo)
	communityHandler := handlers.NewCommunityHandler(communityService)

	r := mux.NewRouter()
	r = routes.Config(r, userHandler, postHandler, communityHandler)
	fmt.Println("Server has started")

	handler := util.EnableCORS(r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), handler))
}
