package main

import (
	"bufio"
	"bytes"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"
)

/*
	{
        "id": "usr_694a234a-2df6-41bc-9c19-a17fa0554710",
        "username": "sometimessay",
        "displayName": "＂Taciturn＂",
        "currentAvatarImageUrl": "https://media.discordapp.net/attachments/484191882495262731/582015797271068700/MakiPonytailBG.png",
        "currentAvatarThumbnailImageUrl": "https://d348imysud55la.cloudfront.net/thumbnails/2941037663.thumbnail-500.png",
        "last_platform": "standalonewindows",
        "tags": [
            "system_avatar_access",
            "system_world_access",
            "system_trust_basic",
            "system_trust_known",
            "system_feedback_access",
            "system_trust_trusted",
            "system_trust_veteran"
        ],
        "developerType": "none",
        "status": "offline",
        "statusDescription": "",
        "friendKey": "abe857e9a2d59b546e376f122f70329d",
        "last_login": "2020-08-13T00:37:41.144Z",
        "isFriend": true,
        "location": "offline"
	},
*/

// Player - no
type Player struct {
	Id                             string   `json:"id"`
	Username                       string   `json:"username"`
	DisplayName                    string   `xml:"displayName"`
	CurrentAvatarImageUrl          string   `xml:"currentAvatarImageUrl"`
	CurrentAvatarThumbnailImageUrl string   `xml:"currentAvatarThumbnailImageUrl"`
	LastPlatform                   string   `xml:"last_platform"`
	Tags                           []string `json:"tags"`
}

type PlayerList struct {
	players []Player
}

type TwoFactor struct {
	Code string `json:"code"`
}

func main() {

	// http client is technically unecessary
	client := http.Client{}

	// Use terminal.ReadPassword to hide password input
	// Could also create a local xml or json file with the password
	// Load the file to hide any verbal output/input of password
	fmt.Println("Enter your password (Type carefully or copy/paste and press [Enter])")
	password, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}

	// No "good" reason to hide twofactor input
	// bufio allows use to read input from the console
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your two-factor code")
	// Read up to the \n character
	codeStr, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	// Have to cut the \r\n characters out of our string
	codeStr = strings.Replace(codeStr, "\r\n", "", 1)

	// Create a json representation of the twofactor input
	// twofactor expects a string value
	twoFactor := TwoFactor{Code: codeStr}
	data, err := json.Marshal(twoFactor)
	if err != nil {
		panic(err)
	}

	// A buffer allows us to create an io.Reader object for data
	// AKA Golang fancy stuff
	buffer := bytes.NewBuffer(data)
	//u, err := url.Parse("https://api.vrchat.cloud/api/1/auth/twofactorauth/totp/verify")
	if err != nil {
		panic(err)
	}

	// Create a web request to twofactor api URL
	req, err := http.NewRequest("POST", "https://api.vrchat.cloud/api/1/auth/twofactorauth/totp/verify", buffer)
	if err != nil {
		panic(err)
	}
	// Set the basic digest information for username/password
	req.SetBasicAuth("Stachio", string(password))
	// Set the content type so the server knows we're sending JSON
	req.Header.Set("Content-Type", "application/json")

	// Destroy the password []byte variable
	// So our password doesn't live in memory longer than it needs to
	rand.Read(password)

	// Perform the web request
	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	// Read the body of the response
	data, err = ioutil.ReadAll(res.Body)
	// Have to close the body since it's an open stream
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// Do a test request on /auth/user (Returns a buttload of info)
	req, err = http.NewRequest("GET", "https://api.vrchat.cloud/api/1/auth/user", nil)
	if err != nil {
		panic(err)
	}
	// Scoop the auth cookie from the previous response
	// The auth cookie allows our new request to act as an authenticated request
	var authCookie *http.Cookie
	cookies := res.Cookies()
	for _, cookie := range cookies {
		if cookie.Name == "auth" {
			fmt.Println("Auth cookie", cookie)
			authCookie = cookie
			break
		}
	}

	// Add the auth cookie
	req.AddCookie(authCookie)
	// Perform the web request
	res, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	// Read the body of the response
	data, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	// Create a new friends list request
	// Do most of the same stuff as above (A function would help here)
	req, err = http.NewRequest("GET", "https://api.vrchat.cloud/api/1/auth/user/friends?apiKey=JlE5Jldo5Jibnk5O5hTx6XVqsJu4WJ26&offline=true", nil)
	if err != nil {
		panic(err)
	}
	// Add the auth cookie
	req.AddCookie(authCookie)

	// Perform the web request
	res, err = client.Do(req)
	if err != nil {
		panic(err)
	}

	// Read the body of the response
	// data is now our friends list
	data, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))

	/*
		data, err = ioutil.ReadFile("test.json")
		if err != nil {
			panic(err)
		}
	*/

	var players []Player
	err = json.Unmarshal(data, &players)
	if err != nil {
		panic(err)
	}

	/* for _, player := range players {
		fmt.Println(player.DisplayName, player.Tags)
	} */

	for _, player := range players {
		if len(player.Tags) == 0 {
			fmt.Println(player.DisplayName, player.LastPlatform)
		}
	}

	fmt.Println()
	for _, player := range players {
		for _, tag := range player.Tags {
			if tag == "system_trust_known" {
				fmt.Println(player.DisplayName)
				break
			}
		}
	}

	// 00111100 = byte
	// 0 = bit
	// 00110101

	apple := "apple"
	var firstLetter byte
	firstLetter = apple[0]
	fmt.Println(firstLetter)
}
