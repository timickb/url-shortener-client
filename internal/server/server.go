package server

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	rice "github.com/GeertJohan/go.rice"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	port         int
	api          string
	router       *mux.Router
	frontendPath string
	logger       *logrus.Logger
}

func New(logger *logrus.Logger, port int, api string, frontendPath string) *Server {
	srv := &Server{
		port:         port,
		api:          api,
		frontendPath: frontendPath,
		logger:       logger,
		router:       mux.NewRouter(),
	}

	srv.router.HandleFunc("/{shortening:[a-zA-Z0-9_]+}", srv.queryHandler())

	srv.router.PathPrefix("/").
		Handler(http.FileServer(rice.MustFindBox(frontendPath).HTTPBox()))

	return srv
}

func (s *Server) Start() {
	s.logger.Infof("Starting server on port %d", s.port)
	http.ListenAndServe(":"+strconv.Itoa(s.port), s.router)
}

func (s *Server) queryHandler() func(http.ResponseWriter, *http.Request) {
	type response struct {
		Original string
	}

	return func(w http.ResponseWriter, r *http.Request) {
		shortening := mux.Vars(r)["shortening"]

		if shortening == "" {
			http.FileServer(rice.MustFindBox(s.frontendPath).HTTPBox())
			return
		}

		s.logger.Infof("Requested shortening %s", shortening)

		// request the API
		req := fmt.Sprintf(s.api+"restore?hash=%s", shortening)
		resp, err := http.Get(req)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Server error"))
			return
		}

		respData := &response{}
		bodyBytes, err := io.ReadAll(resp.Body)
		json.Unmarshal(bodyBytes, respData)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Server error"))
			return
		}

		http.Redirect(w, r, respData.Original, http.StatusPermanentRedirect)
	}
}
