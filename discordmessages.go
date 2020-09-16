package main

import (
    "bytes"
    "encoding/json"
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
    // Discord Notification WebHook
    // To hardcode your webhook just remove "os.Getenv" below and add your Webhook in place of "D_NOTIFICATION_WH"
    whurl := os.Getenv("D_NOTIFICATION_WH")

    c := strings.Join(os.Args[1:], "\n")
    // format make sure the data send via the webhook is placed inside a code box
    format := ("```" + c + "```")
    m := &Message{
        Content: format,
    }

    j, _ := json.Marshal(m)

    req, err := http.NewRequest("POST", whurl, bytes.NewBuffer(j))
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
