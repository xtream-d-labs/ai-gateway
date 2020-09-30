package http

import (
	"compress/gzip"
	"compress/zlib"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-openapi/swag"
	"github.com/scaleshift/scaleshift/api/src/auth"
	"github.com/scaleshift/scaleshift/api/src/config"
	"github.com/scaleshift/scaleshift/api/src/db"
	"github.com/scaleshift/scaleshift/api/src/log"
)

// Wrap wraps HTTP request handler
func Wrap(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case eqauls(r, "/health"):
			w.WriteHeader(http.StatusOK)

		case eqauls(r, "/version"):
			fmt.Fprintf(w, "%s", config.Config.APIVersion)

		default:
			proc := time.Now()

			if config.Config.AllowCORS {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Header().Set("Access-Control-Allow-Methods", "GET,POST,DELETE,PUT,PATCH,HEAD")
				w.Header().Set("Access-Control-Allow-Headers", "Origin,Content-Type,Authorization")
				w.Header().Set("Access-Control-Expose-Headers", "Content-Type,Authorization,Date")
				w.Header().Set("Access-Control-Max-Age", "86400")
			}
			if strings.EqualFold(r.Method, http.MethodOptions) {
				w.WriteHeader(http.StatusOK)
				return
			}
			if config.Config.SecuredTransport {
				if r.Header.Get("X-Forwarded-Proto") == "http" {
					http.Redirect(w, r, config.Config.APIEndpoint+r.RequestURI, http.StatusMovedPermanently)
					w.Header().Add("Strict-Transport-Security", "max-age=31536000;") // 1 year
					return
				}
			}
			ioWriter := w.(io.Writer)
			if encodings, found := header(r, "Accept-Encoding"); found && config.Config.ContentEncoding {
				for _, encoding := range splitCsvLine(encodings[0]) {
					if encoding == "gzip" {
						w.Header().Set("Content-Encoding", "gzip")
						g := gzip.NewWriter(w)
						defer g.Close()
						ioWriter = g
						break
					}
					if encoding == "deflate" {
						w.Header().Set("Content-Encoding", "deflate")
						z := zlib.NewWriter(w)
						defer z.Close()
						ioWriter = z
						break
					}
				}
			}
			config.Config.Init()

			if sess, err := auth.RetrieveSession(r); err == nil && sess != nil {
				creds := auth.FindCredentials(sess.DockerUsername)
				if !swag.IsZero(creds.Base.DockerRegistry) {
					config.Config.DockerRegistryEndpoint = creds.Base.DockerRegistry
				}
				if !swag.IsZero(creds.Base.DockerHostname) {
					config.Config.DockerRegistryHostName = creds.Base.DockerHostname
				}
				if !swag.IsZero(creds.Base.DockerUsername) {
					config.Config.DockerRegistryUserName = creds.Base.DockerUsername
				}
				if !swag.IsZero(creds.Base.RescalePlatform) {
					config.Config.RescaleEndpoint = creds.Base.RescalePlatform
				}
				if !swag.IsZero(creds.Base.RescaleKey) {
					config.Config.RescaleAPIToken = creds.Base.RescaleKey
				}
			}
			writer := overrideWriter(w, ioWriter)
			handler.ServeHTTP(writer, r)

			if writer.Header().Get("Content-Type") == "application/json" {
				writer.Header().Set("Content-Type", "application/json; charset=UTF-8")
			}
			writer.Header().Add("X-Frame-Options", "SAMEORIGIN")
			writer.Header().Add("X-XSS-Protection", "1; mode=block")
			writer.Header().Add("X-Content-Type-Options", "nosniff")
			writer.Header().Add("Content-Security-Policy", "default-src 'none';style-src 'self' 'unsafe-inline';")

			if config.Config.AccessLog {
				addr := r.RemoteAddr
				if ip, found := header(r, "X-Forwarded-For"); found {
					addr = ip[0]
				}
				log.Info(fmt.Sprintf("[%s] %.3f %d %s %s",
					addr, time.Since(proc).Seconds(),
					writer.status, r.Method, r.URL), nil, nil)
			}
		}
	})
}

// ServerShutdown wraps the HTTP server shutdown event
func ServerShutdown() {
	db.ShutdownQueue()
	db.ShutdownCache()
	db.Shutdown()
}

func eqauls(r *http.Request, url string) bool {
	return url == r.URL.Path
}

func header(r *http.Request, key string) (values []string, found bool) {
	if r.Header == nil {
		return nil, false
	}
	for k, v := range r.Header {
		if strings.EqualFold(k, key) && len(v) > 0 {
			return v, true
		}
	}
	return nil, false
}

func splitCsvLine(data string) []string {
	splitted := strings.SplitN(data, ",", -1)
	parsed := make([]string, len(splitted))
	for i, val := range splitted {
		parsed[i] = strings.TrimSpace(val)
	}
	return parsed
}
