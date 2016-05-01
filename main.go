package main

import (

	."github.com/Abdul2/dbupload/dboperations"
	."github.com/Abdul2/dbupload/helper"
	."github.com/Abdul2/dbupload/data"

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
	var jsoncontent []Game
	var filesname *[]string
	var err error

	Dboperations("deletetable")
	Dboperations("createtable")

	filesname, err = filenamesindir("datasource")

	if err != nil {

		Error.Println(err.Error())
	}

	for counter, value := range *filesname {

		Info.Printf("file number %d is called %s", counter, value)

		bytecontent = Readfileintobyteslice(value) //read file content into byte slice

		jsoncontent = Loadintostruct(bytecontent) //read into the byte slice into struct

		Info.Println(jsoncontent)

		Dboperations("insertraws")

	}
}

//filenameindir returns a slice with fully qualified file names of a
//givev directory.
func filenamesindir(dir string) (*[]string, error) {


	Info.Println(`filenamesindir() started`)

	var err error
	var filesinfo []os.FileInfo
	var filenames = make([]string,0)
	var dirfqn string
	var currentdir string


	currentdir, err = os.Getwd()

	if err != nil{

		Error.Println("\n\n\n cant get the working dir \n")
	}

	dirfqn = fmt.Sprintf(`%s/%s`,currentdir,dir)


	fmt.Sprintf(dirfqn)

	filesinfo, err = ioutil.ReadDir(dirfqn)

	if err != nil {

		Error.Println("cant read the source dir")
		return nil, errors.New("cant read source dir")
	}

	for f := range filesinfo{


		filefqn := fmt.Sprintf(`%s/%s`,dirfqn,filesinfo[f].Name())

		Info.Println (filesinfo[f].Name())
		filenames = append(filenames, filefqn)
	}

	return &filenames,nil
}
