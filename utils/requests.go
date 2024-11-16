package utils

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"go-jocy/config"
)

type Request struct {
	client    *resty.Client
	AuthToken string
}

func New(AuthToken string) *Request {
	client := resty.New()
	client.Debug = true

	client.SetHeaderVerbatim("User-Agent", "Dart/2.17 (dart:io)")
	// client.SetHeaderVerbatim("ts", "1720943886093")
	client.SetHeaderVerbatim("ts", "1720943886")
	client.SetHeaderVerbatim("x-version", "2020-09-17")
	client.SetHeaderVerbatim("appid", "4150439554430529")
	client.SetHeaderVerbatim("authentication", "HPNGF8PeCIjBOsHyrnnFuRhGF2immEFK7SOOT1D4+is+BNfhx82bTZrRYJ6rswOBSStD6M2oFrvkfQtSL6xGCOAxOx42pB34/ZyV+5TntqS6hnqAt4Xn/wHOWItBdU0/qJiwOg99FjdD3UwXAaZTig==")

	if AuthToken != "" {
		client.SetHeaderVerbatim("x-token", AuthToken)
	}
	return &Request{
		client:    client,
		AuthToken: AuthToken,
	}
}

func (r *Request) Get(url string, params map[string]string) (*resty.Response, error) {
	config.GinLOG.Debug("GET: " + url)
	config.GinLOG.Debug(fmt.Sprintf("params: %v", params))

	return r.client.R().
		SetQueryParams(params).
		Get(url)
}

func (r *Request) Post(url string) (*resty.Response, error) {
	return r.client.R().
		Post(url)
}
