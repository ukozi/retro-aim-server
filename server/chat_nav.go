package server

import (
	"encoding/binary"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/mkaminski/goaim/oscar"
	"io"
	"time"
)

const (
	ChatNavErr                 uint16 = 0x0001
	ChatNavRequestChatRights          = 0x0002
	ChatNavRequestExchangeInfo        = 0x0003
	ChatNavRequestRoomInfo            = 0x0004
	ChatNavRequestMoreRoomInfo        = 0x0005
	ChatNavRequestOccupantList        = 0x0006
	ChatNavSearchForRoom              = 0x0007
	ChatNavCreateRoom                 = 0x0008
	ChatNavNavInfo                    = 0x0009
)

func routeChatNav(sess *Session, cr *ChatRegistry, snac oscar.SnacFrame, r io.Reader, w io.Writer, sequence *uint32) error {
	switch snac.SubGroup {
	case ChatNavErr:
		panic("not implemented")
	case ChatNavRequestChatRights:
		return SendAndReceiveNextChatRights(snac, w, sequence)
	case ChatNavRequestExchangeInfo:
		panic("not implemented")
	case ChatNavRequestRoomInfo:
		return SendAndReceiveRequestRoomInfo(cr, snac, r, w, sequence)
	case ChatNavRequestMoreRoomInfo:
		panic("not implemented")
	case ChatNavRequestOccupantList:
		panic("not implemented")
	case ChatNavSearchForRoom:
		panic("not implemented")
	case ChatNavCreateRoom:
		return SendAndReceiveCreateRoom(sess, cr, snac, r, w, sequence)
	case ChatNavNavInfo:
		panic("not implemented")
	}
	return nil
}

type ChatCookie struct {
	Cookie []byte
	SessID string
}

func (s *ChatCookie) Read(r io.Reader) error {
	var l uint16
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return err
	}
	s.Cookie = make([]byte, l)
	if _, err := r.Read(s.Cookie); err != nil {
		return err
	}
	if err := binary.Read(r, binary.BigEndian, &l); err != nil {
		return err
	}
	buf := make([]byte, l)
	if _, err := r.Read(buf); err != nil {
		return err
	}
	s.SessID = string(buf)
	return nil
}

func (s ChatCookie) Write(w io.Writer) error {
	if err := binary.Write(w, binary.BigEndian, uint16(len(s.Cookie))); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, s.Cookie); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, uint16(len(s.SessID))); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, []byte(s.SessID)); err != nil {
		return err
	}
	return nil
}

func SendAndReceiveNextChatRights(snac oscar.SnacFrame, w io.Writer, sequence *uint32) error {
	fmt.Printf("sendAndReceiveNextChatRights read SNAC frame: %+v\n", snac)

	snacFrameOut := oscar.SnacFrame{
		FoodGroup: CHAT_NAV,
		SubGroup:  ChatNavNavInfo,
	}

	snacPayloadOut := oscar.SNAC_0x0D_0x09_ChatNavNavInfo{
		TLVRestBlock: oscar.TLVRestBlock{
			TLVList: oscar.TLVList{
				{
					TType: 0x02,
					Val:   uint8(10),
				},
				{
					TType: 0x03,
					Val: oscar.SNAC_0x0D_0x09_TLVExchangeInfo{
						Identifier: 4,
						TLVBlock: oscar.TLVBlock{
							TLVList: oscar.TLVList{
								{
									TType: 0x0002,
									Val:   uint16(0x0010),
								},
								{
									TType: 0x00c9,
									Val:   uint16(15),
								},
								{
									TType: 0x00d3,
									Val:   "default Exchange",
								},
								{
									TType: 0x00d5,
									Val:   uint8(2),
								},
								{
									TType: 0xd6,
									Val:   "us-ascii",
								},
								{
									TType: 0xd7,
									Val:   "en",
								},
								{
									TType: 0xd8,
									Val:   "us-ascii",
								},
								{
									TType: 0xd9,
									Val:   "en",
								},
							},
						},
					},
				},
			},
		},
	}

	return writeOutSNAC(snac, snacFrameOut, snacPayloadOut, sequence, w)
}

func SendAndReceiveCreateRoom(sess *Session, cr *ChatRegistry, snac oscar.SnacFrame, r io.Reader, w io.Writer, sequence *uint32) error {
	fmt.Printf("SendAndReceiveCreateRoom read SNAC frame: %+v\n", snac)

	snacPayloadIn := oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate{}
	if err := oscar.Unmarshal(&snacPayloadIn, r); err != nil {
		return err
	}

	name, hasName := snacPayloadIn.GetString(0x00d3)
	if !hasName {
		return errors.New("unable to find chat name")
	}

	room := ChatRoom{
		ID:             uuid.New().String(),
		SessionManager: NewSessionManager(),
		CreateTime:     time.Now(),
		Name:           name,
	}
	cr.Register(room)

	// add user to chat room
	room.SessionManager.NewSessionWithSN(sess.ID, sess.ScreenName)

	snacFrameOut := oscar.SnacFrame{
		FoodGroup: CHAT_NAV,
		SubGroup:  ChatNavNavInfo,
	}
	snacPayloadOut := oscar.SNAC_0x0D_0x09_ChatNavNavInfo{
		TLVRestBlock: oscar.TLVRestBlock{
			TLVList: oscar.TLVList{
				{
					TType: 0x04,
					Val: oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate{
						Exchange:       4,
						Cookie:         room.ID,
						InstanceNumber: 100,
						DetailLevel:    2,
						TLVBlock: oscar.TLVBlock{
							TLVList: room.TLVList(),
						},
					},
				},
			},
		},
	}

	return writeOutSNAC(snac, snacFrameOut, snacPayloadOut, sequence, w)
}

func SendAndReceiveRequestRoomInfo(cr *ChatRegistry, snac oscar.SnacFrame, r io.Reader, w io.Writer, sequence *uint32) error {
	fmt.Printf("SendAndReceiveRequestRoomInfo read SNAC frame: %+v\n", snac)

	snacPayloadIn := oscar.SNAC_0x0D_0x04_ChatNavRequestRoomInfo{}
	if err := oscar.Unmarshal(&snacPayloadIn, r); err != nil {
		return err
	}

	room, err := cr.Retrieve(string(snacPayloadIn.Cookie))
	if err != nil {
		return err
	}

	snacFrameOut := oscar.SnacFrame{
		FoodGroup: CHAT_NAV,
		SubGroup:  ChatNavNavInfo,
	}

	snacPayloadOut := oscar.SNAC_0x0D_0x09_ChatNavNavInfo{
		TLVRestBlock: oscar.TLVRestBlock{
			TLVList: oscar.TLVList{
				{
					TType: 0x04,
					Val: oscar.SNAC_0x0E_0x02_ChatRoomInfoUpdate{
						Exchange:       4,
						Cookie:         room.ID,
						InstanceNumber: 100,
						DetailLevel:    2,
						TLVBlock: oscar.TLVBlock{
							TLVList: room.TLVList(),
						},
					},
				},
			},
		},
	}

	return writeOutSNAC(snac, snacFrameOut, snacPayloadOut, sequence, w)
}