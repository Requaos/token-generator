package main

import (
	"fmt"
	"time"

	"github.com/spf13/viper"

	"github.com/dgrijalva/jwt-go"
	"github.com/satori/go.uuid"
)

func main() {
	settings := getSettings()
	fmt.Println("Generating JWT Token...")
	apiKey := settings["key"]
	apiSecret := []byte(settings["secret"])
	tenant := settings["tenantID"]
	applicationID := settings["applicationID"]
	// Time token was obtained at
	obtainedAt := time.Now().Unix()
	// Universally Unique Identifier
	v4uuid := uuid.NewV4()
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":         obtainedAt,
		"iss":         apiKey,
		"sub":         applicationID,
		"callbackUrl": settings["callBackURL"],
		"tenantId":    tenant,
		"jti":         v4uuid,
		"exp":         obtainedAt + int64(3600),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(apiSecret)

	if err == nil {
		fmt.Printf("%s: \n%s\n", "iat", time.Unix(obtainedAt, 0).Format("03:04:05 PM 01-02-2006"))
		fmt.Printf("%s: \n%s\n", "exp", time.Unix(obtainedAt+int64(3600), 0).Format("03:04:05 PM 01-02-2006"))
		fmt.Printf("%s: \n%s\n", "jti", v4uuid)
		fmt.Println(tokenString)
	}
}

func getSettings() map[string]string {
	// example file: secrets.toml
	// [settings]
	// key = "ZXZVBHV3GW9UCGF5CM9V"
	// secret = "IjcDChcC/e5Hy67Pwr0acelPN+SdZG6paSius9Sv"
	// tenantID = "20edcfd1-1727-4c47-b707-239b2383ce4d"
	// applicationID = "3f5271d5-7ed7-4359-a1e8-fcca94fd413f"
	// callBackURL = "http://localhost:4300"
	viper.SetConfigName("secrets")
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
		panic(err)
	}
	return viper.GetStringMapString("settings")
}
