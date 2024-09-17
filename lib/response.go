package gopress

import (
	"encoding/json"
	"fmt"
	"net"
	"reflect"
	"strings"
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

//	Summary:
//		Build response with head, headers and body then sends it to the client.
func (res *Response) Send(body string, statusCode int) {
	res.Body = body;
	var headers = res.buildHeaders(statusCode, "txt");
	var response = fmt.Sprintf("%s\r\n%s", headers, body)
	var client = *res.client
  client.Write([]byte(response));
}

//	Summary:
//		Attempt to parse the given body to a json object.
func (res *Response) Json(body any, statusCode int) {
	jsonBody, err := json.Marshal(body);
	if err != nil { fmt.Println("Not able to parse the response body to json.") }

	newJsonBody := strings.ReplaceAll(string(jsonBody), "\\", "")
	/* newJsonBody = newJsonBody[1 : len(newJsonBody) - 1] */

	var headers = res.buildHeaders(statusCode, "json");
	var response = fmt.Sprintf("%s\r\n%s", headers, newJsonBody)
	res.Headers.ContentLength = len(string(jsonBody));

	var client = *res.client
	client.Write([]byte(response));
}

//	Summary:
//		Builds basic response struct.
func buildResponse(client *net.Conn, request Request) (*Response) {
	return &Response{
		client: client,
		Protocol: request.Protocol,
		Headers: ResponseHeaders{
			Server: "Gopress",
			Connection: request.Headers.Connection,
			CacheControl: "no-cache",
			AccessControlAllowOrigin: "*",
			XPoweredBy: "Go",
		},
	}
}

//	Summary:
//		Fill the header with additional information and converts it to plain text.
func (res *Response) buildHeaders(statusCode int, contentType string) (string) {
	StatusMessage := httpStatusCodes[statusCode]
	
	res.Headers.ContentType = contentTypes[contentType]
  res.Headers.ContentLength = len([]byte(res.Body));
	res.Headers.Date = time.Now().Format("Mon, 02 Jan 2006 15:04:05 GMT");

  var head = fmt.Sprintf("%s %d %s\n", strings.TrimSpace(res.Protocol), statusCode, StatusMessage);
  var headers = res.Headers.toPlainText();

	return fmt.Sprintf("%s%s", head, headers);
}

//	Summary:
//		Convert the response headers to plain-text to be sent back to the client.
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