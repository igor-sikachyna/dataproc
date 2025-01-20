package dataproc

import (
// fmt "fmt"

// "cosmossdk.io/errors"
// sdk "github.com/cosmos/cosmos-sdk/types"
)

func (storedCode StoredCode) CompileCode() (runtime string, err error) {
	// board, errBoard := rules.Parse(storedGame.Board)
	// if errBoard != nil {
	// 	return nil, errors.Wrapf(errBoard, ErrGameNotParseable.Error())
	// }
	// board.Turn = rules.StringPieces[storedGame.Turn].Player
	// if board.Turn.Color == "" {
	// 	return nil, errors.Wrapf(fmt.Errorf("turn: %s", storedGame.Turn), ErrGameNotParseable.Error())
	// }
	// return board, nil
	return "", nil
}

func (storedCode StoredCode) Validate() (err error) {
	// _, err = storedGame.GetBlackAddress()
	// if err != nil {
	// 	return err
	// }
	// _, err = storedGame.GetRedAddress()
	// if err != nil {
	// 	return err
	// }
	// _, err = storedGame.ParseGame()
	// return err
	return nil
}
