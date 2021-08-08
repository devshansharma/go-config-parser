package config

import (
	"fmt"
	"reflect"
	"encoding/json"
	"os"
	// "errors"
)


func isStruct (i interface{}) bool {
	if reflect.ValueOf(i).Type().Kind() == reflect.Struct {
		return true
	}
	return false
}

func isPtr (i interface{}) bool {
	if reflect.ValueOf(i).Type().Kind() == reflect.Ptr {
		return true
	}
	return false
}


// Parse for parsing struct and filling it as per available environmental variables
func Parse (ifc interface{}) error {
	envOptions := make([]string, 0)	
	envMap := make(map[string]string)

	// if is pointer than take values from elem
	ifcStruct := reflect.ValueOf(ifc).Type()
	if isPtr(ifc) {
		ifcStruct = reflect.ValueOf(ifc).Elem().Type()
		// fmt.Printf("%+v\n",ifcStruct)
	}

	// save json tags to slice, so that we can loop through
	// environment variables	
	for i := 0; i < ifcStruct.NumField(); i++ {
		sf := ifcStruct.Field(i)
		if val, ok := sf.Tag.Lookup("json"); ok {
			envOptions = append(envOptions, val)
		}
	}

	// loop through environment variables and save it in map
	// to do json encoding
	for _, v := range envOptions {
		if ans, ok := os.LookupEnv(v); ok {
			envMap[v] = ans
		}
	}	

	// encode map and decode to struct as per JSON tag names
	data, err := json.Marshal(&envMap)
	if err != nil {
		fmt.Println("Error while marshaling to json", err.Error())
	}
	if err = json.Unmarshal(data, &ifc); err != nil {
		fmt.Println("Error while unmarshaling data", err.Error())
	}		

	return nil
}