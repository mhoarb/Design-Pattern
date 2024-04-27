package main

import "fmt"

type WalletFacade struct {
	account      *Account
	wallet       *Wallet
	securityCode *SecurityCode
	notification *Notification
	ledger       *Ledger
}

func NewWaalletFacade(accountID string, code int) *WalletFacade {
	fmt.Println("starting create walletFacade")
	walletFacade := &WalletFacade{
		account:      NewAccount(accountID),
		wallet:       NewWallet(),
		securityCode: NewSwcurityCode(code),
		notification: &Notification{},
		ledger:       &Ledger{},
	}
	fmt.Println("account created")
	return walletFacade
}
func (w *WalletFacade) AddMoneyToWallet(accountID string, code int, amount int) error {
	fmt.Println("starting create account")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.CheckCode(code)
	if err != nil {
		return err
	}
	w.wallet.CreditBalance(amount)
	w.notification.WalletCreditNotif()
	w.ledger.MakeEntry(accountID, "credit", amount)
	return nil
}

func (w *WalletFacade) deductMoneyfromWallet(accountID string, code, amount int) error {
	fmt.Println("starting debit money from wallet")
	err := w.account.CheckAccount(accountID)
	if err != nil {
		return err
	}
	err = w.securityCode.CheckCode(code)
	if err != nil {
		return err
	}
	err = w.wallet.DebitBalance(amount)
	if err != nil {
		return err
	}
	w.notification.WalletDebitNotif()
	w.ledger.MakeEntry(accountID, "debit", amount)
	return nil

}
