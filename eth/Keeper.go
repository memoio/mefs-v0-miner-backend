package eth

type KeeperInfo struct{
	Addr string `json:"addr"`
	ProviderList []ProviderInfo `json:"providerList"`
}