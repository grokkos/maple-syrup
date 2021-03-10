package controllers

import "github.com/grokkos/maple-syrup/api/responses"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/roundup", responses.SetMiddlewareJSON(s.CreateRoundup)).Methods("POST")

}
