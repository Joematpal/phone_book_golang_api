package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Decode and marshals the data
func Decode(r *http.Request, v interface{}) error {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	rdr := ioutil.NopCloser(bytes.NewBuffer(buf))
	r.Body = ioutil.NopCloser(bytes.NewBuffer(buf))

	if err := json.NewDecoder(rdr).Decode(v); err != nil {
		fmt.Println("json", err)
		return err
	}
	return nil
}
