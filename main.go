package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

func doGraphQLQuery(url, q string) (map[string]interface{}, error) {
	query := map[string]interface{}{
		"query": q,
	}
	var buf bytes.Buffer
	e := json.NewEncoder(&buf)
	err := e.Encode(&query)
	if err != nil {
		return nil, err
	}

	r, err := http.Post(url, "application/json", &buf)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()
	d := json.NewDecoder(r.Body)

	var m map[string]interface{}
	err = d.Decode(&m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func cleanQuery(s string) string {
	// s = strings.ReplaceAll(s, "\n", `\n`)
	// s = strings.ReplaceAll(s, "\t", `\t`)
	// s = strings.ReplaceAll(s, "  ", ` `)

	return s
}

func main() {
	flag.Parse()

	url := flag.Arg(0)
	query := flag.Arg(1)

	query = cleanQuery(query)

	res, err := doGraphQLQuery(url, query)
	if err != nil {
		fmt.Println("error doing GraphQL query:", err)
		return
	}

	b, _ := json.Marshal(&res)
	fmt.Println(string(b))
}
