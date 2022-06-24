package stoplight_elements

import (
	"embed"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-openapi/spec"
	"github.com/pkg/errors"

	"github.com/far4599/swagger-openapiv2-merge/pkg/file"
	"github.com/far4599/swagger-openapiv2-merge/pkg/marshaller"
)

//go:embed static/*.css static/*.js
var staticFS embed.FS

//go:embed static/index_template.html
var indexTemplate []byte

type elementsServer struct {
	specFilePath string
	httpPort     string
	hostname     string
	openUrl      bool
}

func NewServer(specFilePath string) *elementsServer {
	return &elementsServer{
		specFilePath: specFilePath,
	}
}

func (e *elementsServer) WithHostname(hostname string) *elementsServer {
	e.hostname = hostname
	return e
}

func (e *elementsServer) WithPort(port string) *elementsServer {
	e.httpPort = port
	return e
}

func (e *elementsServer) WithOpenURL(openUrl bool) *elementsServer {
	e.openUrl = openUrl
	return e
}

func (e *elementsServer) newRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(middleware.Compress(9))
	r.Use(middleware.Recoverer)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(indexTemplate)
	})

	r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		var swagger spec.Swagger

		err := file.NewFileReader(e.specFilePath).Read(&swagger)
		if err != nil {
			err = errors.Wrapf(err, "failed to get spec from file '%s'", e.specFilePath)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		if len(e.hostname) > 0 {
			swagger.Host = e.hostname
		}

		fileContent, err := marshaller.Marshal(swagger, marshaller.JSONFormat)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(fileContent)
	})

	r.Handle("/static/*", CacheStatic(http.FileServer(http.FS(staticFS))))

	return r
}

func (e elementsServer) Run() (err error) {
	url := fmt.Sprintf("http://localhost:%s", e.httpPort)

	if e.openUrl {
		go func() {
			if err != nil {
				return
			}

			time.Sleep(1 * time.Second)

			err = openUrl(url)
			if err != nil {
				fmt.Printf("failed to open url: %v\n", err)
			}
		}()
	}

	fmt.Printf("Serving swagger UI at: %s - You may open this link in browser.\n", url)
	fmt.Println("Press Ctrl+C to stop the server.")
	return http.ListenAndServe(":"+e.httpPort, e.newRouter())
}

func CacheStatic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Cache-Control", "public, max-age=31556952")

		next.ServeHTTP(w, r)
	})
}

func openUrl(url string) error {
	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	}

	return fmt.Errorf("unsupported platform")
}
