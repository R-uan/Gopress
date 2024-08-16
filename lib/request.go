package gopress

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