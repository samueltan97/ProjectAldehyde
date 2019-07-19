package main

import (
	"fmt"
    "log"
	//"io/ioutil"
    "net/http"
	"strings"
	"encoding/json"
)

func main() {
	http.HandleFunc("/capitalize", handler)
	http.Handle("/", http.FileServer(http.Dir("./static")))
	log.Fatal(http.ListenAndServe(":3001", nil))
}

type Text struct {
	ToBeCapitalized string `json:"toBeCapitalized"`
}

func handler(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var data Text
		err := decoder.Decode(&data)
		if err != nil {
			fmt.Print(err)
		}
		fmt.Fprintf(w, capitalize(data.ToBeCapitalized));
}

func capitalize(preCaps string) string {
	preCapsArray := strings.Fields(preCaps);
	arrayLength := len(preCapsArray);
	postCapsArray := make([]string, arrayLength); 
	for i := 0; i < len(preCapsArray); i++ {
		postCapsArray[i] = strings.ToUpper(string(preCapsArray[i][0])) + trimLeftChar(string(preCapsArray[i]));
	}

	postCaps := strings.Join(postCapsArray, " ");
	return postCaps;
}

func trimLeftChar(s string) string {
    for i := range s {
        if i > 0 {
            return s[i:]
        }
    }
    return s[:0]
}