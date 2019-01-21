package parser

import (
	"errors"
	"github.com/xbox1994/xviper/constant"
	"net/url"
)

func Parse(configUrl string) (*url.URL, error) {
	resultUrl, err := url.Parse(configUrl)
	if err != nil {
		return nil, err
	}

	switch resultUrl.Scheme {
	case constant.UrlPrefixFile:
	case constant.UrlPrefixEtcd:
	case constant.UrlPrefixConsul:
	default:
		return nil, errors.New("only support " +
			constant.UrlPrefixFile + constant.UrlSeparator +
			constant.UrlPrefixEtcd + constant.UrlSeparator +
			constant.UrlPrefixConsul + " now")
	}

	return resultUrl, nil
}
