func (b *BroadcasterAccount) GetAccountWithHeight(clienCtx client.Context, addr sdk.AccAddress) (client.Account, int64, error) {
	var header metadata.MD
	queryClient := authtypes.NewQueryClient(b.rpcClient)
	res, err := queryClient.Account(clienCtx.CmdContext, &authtypes.QueryAccountRequest{Address: b.addressString}, grpc.Header(&header))
	if err != nil {
		return nil, 0, err
	}
//Code Fix: Replace the hardcoded b.addressString with the dynamic addr.String():
	//res, err := queryClient.Account(clientCtx.CmdContext, &authtypes.QueryAccountRequest{Address: addr.String()}, grpc.Header(&header))
	//if err != nil {