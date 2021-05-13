// Should only use structs and methods that are prefixed with the word
// Safe concurrently, anything else is not safe to use concurrently
package common

import (
    "fmt"
    "sync"
    "math/big"
    "strings"

    common "github.com/ethereum/go-ethereum/common"
	"github.com/json-iterator/go"
)

var json = jsoniter.ConfigFastest

// Stringer implementations for types
func (buy Trade) String() string {
    return fmt.Sprintf(" Hash: %v\n IsSell: %v\n Value: %v\n TokenContractAddress: %v\n" +
                       "TokenPairAddress: %v\n AmountIn: %v\n AmountInMax %v\n AmountOut: %v\n" +
                       "AmountOutMin: %v\n Gas Price: %v\n", buy.TxHash,buy.Sell,  
                       buy.ValueETH, buy.TokenContractAddress, buy.TokenPairAddress, 
                       buy.AmountIn, buy.AmountInMax, buy.AmountOut, buy.AmountOutMin, buy.TxGasPrice)
}
// NOTE, this is a crappy function and should not be used in production
// it should only be used for debugging
func (pendingTx *SafeOwnedPendingTransactions) SafePrintPending() int {
    pendingTx.PendingTxMu.Lock()
    defer pendingTx.PendingTxMu.Unlock()

    fmt.Printf("TxHashes: %v\n TokenContractAddresses: %v\n",
               pendingTx.TxHashes, 
               pendingTx.TokenContractAddresses)
    return 0
}

// This is the struct that holds the tokens that we actually own
type SafeHeldTokens struct {
    HtMu                        sync.Mutex 
    HeldTokensContractAddresses []string
}

// Checks if we own a given token, returns true if we do and the index
// of the token in the arrays of SafeHeldTokens
// Note reciever must be a pointer to avoid copying the mutex lock
func (heldTokens *SafeHeldTokens) SafeContains(tokenAddress string) (bool, int) {
    heldTokens.HtMu.Lock()

    for idx, val := range heldTokens.HeldTokensContractAddresses {

        if strings.ToLower(val) == strings.ToLower(tokenAddress) {
            heldTokens.HtMu.Unlock()
            return true, idx
        }

    }

    heldTokens.HtMu.Unlock()
    return false, -1
}

func (heldTokens *SafeHeldTokens) SafeAddToken(tokenContractAddress string) error {
    heldTokens.HtMu.Lock()

    heldTokens.HeldTokensContractAddresses = append(heldTokens.HeldTokensContractAddresses, 
                                                    tokenContractAddress)
    heldTokens.HtMu.Unlock()
    return nil
}
func (heldTokens *SafeHeldTokens) SafeRemoveToken(tokenContractAddress string) error {
    heldTokens.HtMu.Lock()
    for idx, val := range heldTokens.HeldTokensContractAddresses {

        if val == tokenContractAddress {
            heldTokens.HeldTokensContractAddresses = append(heldTokens.HeldTokensContractAddresses[:idx], 
                                                            heldTokens.HeldTokensContractAddresses[idx+1:]...)
            heldTokens.HtMu.Unlock()
            return nil
        }
    }

    heldTokens.HtMu.Unlock()
    return nil
}

// struct representing our submitted transactions that have 
// not been confirmed/ failed
type SafeOwnedPendingTransactions struct {
    PendingTxMu             sync.Mutex
    TxHashes                []string
    TokenContractAddresses  []string
}


func (transaction *SafeOwnedPendingTransactions) SafeAddOwnedPendingTx(ourTxHash string ,tradeData TxData, nonce *big.Int) {
    transaction.PendingTxMu.Lock()
    transaction.TokenContractAddresses = append(transaction.TokenContractAddresses, tradeData.TokenContractAddress.Hex())
    transaction.TxHashes               = append(transaction.TxHashes, ourTxHash)
    transaction.PendingTxMu.Unlock()
}

func (tx *SafeOwnedPendingTransactions) Contains(tokenAddress string) (bool, int) {
    tx.PendingTxMu.Lock()
    for i, v := range tx.TokenContractAddresses {
        if v == tokenAddress {
            tx.PendingTxMu.Unlock()
            return true, i
        }
    }
    tx.PendingTxMu.Unlock()
    return false, -1
}

// The reason this function removes items from it's respective slices like this is because
// using a techinque similar to the one above could prevent the GC from cleaning up
// the deleted item from the slice, as the original slice may still hold a reference. 
// This is why we need to zero out the value like this, to make sure it is garbage collected
func (transaction *SafeOwnedPendingTransactions) SafeRemoveOwnedPendingTx(idx int) {
    transaction.PendingTxMu.Lock()

    copy(transaction.TxHashes[idx:], transaction.TxHashes[idx+1:])
    transaction.TxHashes[len(transaction.TxHashes)-1] = ""
    transaction.TxHashes = transaction.TxHashes[:len(transaction.TxHashes) -1]

    copy(transaction.TokenContractAddresses[idx:], transaction.TokenContractAddresses[idx+1:])
    transaction.TokenContractAddresses[len(transaction.TokenContractAddresses)-1] = ""
    transaction.TokenContractAddresses = transaction.TokenContractAddresses[:len(transaction.TokenContractAddresses) -1]

    transaction.PendingTxMu.Unlock()
}
// Returns true if pendingtx contains data about the incoming trade and status is confirmed
// and returns the index of the trade in SafeOwnedPendingTransactions
func (tx *SafeOwnedPendingTransactions) SafeConfirmedOwnedTxHandler(txStatus string,txHash string) (bool, int) {
    tx.PendingTxMu.Lock()
    for i, v := range tx.TxHashes {
        fmt.Println("Inside SafeConfirmedOwnedTxHandler")

        fmt.Printf(" txData.TxHash: %v\n", txHash)
        fmt.Printf(" txData.Status: %v\n", txStatus)

        if strings.ToLower(v) == strings.ToLower(txHash) && txStatus == "confirmed"{
            fmt.Println("=========GOT WHERE YOU NEED TO BE!!=======")
            tx.PendingTxMu.Unlock()
            return true, i
        }

    }
    tx.PendingTxMu.Unlock()
    return false, -1
}

type SafeBumptedTxs struct {
    BumpedMu        sync.Mutex
    TxHashes        []string
    TargetAddresses []string
}

func (txBump *SafeBumptedTxs) SafeAddBumpedTx(hash string, address string) {
    txBump.BumpedMu.Lock()
    txBump.TxHashes        = append(txBump.TxHashes, hash)
    txBump.TargetAddresses = append(txBump.TargetAddresses, address)
    txBump.BumpedMu.Unlock()
}

type SafeNonce struct {
    NonceMu sync.Mutex
    Nonce   *big.Int
}

func (nonce *SafeNonce) SafeIncrementNonce() {
    nonce.NonceMu.Lock()
    nonce.Nonce.Add(nonce.Nonce, big.NewInt(1))
    nonce.NonceMu.Unlock()
}

type Trade struct {
    Token0               common.Address
    TokenContractAddress common.Address
    TokenPairAddress     common.Address
    TxGasPrice           *big.Int
    AmountToken          *big.Int
    AmountWETH           *big.Int
    TxHash               string
    ValueETH             string
    AmountIn             *big.Int
    AmountInMax          *big.Int
    AmountOut            *big.Int
    AmountOutMin         *big.Int
    Sell                 bool
}

type TxData struct {
    ReserveIn            *big.Int
    ReserveOut           *big.Int
    Status               string
    From                 string
    TimeStamp            string
    Deadline             *big.Int
    Token0               common.Address
    Reserve              Reserves
    TokenContractAddress common.Address
    TokenPairAddress     common.Address
    Path                 []string
    TxGasPrice           *big.Int
    AmountToken          *big.Int
    AmountWETH           *big.Int
    TxHash               string
    AmountIn             *big.Int
    AmountInMax          *big.Int
    AmountOut            *big.Int
    AmountOutMin         *big.Int
    Sell                 bool

}

func (tx TxData) String() string {
    return fmt.Sprintf(" deadline: %v\n timeStamp: %v\n Status: %v\n From: %v\n Token0: %v\n TokenContractAddress: %v\n" + 
    " TokenPairAddress: %v\n Path: %v\n TxGasPrice: %v\n AmountToken: %v\n" +
    " AmountWETH: %v\n TxHash: %v\n AmountIn: %v\n AmountInMax: %v\n" + 
    " AmountOut: %v\n AmountOutMin: %v\n Sell: %v\n Reserves: %v\n",tx.Deadline, tx.TimeStamp, tx.Status, tx.From,
    tx.Token0, tx.TokenContractAddress, tx.TokenPairAddress, tx.Path,tx.TxGasPrice, 
    tx.AmountToken, tx.AmountWETH, tx.TxHash, tx.AmountIn, tx.AmountInMax, 
    tx.AmountOut, tx.AmountOutMin, tx.Sell, tx.Reserve)
}

type Reserves struct {
    Token0 *big.Int
    Token1 *big.Int
}

type SafeTargetFirstTxs struct {
    TargetMu       sync.Mutex
    Addresses      []string
    TokenAddresses []string
    TxHashes       []string
}

func (targetFirstTxs *SafeTargetFirstTxs) Contains(address string, hash string) (bool, int) {
    targetFirstTxs.TargetMu.Lock()
    for i, v := range targetFirstTxs.Addresses {

        if v == address && targetFirstTxs.TxHashes[i] == hash {
            targetFirstTxs.TargetMu.Unlock()
            return true, i

        } else if v == address && targetFirstTxs.TxHashes[i] == hash {
            targetFirstTxs.TargetMu.Unlock()
            return false, -1
        }
    }
    targetFirstTxs.TargetMu.Unlock()
    return false, -1
}

func (targetFirstTxs *SafeTargetFirstTxs) SafeCheckBump(status string, address string, hash string) bool {
    targetFirstTxs.TargetMu.Lock()
    for i, v := range targetFirstTxs.Addresses {
        if v == address && targetFirstTxs.TxHashes[i] == hash && status == "failed"{
            targetFirstTxs.TargetMu.Unlock()
            return true

        } else if v == address && targetFirstTxs.TxHashes[i] == hash && status == "confirmed"{
            targetFirstTxs.TargetMu.Unlock()
            return false
        }
    }
    fmt.Println("!!!!!!Didn't find this transaction in SafeTargetFirstTxs struct!!!!!")
    targetFirstTxs.TargetMu.Unlock()
    return false
}


func (targetFirstTxs *SafeTargetFirstTxs) SafeAddTargetFirstTx(tokenContractAddress,
                                                               address string, 
                                                               txHash string) {
    targetFirstTxs.TargetMu.Lock()
    targetFirstTxs.Addresses      = append(targetFirstTxs.Addresses, address)
    targetFirstTxs.TxHashes       = append(targetFirstTxs.TxHashes, txHash)
    targetFirstTxs.TokenAddresses = append(targetFirstTxs.TokenAddresses, tokenContractAddress)
    targetFirstTxs.TargetMu.Unlock()
}

func (targetAddresses *SafeTargetFirstTxs) SafeRemoveTargetFirstTx(address string) error  {
    targetAddresses.TargetMu.Lock()
    for i, val := range targetAddresses.Addresses {
        if val == address {
            targetAddresses.Addresses = append(targetAddresses.Addresses[:i],
                                               targetAddresses.Addresses[i+1:]...)
            targetAddresses.TxHashes  = append(targetAddresses.TxHashes[:i],
                                               targetAddresses.TxHashes[i+1:]...)
            targetAddresses.TargetMu.Unlock()
            return nil

        } 
    }
    targetAddresses.TargetMu.Unlock()
    return nil
}

type SafeTargetSecondTxs struct {
    TargetMu  sync.Mutex
    Addresses []string
    TxHashes  []string
}

func (targetAddresses *SafeTargetSecondTxs) SafeAddTargetSecondTx(address string, txHash string) {
    targetAddresses.TargetMu.Lock()
    targetAddresses.Addresses = append(targetAddresses.Addresses, address)
    targetAddresses.TxHashes  = append(targetAddresses.TxHashes, txHash)
    targetAddresses.TargetMu.Unlock()
}

func (targetAddresses *SafeTargetSecondTxs) SafeRemoveTargetSecondTx(address string) error  {
    targetAddresses.TargetMu.Lock()
    for i, val := range targetAddresses.Addresses {
        if val == address {
            targetAddresses.Addresses = append(targetAddresses.Addresses[:i],
                                               targetAddresses.Addresses[i+1:]...)
            targetAddresses.TxHashes  = append(targetAddresses.TxHashes[:i],
                                               targetAddresses.TxHashes[i+1:]...)
            targetAddresses.TargetMu.Unlock()
            return nil

        } 
    }
    targetAddresses.TargetMu.Unlock()
    return nil
}

func (targetAddresses *SafeTargetFirstTxs) IsTargetSecondBuy(tokenContractAddress string,
                                                             address string) bool {
    targetAddresses.TargetMu.Lock()
    for i, val := range targetAddresses.Addresses {

        if strings.ToLower(val) == address && 
            strings.ToLower(targetAddresses.TokenAddresses[i]) == tokenContractAddress {
            fmt.Println("Got into IsTargetSecondBuy")
            fmt.Printf("i: %v\n val: %v\n address(from args): %v\n" +
                       "targetAddresses.TokenAddresses[i]: %v\n tokenContractAddress(from args): %v\n",
                       i, val, address, targetAddresses.TokenAddresses[i], tokenContractAddress)
            targetAddresses.TargetMu.Unlock()
            return true
        } 

    }
    targetAddresses.TargetMu.Unlock()
    return false
}


type BBBuy struct {
    Token0               string
    TxGasPrice           *big.Int
    TxHash               string 
    TokenContractAddress string
    TokenPairAddress     string
    Sell                 bool
}

type MonitoredTokens struct {

}

