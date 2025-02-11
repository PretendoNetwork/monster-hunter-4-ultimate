package nex

import (
	"github.com/PretendoNetwork/monster-hunter-4-ultimate/database"
	"github.com/PretendoNetwork/monster-hunter-4-ultimate/globals"
	"github.com/PretendoNetwork/nex-go/v2/types"
	match_making_types "github.com/PretendoNetwork/nex-protocols-go/v2/match-making/types"
	nex_matchmake_extension "github.com/PretendoNetwork/monster-hunter-4-ultimate/nex/matchmake-extension"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	common_match_making "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making"
	common_match_making_ext "github.com/PretendoNetwork/nex-protocols-common-go/v2/match-making-ext"
	common_matchmake_extension "github.com/PretendoNetwork/nex-protocols-common-go/v2/matchmake-extension"
	common_nat_traversal "github.com/PretendoNetwork/nex-protocols-common-go/v2/nat-traversal"
	common_secure "github.com/PretendoNetwork/nex-protocols-common-go/v2/secure-connection"
	match_making "github.com/PretendoNetwork/nex-protocols-go/v2/match-making"
	match_making_ext "github.com/PretendoNetwork/nex-protocols-go/v2/match-making-ext"
	matchmake_extension "github.com/PretendoNetwork/nex-protocols-go/v2/matchmake-extension"
	nat_traversal "github.com/PretendoNetwork/nex-protocols-go/v2/nat-traversal"
	secure "github.com/PretendoNetwork/nex-protocols-go/v2/secure-connection"
)

func registerCommonSecureServerProtocols() {
	secureProtocol := secure.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(secureProtocol)
	common_secure.NewCommonProtocol(secureProtocol)

	natTraversalProtocol := nat_traversal.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(natTraversalProtocol)
	common_nat_traversal.NewCommonProtocol(natTraversalProtocol)

	matchmakingManager := common_globals.NewMatchmakingManager(globals.SecureEndpoint, database.Postgres)

	matchMakingProtocol := match_making.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingProtocol)
	commonMatchMakingProtocol := common_match_making.NewCommonProtocol(matchMakingProtocol)
	commonMatchMakingProtocol.SetManager(matchmakingManager)

	matchMakingExtProtocol := match_making_ext.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchMakingExtProtocol)
	commonMatchMakingExtProtocol := common_match_making_ext.NewCommonProtocol(matchMakingExtProtocol)
	commonMatchMakingExtProtocol.SetManager(matchmakingManager)

	matchmakeExtensionProtocol := matchmake_extension.NewProtocol()
	globals.SecureEndpoint.RegisterServiceProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol := common_matchmake_extension.NewCommonProtocol(matchmakeExtensionProtocol)
	commonMatchmakeExtensionProtocol.SetManager(matchmakingManager)
	matchmakeExtensionProtocol.SetHandlerGetMyBlockList(nex_matchmake_extension.GetMyBlockList)
	commonMatchmakeExtensionProtocol.CleanupMatchmakeSessionSearchCriterias = func(searchCriterias types.List[match_making_types.MatchmakeSessionSearchCriteria]) {
		for i := range searchCriterias {
			// * Index 2 is related to the region. Clear it to allow MH4U and MH4G users to play together
			searchCriterias[i].Attribs[2] = ""
		}
	}
}
