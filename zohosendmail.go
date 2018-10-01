// Package zohosendmail is a simple library to send an
// email using the zoho mail api.
// This is usefull for example if your hosting smtp port
// it's blocked.
package zohosendmail

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

var (
	ErrResponseInvalid  = errors.New("Zoho response is invalid!")
	ErrMapEntryNotFound = errors.New("Zoho response data not found!")
)

type zohoEmail struct {
	FromAddress string `json:"fromAddress"`
	ToAddress   string `json:"toAddress"`
	Subject     string `json:"subject"`
	Content     string `json:"content"`
}

// The main object struct.
// Holds the data needed to use the Zoho mail api.
type ZohoMailSender struct {
	// This is the API auth token, to get one
	// login in your zoho mail and go to
	// https://accounts.zoho.com/apiauthtoken/create?SCOPE=ZohoMail/ZohoMailAPI
	// TODO:
	//	Is there a way to get one directly with the api?
	ZohoAuthToken string

	// This is the account id.
	// It's used to call the api rest address
	// i.e https://mail.zoho.com/api/accounts/ + accountId + /messages
	// this is fetched directly from the api with the zoho auth token.
	ZohoAccountId string

	// This is the mail address of the account
	// Fetched directly from the api
	ZohoMailAddress string
}

// Creates a new ZohoMailSender object
func New(zohoAuthToken string) (*ZohoMailSender, error) {
	o := new(ZohoMailSender)

	//
	o.ZohoAuthToken = zohoAuthToken

	// Try to get the account info.
	if err := o.zohoGetAccountInfo(); err != nil {
		return nil, err
	}
	return o, nil
}

func (m *ZohoMailSender) zohoGetAccountInfo() error {
	req, _ := http.NewRequest("GET", "https://mail.zoho.com/api/accounts", nil)
	req.Header.Set("Authorization", "Zoho-authtoken "+m.ZohoAuthToken)

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	//
	var response map[string]interface{}

	if err := json.Unmarshal(body, &response); err != nil {
		return err
	}

	// Make sure the data key exists
	if _, present := response["data"]; !present {
		return ErrMapEntryNotFound
	}

	dataMap := response["data"].([]interface{})
	data := dataMap[0].(map[string]interface{})

	//
	if _, present := data["accountId"]; !present {
		return ErrMapEntryNotFound
	}

	//
	m.ZohoAccountId = data["accountId"].(string)

	// TODO:
	//	Is this value correct?
	// 	We've data.primaryEmailAddress, data.incomingUserName
	//	data.sendMailDetails[index].fromAddress
	if _, present := data["primaryEmailAddress"]; !present {
		return ErrMapEntryNotFound
	}

	m.ZohoMailAddress = data["primaryEmailAddress"].(string)

	//
	return nil
}

// Send an email using the zoho api
// https://www.zoho.com/mail/help/api/post-send-an-email.html
func (m *ZohoMailSender) SendMail(to, subject, content string) error {
	// Build the request json object
	apiRequest, _ := json.Marshal(zohoEmail{
		m.ZohoMailAddress,
		to,
		subject,
		content,
	})
	req, _ := http.NewRequest("POST", "https://mail.zoho.com/api/accounts/"+m.ZohoAccountId+"/messages", bytes.NewBuffer(apiRequest))
	req.Header.Set("Authorization", "Zoho-authtoken "+m.ZohoAuthToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("Response code is %v", resp.StatusCode)
	}

	defer resp.Body.Close()

	return nil
}
