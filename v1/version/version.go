package version

import (
	"database/sql"

	"github.com/gorilla/mux"
	//postgres
	_ "github.com/lib/pq"
)

// V1 route struct
type V1 struct {
	DB     *sql.DB
	Router *mux.Router
}
