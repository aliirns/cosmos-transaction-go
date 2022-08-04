# cosmos-transaction-go

![](https://img.shields.io/badge/golang-1.18+-blue.svg?style=flat)

cosmos-transaction-go performs grpc transactions with [Cosmos-SDK ](https://github.com/cosmos/cosmos-sdk) chains

# Installation

```
go get github.com/aliirns/cosmos-transaction-go

```

# Quick Start

```

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
    pylonsApp "github.com/Pylons-tech/pylons/app"
)

const (
	grpcURL  = "127.0.0.1:9090"
	address  = "<address>"
	privKey  = "<exported private key>"
	chainID  = "pylons-testnet-1"
)



func main(){
    config := pylonsApp.DefaultConfig()
    types.NewMsgSend(sdk.AccAddress(addr), sdk.AccAddress(receiver), coins)
    res, err := CosmosTx(address, privKey, grpcURL, msg, chainID)
    fmt.Println(res.TxResponse.Code)
}
```

## Exporting Private Key in Hex

`chaind keys export alii --unsafe --unarmored-hex`

# License

cosmos-transaction-go is licensed under the MIT.
