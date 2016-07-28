package main

import (
    "fmt"
    "mime/multipart"
    "log"
)

const apiSendMessage string = "https://api.telegram.org/bot%s/sendMessage"
const apiSetWebhook string  = "https://api.telegram.org/bot%s/setWebhook"

func TelegramSetWebhook(token string, hookUrl string) error {
    //TODO Add bot ID
    DbLog(nil, INFO, "webhook", fmt.Sprintf("Registering %s", hookUrl))
    log.Printf("Registering %s", hookUrl)

    apiUrl := fmt.Sprintf(apiSetWebhook, token)

    // Create payload
    payload, contentType, err := ApiCreateFilePayload("certificate", BotifyCertificateAbsPath, func(writer *multipart.Writer) {
        writer.WriteField("url", hookUrl)
    })
    if(err != nil) {
        DbLog(nil, ERROR, "webhook", fmt.Sprintf("Failed to generate payload: %s", err))
        return err
    }

    var response ApiBaseResponse
    err = ApiPerform(apiUrl, payload, contentType, &response)
    if(err != nil) {
        DbLog(nil, ERROR, "webhook", fmt.Sprintf("Failed to register: %s", err))
        return err
    }

    return nil
}
