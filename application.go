package main

import (
	"fmt"
	"github.com/marcelblijleven/SuperMarketApp/basket"
	"github.com/marcelblijleven/SuperMarketApp/inventory"
	"log"
	"os"
	"runtime/debug"
)

// application is the main interaction point of the application
type application struct {
	infoLog          *log.Logger
	errorLog         *log.Logger
	basketManager    basket.Manager
	inventoryManager inventory.Manager
}

// newApplication returns a new application with initialised logs
func newApplication() *application {
	return &application{
		infoLog:  log.New(os.Stdout, "INFO:\t", log.Ldate|log.Ltime),
		errorLog: log.New(os.Stdout, "ERROR:\t", log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (a *application) LogInfo(v ...interface{}) {
	a.infoLog.Println(v...)
}

func (a *application) LogInfoFormatted(format string, v ...interface{}) {
	a.infoLog.Printf(format, v...)
}

func (a *application) LogError(err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	if err = a.errorLog.Output(2, trace); err != nil {
		a.errorLog.Println(err.Error())
	}
}
