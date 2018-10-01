// Copyright (c) 2018 Guido Pola
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.
package main

import (
	"fmt"
	"os"

	"github.com/guidopola/zohosendmail"
)

// This is the API auth token, to get one
// login in your zoho mail and go to
// https://accounts.zoho.com/apiauthtoken/create?SCOPE=ZohoMail/ZohoMailAPI
// to create an auth token.
const ZohoAuthToken string = ""

func main() {
	zm, err := zohosendmail.New(ZohoAuthToken)

	if err != nil {
		fmt.Printf("Error creating Zoho mail object: %v\n", err)
		os.Exit(0)
	}

	err := zm.SendMail("example@mail.com", "Test subject", "This is a test!!")

	if err != nil {
		fmt.Printf("Error sending mail: %v", err)
	}
}
