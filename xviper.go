package xviper

import (
	"github.com/spf13/viper"
	"time"
)

var v *viper.Viper

func Get(key string) interface{}                             { return v.Get(key) }
func GetString(key string) string                            { return v.GetString(key) }
func GetBool(key string) bool                                { return v.GetBool(key) }
func GetInt(key string) int                                  { return v.GetInt(key) }
func GetInt32(key string) int32                              { return v.GetInt32(key) }
func GetInt64(key string) int64                              { return v.GetInt64(key) }
func GetFloat64(key string) float64                          { return v.GetFloat64(key) }
func GetTime(key string) time.Time                           { return v.GetTime(key) }
func GetDuration(key string) time.Duration                   { return v.GetDuration(key) }
func GetStringSlice(key string) []string                     { return v.GetStringSlice(key) }
func GetStringMap(key string) map[string]interface{}         { return v.GetStringMap(key) }
func GetStringMapString(key string) map[string]string        { return v.GetStringMapString(key) }
func GetStringMapStringSlice(key string) map[string][]string { return v.GetStringMapStringSlice(key) }
func GetSizeInBytes(key string) uint                         { return v.GetSizeInBytes(key) }
func GetViper() *viper.Viper                                 { return v }
