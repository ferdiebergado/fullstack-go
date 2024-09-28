package debug

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"reflect"
)

func DumpStruct(s interface{}) {
	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		fmt.Printf("%s: %v\n", field.Name, value.Interface())
	}
}

func DumpJsonBody(w http.ResponseWriter, r *http.Request) {
	// Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Unable to read request body", http.StatusBadRequest)
		return
	}
	// defer r.Body.Close() // Ensure the body is closed after reading

	// Dump the raw JSON as a string
	rawJSON := string(body)

	// Print the raw JSON to the server console
	fmt.Println("Raw JSON request body:", rawJSON)
}

func DumpRequestBody(r *http.Request) {
	buf, _ := io.ReadAll(r.Body)
	fmt.Println("Request body:", string(buf))
	r.Body = io.NopCloser(bytes.NewBuffer(buf)) // Restore the body for further use
}

func RequestBodyMap(w http.ResponseWriter, r *http.Request) {
	var dataMap map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&dataMap)
	if err != nil {
		log.Printf("JSON decoding error: %v\n", err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	log.Printf("Decoded map: %+v\n", dataMap)

}
