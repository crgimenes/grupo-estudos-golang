package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func chkMap(payload map[string]interface{}) (ret bool) {
	limit, lmtOk := payload["Limit"]
	value, valOk := payload["Value"]
	if lmtOk && valOk {
		// Se qualquer valor passar do limite retorna true
		if value.(float64) < limit.(float64) {
			ret = true
			return
		}
	}

	for _, v := range payload {
		//fmt.Printf("%v %T: %v\n", k, v, v)
		switch v.(type) {
		case []interface{}:
			ret = chkSlice(v.([]interface{}))
		case map[string]interface{}:
			ret = chkMap(v.(map[string]interface{}))
		}
		if ret {
			return
		}
	}
	return
}

func chkSlice(pauload []interface{}) (ret bool) {
	for _, v := range pauload {
		//fmt.Printf("%v %T: %v\n", k, v, v)
		switch v.(type) {
		case []interface{}:
			ret = chkSlice(v.([]interface{}))
		case map[string]interface{}:
			ret = chkMap(v.(map[string]interface{}))
		}
		if ret {
			return
		}
	}
	return
}

func main() {
	data, err := os.ReadFile("../payload.json")
	if err != nil {
		fmt.Println(err)
		return
	}

	payload := make(map[string]interface{})
	err = json.Unmarshal(data, &payload)
	if err != nil {
		fmt.Println(err)
		return
	}
	//fmt.Printf("payload %#v\n", payload)

	if chkMap(payload) {
		fmt.Println("Um ou mais itens abaixo do limite")
	}
	fmt.Println("fim")

}
