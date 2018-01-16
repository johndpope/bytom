package account

import (
	"context"
	"time"

	"github.com/bytom/blockchain/txbuilder"
	"github.com/bytom/errors"
)

const defaultReceiverExpiry = 30 * 24 * time.Hour // 30 days

// CreateReceiver creates a new account receiver for an account
// with the provided expiry. If a zero time is provided for the
// expiry, a default expiry of 30 days from the current time is
// used.
func (m *Manager) CreateReceiver(ctx context.Context, accountInfo string) (*txbuilder.Receiver, error) {
	expiresAt := time.Now().Add(defaultReceiverExpiry)

	accountID := accountInfo

	if s, err := m.FindByAlias(ctx, accountInfo); err == nil {
		accountID = s.ID
	}

	cp, err := m.CreateControlProgram(ctx, accountID, false, expiresAt)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	receiver := &txbuilder.Receiver{
		ControlProgram: cp,
		ExpiresAt:      expiresAt,
	}

	return receiver, nil
}

// CreateAddressReceiver creates a new address receiver for an account
func (m *Manager) CreateAddressReceiver(ctx context.Context, accountInfo string) (*txbuilder.Receiver, error) {
	accountID := accountInfo
	if s, err := m.FindByAlias(ctx, accountInfo); err == nil {
		accountID = s.ID
	}

	program, err := m.CreateAddress(ctx, accountID, false)
	if err != nil {
		return nil, err
	}

	return &txbuilder.Receiver{
		ControlProgram: program.ControlProgram,
		Address:        program.Address,
		ExpiresAt:      program.ExpiresAt,
	}, nil
}

func (m *Manager) CreatePubkeyInfo(ctx context.Context, accountInfo string) (*AccountPubkey, error) {
	accountID := accountInfo
	if s, err := m.FindByAlias(ctx, accountInfo); err == nil {
		accountID = s.ID
	}

	accountPubkey, err := m.createPubkey(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return &AccountPubkey{
		Root:      	accountPubkey.Root,
		Pubkey:    	accountPubkey.Pubkey,
		Path: 		accountPubkey.Path,
		Index:		accountPubkey.Index,
	}, nil
}