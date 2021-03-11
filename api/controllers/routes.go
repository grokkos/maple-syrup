package controllers

import "github.com/grokkos/maple-syrup/api/responses"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/user", responses.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/roundup", responses.SetMiddlewareJSON(s.CreateRoundup)).Methods("POST")

	s.Router.HandleFunc("/users", responses.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/roundups", responses.SetMiddlewareJSON(s.GetRoundups)).Methods("GET")
	s.Router.HandleFunc("/batches", responses.SetMiddlewareJSON(s.GetBatches)).Methods("GET")
	s.Router.HandleFunc("/transactions", responses.SetMiddlewareJSON(s.GetTransactions)).Methods("GET")

}
