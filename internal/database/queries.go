package database

import (
	"log/slog"
	"os"
	"register-backend/resources"

	"github.com/qustavo/dotsql"
)

var Queries *dotsql.DotSql

func init() {

	queries, err := dotsql.LoadFromString(resources.Queries)
	if err != nil {
		slog.Error("unable to read emebedded queries", "error", err.Error())
		os.Exit(1)
	}
	Queries = queries
}
