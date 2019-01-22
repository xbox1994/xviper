package parser

import (
	"errors"
	"github.com/xbox1994/xviper/constant"
	"net/url"
)

const (
	UrlPrefixFile   = "file"
	UrlPrefixEtcd   = "etcd"
	UrlPrefixConsul = "consul"
)

func Parse(configUrl string) (*url.URL, error) {
	resultUrl, err := url.Parse(configUrl)
	if err != nil {
		return nil, err
	}

	switch resultUrl.Scheme {
	case UrlPrefixFile:
	case UrlPrefixEtcd:
	case UrlPrefixConsul:
	default:
		return nil, errors.New("only support " +
			UrlPrefixFile + constant.UrlSeparator +
			UrlPrefixEtcd + constant.UrlSeparator +
			UrlPrefixConsul + " now")
	}

	return resultUrl, nil
}
