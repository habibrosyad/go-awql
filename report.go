package awql

import (
	"encoding/csv"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Report download URL.
const reportURL = "https://adwords.google.com/api/adwords/reportdownload/" + version

type reportQuery struct {
	fields []string
	res    *http.Response
	next   []string
	reader *csv.Reader
}

var _ Rows = &reportQuery{}

// Create report query request.
func newReportQuery(c *Client, q string) (*reportQuery, error) {
	req, err := http.NewRequest(
		"POST", reportURL,
		strings.NewReader(url.Values{"__rdquery": {q}, "__fmt": {"CSV"}}.Encode()),
	)
	if err != nil {
		return nil, err
	}

	// https://developers.google.com/adwords/api/docs/guides/reporting#request_headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; param=value")
	req.Header.Add("Accept", "*/*")
	req.Header.Add("clientCustomerId", c.auth.CustomerId)
	req.Header.Add("developerToken", c.auth.DeveloperToken)
	req.Header.Add("skipReportHeader", "true")
	req.Header.Add("skipReportSummary", "true")
	req.Header.Add("skipColumnHeader", "true")

	res, err := c.client.Do(req)
	if err != nil || res.StatusCode != 200 {
		if res.StatusCode != 200 {
			defer res.Body.Close()
			var b []byte
			b, err = ioutil.ReadAll(res.Body)
			if err != nil {
				return nil, err
			}
			err = errors.New(string(b))
		}
		return nil, err
	}

	return &reportQuery{
		fields: getFields(q),
		res:    res,
		reader: csv.NewReader(res.Body),
	}, nil
}

// Iterator for query results.
func (r *reportQuery) Next() bool {
	var err error
	r.next, err = r.reader.Read()

	if err != nil {
		return false
	}
	return true
}

// Get a row.
func (r *reportQuery) Scan() (map[string]interface{}, error) {
	if len(r.next) == 0 {
		return nil, ErrNoRows
	}

	m := make(map[string]interface{})

	for i := range r.fields {
		m[r.fields[i]] = r.next[i]
	}
	return m, nil
}

// Close response body properly.
func (r *reportQuery) Close() {
	r.res.Body.Close()
}
