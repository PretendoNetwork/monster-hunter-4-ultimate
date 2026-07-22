package main

import (
	"crypto/rand"
	"os"
	"strconv"
	"strings"

	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	common_globals "github.com/PretendoNetwork/nex-protocols-common-go/v2/globals"
	"github.com/PretendoNetwork/plogger-go"
	"github.com/joho/godotenv"

	"github.com/PretendoNetwork/monster-hunter-4-ultimate/database"
	"github.com/PretendoNetwork/monster-hunter-4-ultimate/globals"
)

func init() {
	globals.Logger = plogger.NewLogger()

	var err error

	err = godotenv.Load()
	if err != nil {
		globals.Logger.Warning("Error loading .env file")
	}

	postgresURI := os.Getenv("PN_MH4U_POSTGRES_URI")
	authenticationServerPort := os.Getenv("PN_MH4U_AUTHENTICATION_SERVER_PORT")
	secureServerHost := os.Getenv("PN_MH4U_SECURE_SERVER_HOST")
	secureServerPort := os.Getenv("PN_MH4U_SECURE_SERVER_PORT")
	accountGRPCHost := os.Getenv("PN_MH4U_ACCOUNT_GRPC_HOST")
	accountGRPCPort := os.Getenv("PN_MH4U_ACCOUNT_GRPC_PORT")
	accountGRPCAPIKey := os.Getenv("PN_MH4U_ACCOUNT_GRPC_API_KEY")

	if strings.TrimSpace(postgresURI) == "" {
		globals.Logger.Error("PN_MH4U_POSTGRES_URI environment variable not set")
		os.Exit(0)
	}

	kerberosPassword := make([]byte, 0x10)
	_, err = rand.Read(kerberosPassword)
	if err != nil {
		globals.Logger.Error("Error generating Kerberos password")
		os.Exit(0)
	}

	globals.KerberosPassword = string(kerberosPassword)

	globals.AuthenticationServerAccount = nex.NewAccount(types.NewPID(1), "Quazal Authentication", globals.KerberosPassword, false)
	globals.SecureServerAccount = nex.NewAccount(types.NewPID(2), "Quazal Rendez-Vous", globals.KerberosPassword, false)

	if strings.TrimSpace(authenticationServerPort) == "" {
		globals.Logger.Error("PN_MH4U_AUTHENTICATION_SERVER_PORT environment variable not set")
		os.Exit(0)
	}

	if port, err := strconv.Atoi(authenticationServerPort); err != nil {
		globals.Logger.Errorf("PN_MH4U_AUTHENTICATION_SERVER_PORT is not a valid port. Expected 0-65535, got %s", authenticationServerPort)
		os.Exit(0)
	} else if port < 0 || port > 65535 {
		globals.Logger.Errorf("PN_MH4U_AUTHENTICATION_SERVER_PORT is not a valid port. Expected 0-65535, got %s", authenticationServerPort)
		os.Exit(0)
	}

	if strings.TrimSpace(secureServerHost) == "" {
		globals.Logger.Error("PN_MH4U_SECURE_SERVER_HOST environment variable not set")
		os.Exit(0)
	}

	if strings.TrimSpace(secureServerPort) == "" {
		globals.Logger.Error("PN_MH4U_SECURE_SERVER_PORT environment variable not set")
		os.Exit(0)
	}

	if port, err := strconv.Atoi(secureServerPort); err != nil {
		globals.Logger.Errorf("PN_MH4U_SECURE_SERVER_PORT is not a valid port. Expected 0-65535, got %s", secureServerPort)
		os.Exit(0)
	} else if port < 0 || port > 65535 {
		globals.Logger.Errorf("PN_MH4U_SECURE_SERVER_PORT is not a valid port. Expected 0-65535, got %s", secureServerPort)
		os.Exit(0)
	}

	if strings.TrimSpace(accountGRPCHost) == "" {
		globals.Logger.Error("PN_MH4U_ACCOUNT_GRPC_HOST environment variable not set")
		os.Exit(0)
	}

	if strings.TrimSpace(accountGRPCPort) == "" {
		globals.Logger.Error("PN_MH4U_ACCOUNT_GRPC_PORT environment variable not set")
		os.Exit(0)
	}

	accountPort, err := strconv.Atoi(accountGRPCPort)
	if err != nil {
		globals.Logger.Errorf("PN_MH4U_ACCOUNT_GRPC_PORT is not a valid port. Expected 0-65535, got %s", accountGRPCPort)
		os.Exit(0)
	} else if accountPort < 0 || accountPort > 65535 {
		globals.Logger.Errorf("PN_MH4U_ACCOUNT_GRPC_PORT is not a valid port. Expected 0-65535, got %s", accountGRPCPort)
		os.Exit(0)
	}

	if strings.TrimSpace(accountGRPCAPIKey) == "" {
		globals.Logger.Warning("Insecure gRPC server detected. PN_MH4U_ACCOUNT_GRPC_API_KEY environment variable not set")
	}

	common_globals.ConnectToAccountGRPC(accountGRPCHost, uint16(accountPort), accountGRPCAPIKey)

	database.ConnectPostgres()
}
