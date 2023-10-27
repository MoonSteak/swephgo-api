package handlers

import (
	"github.com/MoonSteak/swephgo-api/internal/core"
	"net/http"
	"strconv"
	"strings"
)

func BodiesDegreeHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	jdn, err := strconv.ParseFloat(query.Get("jdn"), 64)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest,
			"error parsing 'jdn' query param. Usage query example: '?bodies=1,2,3&jdn=2452305.5625'")
		return
	}

	var iplArr []int
	for _, body := range strings.Split(query.Get("bodies"), ",") {
		ipl, err := strconv.Atoi(body)
		if err != nil {
			WriteErrorResponse(w, http.StatusBadRequest,
				"error parsing 'bodies' query param. Usage query example: '?bodies=1,2,3&jdn=2452305.5625'")
			return
		}
		iplArr = append(iplArr, ipl)
	}

	degrees := make(map[int]float64, len(iplArr))
	xx := make([]float64, 6)
	for _, ipl := range iplArr {
		err := core.CalcUt(jdn, ipl, xx)
		if err != nil {
			WriteErrorResponse(w, http.StatusBadRequest, "calculation error")
			return
		}
		degrees[ipl] = xx[0]
	}
	WriteDataResponse(w, degrees)
}

func JdnByOffsetHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	ipl, err := strconv.Atoi(query.Get("body"))
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest,
			"error parsing 'body' query param. Usage query example: '?body=0&jdn=2452305.5625&&delta=100.0'")
		return
	}

	jdn, err := strconv.ParseFloat(query.Get("jdn"), 64)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest,
			"error parsing 'jdn' query param. Usage query example: '?body=0&jdn=2452305.5625&&delta=100.0'")
		return
	}

	delta, err := strconv.ParseFloat(query.Get("delta"), 64)
	if err != nil {
		WriteErrorResponse(w, http.StatusBadRequest,
			"error parsing 'jdn' query param. Usage query example: '?body=0&jdn=2452305.5625&&delta=100.0'")
		return
	}

	WriteDataResponse(w, core.FindJDN(jdn, ipl, delta))
}
