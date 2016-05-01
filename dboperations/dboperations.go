
//package dboperations provides operations to create and delete table cpfc
//and insert raws into cpfc
package dboperations



import (
	. "github.com/Abdul2/dbupload/sqlstmnts"
	"encoding/json"
	"log"
	_ "github.com/lib/pq"
	."github.com/Abdul2/dbupload/helper"
	."github.com/Abdul2/dbupload/data"
	"os"

	"io/ioutil"
)



var file *os.File

//Dboperations calls dbtransactions to run dboperations
func Dboperations(op string) {

	Trace.Println("Dboperations started")

	var t DBEntity

	switch op {

	case "deletetable":

		t.Entitytype = "table"
		t.Sqlstmnt = Deletetable
		dbtransactions(t, t.Entitytype)

	case "createtable":

		t.Entitytype = "table"
		t.Sqlstmnt = Createtable
		dbtransactions(t, t.Entitytype)

	case "insertraws":

		t.Entitytype = "raw"
		t.Sqlstmnt = Insertraw
		t.Raws = Structcontent
		dbtransactions(t, t.Entitytype)

	}

}

//dbtransactions follows a simple pattern of prep sql statement and then excute it
func dbtransactions(t DBEntity, entitytype string) {

	switch  entitytype {

	case "table":

		t.Prepstmt(t.Sqlstmnt)
		t.Execstmt(t.Preparedstmnt)

	case "raw":

		t.Insertrawst(t.Sqlstmnt, t.Raws)
		Info.Println("lastrawnumber is ",t.LastInsert)

	}

}

//Loadintostruct reads bytes into a struct
//the structure of the struct object and json must match
func Loadintostruct(contentinbytes []byte) ([]Game) {



	err := json.Unmarshal(contentinbytes, &Structcontent)

	if err != nil {

		Error.Println(err.Error())
		os.Exit(1)
	}

	return Structcontent

}


// Readfileintobyteslice takes in the name of file and returns the
//the file content in a byte slice

func Readfileintobyteslice(filename string) []byte {


	var err  error
	var bytes []byte

	Info.Println("\n file to procesess is \n",filename)

	if bytes, err = ioutil.ReadFile(filename); err != nil{

		Error.Println(err.Error())

	}


	//
	Info.Printf("%s  file content is \n %s",filename, string(bytes))


	return bytes


}



//error

func checkErr(err error, msg string) {
	if err != nil {
		log.Fatalln(msg, err)
	}
}