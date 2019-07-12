package main

import (
		"fmt"
		"os"
		"io/ioutil"
		"strings"
		)

func main() {
	inputFileName := os.Args[1];
	outputFileName := os.Args[2];
	preCapsInBytes, readErr := ioutil.ReadFile(inputFileName);
	if readErr != nil {
        fmt.Print(readErr)
    }
	preCaps := string(preCapsInBytes);
	preCapsArray := strings.Fields(preCaps);
	arrayLength := len(preCapsArray);
	postCapsArray := make([]string, arrayLength); 
	for i := 0; i < len(preCapsArray); i++ {
		postCapsArray[i] = strings.ToUpper(string(preCapsArray[i][0])) + trimLeftChar(string(preCapsArray[i]));
	}

	postCaps := strings.Join(postCapsArray, " ");
	writeErr := ioutil.WriteFile(outputFileName, []byte(postCaps), 0644);
	if writeErr != nil {
		fmt.Print(writeErr)
	}
}

func trimLeftChar(s string) string {
    for i := range s {
        if i > 0 {
            return s[i:]
        }
    }
    return s[:0]
}