package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

type Message struct {
	Content string `json:"content"`
}

func main() {

	// choose type of formatting
	var bold bool
	flag.BoolVar(&bold, "b", false, "Print output in bold")

	var inlinecode bool
	flag.BoolVar(&inlinecode, "i", false, "Print output in an inline code block")

	var codeBlock bool
	flag.BoolVar(&codeBlock, "c", false, "Print output in a code block")

	flag.Parse()

	// hook := os.Getenv("D_NOTIFICATION_WH")
    // add your discord own webhook here
	hook := ("https://discordapp.com/api/webhooks/your-discord-webhook-here")

	var c string
	var format string

	// add formatting to the message
	switch {
	case bold:
		c = strings.Join(os.Args[2:], "\n")
		format = ("**" + c + "**")
	case inlinecode:
		c = strings.Join(os.Args[2:], "\n")
		format = ("`" + c + "`")
	case codeBlock:
		c = strings.Join(os.Args[2:], "\n")
		format = ("```" + c + "```")
	default:
		c = strings.Join(os.Args[1:], "\n")
		format = (c)
	}

	m := &Message{
		Content: format,
	}

	j, _ := json.Marshal(m)

	req, err := http.NewRequest("POST", hook, bytes.NewBuffer(j))
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		panic(err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Println("response Status:", resp.Status)
	log.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("response Body:", string(body))
}
