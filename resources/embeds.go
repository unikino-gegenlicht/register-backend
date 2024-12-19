package resources

import "embed"

//go:embed queries.sql
var QueryFile embed.FS

//go:embed migrations/*.sql
var Migrations embed.FS
