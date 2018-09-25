package executor


import (
	"gitlab.33.cn/chain33/chain33/account"
	"gitlab.33.cn/chain33/chain33/common/address"
	"gitlab.33.cn/chain33/chain33/plugin/dapp/evm/executor/vm/common"
	"gitlab.33.cn/chain33/chain33/types"
	"github.com/pkg/errors"
	"gitlab.33.cn/chain33/chain33/common/db"
)

func (a *action) assetTransfer(transfer *types.AssetsTransfer) (*types.Receipt, error) {
	isPara := types.IsPara()
	if !isPara {
		accDB, err := createAccount(a.db, transfer.Cointoken)
		if err != nil {
			return nil, errors.Wrap(err, "assetTransferToken call account.NewAccountDB failed")
		}
		execAddr := address.ExecAddress(types.ParaX)
		fromAcc := accDB.LoadExecAccount(a.fromaddr, execAddr)
		if fromAcc.Balance < transfer.Amount {
			return nil, types.ErrNoBalance
		}
		toAddr := address.ExecAddress(string(a.tx.Execer))
		clog.Debug("paracross.AssetTransfer not isPara", "execer", string(a.tx.Execer),
			"txHash", common.Bytes2Hex(a.tx.Hash()))
		return accDB.ExecTransfer(a.fromaddr, toAddr, execAddr, transfer.Amount)
	} else {
		paraTitle, err := getTitleFrom(a.tx.Execer)
		if err != nil {
			return nil, errors.Wrap(err, "assetTransferCoins call getTitleFrom failed")
		}
		var paraAcc *account.DB
		if transfer.Cointoken == "" {
			paraAcc, err = NewParaAccount(string(paraTitle), types.CoinsX, "bty", a.db)
		} else {
			paraAcc, err = NewParaAccount(string(paraTitle), types.TokenX, transfer.Cointoken, a.db)
		}
		if err != nil {
			return nil, errors.Wrap(err, "assetTransferCoins call NewParaAccount failed")
		}
		clog.Debug("paracross.AssetTransfer isPara", "execer", string(a.tx.Execer),
			"txHash", common.Bytes2Hex(a.tx.Hash()))
		return assetDepositBalance(paraAcc, transfer.To, transfer.Amount)
	}
}


func (a *action) assetWithdrawCoins(withdraw *types.AssetsWithdraw, withdrawTx *types.Transaction) (*types.Receipt, error) {
	isPara := types.IsPara()
	if !isPara {
		accDB, err := createAccount(a.db, withdraw.Cointoken)
		if err != nil {
			return nil, errors.Wrap(err, "assetWithdrawCoins call account.NewAccountDB failed")
		}
		fromAddr := address.ExecAddress(string(withdrawTx.Execer))
		execAddr := address.ExecAddress(types.ParaX)
		clog.Debug("Paracross.Exec", "AssettWithdraw", withdraw.Amount, "from", fromAddr,
			"to", withdraw.To, "exec", execAddr, "withdrawTx execor", string(withdrawTx.Execer))
		return accDB.ExecTransfer(fromAddr, withdraw.To, execAddr, withdraw.Amount)
	} else {
		paraTitle, err := getTitleFrom(a.tx.Execer)
		if err != nil {
			return nil, errors.Wrap(err, "assetWithdrawCoins call getTitleFrom failed")
		}
		var paraAcc *account.DB
		if withdraw.Cointoken == "" {
			paraAcc, err = NewParaAccount(string(paraTitle), types.CoinsX, "bty", a.db)
		} else {
			paraAcc, err = NewParaAccount(string(paraTitle), types.TokenX, withdraw.Cointoken, a.db)
		}
		if err != nil {
			return nil, errors.Wrap(err, "assetWithdrawCoins call NewParaAccount failed")
		}
		clog.Debug("paracross.assetWithdrawCoins isPara", "execer", string(a.tx.Execer),
			"txHash", common.Bytes2Hex(a.tx.Hash()))
		return assetWithdrawBalance(paraAcc, a.fromaddr, withdraw.Amount)
	}
}

func createAccount(db db.KV, symbol string) (*account.DB, error) {
	var accDB *account.DB
	var err error
	if symbol == "" {
		accDB = account.NewCoinsAccount()
		accDB.SetDB(db)
	} else {
		accDB, err = account.NewAccountDB(types.TokenX, symbol, db)
	}
	return accDB, err
}
