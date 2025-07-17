package workflow

import (
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v5"
	"net/http"
	"workflow-code-test/api/services/workflow/store"
)

type ServiceHandler interface {
	HandleExecuteWorkflow(w http.ResponseWriter, r *http.Request)
	HandleGetWorkflow(w http.ResponseWriter, r *http.Request)
	LoadRoutes(parentRouter *mux.Router, isProduction bool)
}

type Service struct {
	DB    *pgx.Conn
	Store store.Store
}

func NewService(db *pgx.Conn) (ServiceHandler, error) {
	return &Service{
		DB:    db,
		Store: store.NewStore(db),
	}, nil
}

// jsonMiddleware sets the Content-Type header to application/json
func jsonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func (s *Service) LoadRoutes(parentRouter *mux.Router, isProduction bool) {
	router := parentRouter.PathPrefix("/workflows").Subrouter()
	router.StrictSlash(false)
	router.Use(jsonMiddleware)

	router.HandleFunc("/{id}", s.HandleGetWorkflow).Methods("GET")
	router.HandleFunc("/{id}/execute", s.HandleExecuteWorkflow).Methods("POST")

}
