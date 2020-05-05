package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// b2_authorize_account
// https://www.backblaze.com/b2/docs/b2_authorize_account.html
func authAccount(idPrivate string, keyPrivate string) (loginToken string) {
	url := fmt.Sprintf("%s/b2api/v2/b2_authorize_account", apiUrl)
	timeout := 5 * time.Second
	client := http.Client{Timeout: timeout}

	// create request, attach basic auth
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("!!! could not create auth request, err:", err)
	}
	req.SetBasicAuth(idPrivate, keyPrivate)

	// make the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("!!! could not make auth request, err:", err)
	}
	defer resp.Body.Close()

	// was this ok?
	if resp.StatusCode != 200 {
		log.Fatal("!!! didn't get 200 while logging in, got:", resp.StatusCode)
	}

	// read in the response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("!!! malformed auth response? err:", err)
	}
	j := string(respBody)

	// get download token
	loginTokenPath := "authorizationToken"
	token := gjson.Get(j, loginTokenPath).Str
	return token
}

// b2_get_download_authorization
// https://www.backblaze.com/b2/docs/b2_get_download_authorization.html
func getAuthorization(loginAuthToken string, bucket string) (downloadToken string) {
	url := fmt.Sprintf("%s/b2api/v2/b2_get_download_authorization", apiUrl)
	timeout := 5 * time.Second
	client := http.Client{Timeout: timeout}

	// build the object to POST
	type postObject struct {
		BucketId               string `json:"bucketId"`
		FileNamePrefix         string `json:"fileNamePrefix"`
		ValidDurationInSeconds int    `json:"validDurationInSeconds"`
	}
	reqObject := postObject{
		BucketId:               bucket,
		ValidDurationInSeconds: 604800, // 1 week in seconds
		FileNamePrefix:         "",
	}

	// create the json object
	jsonObject, err := json.Marshal(reqObject)
	if err != nil {
		log.Fatal("!!! malformed token reqObject? err:", err)
	}

	// create request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonObject))
	if err != nil {
		log.Fatal("!!! could not create token request, err:", err)
	}
	req.Header.Set("Authorization", loginAuthToken)

	// make the request
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal("!!! could not make token request, err:", err)
	}
	defer resp.Body.Close()

	// was this ok?
	if resp.StatusCode != 200 {
		log.Fatal("!!! didn't get 200 while getting download token, got:", resp.StatusCode)
	}

	// read in the response
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("!!! malformed token response? err:", err)
	}
	j := string(respBody)

	// get download token
	loginTokenPath := "authorizationToken"
	token := gjson.Get(j, loginTokenPath).Str
	return token
}
