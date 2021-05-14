package common

import (
    "math/big"
    "fmt"
	"github.com/ethereum/go-ethereum/common"
)

type ExactInputSingleParams struct {
    FuncName            string
    TokenIn             common.Address "json:\"tokenIn\""; 
    TokenOut            common.Address "json:\"tokenOut\""; 
    Fee                 *big.Int "json:\"fee\""; 
    Recipient           common.Address "json:\"recipient\""; 
    Deadline            *big.Int "json:\"deadline\""; 
    AmountIn            *big.Int "json:\"amountIn\""; 
    AmountOutMinimum    *big.Int "json:\"amountOutMinimum\""; 
    SqrtPriceLimitX96   *big.Int "json:\"sqrtPriceLimitX96\"" ;
}

func (exactInputSingle ExactInputSingleParams) String() string {
    return fmt.Sprintf("Function: %v\n Parameters: \n  TokenIn: %v\n  TokenOut: %v\n  Fee: %v\n  Recipient: %v\n  Deadline: %v\n" + 
                       "  AmountIn: %v\n  AmountOutMinimum: %v\n  SqrtPriceLimitX96: %v\n",
                       exactInputSingle.FuncName, exactInputSingle.TokenIn, exactInputSingle.TokenOut, exactInputSingle.Fee,
                       exactInputSingle.Recipient, exactInputSingle.Deadline, exactInputSingle.AmountIn,
                       exactInputSingle.AmountOutMinimum, exactInputSingle.SqrtPriceLimitX96)
}

type ExactInputBytesPath struct {
    Path []uint8 "json:\"path\""; 
    Recipient common.Address "json:\"recipient\""; 
    Deadline *big.Int "json:\"deadline\""; 
    AmountIn *big.Int "json:\"amountIn\""; 
    AmountOutMinimum *big.Int "json:\"amountOutMinimum\"" 
}

type ExactInputStringPath struct {
    FuncName            string
    Path                []string "json:\"path\""; 
    Recipient           common.Address "json:\"recipient\""; 
    Deadline            *big.Int "json:\"deadline\""; 
    AmountIn            *big.Int "json:\"amountIn\""; 
    AmountOutMinimum    *big.Int "json:\"amountOutMinimum\"" 
}

func (exactInput ExactInputStringPath) String() string {
    return fmt.Sprintf("Function: %v\n Parameters: \n  Path: %v\n  Recipient: %v\n  Deadline: %v\n  AmountIn: %v\n" +
                       "  AmountOutMinimum: %v\n", exactInput.FuncName, exactInput.Path, exactInput.Recipient, 
                       exactInput.Deadline, exactInput.AmountIn, exactInput.AmountOutMinimum)
}

type ExactOutPutBytesPath struct {
    Path            []uint8 "json:\"path\""; 
    Recipient       common.Address "json:\"recipient\""; 
    Deadline        *big.Int "json:\"deadline\""; 
    AmountOut       *big.Int "json:\"amountOut\""; 
    AmountInMaximum *big.Int "json:\"amountInMaximum\""
}

type ExactOutPutStringPath struct {
    FuncName         string
    Path            []string "json:\"path\""; 
    Recipient       common.Address "json:\"recipient\""; 
    Deadline        *big.Int "json:\"deadline\""; 
    AmountOut       *big.Int "json:\"amountOut\""; 
    AmountInMaximum *big.Int "json:\"amountInMaximum\""
}

func (exactOutput ExactOutPutStringPath) String() string {
    return fmt.Sprintf("Function: %v\n Parameters: \n  Path: %v\n  Recipient: %v\n  Deadline: %v\n" +
                       "  AmountOut: %v\n  AmountInMaximum: %v\n", exactOutput.FuncName, 
                       exactOutput.Path, exactOutput.Recipient, exactOutput.Deadline, 
                       exactOutput.AmountOut, exactOutput.AmountInMaximum)
}


type SelfPermit struct {
    FuncName    string
    Token       common.Address
    Value       *big.Int
    Deadline    *big.Int
    V           *big.Int
    R           string
    S           string
}

func (selfPermit SelfPermit) String() string {
    return fmt.Sprintf("Function: %v\n Parameters: \n  Token: %v\n  Value: %v\n  Deadline: %v\n  V: %v\n  R: %v\n  S: %v\n",
                       selfPermit.FuncName, selfPermit.Token, selfPermit.Value, selfPermit.Deadline,
                       selfPermit.V, selfPermit.R, selfPermit.S)
}

