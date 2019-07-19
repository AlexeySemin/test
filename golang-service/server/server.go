package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

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
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      s.negroni,
	}

	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}

func (s *server) registerRoutes() {
	s.router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome!")
	})
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
