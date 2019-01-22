package xviper

import (
	"github.com/spf13/viper"
	"github.com/xbox1994/xviper/engine"
	"github.com/xbox1994/xviper/option"
	"time"
)

func Init(opt *option.Option) error { return engine.Init(opt) }
func Get(key string) interface{}    { return viper.Get(key) }
func GetString(key string) string   { return viper.GetString(key) }
func GetBool(key string) bool       { return viper.GetBool(key) }
func GetInt(key string) int         { return viper.GetInt(key) }
func GetInt32(key string) int32     { return viper.GetInt32(key) }
func GetInt64(key string) int64     { return viper.GetInt64(key) }
func GetFloat64(key string) float64                          { return viper.GetFloat64(key) }
func GetTime(key string) time.Time                           { return viper.GetTime(key) }
func GetDuration(key string) time.Duration                   { return viper.GetDuration(key) }
func GetStringSlice(key string) []string                     { return viper.GetStringSlice(key) }
func GetStringMap(key string) map[string]interface{}         { return viper.GetStringMap(key) }
func GetStringMapString(key string) map[string]string        { return viper.GetStringMapString(key) }
func GetStringMapStringSlice(key string) map[string][]string { return viper.GetStringMapStringSlice(key) }
func GetSizeInBytes(key string) uint                         { return viper.GetSizeInBytes(key) }
func GetViper() *viper.Viper                                 { return viper.GetViper() }
