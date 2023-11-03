# Wikipedia API

Brincando com ler dados da wikipedia

```
curl https://en.wikipedia.org/w/api.php\?action\=opensearch\&format\=json\&search\=golang
```

```go
package main

import (
 "encoding/json"
 "fmt"
 "io/ioutil"
 "net/http"
 "reflect"
 "strings"
)

func main() {
 resp, err := http.Get("https://pt.wikipedia.org/w/api.php?action=opensearch&format=json&search=Go_(linguagem_de_programação)")
 if err != nil {
  fmt.Println(err)
  return
 }
 defer resp.Body.Close()

 body, err := io.ReadAll(resp.Body)
 if err != nil {
  fmt.Println(err)
 }

 //fmt.Println(string(body))

 var m []interface{}

 err = json.Unmarshal(body, &m)
 if err != nil {
  fmt.Println(err)
  return
 }

 r := make(map[string]int)
 walker(m, r)

 for k, v := range r {
  fmt.Println(k, v)
 }

 /*
  b, err := json.MarshalIndent(m, "", "\t")
  if err != nil {
   fmt.Println(err)
   return
  }

  fmt.Println(string(b))
 */
}

// recursividade e reflection na mesma função... programador corajoso!
func walker(m interface{}, r map[string]int) {
 for _, v := range m.([]interface{}) {
  switch reflect.TypeOf(v).Kind() {
  case reflect.String:
   a := strings.Split(v.(string), " ")
   for _, s := range a {
    aux := r[s]
    aux++
    r[s] = aux
   }
  case reflect.Slice:
   walker(v, r)
  default:
   fmt.Println(v, reflect.TypeOf(v))
  }
 }
}
```
