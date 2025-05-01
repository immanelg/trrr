package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"net/http"
	"net/url"
	"os"
	"strings"
)

const apiBaseUrl = "https://translate.googleapis.com/translate_a/single"
const defaultSrc = "auto"
const defaultDst = "en"

func main() {
	n := len(os.Args) 
	src, dst := "", ""
	content := ""
	if n >= 2 {
		src, dst, _ = strings.Cut(os.Args[1], ":")
	}
	if src == "" {
		src = defaultSrc
	}
	if dst == "" {
		dst = defaultDst
	}
	if n >= 3 {
		content = os.Args[2]
	} else if content == "" {
		stat, _ := os.Stdin.Stat()

		if (stat.Mode() & os.ModeCharDevice) == 0 {
			c, err := io.ReadAll(os.Stdin)
			if err != nil && errors.Is(err, io.EOF) {
				fmt.Fprintln(os.Stderr, "cannot read stdin:", err)
				os.Exit(1)
			}
			content = string(c)
		} else {
			fmt.Fprintf(os.Stderr, "neither CONTENT or stdin provided, see usage")
		}
	}

    params := url.Values{}
    params.Add("client", "gtx")
    params.Add("dt", "t")
    params.Add("q", content)
    params.Add("sl", src)
    params.Add("tl", dst)

    requestFullUrl := fmt.Sprintf("%s?%s", apiBaseUrl, params.Encode())
    resp, err := http.Get(requestFullUrl)
    if err != nil {
        fmt.Fprintf(os.Stderr, "cant make request: %s", err)
		os.Exit(1)
        return
    }
    defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot read response body: %s", err)
		return
	}

    var result []any
    if err := json.Unmarshal(body, &result); err != nil {
        fmt.Fprintf(os.Stderr, "api error, cant parse json response:\n%s", body)
        return
    }

	translation := result[0].([]any)[0].([]any)[0] // yeah

    fmt.Print(translation)
}
