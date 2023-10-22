package core

import (
	"fmt"
	"github.com/mshafiee/swephgo"
	log "github.com/sirupsen/logrus"
	"os"
	"strings"
)

var ephePath string
var jplFile string

func LibInfo() {
	sweVer := make([]byte, 12)
	swephgo.Version(sweVer)
	log.Infof("Library used: Swiss Ephemeris v%s\n", sweVer)
}

func LoadEphPath() {
	if ephePath == "" {
		ephePath = os.Getenv("EPHE_PATH")
		if ephePath == "" {
			ephePath = "jpl"
		}
		log.Infof("Use Ephe Path: %s", ephePath)
	}

	if jplFile == "" {
		jplFile = os.Getenv("JPL_FILE")
		if jplFile == "" {
			jplFile = "de440.eph"
		}
		log.Infof("Use Jpl File: %s", jplFile)
	}
}

func CalcUt(tjdUt float64, ipl int, xx []float64) error {
	LoadEphPath()
	swephgo.SetEphePath([]byte(ephePath))
	swephgo.SetJplFile([]byte(jplFile))

	serr := make([]byte, 256)
	swephgo.CalcUt(tjdUt, ipl, swephgo.SeflgJpleph, xx, serr)
	if serr[0] != 0 {
		return fmt.Errorf(strings.ReplaceAll(string(serr), "\n", ""))
	}
	return nil
}
