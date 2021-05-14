package main

import (
	"encoding/hex"
	"fmt"
    "io/ioutil"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/BitBaseBit/UniswapV3-Multicall-Decoder/src/common"
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
	rawData, err   := hex.DecodeString(txInput[2:])
	multicall, err := g_uniV3abi.MethodById(rawData[:4])
	result, err    := multicall.Inputs.UnpackValues(rawData[4:])
	resLen         := len(result[0].([][]byte))

	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < resLen; i++ {
		rawFunc    := result[0].([][]byte)[i]
		funcSig, _ := g_uniV3abi.MethodById(rawFunc[:4])
		inputs, _  := funcSig.Inputs.UnpackValues(rawFunc[4:])
		fmt.Println(funcSig.Name)
		fmt.Println(inputs.([]string))
	}
}


func main() {
	_abiUniV3JSON, err := ioutil.ReadFile("SwapRouter.json")

	if err != nil {
		log.Fatal(err)
	}

	g_uniV3abi, err = abi.JSON(strings.NewReader(string(_abiUniV3JSON)))

	if err != nil {
		log.Fatal(err)
	}

	multicallInputExample := "0xac9650d8000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000c4f3995c67000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000000000000000000000000000000000006781b6d900000000000000000000000000000000000000000000000000000000609e4f45000000000000000000000000000000000000000000000000000000000000001c0ef894b078cce26259b98a233444becebddbea058c9ab3978c9356d19d3bb3f659974f0aed20c079170ba6b6cd9a2ff7085e07e5c1449d3c5369121d7f20874d000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000104414bf389000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb48000000000000000000000000ca1207647ff814039530d7d35df0e1dd2e91fa840000000000000000000000000000000000000000000000000000000000000bb8000000000000000000000000e9a75f81a34a5d1c39fb1ebb8dfc7904cc104b8e00000000000000000000000000000000000000000000000000000000609e4ae0000000000000000000000000000000000000000000000000000000006781b6d900000000000000000000000000000000000000000000002c2d1447368e5d35df000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"

	decodeMulticallInput(multicallInputExample)
}
