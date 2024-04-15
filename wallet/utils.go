package wallet

func IsAllowedWalletType(walletType string) bool {
	for _, v := range AllowedWalletTypes {
		if v == walletType {
			return true
		}
	}
	return false
}
