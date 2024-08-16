package gopress

import (
	"fmt"
	"net"
	"time"
)

type Response struct {
	Protocol      string            // The protocol used (e.g., "HTTP/1.1")
	StatusCode    int               // The HTTP status code (e.g., 200, 404, 500)
	StatusMessage string            // The status message corresponding to the status code (e.g., "OK", "Not Found")
	Headers       ResponseHeaders // A map of headers (e.g., "Content-Type": "application/json")
	Body          []byte            // The response body, typically the content returned by the server
	Client        *net.Conn
}

type ResponseHeaders struct {
  Date            string
  ContentType     string
  ContentLength   int
  Location        string
  Server          string
  SetCookie       string
  CacheControl    string
  Expires         string
  LastModified    string
  ETag            string
  Connection      string
  ContentEncoding string
  TransferEncoding string
  AccessControlAllowOrigin string
  XPoweredBy      string
  CustomHeaders   map[string]string 
}

func (res *Response) Send(body string, statusCode int) {
	res.StatusCode = statusCode;
	res.StatusMessage = httpStatusCodes[statusCode]
  res.Body = []byte(body);
  
  var head = fmt.Sprintf("%s %d %s\n", res.Protocol, res.StatusCode, res.StatusMessage);

  res.Headers.ContentLength = len(res.Body);
	res.Headers.ContentType = contentTypes["html"]
  res.Headers.Date = time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT");

  var response = ParseHeadersToPlainText(&res.Headers);

  var cry = fmt.Sprintf("%s%s\r\n%s", head, response, body);

  println(cry);
}

func (res *Response) AddCustomHeader(key string, value string) {
	if res.Headers.CustomHeaders == nil { res.Headers.CustomHeaders = make(map[string]string)}
	res.Headers.CustomHeaders[key] = value;
}

func (res *Response) Json(body any) {

}