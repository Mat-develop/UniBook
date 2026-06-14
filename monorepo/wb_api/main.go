package main

import (
	"fmt"
	"log"
	"net/http"
	commentRepo "v1/comment/repository"
	commentServ "v1/comment/service"
	commuRepo "v1/community/repository"
	commuServ "v1/community/service"
	postRepo "v1/post/repository"
	postServ "v1/post/service"
	profileRepo "v1/profile/repository"
	profileServ "v1/profile/service"
	tagRepo "v1/tag/repository"
	tagServ "v1/tag/service"
	"v1/users/repository"
	"v1/users/service"
	util "v1/util/cors"
	dbconfig "v1/util/db_config"
	config "v1/util/route_config"
	"v1/v1/handlers"
	"v1/wb_router/routes"

	"github.com/gorilla/mux"
)

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

	tagRepository := tagRepo.NewTagRepository(db)
	tagService := tagServ.NewTagService(tagRepository)
	tagHandler := handlers.NewTagHandler(tagService)

	commentRepository := commentRepo.NewCommentRepository(db)
	commentService := commentServ.NewCommentService(commentRepository)
	commentHandler := handlers.NewCommentHandler(commentService)

	searchHandler := handlers.NewSearchHandler(communityService, postService)

	profileRepository := profileRepo.NewProfileRepository(db)
	profileService := profileServ.NewProfileService(profileRepository)
	profileHandler := handlers.NewProfileHandler(profileService)

	r := mux.NewRouter()
	r = routes.Config(r, userHandler, postHandler, communityHandler, tagHandler, commentHandler, searchHandler, profileHandler)
	fmt.Println("Server has started")

	handler := util.EnableCORS(r)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), handler))
}
