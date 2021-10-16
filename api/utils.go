package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func handleErr(err error, w http.ResponseWriter, statusError int) {
	if err == nil {
		return
	}
	log.Println(err)
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Error to create a new clocker schema"))
}

func convertStringArrayToIntArray(stringArray string) ([]int, error) {
	sArray := strings.Split(stringArray, ",")
	var iArray []int

	for _, str := range sArray {
		iStr, err := strconv.Atoi(str)
		if err != nil {
			return nil, errors.New("Cannt make convert")
		}
		iArray = append(iArray, iStr)
	}

	return iArray, nil
}
