package utils

import (
	"github.com/pomcom/bagoScan/pkg/tools"
	"github.com/pomcom/bagoScan/pkg/tools/nmap"
	"github.com/pomcom/bagoScan/pkg/tools/testssl"
)

// Only place new supported tools need to be added (for the full scan)
func GetSupportedTools() []tools.Tool {
	return []tools.Tool{
		testssl.Testssl{},
		nmap.Nmap{},
	}
}

// map tools from config to implemented tools here?
