package pdu

import (
	"frostmourne/lib/data"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestEnquireLink(t *testing.T) {
	v := NewEnquireLink().(*EnquireLink)
	require.True(t, v.CanResponse())
	v.SequenceNumber = 13

	validate(t,
		v.GetResponse(),
		"0000001080000015000000000000000d",
		data.ENQUIRE_LINK_RESP,
	)

	validate(t,
		v,
		"0000001000000015000000000000000d",
		data.ENQUIRE_LINK,
	)
}