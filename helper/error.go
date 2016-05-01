package helper


import "os"

func CheckError(err error){


	Error.Println(err)
	os.Exit(1)

}