package service

import (
	"encoding/csv"
	"fmt"
	"reflect"
	"votes/model"
)

// structFieldNames returns the names of the fields of a struct type
func StructFieldNames(s interface{}) []string {
	typ := reflect.TypeOf(s)
	names := make([]string, typ.NumField())
	for i := 0; i < typ.NumField(); i++ {
		names[i] = typ.Field(i).Name
	}
	return names
}

// structFieldValues returns the values of the fields of a struct instance
func StructFieldValues(s interface{}) []string {
	val := reflect.ValueOf(s)
	values := make([]string, val.NumField())
	for i := 0; i < val.NumField(); i++ {
		values[i] = FormatValue(val.Field(i))
	}
	return values
}

// formatValue converts a reflect.Value to a string
func FormatValue(v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return v.String()
	case reflect.Int:
		return fmt.Sprintf("%d", v.Int())
	case reflect.Float64:
		return fmt.Sprintf("%.2f", v.Float())
	default:
		return fmt.Sprintf("%v", v)
	}
}

func SaveCsv(territoryChannel <-chan model.TerritoryResults, writer *csv.Writer) {
	first := true
	for n := range territoryChannel {
		res := model.VoteResultFromTR(&n)
		if first {
			header := StructFieldNames(*res)
			if err := writer.Write(header); err != nil {
				panic(err)
			}
			first = false
		}
		row := StructFieldValues(*res)
		if err := writer.Write(row); err != nil {
			panic(err)
		}
	}
}
