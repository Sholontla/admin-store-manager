package cassandra

import (
	"fmt"
	"log"
	"time"

	"github.com/gocql/gocql"
)

type DB interface {
	Query(string, ...interface{}) *gocql.Query
	Close()
}

type CassandraDB struct {
	Session *gocql.Session
}

func (c *CassandraDB) Query(query string, args ...interface{}) *gocql.Query {
	return c.Session.Query(query, args...)
}

func (c *CassandraDB) Close() {
	c.Session.Close()
}

func NewCassandraDB() (DB, error) {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "stores"
	cluster.Consistency = gocql.Quorum
	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "yourusername",
		Password: "yourpassword",
	}
	cluster.Timeout = 10 * time.Second

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, err
	}

	return &CassandraDB{Session: session}, nil
}

func Ping(db DB) {
	// Use the session to execute queries
	var result string
	if err := db.Query("SELECT release_version FROM system.local").Scan(&result); err != nil {
		log.Fatal(err)
	}
	fmt.Println(result)
}
