package types

import (
	"fmt"

	"github.com/tendermint/tendermint/crypto"

	sdk "github.com/cosmos/cosmos-sdk/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/24-host"
)

const (
	// ModuleName defines the IBC transfer name
	ModuleName = "transfer"

	// Version defines the current version the IBC tranfer
	// module supports
	Version = "ics20-1"

	// PortID is the default port id that transfer module binds to
	PortID = "transfer"

	// StoreKey is the store key string for IBC transfer
	StoreKey = ModuleName

	// RouterKey is the message route for IBC transfer
	RouterKey = ModuleName

	// QuerierRoute is the querier route for IBC transfer
	QuerierRoute = ModuleName
)

// PortKey defines the key to store the port ID in store
var PortKey = []byte{0x01}

// GetEscrowAddress returns the escrow address for the specified channel
//
// CONTRACT: this assumes that there's only one bank bridge module that owns the
// port associated with the channel ID so that the address created is actually
// unique.
func GetEscrowAddress(portID, channelID string) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(portID + channelID)))
}

// GetDenomPrefix returns the receiving denomination prefix
func GetDenomPrefix(portID, channelID string) string {
	return fmt.Sprintf("%s/%s/", portID, channelID)
}

// GetModuleAccountName returns the IBC transfer module account name for supply
func GetModuleAccountName() string {
	return fmt.Sprintf("%s/%s", host.ModuleName, ModuleName)
}
