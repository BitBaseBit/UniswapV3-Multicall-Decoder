package tradeExecution

import (
	"bufio"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	trader "github.com/BitBaseBit/FlashGods/goDir/Trader"
	_common "github.com/BitBaseBit/FlashGods/goDir/common"
    info "github.com/BitBaseBit/FlashGods/goDir/common/debugInfo"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/accounts/abi"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-ethereum-hdwallet"
)

var g_abi abi.ABI

type Zk struct {
    D           *ecdsa.PrivateKey 
    FromAddress common.Address
}


const WETHAddress string = "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"
const CONTRACT string = "0x23674777D5c5D8FD59852D9a44FeA45C49A36D5e"


func Pack(abi abi.ABI, name string, args ...interface{}) ([]byte, error) {
	method, exist := abi.Methods[name]
    if !exist {
        return nil, fmt.Errorf("method '%s' not found", name)
    }
    arguments, err := method.Inputs.Pack(args...)
    if err != nil {
        return nil, err
    }
    // Pack up the method ID too if not a constructor and return
    return append(method.ID, arguments...), nil
} 


func SimpleTrade(contractAddr common.Address, client *ethclient.Client, zk Zk) {
    start := time.Now()
    address       := contractAddr
    instance, err := trader.NewUniswapV2Trader(address, client)

    //erc20, err := trader.NewIERC20(common.HexToAddress("0x0de6aD1AE91c3df9aB3B883D860beF3DAF6F7f58"),client)


    //fmt.Println(erc20.BalanceOf(&bind.CallOpts{}, address))

    if err != nil {
        log.Fatal(err)
    }

    d           := zk.D
    pubZk       := d.Public()
    pubZkECDSA  := pubZk.(*ecdsa.PublicKey)
    gasPrice    := big.NewInt(5000000000)
    fromAddress := crypto.PubkeyToAddress(*pubZkECDSA)
    nonce,err   := client.PendingNonceAt(context.Background(), fromAddress)

    if err != nil {
        log.Fatal(err)
    }

    auth, err   := bind.NewKeyedTransactorWithChainID(d, big.NewInt(3))

    if err != nil {
        log.Fatal(err)
    }

    auth.Nonce = big.NewInt(int64(nonce))
    auth.Value = big.NewInt(0)
    auth.GasLimit = uint64(2000000)
    auth.GasPrice = gasPrice

    //Akita Inu token
    tokenaddr := "0x3301Ee63Fb29F863f2333Bd4466acb46CD8323E6"
    weth := "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"

    uniV2Router := "0x7a250d5630B4cF539739dF2C5dAcb4c659F2488D"
    //sushiV2Router := "0xd9e1cE17f2641f24aE83637ab66a2cca9C378B9F"
    
    token0, pairaddr, _ := GetPairAddress(uniV2Router, tokenaddr, weth)

    if err != nil {
        fmt.Println(err)
    }

    //fmt.Println(tx2)

    _tokenaddr := common.HexToAddress(tokenaddr)

    fmt.Printf("tokenaddr: %v\n, token0: %v\n, pairaddr: %v\n", _tokenaddr, token0, pairaddr) 


    oneWETH := big.NewInt(1000000000000000000)

    //fromIsToken0 := (_tokenaddr == token0)

    /*tx, err := instance.Sell(auth, 
                             fromIsToken0,
                             _tokenaddr,
                             pairaddr,
                             big.NewInt(0),
                             big.NewInt(time.Now().Unix() + 250))*/

    tx, err := instance.ConvertETHtoWETH(auth, new(big.Int).Mul(big.NewInt(12), oneWETH))

    if err != nil {
        fmt.Printf("====================================================\n" +
                   " FAILED TO EXECUTE SIMPLETRADE\n" +
                   "====================================================\n")
    }
    elapsed := time.Since(start)
    fmt.Printf("took %v to make tx\n", elapsed)

    fmt.Printf("tx send: %s\n", tx.Hash().Hex())
    

} 


func CheckOwnedHBalance(instance *trader.UniswapV2Trader, client *ethclient.Client) (bool, *big.Int) {
	WETH,_         := NewIERC20(common.HexToAddress("0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"), client)
	balanceWETH, _ := WETH.BalanceOf(&bind.CallOpts{}, common.HexToAddress(CONTRACT))
    WETHLow          := balanceWETH.Cmp(big.NewInt(900000000000000000))

    balanceETH, _ := client.BalanceAt(context.Background(),common.HexToAddress(CONTRACT), nil)
    ETHLow          := balanceETH.Cmp(big.NewInt(400000000000000000))

    if WETHLow <= 0 && ETHLow <= 0 {
        fmt.Println(info.BALANCELOW)
    }

    return true, nil
}

func (zk *Zk) SetWalletFromMnemonic() { 
    fmt.Print("Enter wallet mnemonic: ") 
    inputReader := bufio.NewReader(os.Stdin)
    mnemonic, _ := inputReader.ReadString('\n')
    
    wallet, err := hdwallet.NewFromMnemonic(strings.TrimSpace(mnemonic))

    if err != nil  {
       log.Fatal(err) 
    }

    path         := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, true)

	if err != nil {
		log.Fatal(err)
	}
    
    d, err := wallet.PrivateKey(account)

	if err != nil {
		log.Fatal(err)
	}

    fmt.Printf("wallet is loaded\n")
    zk.D = d
}

func SetContractTrader(client *ethclient.Client, _address string) *trader.UniswapV2Trader {
    address       := common.HexToAddress(_address)
    instance, err := trader.NewUniswapV2Trader(address, client)

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("contract is loaded")
    return instance
}

func BuyTokenWithWETH(zk Zk,
                      client *ethclient.Client,
                      instance *trader.UniswapV2Trader,
                      txData _common.TxData,
                      replaceTx bool,
                      tokenAddress,
                      pairAddress,
                      token0 common.Address,
                      WETHAmountMax_wei *big.Int,
                      gasPrice,
                      duration,
                      reserveIn_wei,
                      reserveOut_wei,
                      amountInTarget_wei,
                      amountOutMinTarget_wei *big.Int,
                      slippageMaxTarget *big.Float,
                      minProfit *big.Int,
                      gasBuy,
                      gasSell *big.Int) (bool, *big.Int,*types.Transaction, error) {

    if (reserveIn_wei.Cmp(big.NewInt(0)) <= 0) || (reserveOut_wei.Cmp(big.NewInt(0)) <= 0) {
       return false, nil, nil, nil 
    }

//    fmt.Printf("amount in from txData: %v\n", txData.AmountIn)
//    fmt.Printf("amount out min target from txData: %v\n", txData.AmountOutMin)
//    fmt.Printf("amount in target: %v\n", amountInTarget_wei)
//    fmt.Printf("amount out min target: %v\n", amountOutMinTarget_wei)
    
    shouldBuy, WETHAmount_wei, reserveInMin_wei := AssessBuy(reserveIn_wei, 
                                                             reserveOut_wei, 
                                                             amountInTarget_wei, 
                                                             amountOutMinTarget_wei, 
                                                             slippageMaxTarget, 
                                                             minProfit, 
                                                             WETHAmountMax_wei,  
                                                             gasPrice, 
                                                             gasBuy, 
                                                             gasSell)

//    fmt.Printf("shouldBuy: %v\n WETHAmount: %v", shouldBuy, WETHAmount_wei)

    if shouldBuy || replaceTx {

        fmt.Println("Now trying to buy")
        deadline  := new(big.Int).Add(big.NewInt(time.Now().Unix()), duration)
        d         := zk.D
        auth, err := bind.NewKeyedTransactorWithChainID(d, big.NewInt(1))

        if err != nil {
            fmt.Printf(info.KEYEDTRANSACTORFAIL, err)
            return false, nil, &types.Transaction{}, nil
        }

        nonce,err   := client.PendingNonceAt(context.Background(), 
                                             common.HexToAddress("0x76216457828c7Db81c888A87F6cBd7dcfA86A907"))

        auth.Nonce    = big.NewInt(int64(nonce))
        auth.Value    = big.NewInt(0)
        auth.GasLimit = uint64(300000)
        auth.GasPrice = gasPrice

        from         := common.HexToAddress(WETHAddress)
        fromIsToken0 := (from == token0)

        tx, err := instance.Buy(auth,
                                fromIsToken0,
                                from,
                                pairAddress,
                                WETHAmount_wei,
                                reserveInMin_wei,
                                deadline)

        if err != nil {
            fmt.Printf(info.BUYFAILED, err)
            return false, nil, &types.Transaction{}, nil
        }

        fmt.Printf("BuyTokenWithWETH executed with tx hash: %v\n", tx.Hash().Hex())

        return true, WETHAmount_wei, tx, nil

    } else {
        return false, nil, nil, nil
    }
}


func SellTokenForWETH(zk Zk,
                      client *ethclient.Client,
                      instance *trader.UniswapV2Trader,
                      tokenAddress,
                      pairAddress,
                      token0 common.Address,
                      reserveInMin_wei,
                      gasPrice,
                      duration *big.Int) (*types.Transaction, error) {
    deadline := new(big.Int).Add(big.NewInt(time.Now().Unix()), duration)
    d           := zk.D
    auth, err   := bind.NewKeyedTransactorWithChainID(d, big.NewInt(1))

    if err != nil {
        log.Fatal(err)
    }

    nonce,err  := client.PendingNonceAt(context.Background(), common.HexToAddress("0x76216457828c7Db81c888A87F6cBd7dcfA86A907"))

    auth.Nonce    = big.NewInt(int64(nonce))
    auth.Value    = big.NewInt(0)
    auth.GasLimit = uint64(300000)
    auth.GasPrice = gasPrice

    fromIsToken0 := (tokenAddress == token0)

    tx, err := instance.Sell(auth,
                             fromIsToken0,
                             tokenAddress,
                             pairAddress,
                             reserveInMin_wei,
                             deadline)
    if err != nil {
			fmt.Printf(info.SELL, err)
        return &types.Transaction{}, nil
    }

    fmt.Printf("SellTokenForWETH executed with tx hash: %v\n", tx.Hash().Hex())

    CheckOwnedHBalance(instance, client)

    return tx, err
}

func ConvertETHtoWETH(zk Zk,
                      instance *trader.UniswapV2Trader,
                      ETHAmount_wei,
                      nonce,
                      gasPrice *big.Int) (*types.Transaction, error){
    d           := zk.D
    auth, err   := bind.NewKeyedTransactorWithChainID(d, big.NewInt(1))

    if err != nil {
        log.Fatal(err)
    }

    auth.Nonce    = nonce
    auth.Value    = big.NewInt(0)
    auth.GasLimit = uint64(200000)
    auth.GasPrice = gasPrice

    tx, err := instance.ConvertETHtoWETH(auth, ETHAmount_wei)

    if err != nil {
        fmt.Printf("failed to execute ConvertETHtoWETH with error: %v\n", err)
    }

    fmt.Printf("ConvertETHtoWETH tx hash: %v\n", tx.Hash().Hex())
    return tx, nil

}


func ConvertWETHtoETH(zk Zk,
                      instance *trader.UniswapV2Trader,
                      WETHAmount_WEI,
                      nonce,
                      gasPrice *big.Int) (*types.Transaction, error){
    d           := zk.D
    auth, err   := bind.NewKeyedTransactorWithChainID(d, big.NewInt(1))

    if err != nil {
        log.Fatal(err)
    }

    auth.Nonce    = nonce
    auth.Value    = big.NewInt(0)
    auth.GasLimit = uint64(200000)
    auth.GasPrice = gasPrice

    tx, err := instance.ConvertETHtoWETH(auth, WETHAmount_WEI)

    if err != nil {
        fmt.Printf("failed to execute transaction with error: %v\n", err)
    }

    fmt.Printf("ConvertWETHtoETH tx hash: %v\n", tx.Hash().Hex())
    return tx, nil
}

func Init(client *ethclient.Client, contractAddress string, chainID *big.Int) Zk {
    zk := Zk{}
    zk.SetWalletFromMnemonic()
    d              := zk.D
    pubZk          := d.Public()
    pubZkECDSA     := pubZk.(*ecdsa.PublicKey)
    zk.FromAddress = crypto.PubkeyToAddress(*pubZkECDSA)
    return zk
}

