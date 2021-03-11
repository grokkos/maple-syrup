package controllers

import "github.com/grokkos/maple-syrup/api/responses"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/roundup", responses.SetMiddlewareJSON(s.CreateRoundup)).Methods("POST")
	s.Router.HandleFunc("/roundups", responses.SetMiddlewareJSON(s.GetRoundups)).Methods("GET")
	s.Router.HandleFunc("/batches", responses.SetMiddlewareJSON(s.GetBatches)).Methods("GET")

}
