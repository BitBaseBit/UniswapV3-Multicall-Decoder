package common

import (
    "math/big"
	"github.com/ethereum/go-ethereum/common"
)

type ExactInputSingleParams struct {
    TokenIn             common.Address "json:\"tokenIn\""; 
    TokenOut            common.Address "json:\"tokenOut\""; 
    Fee                 *big.Int "json:\"fee\""; 
    Recipient           common.Address "json:\"recipient\""; 
    Deadline            *big.Int "json:\"deadline\""; 
    AmountIn            *big.Int "json:\"amountIn\""; 
    AmountOutMinimum    *big.Int "json:\"amountOutMinimum\""; 
    SqrtPriceLimitX96   *big.Int "json:\"sqrtPriceLimitX96\"" ;
}

type ExactInputBytesPath struct {
     Path []uint8 "json:\"path\""; 
    Recipient common.Address "json:\"recipient\""; 
    Deadline *big.Int "json:\"deadline\""; 
    AmountIn *big.Int "json:\"amountIn\""; 
    AmountOutMinimum *big.Int "json:\"amountOutMinimum\"" 
}

type ExactOutPutBytesPath struct {
    Path            []uint8 "json:\"path\""; 
    Recipient       common.Address "json:\"recipient\""; 
    Deadline        *big.Int "json:\"deadline\""; 
    AmountOut       *big.Int "json:\"amountOut\""; 
    AmountInMaximum *big.Int "json:\"amountInMaximum\""
}

type ExactOutPutStringPath struct {
    Path            []string "json:\"path\""; 
    Recipient       common.Address "json:\"recipient\""; 
    Deadline        *big.Int "json:\"deadline\""; 
    AmountOut       *big.Int "json:\"amountOut\""; 
    AmountInMaximum *big.Int "json:\"amountInMaximum\""
}


type SelfPermit struct {
    Token       common.Address
    Value       *big.Int
    Deadline    *big.Int
    V           *big.Int
    R           string
    S           string
}

