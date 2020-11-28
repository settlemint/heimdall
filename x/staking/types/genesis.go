package types

import (
	"encoding/json"
	"errors"

	"github.com/cosmos/cosmos-sdk/codec"

	hmTypes "github.com/maticnetwork/heimdall/types"
)

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	return nil
}

func (gs *GenesisState) MarshalJSON() ([]byte, error) {
	return json.Marshal(gs.String())
}

func (gs *GenesisState) UnMarshalJSON(data []byte) error {
	return json.Unmarshal(data, &gs)
}

// HeimdallValidator converts genesis validator validator to Heimdall validator
func (v *GenesisValidator) HeimdallValidator() hmTypes.Validator {
	return hmTypes.Validator{
		ID:          v.ID,
		PubKey:      v.PubKey,
		VotingPower: int64(v.Power),
		StartEpoch:  v.StartEpoch,
		EndEpoch:    v.EndEpoch,
		Nonce:       v.Nonce,
		Signer:      v.Signer,
	}
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(
	validators []*hmTypes.Validator,
	currentValSet *hmTypes.ValidatorSet,
	stakingSequences []string,
) *GenesisState {
	return &GenesisState{
		Validators:       validators,
		CurrentValSet:    currentValSet,
		StakingSequences: stakingSequences,
	}
}

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return NewGenesisState(nil, &hmTypes.ValidatorSet{}, nil)
}

// ValidateGenesis performs basic validation of bor genesis data returning an
// error for any failed validation criteria.
func ValidateGenesis(data GenesisState) error {
	for _, validator := range data.Validators {
		if !validator.ValidateBasic() {
			return errors.New("Invalid validator")
		}
	}
	for _, sq := range data.StakingSequences {
		if sq == "" {
			return errors.New("Invalid Sequence")
		}
	}

	return nil
}


// GetGenesisStateFromAppState returns x/staking GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONMarshaler, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}

// SetGenesisStateToAppState sets state into app state
func SetGenesisStateToAppState(cdc codec.JSONMarshaler, appState map[string]json.RawMessage, validators []*hmTypes.Validator, currentValSet *hmTypes.ValidatorSet) (map[string]json.RawMessage, error) {
	// set state to staking state
	stakingState := GetGenesisStateFromAppState(cdc, appState)
	stakingState.Validators = validators
	stakingState.CurrentValSet = currentValSet

	appState[ModuleName] = cdc.MustMarshalJSON(stakingState)
	return appState, nil
}
