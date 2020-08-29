//go:generate statik -src=./swagger-ui
package lite

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
	"github.com/shinecloudnet/explorer/backend/conf"
	_ "github.com/shinecloudnet/explorer/backend/lcd/lite/statik"
	"github.com/shinecloudnet/explorer/backend/logger"
)

func RegisterSwaggerUI(r *mux.Router) {
	if conf.Get().Server.CurEnv == conf.EnvironmentDevelop || conf.Get().Server.CurEnv == conf.EnvironmentLocal || conf.Get().Server.CurEnv == conf.EnvironmentQa {
		statikFS, err := fs.New()
		if err != nil {
			panic(err)
		}

		staticServer := http.FileServer(statikFS)
		r.PathPrefix("/swagger-ui/").Handler(http.StripPrefix("/swagger-ui/", staticServer))
		logger.Info(fmt.Sprintf("enalbe swagger ui in %s environment.", conf.Get().Server.CurEnv))

	} else {
		logger.Info(fmt.Sprintf("disable swagger ui in %s environment.", conf.Get().Server.CurEnv))
	}
}
