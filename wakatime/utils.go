package wakatime

import (
	"errors"
	"io/ioutil"
	"os"
	"regexp"

	"github.com/mitchellh/go-homedir"
)

func GetApiKey() (string, error) {
	if k, err := getApiKeyFromFile("~/.wakatime.cfg"); err == nil {
		return k, nil
	}
	if k := os.Getenv("WAKAGO_API_KEY"); len(k) > 0 {
		return k, nil
	}

	return "", errors.New("cannot detect api_key from system")
}

var reApiKey = regexp.MustCompile(`api_key\s*=\s(\S+)`)

func getApiKeyFromFile(filePath string) (string, error) {
	filePath, err := homedir.Expand(filePath)
	if err != nil {
		return "", err
	}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	ar := reApiKey.FindStringSubmatch(string(b))
	if len(ar) != 2 {
		return "", errors.New("invalid wakatime config file")
	}
	return ar[1], nil
}
