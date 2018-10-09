[![GoDoc](https://godoc.org/github.com/guidopola/zohosendmail?status.svg)](https://godoc.org/github.com/guidopola/zohosendmail)
# A Go library to send mails with Zoho.net


## What is this?
This is a simple library for sending emails with the zoho.net REST API.

## Usage
##### Get the auth token.
Log in your zoho mail and go to [Zoho.net API auth token create](https://accounts.zoho.com/apiauthtoken/create?SCOPE=ZohoMail/ZohoMailAPI)
##### Import the module

```go
import "github.com/guidopola/zohosendmail"
```
##### Create the zohosendmail object
```go
zm, err := zohosendmail.New(ZohoAuthToken)
```
##### Send an email

```go
err := zm.SendMail("test.dest@mail.com", "Test subject", "This is a test email body!", nil)
```

##### To add an attachment.

```go
// Read the file contents
content, err := ioutil.ReadFile("./example.zip")

// Upload the attachment to zoho.
attachment, err := zm.UploadAttachment("example.zip", content)

// Send the email.
err := zm.SendMail("test.dest@mail.com", "Test subject", "This is a test email body!", &zohosendmail.MailAttachmentSlice{attachment})
```

##### To add multiple attachments.

```go
err := zm.SendMail("test.dest@mail.com", "Test subject", "This is a test email body!", &zohosendmail.MailAttachmentSlice{attachment, attachment, attachment})
```

#####



License
----
MIT