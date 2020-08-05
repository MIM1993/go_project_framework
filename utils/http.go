package utils

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"
)

func HTTPGet(url string) ([]byte, bool) {
	res, err := http.Get(url)
	var ret []byte
	if err != nil {
		logrus.Warnf("http get %s failed:%s", url, err.Error())
		return []byte(err.Error()), false
	}
	defer res.Body.Close()

	ret, _ = ioutil.ReadAll(res.Body)

	return ret, true
}

func HTTPPost(url string, data []byte) ([]byte, bool) {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Header.Add("Content-Type", "application/json")
	logrus.Debug(string(data))

	resp, err := client.Do(req)
	if err != nil {
		logrus.Warn(err.Error())
		return []byte(err.Error()), false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	logrus.Debug(string(body))
	return body, true
}

func HTTPPostWithTimeout(url string, data []byte, timeout int) ([]byte, bool) {
	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Close = true
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		logrus.Warn(err.Error())
		return []byte(err.Error()), false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	return body, true
}

func HTTPPostJson(url string, data []byte, timeout int) ([]byte, error) {
	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(data))
	req.Close = true
	req.Header.Add("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func HTTPRequestForm(url, method, token string, params map[string]string) ([]byte, bool) {
	var ret []byte
	form, contentType, _ := CreateMultipartFormBody(params)
	if form == nil {
		return ret, false
	}

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, form)
	if token != "" {
		req.Header.Add("token", token)
	}
	req.Close = true
	req.Header.Set("Content-Type", contentType)

	resp, err := client.Do(req)
	if err != nil {
		logrus.Warn(err.Error())
		return ret, false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		return body, false
	}

	return body, true
}

func HTTPRequestFormWithTimeout(url, method, token string, timeout int, params map[string]string) ([]byte, bool) {
	var ret []byte
	form, contentType, _ := CreateMultipartFormBody(params)
	if form == nil {
		return ret, false
	}

	client := &http.Client{
		Timeout: time.Second * time.Duration(timeout),
	}
	req, _ := http.NewRequest(method, url, form)
	if token != "" {
		req.Header.Add("token", token)
	}
	req.Close = true
	req.Header.Set("Content-Type", contentType)

	resp, err := client.Do(req)
	if err != nil {
		logrus.Warn(err.Error())
		return ret, false
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		return body, false
	}

	return body, true
}

// CreateMultipartFormBody CreateMultipartFormBody
func CreateMultipartFormBody(params map[string]string) (*bytes.Buffer, string, error) {
	body := &bytes.Buffer{}
	w := multipart.NewWriter(body)

	// Add fields
	for key, val := range params {
		if key == "file" {
			// Open file
			f, err := os.Open(val)
			if err != nil {
				return nil, "", err
			}
			defer f.Close()

			// Add file fields
			fw, err := w.CreateFormFile(key, val)
			if err != nil {
				return nil, "", err
			}
			if _, err = io.Copy(fw, f); err != nil {
				return nil, "", err
			}
		} else {
			// Add string fields
			fw, err := w.CreateFormField(key)
			if err != nil {
				return nil, "", err
			}
			if _, err = fw.Write([]byte(val)); err != nil {
				return nil, "", err
			}
		}
	}
	w.Close()

	return body, w.FormDataContentType(), nil
}

// GetParam 获取querystring 中的value
func GetParam(m url.Values, key string, defVal string) string {
	_, ok := m[key]
	if ok && len(m[key][0]) > 0 {
		return m[key][0]
	}
	return defVal
}

// GetParamInt 获取querystring 中的int type value
func GetParamInt(m url.Values, key string, defVal int) int {
	str := GetParam(m, key, "")
	if str == "" {
		return defVal
	}

	val, err := strconv.Atoi(str)
	if err != nil {
		return defVal
	}

	return val
}

// GetParamInt 获取querystring 中的int type value
func GetParamInt64(m url.Values, key string, defVal int64) int64 {
	str := GetParam(m, key, "")
	if str == "" {
		return defVal
	}

	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return defVal
	}

	return val
}

// GetImageParam 获取multipart form 中的图片
func GetImageParam(r *http.Request) ([]byte, bool) {
	// Read file
	value := r.MultipartForm.Value
	file := r.MultipartForm.File
	blob := []byte{}
	if val, ok := file["file"]; ok {
		b := new(bytes.Buffer)
		f, errf := val[0].Open()
		defer f.Close()
		if errf != nil {
			return nil, false
		}
		io.Copy(b, f)
		blob = b.Bytes()
	} else if val, ok := value["image"]; ok {
		b, _ := base64.StdEncoding.DecodeString(val[0])
		blob = b
	} else {
		return nil, false
	}

	// Check image
	ftype := http.DetectContentType(blob)
	if !strings.HasPrefix(ftype, "image/") {
		return nil, false
	}

	return blob, true
}

// HTTPRequest send form data http request with timeout and retries
func HTTPRequest(url, method, token string, params map[string]string, timeOut, retransmission int) (respBody []byte, err error) {
	var (
		resp    *http.Response
		retries = 3
	)
	if retransmission > 0 {
		retries = retransmission
	}
	form, contentType, err := CreateMultipartFormBody(params)
	if err != nil {
		return respBody, err
	}
	client := &http.Client{
		Timeout: time.Duration(timeOut) * time.Second,
	}
	req, err := http.NewRequest(method, url, form)
	if err != nil {
		return respBody, err
	}
	if token != "" {
		req.Header.Add("token", token)
	}
	req.Close = true
	if contentType != "" {
		req.Header.Set("Content-Type", contentType)
	}
	for retries > 0 {
		resp, err = client.Do(req)
		if err != nil {
			retries--
			time.Sleep(time.Duration(1) * time.Millisecond)
		} else {
			break
		}
	}
	if err != nil {
		logrus.Warn(err.Error())
		return respBody, err
	}
	if resp != nil {
		defer resp.Body.Close()
		if resp.StatusCode != 200 {
			return respBody, fmt.Errorf("response status code: %d", resp.StatusCode)
		}
		respBody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logrus.Warn(err.Error())
			return respBody, err
		}
		return respBody, nil
	}
	return respBody, fmt.Errorf("no response")
}
