package main

import (
    "bytes"
    "log"
    "fmt"
    "os"
    "io"
    "io/ioutil"
    "mime/multipart"
    "net/http"
)

const apiSendMessage string = "https://api.telegram.org/bot%s/sendMessage"
const apiSetWebhook string  = "https://api.telegram.org/bot%s/setWebhook"

func TelegramSetWebhook(token string, hook_url string) error {
    //var hook_url string := fmt.Sprintf("%s/bot/hook/%s",
    //    BotifyBaseUrl, token)

    log.Printf("Registering URL %s for bot %s\n", hook_url, token)

    target_url := fmt.Sprintf(apiSetWebhook, token)
    log.Printf("Target URL %s\n", target_url)

    body_buffer := &bytes.Buffer{ }
    body_writer := multipart.NewWriter(body_buffer)
    //defer body_writer.Close()

    err := func() error {
        //Certificate part (from file)
        part_writer, _ := body_writer.CreateFormField("certificate")

        file_handle, err := os.Open(BotifyCertificateAbsPath)
        if(err != nil) {
            log.Printf("Failed to open certificate file %s (%s)\n", BotifyCertificateAbsPath, err)
            return err
        }
        defer file_handle.Close()

        _, err = io.Copy(part_writer, file_handle)
        if(err != nil) {
            log.Printf("Failed to copy file contents (%s)\n", err)
            return err
        }

        return nil
    }()
    if(err != nil) {
        return err
    }

    //Other parameters
    body_writer.WriteField("url", hook_url)
    body_writer.Close()

    log.Printf("Performing request...\n")

    response, err := http.Post(target_url, body_writer.FormDataContentType(), body_buffer)
    if(err != nil) {
        return err
    }
    defer response.Body.Close()
    response_body, err := ioutil.ReadAll(response.Body)
    if(err != nil) {
        return err
    }

    log.Printf("Response #%d\n", response.StatusCode)
    log.Printf("%s\n", response_body)

    return nil
}
