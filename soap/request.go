package soap

import (
	"bufio"
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"text/template"
)

type Env struct {
	Namespace string
	Action    string
}

const ENVELOPE = `<?xml version="1.0" encoding="UTF-8"?>
  <s:Envelope xmlns:s="http://schemas.xmlsoap.org/soap/envelope/"
    xmlns:u="{{ .Namespace }}"
    xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/">
    <s:Body>
      <u:{{ .Action }}/>
    </s:Body>
  </s:Envelope>`

func DoRequest(url string, action string) (Envelope, error) {
	params := strings.Split(action, "#")
	var envelope Envelope

	if len(params) == 2 {

		var b bytes.Buffer
		bufWriter := bufio.NewWriter(&b)

		envTemplate, err := template.New("envelope").Parse(ENVELOPE)
		if err != nil {
			return envelope, err
		}

		if err := envTemplate.Execute(bufWriter, Env{Namespace: params[0], Action: params[1]}); err != nil {
			return envelope, err
		}

		if err := bufWriter.Flush(); err != nil {
			return envelope, err
		}

		fmt.Println(string(b.Bytes()))

		client := &http.Client{}
		req, err := http.NewRequest("POST", url, bytes.NewReader(b.Bytes()))
		if err != nil {
			return envelope, err
		}
		req.Header.Add("Content-Type", "text/xml; charset=utf-8")
		req.Header.Add("SOAPAction", action)

		resp, err := client.Do(req)
		if err != nil {
			return envelope, err
		}

		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return envelope, err
		}

		if err := xml.Unmarshal(data, &envelope); err != nil {
			return envelope, err
		}
		return envelope, err
	} else {
		return envelope, errors.New(fmt.Sprintf("Action \"%s\" is malformed."))
	}
}
