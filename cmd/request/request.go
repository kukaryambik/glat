package request

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Params of request
type Params map[string]string

// Conf - request config.
type Conf struct {
	NoTLS                  bool
	Type, Proto, Host, URI string
	Forms, Headers         Params
}

// Split args
func Split(p Params, s string, sep string) Params {
	for _, i := range strings.Split(s, sep) {
		t := strings.Split(i, "=")
		p[t[0]] = t[1]
	}
	return p
}

// Send - main functuon
func Send(opts Conf) (string, error) {

	// Use http or https
	if opts.NoTLS {
		opts.Proto = "http://"
	}

	if opts.Proto == "" {
		opts.Proto = "https://"
	}

	if opts.Type == "" {
		opts.Type = "GET"
	}

	// Make full url
	u := opts.Proto + opts.Host + opts.URI

	val := url.Values{}
	for key, value := range opts.Forms {
		val.Add(key, value)
	}

	// Make request
	client := &http.Client{}
	conf, _ := http.NewRequest(opts.Type, u, strings.NewReader(val.Encode()))

	for key, value := range opts.Headers {
		conf.Header.Add(key, value)
	}

	req, err := client.Do(conf)
	if err != nil {
		return "", err
	}

	// End request
	defer req.Body.Close()

	resp, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	return string(resp), nil

}
