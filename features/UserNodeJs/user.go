package usernodejs

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

const base_url="https://api.flattenbot.site"

func GetTokenHandler(c echo.Context) (string, error) {
	authHeader := c.Request().Header.Get("Authorization")

	if authHeader == "" {
		return "", c.String(http.StatusUnauthorized, "Header Authorization tidak ditemukan")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", c.String(http.StatusUnauthorized, "Token tidak valid")
	}

	token := parts[1]
	return token, nil
}

func GetByIdUser(idUser string,token string) (User, error) {
	link := fmt.Sprintf("%s/users/%s", base_url, idUser)
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		fmt.Printf("Error creating HTTP request: %s\n", err)
		return User{}, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return User{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return User{}, err
	}

	var respData ResponseUserById
	errjson := json.Unmarshal(body, &respData)
	if err != nil {
		fmt.Println("Error:", errjson)
		return User{}, errjson
	}
	userGet := UserByteToResponse(respData.Data)
	return userGet, nil

}

func GetAllUser(token string) ([]User, error) {
	link := fmt.Sprintf("%s/users", base_url)
	
	req, err := http.NewRequest("GET", link, nil)
	if err != nil {
		fmt.Printf("Error creating HTTP request: %s\n", err)
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return nil, err
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return nil, err
	}

	var respData ResponseAllUser
	err = json.Unmarshal(data, &respData)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return nil, err
	}

	var dataPengguna []User
	for _, pengguna := range respData.Data {
		dataPengguna = append(dataPengguna, UserByteToResponse(pengguna))
	}

	return dataPengguna, nil
}

func LoginUser(login Login) (string, error) {

	jsonData, err := json.Marshal(login)
	if err != nil {
		return "", err
	}
	link := fmt.Sprintf("%s/users/login", base_url)
	request, _ := http.NewRequest("POST", link, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Printf("the HTTP request failed with error %s\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)

		var tokenResp ResponseData
		err := json.Unmarshal(data, &tokenResp)
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return "", errors.New("error unmarshal")
		}
		token := tokenResp.Data.Token
		return token, nil
	}
	return "", nil
}