package info

import (
	"strconv"
	"encoding/base32"
	"fmt"
)

type UserBalance struct {
	UserInfo
}

const (
	UserBalance_Encoding_Head        = "[USER'S BALANCE]"
	UserBalance_Encoding_Tail        = "[/USER'S BALANCE]"
	UserBalance_Encoding_Head_Length = len(UserBalance_Encoding_Head)
	UserBalance_Encoding_Tail_Length = len(UserBalance_Encoding_Tail)
)

func (balance *UserBalance) Encode(val uint64) string {
	raw := ([]byte)(UserBalance_Encoding_Head + strconv.FormatUint(val, 10) + UserBalance_Encoding_Tail)
	return base32.HexEncoding.EncodeToString(raw) //or other encoding forms
}

func (balance *UserBalance) Decode(str string) (uint64, error) {
	raw, err := base32.HexEncoding.DecodeString(str)
	if err != nil {
		return 0, fmt.Errorf("Failed in decoding.")
	} else if string(raw[:UserBalance_Encoding_Head_Length]) != UserBalance_Encoding_Head {
		return 0, fmt.Errorf("Data format error: it's not a valid user balance data")
	} else if string(raw[len(raw)-UserBalance_Encoding_Tail_Length:]) != UserBalance_Encoding_Tail {
		return 0, fmt.Errorf("Data format error: it's not a valid user balance data")
	} else {
		val, err := strconv.ParseUint(string(raw[UserBalance_Encoding_Head_Length:len(raw)-UserBalance_Encoding_Tail_Length]), 10, 64)
		if err != nil {
			return 0, fmt.Errorf("Data format error: it's not a valid user balance data")
		} else {
			return val, nil
		}
	}
}

