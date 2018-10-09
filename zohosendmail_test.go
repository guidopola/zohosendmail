package zohosendmail_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"zohosendmail"
)

// How to use it
func Example() {
	// This is the API auth token, to get one
	// login in your zoho mail and go to
	// https://accounts.zoho.com/apiauthtoken/create?SCOPE=ZohoMail/ZohoMailAPI
	// to create an auth token.
	const ZohoAuthToken string = ""

	// Create the zohosendmail object.
	zm, err := zohosendmail.New(ZohoAuthToken)

	if err != nil {
		fmt.Printf("Error creating Zoho mail object: %v\n", err)
		os.Exit(0)
	}

	// try to send a mail
	err = zm.SendMail("example@mail.com", "Test subject", "This is a test!!", nil)

	if err != nil {
		fmt.Printf("Error sending mail: %v", err)
	}
}

func ExampleAttachment() {
	// This is the API auth token, to get one
	// login in your zoho mail and go to
	// https://accounts.zoho.com/apiauthtoken/create?SCOPE=ZohoMail/ZohoMailAPI
	// to create an auth token.
	const ZohoAuthToken string = ""

	// Create the zohosendmail object.
	zm, err := zohosendmail.New(ZohoAuthToken)

	if err != nil {
		fmt.Printf("Error creating Zoho mail object: %v\n", err)
		os.Exit(0)
	}

	content, err := ioutil.ReadFile("./README.md")

	//
	attachment, err := zm.UploadAttachment("README.md", content)

	//
	if err != nil {
		fmt.Printf("Error uploading attachment: %v", err)
	}

	// try to send a mail with an attachment
	err = zm.SendMail("example@mail.com", "Test subject", "This is a test!!", &zohosendmail.MailAttachmentSlice{attachment})

	if err != nil {
		fmt.Printf("Error sending mail: %v", err)
	}
}
