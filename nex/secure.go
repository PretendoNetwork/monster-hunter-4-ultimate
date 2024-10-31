package nex

import (
	"fmt"
	"os"
	"strconv"

	nex "github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/monster-hunter-4-ultimate/globals"
)

func StartSecureServer() {
	globals.SecureServer = nex.NewPRUDPServer()

	globals.SecureEndpoint = nex.NewPRUDPEndPoint(1)
	globals.SecureEndpoint.IsSecureEndPoint = true
	globals.SecureEndpoint.ServerAccount = globals.SecureServerAccount
	globals.SecureEndpoint.AccountDetailsByPID = globals.AccountDetailsByPID
	globals.SecureEndpoint.AccountDetailsByUsername = globals.AccountDetailsByUsername
	globals.SecureEndpoint.DefaultStreamSettings.MaxSilenceTime = 10000
	globals.SecureServer.BindPRUDPEndPoint(globals.SecureEndpoint)

	globals.SecureServer.LibraryVersions.SetDefault(nex.NewLibraryVersion(3, 6, 1))
	globals.SecureServer.ByteStreamSettings.UseStructureHeader = true
	globals.SecureServer.AccessKey = "d44c6198"

	globals.SecureEndpoint.OnData(func(packet nex.PacketInterface) {
		request := packet.RMCMessage()

		fmt.Println("=== Monster Hunter 4 Ultimate - Secure ===")
		fmt.Printf("Protocol ID: %#v\n", request.ProtocolID)
		fmt.Printf("Method ID: %#v\n", request.MethodID)
		fmt.Println("==================================")
	})

	globals.SecureEndpoint.OnError(func(err *nex.Error) {
		globals.Logger.Error(err.Error())
	})

	registerCommonSecureServerProtocols()

	port, _ := strconv.Atoi(os.Getenv("PN_MH4U_SECURE_SERVER_PORT"))

	globals.SecureServer.Listen(port)
}
