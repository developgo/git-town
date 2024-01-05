package configdomain

import (
	"fmt"
	"strconv"

	"github.com/git-town/git-town/v11/src/gohacks"
	"github.com/git-town/git-town/v11/src/messages"
)

type SyncBeforeShip bool

func (self SyncBeforeShip) Bool() bool {
	return bool(self)
}

func (self SyncBeforeShip) String() string {
	return strconv.FormatBool(self.Bool())
}

func NewSyncBeforeShipRef(value bool) *SyncBeforeShip {
	result := SyncBeforeShip(value)
	return &result
}

func ParseSyncBeforeShipRef(value, source string) (*SyncBeforeShip, error) {
	parsed, err := gohacks.ParseBool(value)
	if err != nil {
		return nil, fmt.Errorf(messages.ValueInvalid, source, value)
	}
	token := SyncBeforeShip(parsed)
	return &token, nil
}
