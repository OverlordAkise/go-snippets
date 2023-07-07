package db

import (
	"database/sql"
	_ "modernc.org/sqlite"
    //tracing - db
    "github.com/uptrace/opentelemetry-go-extra/otelsql"
    "go.opentelemetry.io/otel/trace"
    //tracing - for tracing in gin http request context
    "context"
)

var db *sql.DB

type dbstruct struct {
	Key   string `db:"k" json:"k"`
	Value string `db:"v" json:"v"`
}

func InitDatabase(tp trace.TracerProvider) {
	var err error
	db, err = otelsql.Open("sqlite", ":memory:", otelsql.WithTracerProvider(tp))
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS example(k TEXT PRIMARY KEY, v TEXT);")
    if err != nil { panic(err) }
	_, err = db.Exec("INSERT INTO example(k,v) VALUES('key1','value1');")
    if err != nil { panic(err) }
}

func GetColumns(c context.Context) []dbstruct {
	ret := make([]dbstruct,0)
	rows, err := db.QueryContext(c, "SELECT * FROM example")
	if err != nil {
		panic(err)
	}
    defer rows.Close()
    for rows.Next() {
        var r dbstruct
        if err := rows.Scan(&r.Key, &r.Value); err != nil {
            panic(err)
        }
        ret = append(ret,r)
    }
	return ret
}
