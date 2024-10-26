package tools

import "time"

type mockDB struct{}

var mockLoginDetails = map[string]LoginDetails{
	"zaid": {
		AuthToken: "13531ZA",
		Username:  "zaid",
	},
	"animeman": {
		AuthToken: "REZERO0",
		Username:  "animeman",
	},
	"marie": {
		AuthToken: "ABC123",
		Username:  "marie",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"zaid": {
		Coins:    100,
		Username: "zaid",
	},
	"animeman": {
		Coins:    200,
		Username: "animeman",
	},
	"marie": {
		Coins:    300,
		Username: "marie",
	},
}

func (d *mockDB) GetUserLoginDetails(username string) *LoginDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) GetUserCoins(username string) *CoinDetails {
	// simulate DB call
	time.Sleep(time.Second * 1)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return nil
	}

	return &clientData
}

func (d *mockDB) SetupDatabase() error {
	return nil
}
