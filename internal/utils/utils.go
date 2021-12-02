package utils

import (
	"bufio"
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strconv"
	"strings"
)

const urlTemplate = "https://adventofcode.com/2021/day/{}/input"

// Loads puzzle input based on day, must set environment variable AOC_SESSION_COOKIE 
func LoadInput(day string) ([]byte, error) {
	secret := os.Getenv("AOC_SESSION_COOKIE")

	if len(secret) == 0 {
		return nil, errors.New("missing AOC_SESSION_COOKIE")
	}

	var cookies []*http.Cookie

	cookie := &http.Cookie{
		Name: "session",
		Value: secret,
		Path: "/",
		Domain: ".adventofcode.com",
	}
	cookies = append(cookies, cookie)

	jar, _ := cookiejar.New(nil)
	
	inputUrl := strings.Replace(urlTemplate, "{}", day, -1)
	cookieUrl, _ := url.Parse(inputUrl)

	jar.SetCookies(cookieUrl, cookies)

	client := &http.Client{
		Jar: jar,
	}

	req, _ := http.NewRequest("GET", inputUrl, nil)
	resp, err := client.Do(req)

	if err != nil {
        return nil, errors.New("errored when sending request to the server")
    }

    defer resp.Body.Close()
    resp_body, _ := ioutil.ReadAll(resp.Body)

	return resp_body, nil
}

func ByteToInt(input []byte) ([]int, error) {
	var data []int

	r := bytes.NewReader(input)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		data = append(data, num)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

// ByteToString takes a byte slize and returns string slize for everyone scanned line
func ByteToString(input []byte) ([]string, error) {
	var data []string
	
	r := bytes.NewReader(input)
	scanner := bufio.NewScanner(r)

	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return data, nil
}