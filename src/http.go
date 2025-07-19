package src

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/prithvitewatia/gocache/src/common"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type Server struct {
	cache *Cache
}

func NewServer(c *Cache) *Server {
	return &Server{cache: c}
}

func (server *Server) Start(port string) {
	mux := http.NewServeMux()

	mux.HandleFunc("/get", server.handleGet)
	mux.HandleFunc("/set", server.handleSet)
	mux.HandleFunc("/del", server.handleDel)
	mux.HandleFunc("/keys", server.handleGetAllKeys)
	mux.HandleFunc("/ttl", server.handleGetTll)
	mux.HandleFunc("/flushall", server.handleFlushAll)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}
	go func() {
		log.Printf("Starting server on port %s", port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	srv.Shutdown(ctx)
}

func (server *Server) handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "key is empty", http.StatusBadRequest)
	}
	value, found := server.cache.Get(key)
	if !found {
		http.NotFound(w, r)
		return
	}
	err := json.NewEncoder(w).Encode(map[string]interface{}{"value": value})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (server *Server) handleSet(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	var req common.CacheSetRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	server.cache.Set(req.Key, req.Value, req.TTL)
	w.WriteHeader(http.StatusCreated)
}

func (server *Server) handleDel(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "key is empty", http.StatusBadRequest)
	}
	server.cache.Delete(key)
	w.WriteHeader(http.StatusOK)
}

func (server *Server) handleGetAllKeys(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	keys := server.cache.Keys()
	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"keys": keys,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (server *Server) handleGetTll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	key := r.URL.Query().Get("key")
	if key == "" {
		http.Error(w, "key is empty", http.StatusBadRequest)
	}
	ttl, exists := server.cache.TTL(key)
	if !exists {
		ttl = 0
	}
	err := json.NewEncoder(w).Encode(map[string]int64{"ttl": ttl})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (server *Server) handleFlushAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
	server.cache.FlushAll()
	w.WriteHeader(http.StatusOK)
}
