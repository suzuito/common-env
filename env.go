package env

import (
	"os"
	"strconv"
)

// GetenvAsString ...
func GetenvAsString(name string, dflt string) string {
	v := os.Getenv(name)
	if v == "" {
		return dflt
	}
	return v
}

// GetenvAsInt ...
func GetenvAsInt(name string, dflt int) int {
	v := os.Getenv(name)
	ret, err := strconv.Atoi(v)
	if err != nil {
		return dflt
	}
	return ret
}

// GetenvAsInt64 ...
func GetenvAsInt64(name string, dflt int64) int64 {
	v := os.Getenv(name)
	ret, err := strconv.ParseInt(v, 10, 64)
	if err != nil {
		return dflt
	}
	return ret
}

// GetenvAsFloat64 ...
func GetenvAsFloat64(name string, dflt float64) float64 {
	v := os.Getenv(name)
	ret, err := strconv.ParseFloat(v, 64)
	if err != nil {
		return dflt
	}
	return ret
}

// GetenvAsFloat32 ...
func GetenvAsFloat32(name string, dflt float32) float32 {
	v := os.Getenv(name)
	ret, err := strconv.ParseFloat(v, 32)
	if err != nil {
		return dflt
	}
	return float32(ret)
}

// GetenvAsBool ...
func GetenvAsBool(name string, dflt bool) bool {
	v := os.Getenv(name)
	ret, err := strconv.ParseBool(v)
	if err != nil {
		return dflt
	}
	return ret
}
