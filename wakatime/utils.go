package wakatime

import (
	"errors"
	"os"

	"github.com/mitchellh/go-homedir"
	"gopkg.in/ini.v1"
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

func getApiKeyFromFile(filePath string) (string, error) {
	filePath, _ = homedir.Expand(filePath)
	cfg, err := ini.Load(filePath)
	if err != nil {
		return "", err
	}
	k, err := cfg.Section("settings").GetKey("api_key")
	if err != nil {
		return "", err
	}
	return k.String(), nil
}
