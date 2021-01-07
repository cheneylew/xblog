package com

import (
	"strconv"
	"fmt"
)

type XMap map[string]interface{}
func (x *XMap) GetInt64(key string, def int64) int64 {
	v, ok := (*x)[key]
	if !ok {
		return def
	}
	switch v.(type) {
	case string:
		o, _ := strconv.Atoi(v.(string))
		return int64(o)
	case int:
		return int64(v.(int))
	}
	return def
}

func (x *XMap) GetInt(key string, def int) int {
	v, ok := (*x)[key]
	if !ok {
		return def
	}
	switch v.(type) {
	case string:
		o, _ := strconv.Atoi(v.(string))
		return int(o)
	case int:
		return int(v.(int))
	}
	return def
}

func (x *XMap) GetString(key string, def string) string {
	v, ok := (*x)[key]
	if !ok {
		return def
	}
	switch v.(type) {
	case string:
		return v.(string)
	case int:
	case int64:
		return fmt.Sprintf("%v", v)
	}

	return def
}
