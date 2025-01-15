//Proof of Concept - Code Reference and Context
//Here’s a simplified version of the code highlighting the issue:
func (ch *Child) Initialize(
	ctx types.Context,
	processedHeight int64,
	startOutputIndex uint64,
	host hostNode,
	bridgeInfo ophosttypes.QueryBridgeResponse,
	challenger challenger,
) (time.Time, error) {
	ch.host = host         // Potential nil assignment
	ch.challenger = challenger // Potential nil assignment
	ch.registerHandlers()
}

// If host or challenger is nil, any call like ch.host.SomeMethod() or ch.challenger.SomeMethod() will cause a runtime panic.

// Reproducing the Bug -Consider a scenario where the Initialize function is called with nil for host or challenger:
var nilHost hostNode
var nilChallenger challenger

ch := &Child{}
_, err := ch.Initialize(ctx, 100, 0, nilHost, bridgeInfo, nilChallenger)
if err != nil {
    log.Fatalf("Initialization failed: %v", err)
}

// This would panic due to nil pointer dereference
ch.host.SomeMethod()

//If this code runs, it produces the following stack trace:

panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x0 pc=0x123456]

goroutine 1:
main.(*Child).Initialize()
    example.go:20
main.main()
    example.go:30
