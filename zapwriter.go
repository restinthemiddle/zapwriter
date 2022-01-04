package zapwriter

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Writer struct {
	Logger *zap.Logger
}

func (w Writer) LogRequest(request *http.Request) (err error) {
	query := ""
	rawQuery := request.URL.RawQuery
	if len(rawQuery) > 0 {
		query = fmt.Sprintf("?%s", rawQuery)
	}

	headers := make([]string, 0)
	for name, values := range request.Header {
		for _, value := range values {
			headers = append(headers, fmt.Sprintf("%s: %s", name, value))
		}
	}

	bodyString := ""
	if request.ContentLength > 0 {
		bodyBytes, err := ioutil.ReadAll(request.Body)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		bodyString = string(bodyBytes)
	}

	w.Logger.Info("failed to fetch URL",
		zap.String("request_method", request.Method),
		zap.String("scheme", request.URL.Scheme),
		zap.String("http_host", request.URL.Host),
		zap.String("request", request.URL.Path),
		zap.String("args", query),
		zap.Strings("headers", headers),
		zap.String("body", bodyString),
	)

	return nil

	// requestRow := row{time.Now().Format(time.RFC3339Nano), request.Method, request.URL.Scheme, request.URL.Host, request.URL.Path, query, headers, bodyString}

	// m, err := json.Marshal(requestRow)
	// if err != nil {
	// 	log.Fatal(err)
	// 	panic(err)
	// }

	// fmt.Println(string(m))

	// return err
}

func (w Writer) LogResponse(response *http.Response) (err error) {
	query := ""
	rawQuery := response.Request.URL.RawQuery
	if len(rawQuery) > 0 {
		query = fmt.Sprintf("?%s", rawQuery)
	}

	headers := make([]string, 0)
	for name, values := range response.Request.Header {
		for _, value := range values {
			headers = append(headers, fmt.Sprintf("%s: %s", name, value))
		}
	}

	bodyString := ""
	if response.Request.ContentLength > 0 {
		bodyBytes, err := ioutil.ReadAll(response.Request.Body)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}

		response.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

		bodyString = string(bodyBytes)
	}

	w.Logger.Info("failed to fetch URL",
		zap.String("request_method", response.Request.Method),
		zap.String("scheme", response.Request.URL.Scheme),
		zap.String("http_host", response.Request.URL.Host),
		zap.String("request", response.Request.URL.Path),
		zap.String("args", query),
		zap.Strings("headers", headers),
		zap.String("body", bodyString),
		zap.Time("roundtrip_start", response.Request.Context().Value("roundTripStart").(time.Time)),
		zap.Time("roundtrip_end", response.Request.Context().Value("roundTripEnd").(time.Time)),
	)

	return nil

	// title := fmt.Sprintf("RESPONSE - Code: %d\n", response.StatusCode)

	// headers := ""
	// for key, element := range response.Header {
	// 	headers += fmt.Sprintf("%s: %s\n", key, element)
	// }

	// bodyString := ""
	// if response.ContentLength > 0 {
	// 	bodyBytes, err := ioutil.ReadAll(response.Body)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		panic(err)
	// 	}

	// 	response.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	// 	bodyString = fmt.Sprintf("Content: %s\n", string(bodyBytes))
	// }

	// log.Printf("%s%s%s", title, headers, bodyString)

	// return err
}
