package server

import (
	"errors"
	"net/http"

	conf "clown-id/internal/config"
	"clown-id/internal/handlers"
	"clown-id/internal/store"
	"clown-id/internal/store/sqlstore"

	swaggerDocs "clown-id/docs"

	httpSwagger "github.com/swaggo/http-swagger"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	config *conf.Config
	Logger *logrus.Logger
	router *mux.Router
	store  store.Store
}

func New(config *conf.Config) *Server {
	return &Server{
		config: config,
		Logger: logrus.New(),
		router: mux.NewRouter(),
		// store
	}
}

// Function for starting HTTP. Don't start if using struct in tests
func (s *Server) Start() error {
	if err := s.configureLogger(); err != nil {
		return errors.New("Failed to configure logger: " + err.Error())
	}

	s.Logger.Info("Configuring database...")
	if err := s.configureStore(); err != nil {
		return errors.New("Failed to configure database: " + err.Error())
	}

	s.Logger.Info("Configuring routers...")
	s.configureRouter()

	if s.config.SwaggerEnabled {
		s.Logger.Info("Configuring OpenAPI docs...")
		s.configureDocs()
	}

	s.Logger.Info("Server started on: ", s.config.ApiPrefix+s.config.BindAddr)
	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureRouter() {
	s.router = s.router.PathPrefix(s.config.ApiPrefix).Subrouter().StrictSlash(true)
	handlers.RegisterHandlers(s.router, s.store, s.config.Secret)
}

func (s *Server) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}
	s.Logger.SetFormatter(&logrus.TextFormatter{
		DisableColors: false,
		FullTimestamp: true,
	})
	s.Logger.SetLevel(level)
	return nil
}

func (s *Server) configureStore() error {
	store, err := sqlstore.New(s.config.DbConnStr)
	if err != nil {
		return err
	}
	s.store = store
	return nil
}

func (s *Server) configureDocs() {
	s.router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)
	swaggerDocs.SwaggerInfo.BasePath = s.config.ApiPrefix
}
