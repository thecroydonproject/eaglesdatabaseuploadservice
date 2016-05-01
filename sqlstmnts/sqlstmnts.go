package sqlstmnts
import (
	"database/sql"
	"os"
	"log"
	_ "github.com/lib/pq"
	."github.com/Abdul2/dbupload/data"
	."github.com/Abdul2/dbupload/helper"


)





var Structcontent []Game


//sql statements
var Deletetable = "DROP TABLE IF EXISTS cpfc;"
var Createtable = "CREATE TABLE cpfc (uid serial NOT NULL, gamedate varchar(40), team varchar(40), awayorhome varchar(40),competition varchar(40),result varchar(40),score1 varchar(40),score2 varchar(40),UNIQUE(gamedate));"

var Insertraw = "INSERT INTO public.cpfc(gamedate,team,awayorhome,competition,result,score1,score2) VALUES($1,$2,$3,$4,$5,$6,$7) returning uid;"




//use anonymous function to establish connection
var conn = func() sql.DB {


	db, err := sql.Open("postgres",os.Getenv("DATABASE_URL_CPFC"))
	if err != nil {

		Error.Println("could not open db")
		os.Exit(1)
	}

	return *db
}()


//Entity interface of db calls.
// table, raw, schema ... are Entity
type Entity interface {

	Prepstmt(sqlstmnt string) *sql.Stmt
	Execstmt(*sql.Stmt) sql.Result
	insertrawst(Sqlstmnt string, raws []Game) int
}

//Table implements Entity
type DBEntity struct  {

	Sqlstmnt       string
	Preparedstmnt  *sql.Stmt
	Effectedraws int64
	Raws []Game
	Entitytype string //used to distinguish table from raw
	LastInsert int

}

//Prepstmt takes in the sql statement and returns a prepared sql statement
func (t *DBEntity) Prepstmt(sqlstmnt string) *sql.Stmt{

	Trace.Println("Prepstmt() started")

	st, err := conn.Prepare(t.Sqlstmnt)

	if err != nil {

		CheckError(err)
	}


	t.Preparedstmnt = st

	return st
}

//Execstmt takes in sql prepered statement and executes it returning the result
func (t *DBEntity) Execstmt(*sql.Stmt) sql.Result {




	st, err := conn.Exec(t.Sqlstmnt)

	if err != nil {

		log.Fatal(err)
	}

	i,err := st.RowsAffected()

	if err != nil {

		log.Fatal(err)
	}

	t.Effectedraws = i

	return st
}


//insertrawst takes a query string and slice of type Game
//inserts raws in the slice into cpfc table and returns the
//raw number of the last raw
func (t *DBEntity) Insertrawst(Sqlstmnt string, raws []Game)  {

	raws = Structcontent

	var lastInsertId int

	for i := range raws {

		err := conn.QueryRow(Sqlstmnt, raws[i].Gamedate, raws[i].Team, raws[i].Awayorhome, raws[i].Competition, raws[i].Result, raws[i].Score1, raws[i].Score2).Scan(&lastInsertId)

		if err != nil {

			log.Fatal(err.Error())

		}



	}

	t.LastInsert = lastInsertId
}