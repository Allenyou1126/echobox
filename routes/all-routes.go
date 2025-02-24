package routes

import (
	"echobox/frontend"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
	"os"
)

func RegisterRoutes(r *gin.Engine) {
	staticFS, err := fs.Sub(frontend.FS, "dist")
	if err != nil {
		os.Exit(1)
	}
	r.StaticFS("/admin", http.FS(staticFS))
}
