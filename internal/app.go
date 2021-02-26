package internal

import (
	"database/sql"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//App estructura de la app
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

//Initialize metodo que inicializa la BD
func (a *App) Initialize(user, password, dbname string) {

}

//Run metodo que iniciara la app
func (a *App) Run(addr string) {

}
