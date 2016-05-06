package resultsrestservice

import (
	"database/sql"
	"github.com/coopernurse/gorp"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	//"log"
	"os"
	"strconv"
	//D "github.com/Abdul2/dbupload/data"
	H "github.com/Abdul2/dbupload/helper"
)

//establish connection to db or fail early
var dbmap = initDb()

//iniDb connects to db and returns a connection object
func initDb() *gorp.DbMap {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL_CPFC"))

	if err != nil {
		H.CheckError(err)
	}

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}
	return dbmap
}


//Serverstart starts a server listening to port 8000
func Serverstart() {

	//defer connection to database until all db operations are completed
	defer dbmap.Db.Close()
	router := Router()
	router.Run(":9000")
}

//Router
func Router() *gin.Engine {

	router := gin.Default()
	router.GET("/results", allresults)

	return router
}

func allresults(c *gin.Context) {


	var result []struct{

		Uid int64
		Gamedate string
		Team string
		Awayorhome string
		Competition string
		Result string
		Score1 string
		Score2 string



	}



	_, err := dbmap.Select(&result, "select * from cpfc;")

	if err != nil {

		H.CheckError(err)

	}

	content := gin.H{}

	for k, v := range result {
		content[strconv.Itoa(k)] = v
	}
	c.JSON(200, content)

}

