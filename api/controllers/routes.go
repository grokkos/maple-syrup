package controllers

import (
	"github.com/grokkos/maple-syrup/api/middlewares"
)

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")

	s.Router.HandleFunc("/roundups", middlewares.SetMiddlewareJSON(s.CreateRoundup)).Methods("POST")
	s.Router.HandleFunc("/roundups", middlewares.SetMiddlewareJSON(s.GetRoundups)).Methods("GET")

	s.Router.HandleFunc("/batches", middlewares.SetMiddlewareJSON(s.GetBatchesByUser)).Methods("GET") //user id as param /batches?userid=
	s.Router.HandleFunc("/batches", middlewares.SetMiddlewareJSON(s.GetAllBatches)).Methods("GET")

	s.Router.HandleFunc("/transactions", middlewares.SetMiddlewareJSON(s.GetTransactions)).Methods("GET")
}
