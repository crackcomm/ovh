package ovh

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

	"golang.org/x/net/context"
)

var baseURL = "https://eu.api.ovh.com/1.0"

func apiURL(format string, a ...interface{}) string {
	return fmt.Sprintf("%s%s", baseURL, fmt.Sprintf(format, a...))
}

type httpResponse struct {
	resp *http.Response
	err  error
}

type ovhRequest struct {
	opts *Options

	method    string
	query     string
	body      []byte
	timestamp string
}

var plus = []byte("+")

// requestSignature - Calculates a request signature and returns in following pattern
// "$1$" + SHA1_HEX(AS+"+"+CK+"+"+METHOD+"+"+QUERY+"+"+BODY+"+"+TSTAMP)
func (req *ovhRequest) signature() string {
	hash := sha1.New()
	hash.Write([]byte(req.opts.AppSecret))
	hash.Write(plus)
	hash.Write([]byte(req.opts.ConsumerKey))
	hash.Write(plus)
	hash.Write([]byte(req.method))
	hash.Write(plus)
	hash.Write([]byte(req.query))
	hash.Write(plus)
	hash.Write(req.body)
	hash.Write(plus)
	hash.Write([]byte(req.timestamp))
	return fmt.Sprintf("$1$%x", hash.Sum(nil))
}

func httpDo(ctx context.Context, opts *Options, method, url string, buffer io.Reader) (response *http.Response, err error) {
	var body []byte
	if opts.ConsumerKey != "" && buffer != nil {
		body, err = ioutil.ReadAll(buffer)
		if err != nil {
			return nil, err
		}
		buffer = bytes.NewBuffer(body)
	}

	req, err := http.NewRequest(method, url, buffer)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Ovh-Application", opts.AppKey)

	if opts.ConsumerKey != "" {
		r := &ovhRequest{
			opts:      opts,
			method:    method,
			query:     req.URL.String(),
			body:      body,
			timestamp: ovhTimestamp(),
		}

		req.Header.Set("X-Ovh-Timestamp", r.timestamp)
		req.Header.Set("X-Ovh-Consumer", opts.ConsumerKey)
		req.Header.Set("X-Ovh-Signature", r.signature())
	}

	transport := &http.Transport{}
	client := &http.Client{Transport: transport}

	respchan := make(chan *httpResponse, 1)
	go func() {
		resp, err := client.Do(req)
		respchan <- &httpResponse{resp: resp, err: err}
	}()

	select {
	case <-ctx.Done():
		transport.CancelRequest(req)
		<-respchan
		return nil, ctx.Err()
	case r := <-respchan:
		return r.resp, r.err
	}
}

var (
	onceTime sync.Once
	ovhTime  = new(struct {
		LocalUnixTime int64
		OvhTimestamp  int64
	})
)

func ovhTimestamp() string {
	onceTime.Do(func() {
		resp, err := http.Get("https://eu.api.ovh.com/1.0/auth/time")
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// Set response time
		ovhTime.LocalUnixTime = time.Now().Unix()

		t, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		// Set time on ovh server
		timestamp, err := strconv.Atoi(string(t))
		if err != nil {
			panic(err)
		}
		ovhTime.OvhTimestamp = int64(timestamp)
	})
	return fmt.Sprintf("%d", ovhTime.OvhTimestamp+(time.Now().Unix()-ovhTime.LocalUnixTime))
}

func unexpectedStatusError(resp *http.Response) error {
	body, _ := ioutil.ReadAll(resp.Body)
	return fmt.Errorf("Unexpected status code: %v (body=%s)", resp.StatusCode, body)
}
