package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func chkMap(payload map[string]interface{}) (ret bool) {
	lmt, lmtOk := payload["Limit"]
	pct, pctOk := payload["Value"]
	if lmtOk && pctOk {
		if pct.(float64) < lmt.(float64) {
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
	data, err := ioutil.ReadFile("../payload.json")
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
