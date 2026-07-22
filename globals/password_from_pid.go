package globals

import (
	"context"

	pb_account "github.com/PretendoNetwork/grpc/go/account/v2"
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	"github.com/PretendoNetwork/nex-protocols-go/v2/globals"
	"google.golang.org/grpc/metadata"
)

func PasswordFromPID(pid types.PID) (string, uint32) {
	ctx := metadata.NewOutgoingContext(context.Background(), common_globals.GRPCAccountCommonMetadata)

	response, err := common_globals.GRPCAccountClient.GetNEXPassword(ctx, &pb_account.GetNEXPasswordRequest{Pid: uint32(pid)})
	if err != nil {
		globals.Logger.Error(err.Error())
		return "", nex.ResultCodes.RendezVous.InvalidUsername
	}

	return response.Password, 0
}
