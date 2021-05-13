package tradeExecution
//package main

import (
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"
	"encoding/hex"
	_common "github.com/BitBaseBit/FlashGods/goDir/common"
	info "github.com/BitBaseBit/FlashGods/goDir/common/debugInfo"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi"
)

const WETH string        = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"

func GetLowGas(gas70 *big.Int) *big.Int {
	gasPriceConfidence70F := float64(gas70.Int64() + 1)
	gasPriceConfidence70F = gasPriceConfidence70F * 0.60
	gasPriceBig 		  := big.NewInt(int64(gasPriceConfidence70F))
	gasPriceWEI 		  := new(big.Int)

	gasPriceWEI.Mul(gasPriceBig, big.NewInt(1000000000))
	return gasPriceWEI
}


type ResultData struct {
	MethodName    string
	AmountIn      string
	AmountInMax   string
	AmountOut     string
	AmountOutMin  string
	Path          []string
	To            string
	Deadline      string
}

//Returns UniswapV2Router decoded arguments
func GetUniswapTxData(routerVersion int, txData string, uniABI *abi.ABI) (bool, ResultData) {
	var decodedInputs ResultData
	data, err := hex.DecodeString(txData[2:])

	if err != nil {
		fmt.Printf(info.DECODESTRINGFAIL, err)
		return false, decodedInputs 
	}

	if len(data) < 4 {
		return false, ResultData{}
	}

	method, err := uniABI.MethodById(data[:4])

	if err != nil {
		fmt.Printf(info.ABIMETHODID, err)
		return false, decodedInputs 
	}

	result, err := method.Inputs.UnpackValues(data[4:])


	if routerVersion == 2 {
		if method.Name == "swapExactTokensForTokens" || method.Name == "swapExactTokensForTokensSupportingFeeOnTransferTokens" {
			decodedInputs.MethodName   = method.Name
			decodedInputs.AmountIn     	= fmt.Sprintf("%v", result[0])
			decodedInputs.AmountOutMin 	= fmt.Sprintf("%v", result[1])
			//convert result[1] into []common.Address and loop over it
			for _, v := range result[2].([]common.Address) {
				decodedInputs.Path = append(decodedInputs.Path, fmt.Sprintf("%v", v))
			}
			decodedInputs.To       		= fmt.Sprintf("%v", result[3])
			decodedInputs.Deadline 		= fmt.Sprintf("%v", result[4])

		} else if method.Name == "swapTokensForExactTokens" {
			decodedInputs.MethodName  	= method.Name
			decodedInputs.AmountOut   	= fmt.Sprintf("%v", result[0])
			decodedInputs.AmountInMax 	= fmt.Sprintf("%v", result[1])
			for _, v := range result[2].([]common.Address) {
				decodedInputs.Path = append(decodedInputs.Path, fmt.Sprintf("%v", v))
			}
			decodedInputs.To       		= fmt.Sprintf("%v", result[3])
			decodedInputs.Deadline 		= fmt.Sprintf("%v", result[4])

		} else if method.Name == "swapExactETHForTokens" || method.Name == "swapExactETHForTokensSupportingFeeOnTransferTokens" {
			decodedInputs.MethodName 	= method.Name
			for _, v := range result[1].([]common.Address) {
				decodedInputs.Path = append(decodedInputs.Path, fmt.Sprintf("%v", v))
			}
			decodedInputs.To       		= fmt.Sprintf("%v", result[2])
			decodedInputs.Deadline 		= fmt.Sprintf("%v", result[3])

		} else if method.Name == "swapTokensForExactETH" {
			decodedInputs.MethodName  	= method.Name
			decodedInputs.AmountOut   	= fmt.Sprintf("%v", result[0])
			decodedInputs.AmountInMax 	= fmt.Sprintf("%v", result[1])
			for _, v := range result[2].([]common.Address) {
				decodedInputs.Path = append(decodedInputs.Path, fmt.Sprintf("%v", v))
			}
			decodedInputs.To       		= fmt.Sprintf("%v", result[3])
			decodedInputs.Deadline 		= fmt.Sprintf("%v", result[4])

		} else if method.Name == "swapExactTokensForETH" || method.Name == "swapExactTokensForETHSupportingFeeOnTransferTokens" {
			decodedInputs.MethodName   = method.Name
			decodedInputs.AmountIn     = fmt.Sprintf("%v", result[0])
			decodedInputs.AmountOutMin = fmt.Sprintf("%v", result[1])
			for _, v := range result[2].([]common.Address) {
				decodedInputs.Path = append(decodedInputs.Path, fmt.Sprintf("%v", v))
			}
			decodedInputs.To       		= fmt.Sprintf("%v", result[3])
			decodedInputs.Deadline 		= fmt.Sprintf("%v", result[4])

		} else if method.Name == "swapETHForExactTokens" {
			decodedInputs.MethodName 	= method.Name
			decodedInputs.AmountOut  	= fmt.Sprintf("%v", result[0])
			for _, v := range result[1].([]common.Address) {
				decodedInputs.Path = append(decodedInputs.Path, fmt.Sprintf("%v", v))
			}
			decodedInputs.To       		= fmt.Sprintf("%v", result[2])
			decodedInputs.Deadline 		= fmt.Sprintf("%v", result[3])

		} else {
			return false, decodedInputs
		}

	} else if routerVersion == 3 {
		if method.Name == "multicall" {
			resLen := len(result[0].([][]byte))

			for i := 0; i < resLen; i++ {
				rawFunc    := result[0].([][]byte)[i]
				funcSig, _ := uniABI.MethodById(rawFunc[:4])
				inputs, _  := funcSig.Inputs.UnpackValues(rawFunc[4:])
				fmt.Println(funcSig.Name)
				fmt.Println(inputs)
			}

		} else if method.Name == "exactInputSingle"{

		} else if method.Name == "exactInput" {

		} else if method.Name == "exactOutputSingle" {

		} else if method.Name == "exactOutput" {

		}
	}

	return true, decodedInputs
}

// first big int is reserveIn, second is reserveOut
func GetReservesOrdering(token0 string, res0 *big.Int, res1 *big.Int) (*big.Int, *big.Int) {

		if strings.ToLower(token0) == strings.ToLower(WETH) {
			return res0, res1

		} else {
			return res1, res0
		}
}

func GetDuration(txData _common.TxData) *big.Int {
		_timeStamp, err := time.Parse(time.RFC3339, txData.TimeStamp)

		if err != nil {
			fmt.Printf(info.PARSETIMESTAMP,err)
			return nil
		}

		timeStamp := _timeStamp.Unix()
		deadline  := txData.Deadline
		
		//our duration is 95pct of target's
		duration := new(big.Int).Sub(deadline,  big.NewInt(timeStamp))
		duration.Mul(duration, big.NewInt(95)).Quo(duration, big.NewInt(100))

		if duration.Cmp(big.NewInt(0)) <= 0 {
			fmt.Printf(info.NEGATIVEDURATION + "\n txHash: %v\n", txData.TxHash)
			return nil
		}

		return duration
}

// Get the address of a uniswap pair given two contract addresses
// The first string returned is token0
func GetPairAddress(contractAddress, token, otherToken string) (common.Address, common.Address, error) {
	otherTokenInt := new(big.Int)
	otherTokenInt.SetString(otherToken[2:], 16)
	factoryffString := ""
	initCodeString := ""
	
	if contractAddress == "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D" {
		//0xff + UniswapV2 Factory
		factoryffString = "0xff5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f"
		initCodeString  = "96e8ac4277198ff8b6f785478aa9a39f403cb768dd02cbee326c3e7da348845f"
	} else if contractAddress == "0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F" {
		//0xff + SushiswapV2 Factory
		factoryffString = "0xffC0AEe478e3658e2610c5F7A4A2E1777cE9e4f2Ac"
		initCodeString = "e18a34eb0e04b04f7a0ac29a6e80748dca96319b42c54d679cb821dca90c6303"
	}

	if factoryffString == "" || initCodeString == "" {
		fmt.Printf(info.GETPAIRFAILED,"GetPairAdddress: failed set factoryffString or initCodeString" )
		return common.Address{}, common.Address{}, errors.New("GetPairAdddress: failed set factoryffString or initCodeString")
	}

	// remove 0x
	tokenNorm := token[2:]
	// Get big.Int to compare
	tokenInt := new(big.Int)
	tokenInt.SetString(tokenNorm, 16)

	// Cmp returns -1 if tokenInt < otherTokenInt
	i := tokenInt.Cmp(otherTokenInt)

	if i == -1 {
		token0String := token
		token1String := otherToken
		token01String := token0String + token1String[2:]
		token01, err := hexutil.Decode(token01String)

		if err != nil {
			return common.Address{}, common.Address{}, err
		}

		salt := crypto.Keccak256(token01)
		finalString := factoryffString + hexutil.Encode(salt)[2:] + initCodeString
		pairAddress, err := hexutil.Decode(finalString)

		if err != nil {
			return common.Address{}, common.Address{}, err
		}

		pairAddressFinal, err := hexutil.Decode(hexutil.Encode(crypto.Keccak256(pairAddress)))

		if err != nil {
			return common.Address{}, common.Address{}, err
		}

		addrBytes := common.BytesToAddress(pairAddressFinal)
		return common.HexToAddress(token), addrBytes, nil

	} else if i == 1 {
		token0String := otherToken
		token1String := token
		token01String := token0String + token1String[2:]
		token01, err := hexutil.Decode(token01String)

		if err != nil {
			return common.Address{}, common.Address{}, err
		}

		salt := crypto.Keccak256(token01)
		finalString := factoryffString + hexutil.Encode(salt)[2:] + initCodeString
		pairAddress, err := hexutil.Decode(finalString)

		if err != nil {
			return common.Address{}, common.Address{}, err
		}

		pairAddressFinal, err := hexutil.Decode(hexutil.Encode(crypto.Keccak256(pairAddress)))

		if err != nil {
			return common.Address{}, common.Address{}, err
		}

		addrBytes := common.BytesToAddress(pairAddressFinal)

		return common.HexToAddress(otherToken), addrBytes, nil

	} else {
		return common.Address{}, common.Address{}, errors.New("Something went wrong, WETHint & token were equal")
	}
}

func EtherToWei(eth *big.Float) *big.Int {
    eth.Mul(eth, big.NewFloat(math.Pow10(18)))
    wei := new(big.Int)
    eth.Int(wei)
    return wei
}

func WeiToEther(wei *big.Int) *big.Float {
	weiF := big.NewFloat(0).SetInt(wei)
	weiF.Quo(weiF, big.NewFloat(math.Pow10(18)))
	return weiF
}

func gamma(x *big.Float) *big.Float {
    result := new(big.Float)
    return result.Mul(x, big.NewFloat(997)).Quo(result, big.NewFloat(1000))
}

func GetAmountOut(reserveIn, reserveOut, amountIn *big.Float) *big.Float {
    result := new(big.Float)
    return result.Mul(gamma(amountIn), reserveOut).Quo(result, new(big.Float).Add(reserveIn, gamma(amountIn)))
}


func GetAmountIn(reserveIn, reserveOut, amountOut *big.Float) *big.Float {
    result := new(big.Float)
    return result.Mul(amountOut, reserveIn).Quo(result, gamma(new(big.Float).Sub(reserveOut, amountOut)))
}

//returns slippage in percentage
func Slippage(reserveIn, reserveOut, amountIn, amountOutMin*big.Float) (*big.Float, *big.Float) {
    amountOut := GetAmountOut(reserveIn, reserveOut, amountIn)
    slippage  := new(big.Float).Quo(new(big.Float).Mul(new(big.Float).Sub(amountOut, amountOutMin), big.NewFloat(100)), amountOut)
    return slippage, amountOut
}

func GetAmountOutMin(amountOut, slippage *big.Float) *big.Float {
	amountOutMin := new(big.Float).Mul(new(big.Float).Sub(big.NewFloat(1.0), slippage), amountOut)
	return amountOutMin
}

//returns minimum amount of tokenIn needed to ensure a target gets < amountOutMinTarget of tokenOut
func OptimalInputAmount(reserveInPrev, 
					    reserveOutPrev, 
						amountInTarget,
    					amountOutMinTarget *big.Float) *big.Float {

    a := gamma(gamma(amountOutMinTarget))
    b := new(big.Float).Mul(gamma(reserveInPrev), new(big.Float).Mul(big.NewFloat(2), amountOutMinTarget)) 
    b.Add(b, new(big.Float).Mul(gamma(gamma(amountInTarget)), amountOutMinTarget))
    c := new(big.Float).Mul(reserveInPrev, new(big.Float).Mul(reserveInPrev, amountOutMinTarget))
    c.Sub(c, new(big.Float).Mul(gamma(amountInTarget), new(big.Float).Mul(reserveInPrev, reserveOutPrev)))
    c.Add(c, new(big.Float).Mul(gamma(amountInTarget), new(big.Float).Mul(amountOutMinTarget, reserveInPrev)))
    delta := new(big.Float).Sub(new(big.Float).Mul(b, b), new(big.Float).Mul(big.NewFloat(4), new(big.Float).Mul(a, c)))
    result := new(big.Float).Quo(new(big.Float).Sub(new(big.Float).Sqrt(delta), b), new(big.Float).Mul(big.NewFloat(2), a))
    return result
}

func GetReserveInMinSell(reserveOutMax, amountInBuy, amountOutBuy, minProfit *big.Float) *big.Float {
	amountInMinSell := new(big.Float).Add(amountInBuy, minProfit)
	//we want to sell amountOutBuy - 1 worth of token
	amountOutBuy.Sub(amountOutBuy, big.NewFloat(1))
	aux1 := new(big.Float).Add(new(big.Float).Quo(big.NewFloat(1), gamma(amountOutBuy)), new(big.Float).Quo(amountInMinSell, big.NewFloat(4)))
	auxSqrt := new(big.Float).Sqrt(new(big.Float).Mul(amountInMinSell, aux1))
	ratio := new(big.Float).Mul(big.NewFloat(0.5), new(big.Float).Add(amountInMinSell, auxSqrt))
	return new(big.Float).Mul(ratio, reserveOutMax)
}

func SimulateProfitBuy(reserveInStart, 
					   reserveOutStart, 
					   amountInTarget, 
    				   amountOutMinTarget, 
					   minProfit, 
					   amountInBot,
					   amountInBotMax,
					   gasPrice *big.Float, 
					   gasBuy, 
					   gasSell *big.Int) (bool, *big.Float, *big.Float, *big.Float) {

    reserveIn  := big.NewFloat(0)
    reserveOut := big.NewFloat(0)
    tolerance  := big.NewFloat(0)

    //If optimal execution amount > our budget, we do not execute
    if amountInBot.Cmp(amountInBotMax) > 0 {
       return false, nil, nil, nil

    } 
    //we allow reserves to be reserveIn - amountInBotMax
    //at execution time,
    //because this will still knock-off our target's buy
    tolerance = amountInBot
    

    amountOutBot := GetAmountOut(reserveInStart, reserveOutStart, amountInBot)
    reserveIn.Add(reserveInStart, amountInBot)
    reserveOut.Sub(reserveOutStart, amountOutBot)
    //Target BUY
    amountOutTarget := GetAmountOut(reserveIn, reserveOut, amountInTarget)
    reserveIn.Add(reserveIn, amountInTarget)
    reserveOut.Sub(reserveOut, amountOutTarget)
    //Bot SELL
    amountInBotSell := GetAmountOut(reserveOut, reserveIn, amountOutBot)
	gasBuyFloat, _  := new(big.Float).SetString(gasBuy.String())
    gasCostBuy      := new(big.Float).Mul(gasPrice, gasBuyFloat)
	gasSellFloat, _ := new(big.Float).SetString(gasBuy.String())
    gasCostSell     := new(big.Float).Mul(gasPrice, gasSellFloat)
    profit          := new(big.Float).Sub(new(big.Float).Sub(amountInBotSell, gasCostBuy), new(big.Float).Add(amountInBot, gasCostSell))

//    return true, amountInBot, tolerance, profit
    if new(big.Float).Sub(profit, minProfit).Cmp(big.NewFloat(0)) > 0 {
        return true, amountInBot, tolerance, profit

    } else {
        return false, nil, nil, nil 
    }
}

func AssessBuy(reserveIn, 
			   reserveOut, 
			   amountInTarget,
    		   amountOutMinTarget *big.Int, 
			   slippageMaxTarget *big.Float, 
			   minProfit, 
			   amountInBotMax,
    		   gasPrice *big.Int, 
			   gasBuy, 
			   gasSell *big.Int) (bool, *big.Int, *big.Int) {

	lowSlippage                          := false
	reserveInFloat                       := big.NewFloat(0).SetInt(reserveIn)
	reserveOutFloat                      := big.NewFloat(0).SetInt(reserveOut)
	amountInTargetFloat                  := big.NewFloat(0).SetInt(amountInTarget)
	amountOutMinTargetFloat              := big.NewFloat(0).SetInt(amountOutMinTarget)
	targetSlippage, _ := Slippage(reserveInFloat, 
										             reserveOutFloat, 
										             amountInTargetFloat, 
										             amountOutMinTargetFloat)

	if (targetSlippage.Cmp(slippageMaxTarget) <= 0) && (targetSlippage.Cmp(big.NewFloat(0)) >= 0) {
		lowSlippage = true
	} 

	//filter for profitable transactions with low slippage (so we do not get front-runned)    
	if lowSlippage {
		minProfitFloat      := big.NewFloat(0).SetInt(minProfit)
		amountInBotMaxFloat := big.NewFloat(0).SetInt(amountInBotMax)
		gasPriceFloat       := big.NewFloat(0).SetInt(gasPrice)
		
    	//Bot BUY
        amountInBotFloat := OptimalInputAmount(reserveInFloat, reserveOutFloat, amountInTargetFloat, amountOutMinTargetFloat)
	
		profitable, amountInBotFloat, toleranceFloat, profitFloat := SimulateProfitBuy(reserveInFloat, 
																		               reserveOutFloat, 
																		               amountInTargetFloat, 
																		               amountOutMinTargetFloat, 
																	   	               minProfitFloat, 
																   		               amountInBotFloat,
																		               amountInBotMaxFloat, 
																		               gasPriceFloat, 
																		               gasBuy, 
																		               gasSell)
	
		//if first simulation is profitable, we set amountInBot=amountInBotMax and decrease by 20%
		//until it yields the most profit
		if profitable {
			factor := 1.0
		    for i := 0; i < 5; i++ {
				_profitable, _amountInBotFloat, _toleranceFloat, _profitFloat := SimulateProfitBuy(reserveInFloat,
																				   			       reserveOutFloat,
																		          	   	           amountInTargetFloat,
																			                       amountOutMinTargetFloat,
																			                       minProfitFloat,
																			                       new(big.Float).Mul(big.NewFloat(factor), amountInBotMaxFloat),
																							       amountInBotMaxFloat,
																			   	                   gasPriceFloat,
																			                       gasBuy,
																			                       gasSell)
			   if _profitable && (_profitFloat.Cmp(profitFloat) > 0) {
				   amountInBotFloat = _amountInBotFloat
				   toleranceFloat   = _toleranceFloat
				   profitFloat      = _profitFloat
			        //fmt.Printf("Factor: %v || Bot spends: %v || makes profit: %v\n", factor, amountInBotFloat, profitFloat)
					factor -= 0.2
				} else {
					break
				}
			}

			reserveInMinInter := new(big.Float).Sub(reserveInFloat, toleranceFloat)
			reserveInMin      := new(big.Int)
			reserveInMinInter.Int(reserveInMin)

			if reserveInFloat.Cmp(toleranceFloat) == -1 {
				reserveInMin = big.NewInt(0)
			}

			amountInBotInter  := amountInBotFloat
			amountInBot       := new(big.Int)
			amountInBotInter.Int(amountInBot)

			return true, amountInBot, reserveInMin

		} else {
			return false, nil, nil
		}
		
	} else {
		return false, nil, nil 
	}
}


func AssessSell(reserveIn, reserveOut, amountInTarget, 
    minProfit, amountInBot,  amountOutBot,
    gasPrice, gasBuy, gasSell *big.Int) bool {
	reserveInFloat          := big.NewFloat(0).SetInt(reserveIn)
	reserveOutFloat         := big.NewFloat(0).SetInt(reserveOut)
	amountInTargetFloat     := big.NewFloat(0).SetInt(amountInTarget)
	minProfitFloat          := big.NewFloat(0).SetInt(minProfit)
	amountInBotFloat        := big.NewFloat(0).SetInt(amountInBot)
	amountOutBotFloat       := big.NewFloat(0).SetInt(amountOutBot)
	gasPriceFloat           := big.NewFloat(0).SetInt(gasPrice)
	gasBuyFloat             := big.NewFloat(0).SetInt(gasBuy)
	gasSellFloat            := big.NewFloat(0).SetInt(gasSell)
    //Target BUY
    amountOutTarget := GetAmountOut(reserveInFloat, reserveOutFloat, amountInTargetFloat)
    reserveInFloat.Add(reserveInFloat, amountInTargetFloat)
    reserveOutFloat.Sub(reserveOutFloat, amountOutTarget)
    //Bot SELL
    amountInBotSell := GetAmountOut(reserveOutFloat, reserveInFloat, amountOutBotFloat)
    gasCostBuy := new(big.Float).Mul(gasPriceFloat, gasBuyFloat)
    gasCostSell := new(big.Float).Mul(gasPriceFloat, gasSellFloat)
    profit := new(big.Float).Sub(new(big.Float).Sub(amountInBotSell, gasCostBuy), new(big.Float).Add(amountInBotFloat, gasCostSell))

    if new(big.Float).Sub(profit, minProfitFloat).Cmp(big.NewFloat(0)) >= 0 {
        return true 
    } else {
        return false 
    }
}

