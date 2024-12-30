package resources

import "embed"

//go:embed queries.sql
var Queries string

//go:embed migrations/*.sql
var Migrations embed.FS
