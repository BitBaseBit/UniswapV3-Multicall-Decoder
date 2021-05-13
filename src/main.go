package main

import (
	"encoding/hex"
	"fmt"
    "io/ioutil"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)


var (
	g_uniV3abi 			abi.ABI
)


// Uniswap Router V3 Sigs
const (
    g_uniV3Router     		   string = "0xE592427A0AEce92De3Edee1F18E0157C05861564"
	g_uniswapV3SwapCallback    string = "uniswapV3SwapCallback"

	g_exactInputSingle 		   string = "exactInputSingle"
	g_exactInput 			   string = "exactInput"

	g_exactOutputSingle        string = "exactOutputSingle"
	g_exactOutput 			   string = "exactOutput"

	g_uniswapV3SwapCallbackSig string = "fa461e33"
	g_multicallSig 			   string = "ac9650d8"

	g_exactInputSingleSig 	   string = "414bf389"
	g_exactInputSig 		   string = "c04b8d59"
	
	g_exactOutputSingleSig     string = "db3e2198"
	g_exactOutputSig 		   string = "f28c0498"

)

func decodeMulticallInput(txInput string) {
	rawData, err          := hex.DecodeString(txInput[2:])
	multicall, err        := g_uniV3abi.MethodById(rawData[:4])
	result, err           := multicall.Inputs.UnpackValues(rawData[4:])
	resLen                := len(result[0].([][]byte))

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < resLen; i++ {
		rawFunc    := result[0].([][]byte)[i]
		funcSig, _ := g_uniV3abi.MethodById(rawFunc[:4])
		inputs, _  := funcSig.Inputs.UnpackValues(rawFunc[4:])
		fmt.Println(funcSig.Name)
		fmt.Println(inputs)
	}
}


func main() {
	_abiUniV3JSON, err := ioutil.ReadFile("V3RouterDirectDownload.json")

	if err != nil {
		log.Fatal(err)
	}

	g_uniV3abi, err = abi.JSON(strings.NewReader(string(_abiUniV3JSON)))

	if err != nil {
		log.Fatal(err)
	}

	multicallInputExample := "0xac9650d800000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001800000000000000000000000000000000000000000000000000000000000000104414bf389000000000000000000000000761d38e5ddf6ccf6cf7c55759d5210750b5d60f3000000000000000000000000c02aaa39b223fe8d0a0e5c4f27ead9083c756cc20000000000000000000000000000000000000000000000000000000000002710000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000609be1ca0000000000000000000000000000000000000000048e44e27c7fd9f3e93b0e4300000000000000000000000000000000000000000000000008c26703b6179908000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000004449404b7c00000000000000000000000000000000000000000000000008c26703b617990800000000000000000000000052f92616f26c7879740632d5a9c7ca07bacdc03a00000000000000000000000000000000000000000000000000000000"

	decodeMulticallInput(multicallInputExample)
}
