package sql

import "net/http"



func Queries(w* http.ResponseWriter, r* http.Request){
	if r.Method != http.MethodPost {
		
	}
}
// here i am try to run the sql queries from user to postgres docker container
// i have created a table named users in the database named mydb
// you can create your own table and database and run the queries accordingly