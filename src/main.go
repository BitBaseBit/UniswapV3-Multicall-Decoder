package main

import (
	"encoding/hex"
	"fmt"
    "io/ioutil"
	"math/big"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	_common "github.com/BitBaseBit/UniswapV3-Multicall-Decoder/src/common"
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
		fmt.Printf("funcSig: %v\n", funcSig)
		fmt.Printf("inputs: %v\n", inputs)
		if funcSig.Name == "exactInputSingle" {
			var exactInputSingle _common.ExactInputSingleParams
			exactInputSingle = inputs[0].(struct { TokenIn common.Address "json:\"tokenIn\""; 
												   TokenOut common.Address "json:\"tokenOut\"";
												   Fee *big.Int "json:\"fee\"";
												   Recipient common.Address "json:\"recipient\"";
												   Deadline *big.Int "json:\"deadline\"";
												   AmountIn *big.Int "json:\"amountIn\"";
												   AmountOutMinimum *big.Int "json:\"amountOutMinimum\"";
												   SqrtPriceLimitX96 *big.Int "json:\"sqrtPriceLimitX96\"" 
											     })
			 fmt.Printf("exactInputSingle: %v\n", exactInputSingle)

		} else if funcSig.Name == "selfPermit" {
			rArray := inputs[4].([32]byte)
			sArray := inputs[5].([32]byte)

			selfPermit := _common.SelfPermit {

				Token: 		inputs[0].(common.Address),
				Value: 		inputs[1].(*big.Int),
				Deadline: 	inputs[2].(*big.Int),
				V:     		inputs[2].(*big.Int),
				R: 	   		hex.EncodeToString(rArray[:]),
				S: 	   		hex.EncodeToString(sArray[:]),
			}
			fmt.Printf("selfPermit: %v\n", selfPermit)
		} else if funcSig.Name == "exactOutput" {
			var exactOutputBytesPath _common.ExactOutPutBytesPath
//			var exactOutputStringPath _common.ExactOutPutStringPath
			exactOutputBytesPath = inputs[0].(struct {
										Path []uint8 "json:\"path\""; 
										Recipient common.Address "json:\"recipient\""; 
										Deadline *big.Int "json:\"deadline\""; 
										AmountOut *big.Int "json:\"amountOut\""; 
										AmountInMaximum *big.Int "json:\"amountInMaximum\"";
									})
			pathBytes := hex.EncodeToString([]byte(exactOutputBytesPath.Path))
			pathLen := len(pathBytes)/44
			var pathString []string
			fmt.Println(pathLen)
			if pathLen < 4 {
				for idx := 0; idx < pathLen; idx++ {
					if idx == 0 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][:40])
					} else if idx == pathLen - 1 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][4:40])
					} else {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][2:40])
					}
				}
			} else {
				for idx := 0; idx < pathLen; idx++ {
					if idx == 0 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][:40])
					} else if idx == pathLen - 1 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][4:40])
					} else {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][2:40])
					}
				}
			}
			fmt.Println(pathString)
		} else if funcSig.Name == "exactInput" {
			var exactInputBytesPath _common.ExactInputBytesPath
			exactInputBytesPath = inputs[0].(struct { Path []uint8 "json:\"path\""; 
													  Recipient common.Address "json:\"recipient\""; 
													  Deadline *big.Int "json:\"deadline\""; 
													  AmountIn *big.Int "json:\"amountIn\""; 
													  AmountOutMinimum *big.Int "json:\"amountOutMinimum\"" 
													 })
			
			pathBytes := hex.EncodeToString([]byte(exactInputBytesPath.Path))
			fmt.Println(pathBytes)
			pathLen   := len(pathBytes)/44
			var pathString []string
			if pathLen < 4 {
				for idx := 0; idx < pathLen; idx++ {
					if idx == 0 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][:40])
					} else if idx == pathLen - 1 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][4:40])
					} else {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][2:40])
					}
				}
			} else {
				for idx := 0; idx < pathLen; idx++ {
					if idx == 0 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][:40])
					} else if idx == pathLen - 1 {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][4:40])
					} else {
						pathString = append(pathString, pathBytes[idx*44:(idx+1)*44][2:40])
					}
				}
			}
			fmt.Println(pathString)

		}
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

	multicallInputExample := "0xac9650d8000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000040000000000000000000000000000000000000000000000000000000000000014000000000000000000000000000000000000000000000000000000000000000c4f3995c67000000000000000000000000a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480000000000000000000000000000000000000000000000000000000253640b4000000000000000000000000000000000000000000000000000000000609e9170000000000000000000000000000000000000000000000000000000000000001baf4db1e4fa0fbc6177b5b869da6d37b0b1da5761c49f06a54d1402ae6a4f48f6694a2db5417c9ec8d718371d77bf7f44981e1392f20272d79e68fdab6fc78957000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000144c04b8d59000000000000000000000000000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000a000000000000000000000000055d119b473cdefec64fa51d7d3fc92bb59087cac00000000000000000000000000000000000000000000000000000000609e8cc00000000000000000000000000000000000000000000000000000000253640b4000000000000000000000000000000000000000000000001acae98526b35cd0c30000000000000000000000000000000000000000000000000000000000000042a0b86991c6218b36c1d19d4a2e9eb0ce3606eb480001f4c02aaa39b223fe8d0a0e5c4f27ead9083c756cc2000bb86dea81c8171d0ba574754ef6f8b412f2ed88c54d00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"

	decodeMulticallInput(multicallInputExample)
}
