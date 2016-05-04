package resultsrestservice

import (
	"database/sql"
//"fmt"
	"github.com/coopernurse/gorp"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"os"
	"strconv"
	D "github.com/Abdul2/dbupload/data"
)

//establish connection to db or fail early
var dbmap = initDb()


func initDb() *gorp.DbMap {

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL_CPFC"))
	checkErr(err, "can open db")


	//db, err := sql.Open("postgres", dbUrl)

	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.PostgresDialect{}}


	return dbmap
}

//checkErr is a helper function to deal with errors
func checkErr(err error, msg string) {
	if err != nil {
		//log.Fatalln(msg, err)

		log.Print(err.Error())
	}

}

//main is programme entry point
func Serverstart() {

	//defer connection to database until all db operations are completed
	defer dbmap.Db.Close()
	router := Router()
	router.Run(":8000")
}

func Router() *gin.Engine {

	router := gin.Default()
	router.GET("/results", allresults)

	return router
}

func allresults(c *gin.Context) {

	var result []D.Game

	_, err := dbmap.Select(&result, "select * from cpfc;")

	checkErr(err, "Select failed")

	content := gin.H{}

	for k, v := range result {
		content[strconv.Itoa(k)] = v
	}
	c.JSON(200, content)

}

