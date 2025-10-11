package tui

const (
	FIBER string = "fiber"
	GIN   string = "gin"
	CHI   string = "chi"
)

var (
	frameWorkChoises = []string{FIBER, GIN, CHI}
)

const (
	MYSQL_SQLX      string = "mysql"
	MARIADB_SQLX    string = "mariadb"
	POSTGRESQL_PGX  string = "postgresql"
	MYSQL_GORM      string = "gorm-mysql"
	POSGTRESQL_GORM string = "gorm-pg"
)

var (
	databaseDriverChoises = []string{MYSQL_SQLX, MARIADB_SQLX, POSTGRESQL_PGX, MYSQL_GORM, POSGTRESQL_GORM}
)
