package helper

import (

	"log"
	"os"
)



var Trace,  Info,  Warning, Error =log.New(os.Stdout,"TRACE",log.Ldate|log.Ltime|log.Lshortfile), log.New(os.Stdout,"INFO",log.Ldate|log.Ltime|log.Lshortfile) ,log.New(os.Stdout,"Warning",log.Ldate|log.Ltime|log.Lshortfile),log.New(os.Stdout,"Error",log.Ldate|log.Ltime|log.Lshortfile)