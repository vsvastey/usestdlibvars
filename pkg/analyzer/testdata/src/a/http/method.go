// Code generated by usestdlibvars, DO NOT EDIT.

package http_test

import "net/http"

var (
	_ = "CONNECT"
	_ = "DELETE"
	_ = "GET"
	_ = "HEAD"
	_ = "OPTIONS"
	_ = "PATCH"
	_ = "POST"
	_ = "PUT"
	_ = "TRACE"
)

const (
	_ = "CONNECT"
	_ = "DELETE"
	_ = "GET"
	_ = "HEAD"
	_ = "OPTIONS"
	_ = "PATCH"
	_ = "POST"
	_ = "PUT"
	_ = "TRACE"
)

func _() {
	_, _ = http.NewRequest("CONNECT", "", http.NoBody) // want `"CONNECT" can be replaced by http\.MethodConnect`
	_, _ = http.NewRequest("DELETE", "", http.NoBody)  // want `"DELETE" can be replaced by http\.MethodDelete`
	_, _ = http.NewRequest("GET", "", http.NoBody)     // want `"GET" can be replaced by http\.MethodGet`
	_, _ = http.NewRequest("HEAD", "", http.NoBody)    // want `"HEAD" can be replaced by http\.MethodHead`
	_, _ = http.NewRequest("OPTIONS", "", http.NoBody) // want `"OPTIONS" can be replaced by http\.MethodOptions`
	_, _ = http.NewRequest("PATCH", "", http.NoBody)   // want `"PATCH" can be replaced by http\.MethodPatch`
	_, _ = http.NewRequest("POST", "", http.NoBody)    // want `"POST" can be replaced by http\.MethodPost`
	_, _ = http.NewRequest("PUT", "", http.NoBody)     // want `"PUT" can be replaced by http\.MethodPut`
	_, _ = http.NewRequest("TRACE", "", http.NoBody)   // want `"TRACE" can be replaced by http\.MethodTrace`
}

func _() {
	_, _ = http.NewRequestWithContext(nil, "CONNECT", "", http.NoBody) // want `"CONNECT" can be replaced by http\.MethodConnect`
	_, _ = http.NewRequestWithContext(nil, "DELETE", "", http.NoBody)  // want `"DELETE" can be replaced by http\.MethodDelete`
	_, _ = http.NewRequestWithContext(nil, "GET", "", http.NoBody)     // want `"GET" can be replaced by http\.MethodGet`
	_, _ = http.NewRequestWithContext(nil, "HEAD", "", http.NoBody)    // want `"HEAD" can be replaced by http\.MethodHead`
	_, _ = http.NewRequestWithContext(nil, "OPTIONS", "", http.NoBody) // want `"OPTIONS" can be replaced by http\.MethodOptions`
	_, _ = http.NewRequestWithContext(nil, "PATCH", "", http.NoBody)   // want `"PATCH" can be replaced by http\.MethodPatch`
	_, _ = http.NewRequestWithContext(nil, "POST", "", http.NoBody)    // want `"POST" can be replaced by http\.MethodPost`
	_, _ = http.NewRequestWithContext(nil, "PUT", "", http.NoBody)     // want `"PUT" can be replaced by http\.MethodPut`
	_, _ = http.NewRequestWithContext(nil, "TRACE", "", http.NoBody)   // want `"TRACE" can be replaced by http\.MethodTrace`
}

func _() {
	_ = &http.Request{
		Method: "CONNECT", // want `"CONNECT" can be replaced by http\.MethodConnect`
	}
	_ = &http.Request{
		Method: "DELETE", // want `"DELETE" can be replaced by http\.MethodDelete`
	}
	_ = &http.Request{
		Method: "GET", // want `"GET" can be replaced by http\.MethodGet`
	}
	_ = &http.Request{
		Method: "HEAD", // want `"HEAD" can be replaced by http\.MethodHead`
	}
	_ = &http.Request{
		Method: "OPTIONS", // want `"OPTIONS" can be replaced by http\.MethodOptions`
	}
	_ = &http.Request{
		Method: "PATCH", // want `"PATCH" can be replaced by http\.MethodPatch`
	}
	_ = &http.Request{
		Method: "POST", // want `"POST" can be replaced by http\.MethodPost`
	}
	_ = &http.Request{
		Method: "PUT", // want `"PUT" can be replaced by http\.MethodPut`
	}
	_ = &http.Request{
		Method: "TRACE", // want `"TRACE" can be replaced by http\.MethodTrace`
	}
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == "CONNECT" { // want `"CONNECT" can be replaced by http\.MethodConnect`
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == "DELETE" { // want `"DELETE" can be replaced by http\.MethodDelete`
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == "GET" { // want `"GET" can be replaced by http\.MethodGet`
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == "HEAD" { // want `"HEAD" can be replaced by http\.MethodHead`
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == "OPTIONS" { // want `"OPTIONS" can be replaced by http\.MethodOptions`
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == "PATCH" { // want `"PATCH" can be replaced by http\.MethodPatch`
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == "POST" { // want `"POST" can be replaced by http\.MethodPost`
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == "PUT" { // want `"PUT" can be replaced by http\.MethodPut`
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == "TRACE" { // want `"TRACE" can be replaced by http\.MethodTrace`
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == http.MethodConnect {
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == http.MethodDelete {
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == http.MethodGet {
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == http.MethodHead {
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == http.MethodOptions {
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == http.MethodPatch {
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == http.MethodPost {
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == http.MethodPut {
		return nil
	}
	return nil
}

func _() error {
	resp, err := http.DefaultClient.Do(&http.Request{})
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.Request.Method == http.MethodTrace {
		return nil
	}
	return nil
}
