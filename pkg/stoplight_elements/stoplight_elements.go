package stoplight_elements

import (
	"embed"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

//go:embed static/*.css static/*.js
var staticFS embed.FS

//go:embed static/index_template.html
var indexTemplate []byte

type elementsServer struct {
	router   chi.Router
	httpPort string
}

func NewServer(filePath string) *elementsServer {
	r := chi.NewRouter()
	r.Use(middleware.Compress(9))
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(indexTemplate)
	})

	r.Get("/swagger.spec", func(w http.ResponseWriter, r *http.Request) {
		fileContent, err := ioutil.ReadFile(filePath)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
		}

		w.Write(fileContent)
	})

	r.Handle("/static/*", CacheStatic(http.FileServer(http.FS(staticFS))))

	return &elementsServer{
		router: r,
	}
}

func (e *elementsServer) WithPort(port string) *elementsServer {
	e.httpPort = port
	return e
}

func (e elementsServer) Run() error {
	fmt.Printf("Serving swagger UI at: http://localhost:%s - You may open this link in browser.\n", e.httpPort)
	fmt.Println("Press Ctrl+C to stop the server.")
	return http.ListenAndServe(":"+e.httpPort, e.router)
}

func CacheStatic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=31556952")

		next.ServeHTTP(w, r)
	})
}
