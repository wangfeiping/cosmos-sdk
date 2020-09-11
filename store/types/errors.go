package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const StoreCodespace = "store"

var (
	ErrInvalidProof = sdkerrors.Register(StoreCodespace, 2, "invalid proof")

	ErrJustMock = sdkerrors.Register(StoreCodespace, 1000, "it's just a mock")
)
