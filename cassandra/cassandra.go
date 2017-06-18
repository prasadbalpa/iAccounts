package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"iAccounts/cfg"
)

func GetClusterHandle(keyspace string) *(gocql.ClusterConfig) {
	fmt.Println("KeySpace: " + keyspace)
	cluster := gocql.NewCluster(cfg.GetCassandraClusters())
	cluster.Keyspace = keyspace
	return (cluster)
}

func GetSessionHandle(cluster_handle *gocql.ClusterConfig) *gocql.Session {
	session, _ := cluster_handle.CreateSession()
	return (session)
}

