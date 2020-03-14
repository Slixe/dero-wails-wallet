package main

import (
	"bufio"
	"fmt"
	"github.com/deroproject/derosuite/config"
	"github.com/deroproject/derosuite/globals"
	"github.com/deroproject/derosuite/walletapi"
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
	"io/ioutil"
	"os"
	"strings"

	//"github.com/deroproject/derosuite/address"
	"encoding/hex"
	"encoding/json"
)

type walletInfo struct {
	Version string

	UnlockedBalance uint64
	LockedBalance uint64
	TotalBalance uint64

	DaemonHeight uint64
	DaemonTopoHeight uint64
	WalletDaemonAddress string

	WalletHeight uint64
	WalletTopoHeight uint64
	WalletInitialHeight int64
	WalletAddress string

	WalletAvailable bool
	WalletComplete bool
	WalletOnline bool
	WalletMixin int
	WalletFeesMultiplier float64
	WalletSyncTime int64
	WalletMinimumTopoHeight int64
}

type integratedAddress struct {
	Address string
	PaymentId string
}

type transferResult struct {
	TxId string
	TxHex string
	Fee string
	Amount string
	InputsSum string
	Change string
}

type addressValid struct {
	Valid bool
	Integrated bool
	Address string
	PaymentId string
	Err string
}

var daemonAddress = "https://wallet.dero.io:443" //default daemon
var walletInstance *walletapi.Wallet

var configPath = ".config"
var wallets []string

func basic() string {
	return "Hello World from Golang!"
}

func createNewWallet(filename string, password string) string {
	if fileExists(filename) {
		return "Wallet file already exist!"
	}

	errorMessage := "error"
	w, err := walletapi.Create_Encrypted_Wallet_Random(filename, password)

	if err == nil {
		errorMessage = "success"
		addWalletFile(filename)
		walletInstance = w
		walletInstance.SetDaemonAddress(daemonAddress)
	} else {
		errorMessage = err.Error()
	}

	return errorMessage
}

func createEncryptedWalletFromRecoveryWords(filename string, password string, seed string) string {
	if fileExists(filename) {
		return "Wallet file already exist!"
	}

	errorMessage := "error"
	w, err := walletapi.Create_Encrypted_Wallet_From_Recovery_Words(filename, password, seed)

	if err == nil {
		errorMessage = "success"
		addWalletFile(filename)
		walletInstance = w
		walletInstance.SetDaemonAddress(daemonAddress)
	} else {
		errorMessage = err.Error()
	}

	return errorMessage
}

func createEncryptedWalletViewOnly(filename string, password string, viewkey string) string {
	if fileExists(filename) {
		return "Wallet file already exist!"
	}

	errorMessage := "error"
	wallet, err := walletapi.Create_Encrypted_Wallet_ViewOnly(filename, password, viewkey)

	if err != nil {
		errorMessage = err.Error()
	} else {
		errorMessage = "success"
		addWalletFile(filename)
		walletInstance = wallet
		walletInstance.SetDaemonAddress(daemonAddress)
	}

	return errorMessage
}

/*
 Params: filename, password
*/
func openEncryptedWallet(filename string, password string) string {
	error_message := "error"

	w, err := walletapi.Open_Encrypted_Wallet(filename, password)
	if err == nil {
		error_message = "success"
		walletInstance = w
		walletInstance.SetDaemonAddress(daemonAddress)
	} else {
		error_message = err.Error()
	}
	return error_message
}

/*
  Generate an integrated address, with a random 8 bytes paymentID
*/
func generateIntegratedAddress() string {
	var result integratedAddress

	if walletInstance != nil {
		i8 := walletInstance.GetRandomIAddress8()
		result.Address = i8.String()
		dst := make([]byte, hex.EncodedLen(len(i8.PaymentID)))
		hex.Encode(dst, i8.PaymentID)
		result.PaymentId = string(dst)
	}

	res, err := json.Marshal(result)
	if err != nil {
		return err.Error()
	}

	return string(res)
}

/*
  Get all informations about the wallet (block height, balance, etc...)
*/
func getInfos() string {
	var result walletInfo
	result.Version = config.Version.String()

	if walletInstance != nil {
		result.UnlockedBalance, result.LockedBalance = walletInstance.Get_Balance()
		result.TotalBalance = result.UnlockedBalance + result.LockedBalance

		result.WalletHeight = walletInstance.Get_Height()
		result.DaemonHeight = walletInstance.Get_Daemon_Height()
        result.WalletTopoHeight = uint64(walletInstance.Get_TopoHeight())
		result.DaemonTopoHeight = uint64(walletInstance.Get_Daemon_TopoHeight())
		
		result.WalletAddress = walletInstance.GetAddress().String()
		result.WalletAvailable = true
		result.WalletComplete = !walletInstance.Is_View_Only()

		result.WalletInitialHeight = walletInstance.GetInitialHeight()
		result.WalletOnline = walletInstance.GetMode()
		result.WalletMixin = walletInstance.GetMixin()
		result.WalletFeesMultiplier = float64(walletInstance.GetFeeMultiplier())
		result.WalletDaemonAddress = walletInstance.Daemon_Endpoint
		
		result.WalletSyncTime = walletInstance.SetDelaySync(0)
        result.WalletMinimumTopoHeight = walletInstance.GetMinimumTopoHeight()
	} else {
		result.UnlockedBalance = uint64(0)
		result.LockedBalance = uint64(0)
		result.TotalBalance = uint64(0)

		result.WalletHeight = uint64(0)
		result.DaemonHeight = uint64(0)
		result.WalletTopoHeight = uint64(0)
		result.DaemonTopoHeight = uint64(0)

		result.WalletInitialHeight = int64(0)

		result.WalletAddress = ""

		result.WalletAvailable = false
		result.WalletComplete = true
		result.WalletOnline = false
		result.WalletMixin = 5
		result.WalletFeesMultiplier = float64(1.5)
		result.WalletDaemonAddress = ""
		result.WalletSyncTime = int64(0)
		result.WalletMinimumTopoHeight = int64(-1)
	}

	res, err := json.Marshal(result)

	if err != nil {
		return err.Error()
	}

	return string(res)
}
/*
  Change wallet mode: (syncing or not)
*/
func setOnlineMode(onlineMode bool) bool {
	var currentState = false

	if walletInstance != nil {
		if onlineMode {
			walletInstance.SetOnlineMode()
		} else {
			walletInstance.SetOfflineMode()
		}
		currentState = walletInstance.GetMode()
	}

	return currentState
}

/*
  Close the current wallet
*/
func closeWallet() bool {
	if walletInstance != nil {
		walletInstance.Close_Encrypted_Wallet()
		walletInstance = nil
	}

	return walletInstance == nil
}
/*
  Get the seed in language selected
*/
func getSeedInLanguage(language string) string {
	seed := "Some error occurred"
	if walletInstance != nil && len(language) > 1 {
		seed = walletInstance.GetSeedinLanguage(language)
	}

	return seed
}

func getWallets() []string { //return available wallets
	return wallets
}

func addWalletFile(walletName string) bool {
	if fileExists(walletName) {
		wallets = append(wallets, walletName)
		return true
	}

	return false
}

func removeWalletFile(walletName string) bool {
	if fileExists(walletName) {
		/*
		index := sort.SearchStrings(wallets, walletName)
		wallets = append(wallets[:index], wallets[index+1:]...)
		*/

		for i, w := range wallets {
			if w == walletName {
				wallets = append(wallets[:i], wallets[i+1:]...)
			}
		}
		return os.Remove(walletName) == nil
	}

	return false
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func saveWallets() {
	file, err := os.Create(configPath)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, line := range wallets {
		fmt.Fprintln(writer, line)
	}

	err = writer.Flush()

	if err != nil {
		fmt.Println(err.Error())
	}
}

func loadWallets() {
	if !fileExists(configPath) {
		file, err := os.Create(configPath)

		if err != nil {
			fmt.Println(err.Error())
		} else {
			file.Close()
		}
		return
	}

	content, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	wallets = strings.Split(string(content), "\n")
}

func main() {
  	//DERO Wallet
	globals.Arguments = map[string]interface{}{}
	globals.Arguments["--testnet"] = false
	globals.Config = config.Mainnet

	fmt.Println("Loading available wallets...")
	loadWallets()

	//Wails
	html := mewn.String("./frontend/dist/index.html")
	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "DERO Wallet",
		HTML: html,
		JS:     js,
		CSS:    css,
		Colour: "#131313",
		Resizable: true,
		DisableInspector: false,
	})

	app.Bind(basic)
	app.Bind(createNewWallet)
	app.Bind(createEncryptedWalletFromRecoveryWords)
	app.Bind(createEncryptedWalletViewOnly)
	app.Bind(openEncryptedWallet)
	app.Bind(generateIntegratedAddress)
	app.Bind(getInfos)
	app.Bind(setOnlineMode)
	app.Bind(closeWallet)
	app.Bind(getSeedInLanguage)
	app.Bind(getWallets)
	app.Bind(removeWalletFile)

	app.Run()

	if walletInstance != nil {
		walletInstance.Close_Encrypted_Wallet()
	}

	fmt.Println("Saving wallets...")
	saveWallets()
}