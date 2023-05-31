package sms

import (
	"cl/pkg/bootstrap/http/misc/response"
	"log"
	"net/http"
	"net/url"
	"time"
)

// SendSMS ...
func (p *provider) Send(phone, text string) (err error) {
	// Init http client for make request to sms center.
	// Set a timeout of 25+ seconds to ensure that the
	// response from the SMS center has been processed.
	var client = &http.Client{
		Timeout: 40 * time.Second,
	}

	// Сollect all the parameters for the request from
	// the config file, set the text and phone number.
	params := map[string]string{
		"user": p.config.SMS.User,
		"pass": p.config.SMS.Password,
		"from": p.config.SMS.From,
		"to":   phone,
		"txt":  text,
	}

	urlLink, err := url.Parse(p.config.SMS.URL)
	if err != nil {
		p.logger.Error(err)
		return err
	}

	// Create an empty instance of url.Values.
	// Adding previously init params. Encode other characters.
	urlParams := url.Values{}
	for k, v := range params {
		urlParams.Add(k, v)
	}
	urlLink.RawQuery = urlParams.Encode()

	// Сreate a new request. Specify the GET method,
	// convert url.URL to a string, pass it as the second arg.
	// The third arg is nil because the GET request has no body.
	request, err := http.NewRequest(http.MethodGet, urlLink.String(), nil)
	if err != nil {
		p.logger.Error(err)
		return err
	}

	resp, err := client.Do(request)
	if err != nil {
		p.logger.Error(err)
		return err
	}

	log.Println(resp.Request.URL, "\n", resp)

	if resp.StatusCode != http.StatusOK {
		p.logger.Error(err, resp)
		return response.ErrSomethingWentWrong
		// return errors.New("code, Status: %v, %s", response.StatusCode, response.Status)
	}

	return nil
}
