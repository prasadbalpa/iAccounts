package cfg

import (
	"iAccounts/utils"
)
//All of the configurations are stored in this file
const (
	CASSANDRA_CLUSTER         = "127.0.0.1"
	ORGANIZATION_KEYSPACE     = "organization"
	ORGUSERMAP_KEYSPACE       = "orgusermap"
	ORGID_KEYSPACE            = "org_"
	HTTP_HEADER_CONTENT_TYPE  = "Content-type"
	HTTP_HEADER_DATATYPE_JSON = "application/json"
	LOGIN_MOBILE_NUMBER       = 1
	LOGIN_OTP_NUMBER          = 2
	LOGIN_SUCCESSFUL          = 3
	HTTPS_SERVER_PORT = ":8443"
	HTTPS_TLS_CERTIFICATE = "/Users/prasadk/go/src/iAccounts/certs/localhost.crt"
	HTTPS_TLS_KEY =  "/Users/prasadk/go/src/iAccounts/certs/localhost.key"

)

func GetCassandraClusters() string {
	return CASSANDRA_CLUSTER
}
func GetOrganizationKeySpace() string {
	return ORGANIZATION_KEYSPACE
}
func GetOrgUserMapKeySpace() string {
	return ORGUSERMAP_KEYSPACE
}
func GetOrgIDKeySpace(orgid string) string {
	return (ORGID_KEYSPACE + utils.FindAndReplace(orgid, "-", "_"))
}
