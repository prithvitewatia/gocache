package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/prithvitewatia/gocache/src/common"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

func RequestGet(key string) (string, bool) {
	basePath := fmt.Sprintf("http://%s:%s/get", Config.ServerHost, Config.ServerPort)
	baseUrl, err := url.Parse(basePath)
	if err != nil {
		log.Fatal(err)
	}
	params := url.Values{}
	params.Add("key", key)
	baseUrl.RawQuery = params.Encode()

	getRequest, err := http.NewRequest("GET", baseUrl.String(), nil)
	if err != nil {
		return "", false
	}
	resp, err := Client.Do(getRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var result common.GetResponse
		if err := json.Unmarshal(body, &result); err != nil {
			log.Fatalf("Invalid reponse from server: %s", err)
		}
		return result.Value, true
	} else if resp.StatusCode == 404 {
		return "", false
	} else {
		log.Fatalf("Request failed with status code %d", resp.StatusCode)
	}
	return "", false
}

func RequestSet(key string, value interface{}, ttl time.Duration) error {
	basePath := fmt.Sprintf("http://%s:%s/set", Config.ServerHost, Config.ServerPort)
	baseUrl, err := url.Parse(basePath)
	if err != nil {
		log.Fatal(err)
	}
	setRequestBody := map[string]interface{}{
		"key":   key,
		"value": value,
	}
	if ttl > 0 {
		setRequestBody["ttl"] = ttl
	}
	jsonPayload, err := json.Marshal(setRequestBody)
	if err != nil {
		return err
	}

	setRequest, err := http.NewRequest("POST",
		baseUrl.String(), bytes.NewBuffer(jsonPayload))
	resp, err := Client.Do(setRequest)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func RequestDelete(key string) error {
	basePath := fmt.Sprintf("http://%s:%s/del", Config.ServerHost, Config.ServerPort)
	baseUrl, err := url.Parse(basePath)
	if err != nil {
		log.Fatal(err)
	}
	params := url.Values{}
	params.Add("key", key)
	baseUrl.RawQuery = params.Encode()
	deleteRequest, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	resp, err := Client.Do(deleteRequest)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}

func RequestKeys() []string {
	basePath := fmt.Sprintf("http://%s:%s/keys", Config.ServerHost, Config.ServerPort)
	baseUrl, err := url.Parse(basePath)
	if err != nil {
		log.Fatal(err)
	}
	keysRequest, err := http.NewRequest("GET", baseUrl.String(), nil)
	resp, err := Client.Do(keysRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var result common.GetKeysResponse
		if err := json.Unmarshal(body, &result); err != nil {
			log.Fatalf("Invalid reponse from server: %s", err)
		}
		return result.Keys
	}
	return nil
}

func RequestTtl(key string) int64 {
	basePath := fmt.Sprintf("http://%s:%s/ttl", Config.ServerHost, Config.ServerPort)
	baseUrl, err := url.Parse(basePath)
	if err != nil {
		log.Fatal(err)
	}
	params := url.Values{}
	params.Add("key", key)
	baseUrl.RawQuery = params.Encode()
	ttlRequest, err := http.NewRequest("GET", baseUrl.String(), nil)
	resp, err := Client.Do(ttlRequest)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		var result common.TtlResponse
		if err := json.Unmarshal(body, &result); err != nil {
			log.Fatalf("Invalid reponse from server: %s", err)
		}
		return result.Ttl
	}
	return 0
}

func RequestFlushAll() {
	basePath := fmt.Sprintf("http://%s:%s/flushall", Config.ServerHost, Config.ServerPort)
	baseUrl, err := url.Parse(basePath)
	if err != nil {
		log.Fatal(err)
	}
	flushRequest, err := http.NewRequest("DELETE", baseUrl.String(), nil)
	_, err = Client.Do(flushRequest)
	if err != nil {
		log.Fatal(err)
	}
}
