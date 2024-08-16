package gopress

// Global variables for the package.

var server HttpServer
var httpMethods HttpMethodHandlers
var httpStatusCodes = map[int]string{
        200: "OK",
        201: "Created",
        202: "Accepted",
        204: "No Content",
        301: "Moved Permanently",
        302: "Found",
        304: "Not Modified",
        400: "Bad Request",
        401: "Unauthorized",
        403: "Forbidden",
        404: "Not Found",
        405: "Method Not Allowed",
        500: "Internal Server Error",
        502: "Bad Gateway",
        503: "Service Unavailable",
    }
		
var contentTypes = map[string]string{
        "html":  "text/html",
        "css":   "text/css",
        "js":    "application/javascript",
        "json":  "application/json",
        "xml":   "application/xml",
        "jpeg":  "image/jpeg",
        "jpg":   "image/jpeg",
        "png":   "image/png",
        "gif":   "image/gif",
        "svg":   "image/svg+xml",
        "pdf":   "application/pdf",
        "zip":   "application/zip",
        "txt":   "text/plain",
        "csv":   "text/csv",
        "mp4":   "video/mp4",
        "mp3":   "audio/mpeg",
        "ogg":   "audio/ogg",
        "wav":   "audio/wav",
        "woff":  "font/woff",
        "woff2": "font/woff2",
    }