package web

import (
	"embed"
	"hub/internal/web/authservice"
	"hub/internal/web/handlers"
	"hub/internal/web/mware"
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
}

// Встроим все статические файлы веб сервера в бинарник
//
//go:embed static/*
var staticContent embed.FS

func New() *App {
	auth := authservice.New()

	// Добавляет пользователей и права
	auth.AddUser("admin", "admin", "admin")
	auth.AddUser("user", "user", "user")

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)

	// Админские страницы
	r.Group(func(r chi.Router) {
		r.Use(mware.ChekAuth(&auth, "admin"))

		r.Mount("/debug", middleware.Profiler())

	})

	// Для пользователей
	r.Group(func(r chi.Router) {
		r.Use(mware.ChekAuth(&auth, "admin", "user"))

		r.Get("/", handlers.Index)
		r.Get("/index.html", handlers.Index)
	})

	// Для всех
	r.Group(func(r chi.Router) {
		r.Handle("/static/*", http.FileServer(http.FS(staticContent)))

		r.Get("/login", handlers.LoginForm)
		r.Post("/login", handlers.Login(&auth))
	})

	return &App{
		chiMux: r,
		auth:   &auth,
	}
}

func (a *App) Run(addr string) error {
	return http.ListenAndServe(addr, a.chiMux)
}
