package oscar

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	ICBMErr                uint16 = 0x0001
	ICBMAddParameters             = 0x0002
	ICBMDelParameters             = 0x0003
	ICBMParameterQuery            = 0x0004
	ICBMParameterReply            = 0x0005
	ICBMChannelMsgTohost          = 0x0006
	ICBMChannelMsgToclient        = 0x0007
	ICBMEvilRequest               = 0x0008
	ICBMEvilReply                 = 0x0009
	ICBMMissedCalls               = 0x000A
	ICBMClientErr                 = 0x000B
	ICBMHostAck                   = 0x000C
	ICBMSinStored                 = 0x000D
	ICBMSinListQuery              = 0x000E
	ICBMSinListReply              = 0x000F
	ICBMSinRetrieve               = 0x0010
	ICBMSinDelete                 = 0x0011
	ICBMNotifyRequest             = 0x0012
	ICBMNotifyReply               = 0x0013
	ICBMClientEvent               = 0x0014
	ICBMSinReply                  = 0x0017
)

func routeICBM(flap *flapFrame, snac *snacFrame, r io.Reader, w io.Writer, sequence uint16) error {
	switch snac.subGroup {
	case ICBMErr:
		panic("not implemented")
	case ICBMAddParameters:
		panic("not implemented")
	case ICBMDelParameters:
		panic("not implemented")
	case ICBMParameterQuery:
		return SendAndReceiveICBMParameterReply(flap, snac, r, w, sequence)
	case ICBMChannelMsgTohost:
		panic("not implemented")
	case ICBMChannelMsgToclient:
		panic("not implemented")
	case ICBMEvilRequest:
		panic("not implemented")
	case ICBMEvilReply:
		panic("not implemented")
	case ICBMMissedCalls:
		panic("not implemented")
	case ICBMClientErr:
		panic("not implemented")
	case ICBMHostAck:
		panic("not implemented")
	case ICBMSinStored:
		panic("not implemented")
	case ICBMSinListQuery:
		panic("not implemented")
	case ICBMSinListReply:
		panic("not implemented")
	case ICBMSinRetrieve:
		panic("not implemented")
	case ICBMSinDelete:
		panic("not implemented")
	case ICBMNotifyRequest:
		panic("not implemented")
	case ICBMNotifyReply:
		panic("not implemented")
	case ICBMClientEvent:
		panic("not implemented")
	case ICBMSinReply:
		panic("not implemented")
	}

	return nil
}

type snacParameterRequest struct {
	maxSlots             uint16
	ICBMFlags            uint32
	maxIncomingICBMLen   uint16
	maxSourceEvil        uint16
	maxDestinationEvil   uint16
	minInterICBMInterval uint32
}

func (s *snacParameterRequest) write(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, s.maxSlots); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, s.ICBMFlags); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, s.maxIncomingICBMLen); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, s.maxSourceEvil); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, s.maxDestinationEvil); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, s.minInterICBMInterval); err != nil {
		return err
	}
	return nil
}

func SendAndReceiveICBMParameterReply(flap *flapFrame, snac *snacFrame, _ io.Reader, w io.Writer, sequence uint16) error {
	fmt.Printf("sendAndReceiveICBMParameterReply read SNAC frame: %+v\n", snac)

	snacFrameOut := snacFrame{
		foodGroup: ICBM,
		subGroup:  ICBMParameterReply,
	}
	snacPayloadOut := &snacParameterRequest{
		maxSlots:             100,
		ICBMFlags:            0x00000001,
		maxIncomingICBMLen:   8000,
		maxSourceEvil:        999,
		maxDestinationEvil:   999,
		minInterICBMInterval: 100,
	}

	return writeOutSNAC(flap, snacFrameOut, snacPayloadOut, sequence, w)
}