package web

import (
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

func New() *App {
	auth := authservice.New()

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

		r.Get("/", handlers.MainForm)
	})

	// Для всех
	r.Group(func(r chi.Router) {
		fs := http.FileServer(http.Dir("static"))
		r.Handle("/static/*", http.StripPrefix("/static/", fs))

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
