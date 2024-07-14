package web

import (
	"embed"
	"hub/internal/mstore"
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
	mstore *mstore.MStore
}

// Встроим все статические файлы веб сервера в бинарник
//
//go:embed static
var staticContent embed.FS

func New(mstore *mstore.MStore) *App {
	authService := authservice.New()

	// Добавляет пользователей и права
	authService.AddUser("admin", "admin", "admin")
	authService.AddUser("user", "user", "user")

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.NoCache)

	// Админские страницы
	r.Group(func(r chi.Router) {
		r.Use(mware.ChekAuth(&authService, "admin"))

		r.Mount("/profiler", middleware.Profiler())
		r.Get("/debug", handlers.Debug)

	})

	// Для пользователей
	r.Group(func(r chi.Router) {
		r.Use(mware.ChekAuth(&authService, "admin", "user"))

		r.Get("/", handlers.Index)
		r.Get("/index", handlers.Index)
		r.Get("/goods", handlers.GoodsGet(mstore))
	})

	// Для всех
	r.Group(func(r chi.Router) {
		r.Handle("/static/*", http.FileServer(http.FS(staticContent)))

		r.Get("/login", handlers.LoginGet)
		r.Post("/login", handlers.LoginPost(&authService))
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
