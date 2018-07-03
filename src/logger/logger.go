package logger

import (
	"github.com/plivo/go-plivolog/plivolog"
)

// Log - log object
var Log *plivolog.PlivoLogger

// Init function for initializing the logger
func Init(env string) {

	Log, _ = plivolog.New()
	Log.Infof("Initialized Plivo Logger for %s", env)

}
