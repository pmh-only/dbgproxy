package main

import (
	"fmt"
	"log"
	"time"

	realip "github.com/ferluci/fast-realip"
	"github.com/sugawarayuuta/sonnet"
	"github.com/valyala/fasthttp"
)

func createProxyHandler(client *fasthttp.HostClient) func(ctx *fasthttp.RequestCtx) {
	return func(ctx *fasthttp.RequestCtx) {
		requestTime := time.Now()
		requestBody := ctx.Request.Body()

		responseDump := fasthttp.AcquireResponse()
		defer fasthttp.ReleaseResponse(responseDump)

		err := client.Do(&ctx.Request, responseDump)
		if err != nil {
			log.Println(err)
			return
		}
		responseTime := time.Now()
		responseDump.CopyTo(&ctx.Response)

		// for debug purpose
		if DEBUG_DISABLE_LOGGING {
			return
		}

		requestHeaders := []HeaderLogStructure{}
		responseHeaders := []HeaderLogStructure{}

		responseBody := responseDump.Body()

		ctx.Request.Header.VisitAll(func(key, value []byte) {
			requestHeaders = append(requestHeaders, HeaderLogStructure{
				HeaderKey:   string(key),
				HeaderValue: string(value),
			})
		})

		ctx.Response.Header.VisitAll(func(key, value []byte) {
			responseHeaders = append(responseHeaders, HeaderLogStructure{
				HeaderKey:   string(key),
				HeaderValue: string(value),
			})
		})

		latency := int(responseTime.UnixMilli()) - int(requestTime.UnixMilli())
		realIp := realip.FromRequest(ctx)

		alog := fmt.Sprintf("%s %s %d %dms %s %s '%s'",
			string(ctx.Method()),
			string(ctx.Path()),
			ctx.Response.StatusCode(),
			latency,
			formatByte(len(responseBody)),
			realIp,
			string(ctx.UserAgent()))

		logContent := LogStructure{
			ALog: alog,
			Request: RequestLogStructure{
				RequestTime:     requestTime.UTC().String(),
				RequestMethod:   string(ctx.Method()),
				RequestPath:     string(ctx.Path()),
				RequestIP:       ctx.RemoteIP().String(),
				RequestHeaders:  requestHeaders,
				RequestBody:     string(requestBody),
				RequestBodySize: formatByte(len(requestBody)),
			},
			Response: ResponseLogStructure{
				ResponseTime:     responseTime.UTC().String(),
				ResponseCode:     ctx.Response.StatusCode(),
				ResponseHeaders:  responseHeaders,
				ResponseBody:     string(responseBody),
				ResponseBodySize: formatByte(len(responseBody)),
			},
			LatencyMS: latency,
		}

		logEncoded, err := sonnet.Marshal(logContent)
		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(logEncoded))
	}
}
