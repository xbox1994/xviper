package parser

import (
	"errors"
	"github.com/xbox1994/xviper/constant"
	"github.com/xbox1994/xviper/reader"
	"net/url"
)

func Parse(configUrl string) (*url.URL, error) {
	resultUrl, err := url.Parse(configUrl)
	if err != nil {
		return nil, err
	}

	switch resultUrl.Scheme {
	case reader.File:
	case reader.Etcdv3:
	case reader.Consul:
	default:
		return nil, errors.New("only support " +
			reader.File + constant.UrlSeparator +
			reader.Etcdv3 + constant.UrlSeparator +
			reader.Consul + " now")
	}

	return resultUrl, nil
}
