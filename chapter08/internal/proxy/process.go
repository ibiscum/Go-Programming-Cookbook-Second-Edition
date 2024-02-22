package proxy

import (
	"bytes"
	"net/http"
	"net/url"
)

// ProcessRequest modifies the request in accordance
// with Proxy settings
func (p *Proxy) ProcessRequest(r *http.Request) error {
	proxyURLRaw := p.BaseURL + r.URL.String()

	proxyURL, err := url.Parse(proxyURLRaw)
	if err != nil {
		return err
	}
	r.URL = proxyURL
	r.Host = proxyURL.Host
	r.RequestURI = ""
	return nil
}

// CopyResponse takes the client response and writes everything
// to the ResponseWriter in the original handler
func CopyResponse(w http.ResponseWriter, resp *http.Response) {
	var out bytes.Buffer
	_, err := out.ReadFrom(resp.Body)
	if err != nil {
		panic(err)
	}

	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	w.WriteHeader(resp.StatusCode)
	_, err = w.Write(out.Bytes())
	if err != nil {
		panic(err)
	}
}
