package mysql

import (
	"strconv"
	"reflect"
)

type MapModel map[string]interface{}

const DefaultSTRVAL = ""
const DefaultFLTVAL = 0

// use the struct tag mapping mysql field
const ModelTag = "db"
const DbIdTag = "id"

func (this MapModel) GetAttrString(k string) (bool, string) {
	if val, ok := this[k]; ok {
		if vStr, ok := val.(string); ok {
			return true, vStr
		}
	}

	return false, DefaultSTRVAL
}

func (this MapModel) GetAttrFloat(k string) (bool, float64) {
	ok, str := this.GetAttrString(k)
	if !ok {
		return false, DefaultFLTVAL
	}

	flt, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false, DefaultFLTVAL
	}

	return true, flt
}

func (this MapModel) GetAttrInt(k string) (bool, int64) {
	ok, str := this.GetAttrString(k)
	if !ok {
		return false, DefaultFLTVAL
	}

	i, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return false, DefaultFLTVAL
	}

	return true, i
}

func (this MapModel) GetAttrUint(k string) (bool, uint64) {
	ok, str := this.GetAttrString(k)
	if !ok {
		return false, DefaultFLTVAL
	}

	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return false, DefaultFLTVAL
	}

	return true, i
}

// MAP TO MODEL
func (this MapModel) Load(v interface{}) {
	t := reflect.TypeOf(v)

	rv := reflect.ValueOf(v)
	rv = reflect.Indirect(rv)

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)

		if !field.CanSet() {
			continue
		}

		//Get the field tag value
		tag := t.Elem().Field(i).Tag.Get(ModelTag)

		switch field.Kind() {
		case reflect.String:
			if ok, val := this.GetAttrString(tag); ok {
				field.SetString(val)
			}
		case reflect.Uint:
			fallthrough
		case reflect.Uint8:
			fallthrough
		case reflect.Uint16:
			fallthrough
		case reflect.Uint32:
			fallthrough
		case reflect.Uint64:
			if ok, val := this.GetAttrUint(tag); ok {
				field.SetUint(val)
			}
		case reflect.Int:
			fallthrough
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			if ok, val := this.GetAttrInt(tag); ok {
				field.SetInt(val)
			}
		case reflect.Struct:

		}
	}
}

// MODEL TO MAP
func Unload(v interface{}, isFilterId bool) map[string]string {
	t := reflect.TypeOf(v)

	rv := reflect.ValueOf(v)
	rv = reflect.Indirect(rv)

	m := make(map[string]string)
	//records := make(map[string]bool)

	for i := 0; i < rv.NumField(); i++ {
		field := rv.Field(i)

		//Get the field tag value
		tag := t.Elem().Field(i).Tag.Get(ModelTag)

		if isFilterId && tag == DbIdTag {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			m[tag] = field.String()
		case reflect.Uint:
			fallthrough
		case reflect.Uint8:
			fallthrough
		case reflect.Uint16:
			fallthrough
		case reflect.Uint32:
			fallthrough
		case reflect.Uint64:
			m[tag] = strconv.FormatUint(field.Uint(), 10)
		case reflect.Int:
			fallthrough
		case reflect.Int8:
			fallthrough
		case reflect.Int16:
			fallthrough
		case reflect.Int32:
			fallthrough
		case reflect.Int64:
			m[tag] = strconv.FormatInt(field.Int(), 10)
		case reflect.Struct:
			//var base Base = *(*Base)(unsafe.Pointer(&field))
			//records = base.Records
		}
	}

	//for k, _ := range m {
	//	if _, ok := records[k]; !ok {
	//		delete(m, k)
	//	}
	//}

	return m
}
