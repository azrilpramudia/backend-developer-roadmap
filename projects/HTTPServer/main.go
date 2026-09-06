package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/azrilpramudia/backend-developer-roadmap/internal/users"
)

type UserData struct {
	FirstName string
	LastName string
	Email string
}

type server struct	{
	userManager *users.Manager
}

func main() {
	manager := users.NewManager()

	s := server{
		userManager: manager,
	}
	mux := http.NewServeMux()

	httpServer := &http.Server{
		Addr: ":3000",
		Handler: mux,
	}


	mux.HandleFunc("GET /{$}", handleRoot)
	mux.HandleFunc("GET /goodbye", handleGoodbye)
	mux.HandleFunc("GET /hello/", handleHelloParameterized)
	mux.HandleFunc("GET /responses/{user}/hello/", handleUserResponsesHello)
	mux.HandleFunc("GET /user/hello/", s.handleHelloHeader)
	mux.HandleFunc("POST /json/", handleJSON)
	mux.HandleFunc("POST /add-user", s.addUser)
	mux.HandleFunc("GET /get-user", s.getUser)

	go func() {
		slog.Info("starting server...")
		err := httpServer.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("HTTP server error", "err", err)
			os.Exit(1)
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		sigChan := make(chan os.Signal, 1)
		signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

		<-sigChan
		slog.Info("shutting down server")

		shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer shutdownCancel()

		err := httpServer.Shutdown(shutdownCtx)
		if err != nil {
			slog.Error("error shutting down HTTP server", "err", err)
		}

		defer manager.Shutdown()
	}()

	wg.Wait()
	slog.Info("server shutdown complete")
}

func (s *server) addUser(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, fmt.Sprintf("unsupported Content-Type header: %q", contentType), http.StatusUnsupportedMediaType)
		return
	}
	requestBody := http.MaxBytesReader(w, r.Body, 1048576)

	decoder := json.NewDecoder(requestBody)
	decoder.DisallowUnknownFields()

	var u UserData

	err := decoder.Decode(&u)
	if err != nil {
		slog.Error("error decoding addUsers request body", "err", err)
		http.Error(w, "bad response body", http.StatusBadRequest)
		return
	}

	err = s.userManager.AddUser(u.FirstName, u.LastName, u.Email)
	if err != nil {
		http.Error(w, fmt.Sprintf("error adding user: %v\n", err), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (s *server) getUser(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, fmt.Sprintf("unsupported Content-Type header: %q", contentType), http.StatusUnsupportedMediaType)
		return
	}
	requestBody := http.MaxBytesReader(w, r.Body, 1048576)

	decoder := json.NewDecoder(requestBody)
	decoder.DisallowUnknownFields()

	var u UserData

	err := decoder.Decode(&u)
	if err != nil {
		slog.Error("error decoding getUsers request body", "err", err)
		http.Error(w, "bad response body", http.StatusBadRequest)
		return
	}

	user, err := s.userManager.GetUserByName(u.FirstName, u.LastName)
	if err != nil {
		if errors.Is(err, users.ErrNoResultFound){
			http.Error(w, "no users found", http.StatusNotFound)
		} else {
			slog.Error("error retrieving user", "err", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	converted := convertUserToUserData(user)

	marshalled, err := json.Marshal(converted)
	if err != nil {
		slog.Error("error marshalling getUser response", "err", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(marshalled)
	if err != nil {
		slog.Error("error writing GetUser response body", "err", err)
	}
}

func (s *server) handleHelloHeader(w http.ResponseWriter, r *http.Request) {
	firstName := r.Header.Get("userFirst")
	if firstName == "" {
		http.Error(w, "invalid first name provided", http.StatusBadRequest)
		return
	}

	lastName := r.Header.Get("userLast")

	if lastName == "" {
		http.Error(w, "invalid last name provided", http.StatusBadRequest)
		return
	}

	user, err := s.userManager.GetUserByName(firstName, lastName)
	if err != nil {
		if errors.Is(err, users.ErrNoResultFound){
			http.Error(w, "no users found", http.StatusNotFound)
		} else {
			slog.Error("error retrieving user", "err", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	result := fmt.Sprintf("Hello, %s %s! Your email is: %s\n", user.FirstName, user.LastName, user.Email.Address)

	_, err = w.Write([]byte(result))
	if err != nil {
		slog.Error("error writing response body", "err", err)
	}
}

func handleRoot(w http.ResponseWriter, _ *http.Request) {
	_, err := w.Write([]byte("Welcome to our Server\n"))
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}
}

func handleGoodbye(w http.ResponseWriter, _ *http.Request) {
	wc, err := w.Write([]byte("Goodbye\n"))
	if err != nil {
		slog.Error("error writing response", "err", err)
		return
	}

	fmt.Printf("%d bytes written\n", wc)
}

func handleHelloParameterized(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	userList := params["user"]

	username := "User"
	if len(userList) > 0 {
		username = userList[0]
	}

	handleHello(w, username)
}

func handleUserResponsesHello(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("user")

	handleHello(w, username)
}

func handleJSON(w http.ResponseWriter, r *http.Request) {
	byteData, err := io.ReadAll(r.Body)
	if err != nil {
		slog.Error("error reading request body", "err", err)
		http.Error(w, "bad request body", http.StatusBadRequest)
		return
	}

	if len(byteData) < 1 {
		slog.Warn("request body is empty")
		http.Error(w, "bad request body", http.StatusBadRequest)
		return
	}

	var reqData UserData
	err = json.Unmarshal(byteData, &reqData)
	if err != nil {
		slog.Error("error unmarshalling request body", "err", err)
		http.Error(w, "error parsing request JSON", http.StatusBadRequest)
		return
	}

	if reqData.FirstName == "" {
		http.Error(w, "invalid username provided", http.StatusBadRequest)
		return
	}

	handleHello(w, reqData.FirstName)
}

func handleHello(w http.ResponseWriter, username string) {
	var output bytes.Buffer
	output.WriteString("Hello, ")
	output.WriteString(username)
	output.WriteString("!\n")

	_, err := w.Write(output.Bytes())
	if err != nil {
		slog.Error("error writing response body", "err", err)
		return
	}
}

func convertUserToUserData(u *users.User) *UserData {
	converted := UserData{
		FirstName: u.FirstName,
		LastName: u.LastName,
		Email: u.Email.Address,
	}
	
	return &converted
}