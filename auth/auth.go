package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"os/exec"
	"os/user"
	"path"

	"golang.org/x/oauth2"
)

type memoConfig struct {
	ClientID     string `json:"clientID"`
	ClientSecret string `json:"clientSecret"`
}

func TokenSourceFromConfig(ctx context.Context) (oauth2.TokenSource, error) {
	config, err := loadConfig()
	if err != nil {
		return nil, err
	}

	ch := make(chan string)
	state := genState()

	shutdown, port, err := startAuthServer(ch, state)
	if err != nil {
		return nil, err
	}

	oauthConfig := oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		RedirectURL:  fmt.Sprintf("http://localhost:%v", port),
		Scopes:       []string{"https://www.googleapis.com/auth/devstorage.read_write"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://accounts.google.com/o/oauth2/v2/auth",
			TokenURL: "https://www.googleapis.com/oauth2/v4/token",
		},
	}

	authCodeURL := oauthConfig.AuthCodeURL(state)

	fmt.Println(authCodeURL)
	fmt.Println("Open the URL in your browser if it fails to open automatically")
	exec.Command("open", authCodeURL).Start()

	code := <-ch
	shutdown(ctx)

	token, err := oauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("Auth server failed OAuth exchange: %v", err)
	}

	return oauthConfig.TokenSource(ctx, token), nil
}

func loadConfig() (config memoConfig, err error) {
	config = memoConfig{}

	usr, err := user.Current()
	if err != nil {
		return
	}

	contents, err := ioutil.ReadFile(path.Join(usr.HomeDir, ".memoconfig"))
	if err != nil {
		return
	}

	err = json.Unmarshal(contents, &config)
	if err != nil {
		return
	}

	fmt.Println(config)

	return
}

func startAuthServer(ch chan<- string, state string) (func(context.Context) error, int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return nil, 0, fmt.Errorf("Auth server failed to listen: %v", err)
	}

	port := listener.Addr().(*net.TCPAddr).Port

	server := http.Server{
		Handler: authHandler{
			ch:    ch,
			state: state,
		},
	}

	errCh := make(chan error)

	go func() {
		if err := server.Serve(listener); err != http.ErrServerClosed {
			errCh <- fmt.Errorf("Auth server failed unexpectedly: %v", err)
		} else {
			errCh <- nil
		}
	}()

	shutdown := func(ctx context.Context) error {
		if err := server.Shutdown(ctx); err != nil {
			return fmt.Errorf("Auth server failed shutdown: %v", err)
		}

		return <-errCh
	}

	return shutdown, port, nil
}

var hexRunes = []rune("0123456789abcdef")

func genState() string {
	state := make([]rune, 32)
	for i := range state {
		state[i] = hexRunes[rand.Intn(16)]
	}
	return string(state)
}

type authHandler struct {
	ch    chan<- string
	state string
}

func (handler authHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if handler.state != r.URL.Query().Get("state") {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("<!DOCTYPE html>\n<html>\n<body>\n<h1>OAuth 2.0: Bad State</h1>\n</body>\n</html>"))
	} else if errMsg := r.URL.Query().Get("error"); errMsg != "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("<!DOCTYPE html>\n<html>\n<body>\n<h1>OAuth 2.0: %v</h1>\n</body>\n</html>", errMsg)))
	} else {
		handler.ch <- r.URL.Query().Get("code")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("<!DOCTYPE html>\n<html>\n<body>\n<h1>OAuth 2.0: Success</h1>\n</body>\n</html>"))
	}
}
