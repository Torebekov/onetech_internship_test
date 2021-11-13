package acmp

import (
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

func Difficulty(url string) float64 {
	client := &http.Client{}
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return -1.
	}

	req.AddCookie(&http.Cookie{Name: "English", Value: "1"})
	resp, err := client.Do(req)
	if err != nil {
		return -1.
	}

	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return -1.
	}

	respString := string(respData)
	start := strings.Index(respString, "Difficulty: ")
	numStr := ""
	for i := start + 12; true; i++ {
		if !(respString[i] >= '0' && respString[i] <= '9') {
			break
		}
		numStr += string(respString[i])
	}

	numDiff, err := strconv.Atoi(numStr)
	if err != nil || numDiff == 0 {
		return -1.
	}

	return float64(numDiff)
}
