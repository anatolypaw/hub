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

type App struct {
	chiMux *chi.Mux
	auth   *authservice.Auth
	mstore *mstore.MStore
}

//
//go:embed all:webpanel/dist
var fsys embed.FS

func New(mstore *mstore.MStore) *App {
	reactAppFolder, err := fs.Sub(fsys, "webpanel/dist")
	if err != nil {
		log.Fatal(err)
	}
	fileServer := http.FileServerFS(reactAppFolder)

	authService := authservice.New()

	// Добавляет пользователей и права
	authService.AddUser("admin", "admin", "admin")
	authService.AddUser("user", "user", "user")

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Для всех
	r.Group(func(r chi.Router) {
		r.Handle("/*", fileServer)
	})

	// Только для админов
	r.Group(func(r chi.Router) {
		r.Use(mware.ChekAuth(&authService, "admin"))

		r.Mount("/profiler", middleware.Profiler())
		r.Get("/debug", handlers.Debug)

		r.Get("/api/index", handlers.Index)
		r.Get("/api/goods", handlers.GoodsGet(mstore))
	})

	// Для пользователей и админов
	r.Group(func(r chi.Router) {
		r.Use(mware.ChekAuth(&authService, "admin", "user"))
	})

	return &App{
		chiMux: r,
		auth:   &authService,
		mstore: mstore,
	}
}

func (a *App) Run(addr string) error {
	return http.ListenAndServe(addr, a.chiMux)
}
