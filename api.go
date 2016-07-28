package main

import (
    "errors"
    "fmt"
    "bytes"
    "log"
    "os"
    "io"
    "io/ioutil"
    "mime/multipart"
    "net/http"
    "encoding/json"
)

var sharedClient *http.Client = nil

/*** API RESPONSE MODELS ***/
type ApiBaseResponse struct {
    Ok bool
    Result interface{}
    Description string
}

func getClient() *http.Client {
    if(sharedClient == nil) {
        sharedClient = &http.Client {
            Jar: nil,
        }
    }

    return sharedClient
}

func loadFileIntoMultipart(writer *multipart.Writer, fileField string, filePath string) error {
    part_writer, _ := writer.CreateFormField(fileField)

    file_handle, err := os.Open(filePath)
    if(err != nil) {
        log.Printf("Failed to open file %s (%s)\n", filePath, err)
        return err
    }
    defer file_handle.Close()

    _, err = io.Copy(part_writer, file_handle)
    if(err != nil) {
        log.Printf("Failed to copy file contents (%s)\n", err)
        return err
    }

    return nil
}

func ApiCreateFilePayload(fileField string, filePath string, processor func(*multipart.Writer)) (io.Reader, string, error) {
    var err error

    body_buffer := &bytes.Buffer { }
    body_writer := multipart.NewWriter(body_buffer)
    defer body_writer.Close()

    err = loadFileIntoMultipart(body_writer, fileField, filePath)
    if(err != nil) {
        return nil, "", err
    }

    if(processor != nil) {
        processor(body_writer)
    }

    body_writer.Close()

    return body_buffer, body_writer.FormDataContentType(), nil
}

func ApiPerform(url string, payload io.Reader, contentType string, output interface{}) error {
    client := getClient()

    request, err := http.NewRequest("POST", url, payload)
    if(err != nil) {
        return err
    }
    request.Header.Add("Content-Type", contentType)
    request.Header.Add("Accept", "application/json")
    request.Header.Add("User-Agent", HttpUserAgent)

    response, err := client.Do(request)
    if(err != nil) {
        return err
    }
    defer response.Body.Close()

    if(response.StatusCode != 200) {
        return errors.New(fmt.Sprintf("HTTP request failed with status code %d", response.StatusCode))
    }

    switch o := output.(type) {
        case *string:
            //Decode as string
            responseBytes, err := ioutil.ReadAll(response.Body)
            if(err != nil) {
                return err
            }
            *o = string(responseBytes)
            return nil

        default:
            //Decode as JSON to target value
            return json.NewDecoder(response.Body).Decode(o)
    }
}
