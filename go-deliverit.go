package godeliverit

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

const (
	defaultSendTimeout = time.Second * 30
	baseURL            = "https://cloud.deliverit.com.au"
	authURL            = "reports/v1.1/api/web/v1/auths"
	storeURL           = "reports/v1.1/api/web/v1/stores"
	storeordersURL     = "reports/v1.1/api/web/v1/orders"
)

// Deliverit The main struct of this package
type Deliverit struct {
	ClientSecret string
	Timeout      time.Duration
}

// NewClient will create a Deliverit client with default values
func NewClient() *Deliverit {
	return &Deliverit{
		Timeout: defaultSendTimeout,
	}
}

func checkRedirectFunc(req *http.Request, via []*http.Request) error {
	if req.Header.Get("Authorization") == "" {
		req.Header.Add("Authorization", via[0].Header.Get("Authorization"))
	}
	return nil
}

// GetAuths will return the authorizations of the Account
func (v *Deliverit) GetAuths(username string, password string) (*Auth, error) {
	client := &http.Client{}
	client.CheckRedirect = checkRedirectFunc

	u, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to build Deliverit auths: %v", err)
	}

	u.Path = authURL
	urlStr := fmt.Sprintf("%v", u)

	r, err := http.NewRequest("GET", urlStr, nil)

	secret := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", username, password)))

	r.Header = http.Header(make(map[string][]string))
	r.Header.Set("Accept", "application/json")
	r.Header.Set("Authorization", fmt.Sprintf("Basic %v", secret))

	res, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("Failed to call Deliverit %v: %v", authURL, err)
	}

	if res.StatusCode == 200 {
		rawResBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("Failed to read Deliverit %v: %v", authURL, err)
		}
		//test
		var resp Auth
		err = json.Unmarshal(rawResBody, &resp)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal Deliverit %v: %v", authURL, err)
		}
		return &resp, nil

	}
	return nil, fmt.Errorf("Failed to get Deliverit %v: %s", authURL, res.Status)
}

// GetStores will return the stores of the Token
func (v *Deliverit) GetStores(auth *Auth) (Stores, error) {
	client := &http.Client{}
	client.CheckRedirect = checkRedirectFunc

	u, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to build Deliverit %v: %v", storeURL, err)
	}

	u.Path = storeURL
	urlStr := fmt.Sprintf("%v", u)

	r, err := http.NewRequest("GET", urlStr, nil)

	r.Header = http.Header(make(map[string][]string))
	r.Header.Set("Accept", "application/json")

	data := url.Values{}
	data.Add("key", strconv.Itoa(auth.Key))
	data.Add("token", auth.Token)
	r.URL.RawQuery = data.Encode()

	res, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("Failed to call Deliverit %v: %v", storeURL, err)
	}

	if res.StatusCode == 200 {
		rawResBody, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, fmt.Errorf("Failed to read Deliverit %v: %v", storeURL, err)
		}
		//test
		var resp Stores
		err = json.Unmarshal(rawResBody, &resp)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal Deliverit %v: %v", storeURL, err)
		}
		return resp, nil

	}
	return nil, fmt.Errorf("Failed to get Deliverit %v: %s", storeURL, res.Status)
}

// GetOrders will return the orders of the Token and Store
func (v *Deliverit) GetOrders(auth *Auth, store *Store, startDate string, endDate string) (Orders, error) {
	client := &http.Client{}
	client.CheckRedirect = checkRedirectFunc

	u, err := url.ParseRequestURI(baseURL)
	if err != nil {
		return nil, fmt.Errorf("Failed to build Deliverit %v: %v", storeordersURL, err)
	}

	u.Path = storeordersURL
	urlStr := fmt.Sprintf("%v", u)

	r, err := http.NewRequest("GET", urlStr, nil)

	r.Header = http.Header(make(map[string][]string))
	r.Header.Set("Accept", "application/json")

	data := url.Values{}
	data.Add("storeid", strconv.Itoa(store.ID))
	data.Add("key", strconv.Itoa(auth.Key))
	data.Add("token", auth.Token)
	data.Add("startdate", startDate) //"2019-07-18"
	data.Add("enddate", endDate)
	r.URL.RawQuery = data.Encode()

	res, err := client.Do(r)
	if err != nil {
		return nil, fmt.Errorf("Failed to call Deliverit %v: %v", storeordersURL, err)
	}

	if res.StatusCode == 200 {

		rawResBody, err := ioutil.ReadAll(res.Body)

		if err != nil {
			return nil, fmt.Errorf("Failed to read Deliverit %v: %v", storeordersURL, err)
		}
		//test
		var resp Orders
		err = json.Unmarshal(rawResBody, &resp)
		if err != nil {
			return nil, fmt.Errorf("Failed to unmarshal Deliverit %v: %v", storeordersURL, err)
		}
		return resp, nil

	}
	return nil, fmt.Errorf("Failed to get Deliverit %v: %s", storeordersURL, res.Status)
}
