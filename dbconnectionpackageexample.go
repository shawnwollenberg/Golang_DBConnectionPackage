package dbconnectionpackageexample

// to import in main project ...
//	dbconnect "dbConnectionPackageExample" //import simple db connection string package

import "fmt"

var (
	DBConnString string
	TestDbName   string
)

func init() {
	var dbpassword = "<pwd>"
	var dbserver = "<server>"
	var dbuser = "<username>"
	DBConnString = fmt.Sprintf("server=%s;user id=%s;password=%s", dbserver, dbuser, dbpassword)

	TestDbName = "testDB" //variable to hide DB name for a little extra security
}
