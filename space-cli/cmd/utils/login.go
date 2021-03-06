package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spaceuptech/space-cli/cmd/model"
)

// Login logs the user in
func Login(selectedAccount *model.Account) (*model.LoginResponse, error) {
	requestBody, err := json.Marshal(map[string]string{
		"user": selectedAccount.UserName,
		"key":  selectedAccount.Key,
	})
	if err != nil {
		_ = LogError(fmt.Sprintf("error in login unable to marshal data - %s", err.Error()), nil)
		return nil, err
	}

	resp, err := http.Post(fmt.Sprintf("%s/v1/config/login?cli=true", selectedAccount.ServerURL), "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		_ = LogError(fmt.Sprintf("error in login unable to send http request - %s", err.Error()), nil)
		return nil, err
	}
	defer CloseTheCloser(resp.Body)

	loginResp := new(model.LoginResponse)
	_ = json.NewDecoder(resp.Body).Decode(loginResp)

	if resp.StatusCode != 200 {
		_ = LogError(fmt.Sprintf("error in login got http status code %v with error message - %v", resp.StatusCode, loginResp.Error), nil)
		return nil, fmt.Errorf("error in login got http status code %v with error message - %v", resp.StatusCode, loginResp.Error)
	}
	return loginResp, err
}

// LoginStart take info of the user
func LoginStart(userName, key, url string) error {
	if userName == "None" {
		if err := survey.AskOne(&survey.Input{Message: "Enter username:"}, &userName); err != nil {
			_ = LogError(fmt.Sprintf("error in login start unable to get username - %v", err), nil)
			return err
		}
	}
	if key == "None" {
		if err := survey.AskOne(&survey.Password{Message: "Enter key:"}, &key); err != nil {
			_ = LogError(fmt.Sprintf("error in login start unable to get key - %v", err), nil)
			return err
		}
	}
	account := model.Account{
		UserName:  userName,
		Key:       key,
		ServerURL: url,
	}
	_, err := Login(&account)
	if err != nil {
		_ = LogError(fmt.Sprintf("error in login start unable to login - %v", err), nil)
		return err
	}
	account = model.Account{
		ID:        userName,
		UserName:  userName,
		Key:       key,
		ServerURL: url,
	}
	// write credentials into accounts.yaml file
	if err := StoreCredentials(&account); err != nil {
		_ = LogError(fmt.Sprintf("error in login start unable to check credentials - %v", err), nil)
		return err
	}
	fmt.Printf("Login Successful\n")
	return nil
}
