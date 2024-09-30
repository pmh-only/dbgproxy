package main

type LogStructure struct {
	ALog      string               `json:"a"`
	Request   RequestLogStructure  `json:"request"`
	Response  ResponseLogStructure `json:"response"`
	LatencyMS int                  `json:"latency_ms"`
}

type RequestLogStructure struct {
	RequestTime     string               `json:"request_time"`
	RequestMethod   string               `json:"request_method"`
	RequestPath     string               `json:"request_path"`
	RequestIP       string               `json:"request_ip"`
	RequestHeaders  []HeaderLogStructure `json:"request_headers"`
	RequestBody     string               `json:"request_body"`
	RequestBodySize string               `json:"request_body_size"`
}

type ResponseLogStructure struct {
	ResponseTime     string               `json:"response_time"`
	ResponseCode     int                  `json:"response_code"`
	ResponseHeaders  []HeaderLogStructure `json:"response_headers"`
	ResponseBody     string               `json:"response_body"`
	ResponseBodySize string               `json:"response_body_size"`
}

type HeaderLogStructure struct {
	HeaderKey   string `json:"header_key"`
	HeaderValue string `json:"header_value"`
}
