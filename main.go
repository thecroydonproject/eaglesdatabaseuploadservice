package main

import (

	O "github.com/Abdul2/dbupload/dboperations"
	 H "github.com/Abdul2/dbupload/helper"
	D "github.com/Abdul2/dbupload/data"

)
import (
	"io/ioutil"
	"os"
	"errors"
	"fmt"
)


//main starts the application
func main(){


	var bytecontent []byte
	var jsoncontent []D.Game
	var filesname *[]string
	var err error

	O.Dboperations("deletetable")
	O.Dboperations("createtable")

	filesname, err = filenamesindir("datasource")

	if err != nil {

		H.Error.Println(err.Error())
	}

	for counter, value := range *filesname {

		H.Info.Printf("file number %d is called %s", counter, value)

		bytecontent = O.Readfileintobyteslice(value) //read file content into byte slice

		jsoncontent = O.Loadintostruct(bytecontent) //read into the byte slice into struct

		H.Info.Println(jsoncontent)

		O.Dboperations("insertraws")

	}
}

//filenameindir returns a slice with fully qualified file names of a
//givev directory.
func filenamesindir(dir string) (*[]string, error) {


	H.Info.Println(`filenamesindir() started`)

	var err error
	var filesinfo []os.FileInfo
	var filenames = make([]string,0)
	var dirfqn string
	var currentdir string


	currentdir, err = os.Getwd()

	if err != nil{

		H.Error.Println("\n\n\n cant get the working dir \n")
	}

	dirfqn = fmt.Sprintf(`%s/%s`,currentdir,dir)


	fmt.Sprintf(dirfqn)

	filesinfo, err = ioutil.ReadDir(dirfqn)

	if err != nil {

		H.Error.Println("cant read the source dir")
		return nil, errors.New("cant read source dir")
	}

	for f := range filesinfo{


		filefqn := fmt.Sprintf(`%s/%s`,dirfqn,filesinfo[f].Name())

		H.Info.Println (filesinfo[f].Name())
		filenames = append(filenames, filefqn)
	}

	return &filenames,nil
}
