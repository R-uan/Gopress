package gopress

import (
	"fmt"
	"net"
	"reflect"
	"time"
)

type Response struct {
	client        *net.Conn
	Protocol      string
	Headers       ResponseHeaders
	Body          string
}

type ResponseHeaders struct {
	Date                     string
	ContentType              string
	ContentLength            int
	Location                 string
	Server                   string
	SetCookie                string
	CacheControl             string
	Expires                  string
	LastModified             string
	ETag                     string
	Connection               string
	ContentEncoding          string
	TransferEncoding         string
	AccessControlAllowOrigin string
	XPoweredBy               string
	CustomHeaders            map[string]string
}

func (res *Response) Send(body string, statusCode int) {
	res.Body = body;
	var headers = res.buildHeaders(statusCode);
	var response = fmt.Sprintf("%s%s", headers, body)

	var client = *res.client
  client.Write([]byte(response));
}

func (res *Response) Json(body any) {
	
}

//	Summary:
//		Builds basic response struct.
func buildResponse(client *net.Conn) (*Response) {
	return &Response{
		client: client,
		Protocol: "HTTP/1.1",
		Headers: ResponseHeaders{
			Server: "Gopress/0.1",
			Connection: "keep-alive",
			CacheControl: "no-cache",
			AccessControlAllowOrigin: "*",
			XPoweredBy: "Go(Golang)",
		},
	}
}

func (res *Response) buildHeaders(statusCode int) (string) {
	StatusMessage := httpStatusCodes[statusCode]
	
	res.Headers.ContentType = contentTypes["txt"]
  res.Headers.ContentLength = len([]byte(res.Body));
  res.Headers.Date = time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT");

  var head = fmt.Sprintf("%s %d %s\n", res.Protocol, statusCode, StatusMessage);
  var headers = res.Headers.toPlainText();

	return fmt.Sprintf("%s%s\r\n", head, headers);
}

func (res *Response) AddCustomHeader(key string, value string) {
  if res.Headers.CustomHeaders == nil { res.Headers.CustomHeaders = make(map[string]string)}
  res.Headers.CustomHeaders[key] = value;
}

func (headers *ResponseHeaders) toPlainText() string {
	value := reflect.ValueOf(headers)
	if value.Kind() == reflect.Ptr { value = value.Elem() }
	typ := value.Type()

	var result string

	ignoreEmpty := func(val reflect.Value) bool {
		switch val.Kind() {
		case reflect.String:
			return val.String() == ""
		case reflect.Int:
			return val.Int() == 0
		case reflect.Map:
			return val.Len() == 0
		case reflect.Slice:
			return val.Len() == 0
		default:
			return false
		}
	}

	for i := 0; i < value.NumField(); i++ {
		field := typ.Field(i)
		fieldValue := value.Field(i)

		if ignoreEmpty(fieldValue) { continue }

		if fieldValue.Kind() == reflect.Map {
			result += fmt.Sprintf("%s:\n", field.Name)
			for _, key := range fieldValue.MapKeys() {
				result += fmt.Sprintf("  %s: %v\n", key, fieldValue.MapIndex(key))
			}
		} else {
			value, exists := mapHeaders[field.Name]
			if exists {
				result += fmt.Sprintf("%s: %v\n", value, fieldValue)
			}
		}
	}

	return result
}