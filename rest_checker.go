package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// Ref: https://stackoverflow.com/questions/19038598/how-can-i-pretty-print-json-using-go
// Ref: https://stackoverflow.com/questions/15323767/does-golang-have-if-x-in-construct-similar-to-python

//TODO LIST
// 1. Replace compute_port = 8774 and tenant_id
// 2. Map Service and Region by IDs
// 3. Do it for all Tenants

const KEYSTONE_GET_TOKEN_URL = "http://iam.savitestbed.ca:5000/v2.0/tokens"
const KEYSTONE_GET_ENDPOINT_URL = "http://iam.savitestbed.ca:35357/v2.0/endpoints"
const KEYSTONE_GET_SERVICE_URL = "http://iam.savitestbed.ca:35357/v2.0/OS-KSADM/services"

const CONTENT_TYPE = "application/json; charset=utf-8"

const TENANT_NAME = "xxx"
const USER_NAME = "xxx"
const PASSWORD = "xxx"

type AuthHeader struct {
	Auth struct {
		TenantName          string `json:"tenantName"`
		PasswordCredentials struct {
			Username string `json:"username"`
			Password string `json:"password"`
		} `json:"passwordCredentials"`
	} `json:"auth"`
}

func check(e error) {
	if e != nil {
		panic(e)
		return
	}
}

//dont do this, see above edit
func prettyprint(b []byte) ([]byte, error) {
	var out bytes.Buffer
	err := json.Indent(&out, b, "", "  ")
	return out.Bytes(), err
}

func get_token() string {

	// Create Auth Body for CURL request
	auth := AuthHeader{}
	auth.Auth.TenantName = TENANT_NAME
	auth.Auth.PasswordCredentials.Username = USER_NAME
	auth.Auth.PasswordCredentials.Password = PASSWORD

	// Convert it into bytes
	auth_bytes := new(bytes.Buffer)
	json.NewEncoder(auth_bytes).Encode(auth)

	// Make the POST call
	resp, err := http.Post(KEYSTONE_GET_TOKEN_URL, CONTENT_TYPE, auth_bytes)
	check(err)

	// Store it as string
	body, err := ioutil.ReadAll(resp.Body)
	check(err)

	// Get token
	token, err := jsonparser.GetString(body, "access", "token", "id")
	check(err)

	defer resp.Body.Close()

	return token

}

func get_request(url string, token string) ([]byte, string) {

	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	req, _ := http.NewRequest("GET", url, nil)
	//check(err)

	//req.Header.Add("User-Agent", "python-keystoneclient")
	req.Header.Add("X-Auth-Token", token)

	resp, err := client.Do(req)
	if err != nil {
		err_string := err.Error() + ": 500"
		return nil, err_string
	}

	defer resp.Body.Close()

	resp_body, _ := ioutil.ReadAll(resp.Body)

	//resp_body, _ = prettyprint(resp_body)
	//fmt.Printf("%s", resp_body)

	//fmt.Println(resp.Status)
	//fmt.Println(string(resp_body))

	return resp_body, resp.Status
}

func get_status(url string, token string) {
	//fmt.Println(url)
	_, status := get_request(url, token)

	fmt.Println(url + " -> " + status)
	//fmt.Println(string(resp))
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func main() {

	var token string = get_token()
	endpoints, _ := get_request(KEYSTONE_GET_ENDPOINT_URL, token)

	var url_list []string

	jsonparser.ArrayEach(endpoints, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		url, err := jsonparser.GetString(value, "publicurl")
		check(err)

		if strings.Contains(url, "(") == false && stringInSlice(url, url_list) == false {
			//url_list = append(url_list, (string(url)))
			get_status(url, token)
		}

	}, "endpoints")

	/*
		for _, element := range url_list {
			fmt.Println(element)
		}
	*/

}
