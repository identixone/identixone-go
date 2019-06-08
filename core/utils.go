package core

import (
	"fmt"
	"net/url"
	"os"
	"path"
	"runtime"
	"strings"

	"github.com/identixone/identixone-go/utils"
)

const (
	HttpClientBaseURL = "https://api.identix.one"
)

func baseUrl() string {
	env := os.Getenv("IDENTIXONE_CLIENT_BASE_URL")
	if env != "" {
		return env
	}
	return HttpClientBaseURL
}

func makeUrl(apiPath string, query map[string]interface{}) (string, error) {
	u, err := url.Parse(baseUrl())
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, apiPath)

	if strings.HasSuffix(apiPath, "/") {
		u.Path += "/"
	}

	if query != nil && len(query) > 0 {
		q := u.Query()
		for key, val := range query {
			q.Set(key, fmt.Sprintf("%v", val))
		}
		u.RawQuery = q.Encode()
	}
	return u.String(), nil
}

func makeHeaders(token string) map[string]string {
	return map[string]string{
		"User-Agent":    fmt.Sprintf("identixone-go/%s %s", utils.Version, runtime.Version()),
		"Authorization": fmt.Sprintf("Token %s", token),
	}
}
