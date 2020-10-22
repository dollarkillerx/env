package env

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func FillBase(v interface{}) error {
	ind := reflect.Indirect(reflect.ValueOf(v))
	if reflect.ValueOf(v).Kind() != reflect.Ptr || ind.Kind() != reflect.Struct {
		return fmt.Errorf("only the pointer to a struct is supported")
	}

	err := fillBase(ind)
	if err != nil {
		return err
	}
	return nil
}

func fillBase(ind reflect.Value) error {
	for i := 0; i < ind.NumField(); i++ {
		env := strings.ToUpper(ind.Type().Field(i).Name)
		env, ex := os.LookupEnv(env)
		if !ex {
			continue
		}
		switch ind.Field(i).Kind() {
		case reflect.Struct:
			continue
		case reflect.String:
			ind.Field(i).SetString(env)
		case reflect.Int64:
			if ind.Field(i).String() == "time.Duration" {
				t, err := time.ParseDuration(env)
				if err != nil {
					return err
				}
				ind.Field(i).Set(reflect.ValueOf(t))
			}else {
				iv, err := strconv.ParseInt(env, 10, 64)
				if err != nil {
					return err
				}
				ind.Field(i).SetInt(iv)
			}
		case reflect.Uint:
			uiv, err := strconv.ParseUint(env, 10, 32)
			if err != nil {
				return err
			}
			ind.Field(i).SetUint(uiv)
		case reflect.Uint64:
			uiv, err := strconv.ParseUint(env, 10, 64)
			if err != nil {
				return err
			}
			ind.Field(i).SetUint(uiv)
		case reflect.Float32:
			f32, err := strconv.ParseFloat(env, 32)
			if err != nil {
				return err
			}
			ind.Field(i).SetFloat(f32)
		case reflect.Float64:
			f64, err := strconv.ParseFloat(env, 64)
			if err != nil {
				return err
			}
			ind.Field(i).SetFloat(f64)
		case reflect.Bool:
			b, err := parseBool(env)
			if err != nil {
				return err
			}
			ind.Field(i).SetBool(b)
		case reflect.Slice:
			sep := ";"
			vals := strings.Split(env, sep)
			switch ind.Field(i).Type() {
			case reflect.TypeOf([]string{}):
				ind.Field(i).Set(reflect.ValueOf(vals))
			case reflect.TypeOf([]int{}):
				t := make([]int, len(vals))
				for i, v := range vals {
					val, err := strconv.ParseInt(v, 10, 32)
					if err != nil {
						return err
					}
					t[i] = int(val)
				}
				ind.Field(i).Set(reflect.ValueOf(t))
			case reflect.TypeOf([]int64{}):
				t := make([]int64, len(vals))
				for i, v := range vals {
					val, err := strconv.ParseInt(v, 10, 64)
					if err != nil {
						return err
					}
					t[i] = val
				}
				ind.Field(i).Set(reflect.ValueOf(t))
			case reflect.TypeOf([]uint{}):
				t := make([]uint, len(vals))
				for i, v := range vals {
					val, err := strconv.ParseUint(v, 10, 32)
					if err != nil {
						return err
					}
					t[i] = uint(val)
				}
				ind.Field(i).Set(reflect.ValueOf(t))
			case reflect.TypeOf([]uint64{}):
				t := make([]uint64, len(vals))
				for i, v := range vals {
					val, err := strconv.ParseUint(v, 10, 64)
					if err != nil {
						return err
					}
					t[i] = val
				}
				ind.Field(i).Set(reflect.ValueOf(t))
			case reflect.TypeOf([]float32{}):
				t := make([]float32, len(vals))
				for i, v := range vals {
					val, err := strconv.ParseFloat(v, 32)
					if err != nil {
						return err
					}
					t[i] = float32(val)
				}
			case reflect.TypeOf([]float64{}):
				t := make([]float64, len(vals))
				for i, v := range vals {
					val, err := strconv.ParseFloat(v, 64)
					if err != nil {
						return err
					}
					t[i] = val
				}
				ind.Field(i).Set(reflect.ValueOf(t))
			case reflect.TypeOf([]bool{}):
				t := make([]bool, len(vals))
				for i, v := range vals {
					val, err := parseBool(v)
					if err != nil {
						return err
					}
					t[i] = val
				}
				ind.Field(i).Set(reflect.ValueOf(t))
			}
		}
	}
	return nil
}
