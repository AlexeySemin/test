package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/AlexeySemin/test/golang-service/controllers"
	"github.com/AlexeySemin/test/golang-service/db/postgres"
	"github.com/AlexeySemin/test/golang-service/models"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"github.com/urfave/negroni"
)

type server struct {
	port    int
	router  *mux.Router
	negroni *negroni.Negroni
	db      *gorm.DB
}

func (s *server) Start() {
	s.registerRoutes()
	s.negroni.UseHandler(s.router)

	srv := &http.Server{
		Addr:         fmt.Sprintf("localhost:%v", s.port),
		WriteTimeout: time.Second * 45,
		ReadTimeout:  time.Second * 45,
		IdleTimeout:  time.Second * 60,
		Handler:      s.negroni,
	}

	fmt.Println("Server is listening...")

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func (s *server) registerRoutes() {
	cc := controllers.NewCommonController(s.db)
	dbsac := controllers.NewDBSAController(s.db)
	ssac := controllers.NewSSAController(s.db)

	s.router.HandleFunc("/news", cc.FillNewsDB).Methods(http.MethodPost)
	s.router.HandleFunc("/news", cc.ClearDB).Methods(http.MethodDelete)

	// DB side aggregation
	s.router.HandleFunc("/dbsa/news/min-max-avg-rating", dbsac.GetMinMaxAvgRating).Methods(http.MethodGet)
	s.router.HandleFunc("/dbsa/news/per-month-json-data", dbsac.GetPerMonthJSONData).Methods(http.MethodGet)

	// Server side aggregation
	s.router.HandleFunc("/ssa/news/min-max-avg-rating", ssac.GetMinMaxAvgRating).Methods(http.MethodGet)
	s.router.HandleFunc("/ssa/news/per-month-json-data", ssac.GetPerMonthJSONData).Methods(http.MethodGet)
}

// NewServer init and return new server
func NewServer(port int) (*server, error) {
	syncedModels := getModels()
	db, err := postgres.NewDB(syncedModels)
	if err != nil {
		return nil, err
	}

	return &server{
		port:    port,
		router:  mux.NewRouter(),
		negroni: negroni.Classic(),
		db:      db,
	}, nil
}

func getModels() []interface{} {
	return []interface{}{
		&models.News{},
	}
}
