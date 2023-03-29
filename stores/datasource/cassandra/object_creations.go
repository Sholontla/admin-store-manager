package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
)

func Run() {
	// Connect to Cassandra cluster
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "system"
	session, err := cluster.CreateSession()
	if err != nil {
		fmt.Println("Error connecting to cluster:", err)
		return
	}
	defer session.Close()

	// Create keyspace
	err = createKeyspace(session)
	if err != nil {
		fmt.Println("Error creating keyspace:", err)
		return
	}

	// Create table
	err = createTable(session)
	if err != nil {
		fmt.Println("Error creating table:", err)
		return
	}

	fmt.Println("Keyspace and table created successfully")
}

func createKeyspace(session *gocql.Session) error {
	return session.Query(`CREATE KEYSPACE IF NOT EXISTS my_keyspace
	                      WITH replication = {'class': 'SimpleStrategy', 'replication_factor': 1};`).Exec()
}

func createTable(session *gocql.Session) error {
	return session.Query(`CREATE TABLE IF NOT EXISTS my_keyspace.my_table (
	                      id UUID PRIMARY KEY,
	                      name text,
	                      age int
	                     );`).Exec()
}
