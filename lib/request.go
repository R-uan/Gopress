package gopress

import (
	"gopress/lib/internal"
	"net"
	"strings"
)

type Request struct {
	Path    string
	Method  string
	Headers RequestHeaders
	Body    string
}

type RequestHeaders struct {
	Host            string `json:"host"`
	UserAgent       string `json:"user_agent"`
	Accept          string `json:"accept"`
	AcceptLanguage  string `json:"accept_language"`
	AcceptEncoding  string `json:"accept_encoding"`
	Connection      string `json:"connection"`
	ContentType     string `json:"content_type"`
	ContentLength   int64  `json:"content_length"`
	Authorization   string `json:"authorization"`
	Cookie          string `json:"cookie"`
	Referer         string `json:"referer"`
	CacheControl    string `json:"cache_control"`
	UpgradeInsecure string `json:"upgrade_insecure_requests"`
	IfModifiedSince string `json:"if_modified_since"`
	IfNoneMatch     string `json:"if_none_match"`
	Origin          string `json:"origin"`
	Pragma          string `json:"pragma"`
	XRequestedWith  string `json:"x_requested_with"`
	XForwardedFor   string `json:"x_forwarded_for"`
	XRealIP         string `json:"x_real_ip"`
	Range           string `json:"range"`
}

//	Summary:
//		Maps the http method, path, headers and body to a Request struct.
//
//	Returns:
//		Returns the pointer to said struct.
func extractRequest(rawRequest string) (*Request) {
	
	var rawRequestArray = strings.Split(rawRequest, "\n");
	var requestHead = strings.SplitN(rawRequestArray[0], " ", 3);
	
	var headers = make(map[string]string);
	for line := range rawRequestArray {
		var parts = strings.SplitN(rawRequestArray[line], ":", 2);
		if len(parts) == 2 { headers[strings.TrimSpace(parts[0])] = strings.TrimSpace(parts[1]) }
	}
	
	var Request = Request{
		Path: requestHead[1],
		Method: requestHead[0],
		Body: strings.SplitN(rawRequest, "\r\n", 2)[1],
		Headers: RequestHeaders{
			Host:            headers["Host"],
			UserAgent:       headers["User-Agent"],
			Accept:          headers["Accept"],
			AcceptLanguage:  headers["Accept-Language"],
			AcceptEncoding:  headers["Accept-Encoding"],
			Connection:      headers["Connection"],
			ContentType:     headers["Content-Type"],
			ContentLength:   internal.ParseContentLength(headers["Content-Length"]),
			Authorization:   headers["Authorization"],
			Cookie:          headers["Cookie"],
			Referer:         headers["Referer"],
			CacheControl:    headers["Cache-Control"],
			UpgradeInsecure: headers["Upgrade-Insecure-Requests"],
			IfModifiedSince: headers["If-Modified-Since"],
			IfNoneMatch:     headers["If-None-Match"],
			Origin:          headers["Origin"],
			Pragma:          headers["Pragma"],
			XRequestedWith:  headers["X-Requested-With"],
			XForwardedFor:   headers["X-Forwarded-For"],
			XRealIP:         headers["X-Real-IP"],
			Range:           headers["Range"],
		},
	};
	
	return &Request;
}

//	Summary:
//		Function that handles individual accepted requests on it s own goroutine.
func handleRequests(client net.Conn) {
	defer client.Close();
	var rawRequest = make([]byte, 1024);
	client.Read(rawRequest);

	var req *Request = extractRequest(string(rawRequest));
	var res *Response = buildResponse(&client);
	
	middlewarePipeLine(req, res)
	routing(req, res);
}