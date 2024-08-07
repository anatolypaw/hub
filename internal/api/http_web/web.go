package http_web

import (
	"embed"
	"hub/internal/api/http_web/authservice"
	"hub/internal/api/http_web/handlers"
	"hub/internal/api/http_web/mware"
	"hub/internal/mstore"
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type User struct {
	Username string
	Password string
}

type Webpanel struct {
	chiMux *chi.Mux
	auth   *authservice.Auth
	mstore *mstore.MStore
}

//
//go:embed all:webpanel/build
var fsys embed.FS

func New(mstore *mstore.MStore, version string) *Webpanel {
	webpanelFolder, err := fs.Sub(fsys, "webpanel/build")
	if err != nil {
		log.Fatal(err)
	}
	fileServer := http.FileServerFS(webpanelFolder)

	authService := authservice.New()

	// Добавляет пользователей и права
	authService.AddUser("admin", "admin", "admin")
	authService.AddUser("user", "user", "user")

	r := chi.NewRouter()
	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	})

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Доступно для всех
	r.Group(func(r chi.Router) {
		r.Handle("/*", fileServer)
		r.Post("/api/login", handlers.Login(&authService))
	})

	// Для авторизованных пользователей
	r.Group(func(r chi.Router) {
		r.Use(mware.ChekAuth(&authService, "admin", "user"))
		r.Handle("/", fileServer)
		r.Handle("/api/about", handlers.AboutInfo(version))
		r.Handle("/api/userinfo", handlers.UserInfo(&authService))
	})

	// Только для админов
	r.Group(func(r chi.Router) {
		r.Use(mware.ChekAuth(&authService, "admin"))

		r.Mount("/profiler", middleware.Profiler())

	})

	return &Webpanel{
		chiMux: r,
		auth:   &authService,
		mstore: mstore,
	}
}

func (a *Webpanel) Run(addr string) error {
	return http.ListenAndServe(addr, a.chiMux)
}
