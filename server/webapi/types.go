package webapi

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/mk6i/retro-aim-server/config"

	"github.com/mk6i/retro-aim-server/state"
	"github.com/mk6i/retro-aim-server/wire"
)

type BuddyService interface {
	AddBuddies(ctx context.Context, sess *state.Session, inBody wire.SNAC_0x03_0x04_BuddyAddBuddies) error
	BroadcastBuddyDeparted(ctx context.Context, sess *state.Session) error
	DelBuddies(ctx context.Context, sess *state.Session, inBody wire.SNAC_0x03_0x05_BuddyDelBuddies) error
	RightsQuery(ctx context.Context, inFrame wire.SNACFrame) wire.SNACMessage
}

type ChatService interface {
	ChannelMsgToHost(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x0E_0x05_ChatChannelMsgToHost) (*wire.SNACMessage, error)
}

type ChatNavService interface {
	CreateRoom(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x0E_0x02_ChatRoomInfoUpdate) (wire.SNACMessage, error)
	ExchangeInfo(ctx context.Context, inFrame wire.SNACFrame, inBody wire.SNAC_0x0D_0x03_ChatNavRequestExchangeInfo) (wire.SNACMessage, error)
	RequestChatRights(ctx context.Context, inFrame wire.SNACFrame) wire.SNACMessage
	RequestRoomInfo(ctx context.Context, inFrame wire.SNACFrame, inBody wire.SNAC_0x0D_0x04_ChatNavRequestRoomInfo) (wire.SNACMessage, error)
}

type ICBMService interface {
	ChannelMsgToHost(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x04_0x06_ICBMChannelMsgToHost) (*wire.SNACMessage, error)
	ClientEvent(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x04_0x14_ICBMClientEvent) error
	EvilRequest(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x04_0x08_ICBMEvilRequest) (wire.SNACMessage, error)
	ParameterQuery(ctx context.Context, inFrame wire.SNACFrame) wire.SNACMessage
	ClientErr(ctx context.Context, sess *state.Session, frame wire.SNACFrame, body wire.SNAC_0x04_0x0B_ICBMClientErr) error
}

type OServiceService interface {
	ClientOnline(ctx context.Context, service uint16, bodyIn wire.SNAC_0x01_0x02_OServiceClientOnline, sess *state.Session) error
	IdleNotification(ctx context.Context, sess *state.Session, bodyIn wire.SNAC_0x01_0x11_OServiceIdleNotification) error
	ServiceRequest(ctx context.Context, service uint16, sess *state.Session, frame wire.SNACFrame, bodyIn wire.SNAC_0x01_0x04_OServiceServiceRequest, listener config.Listener) (wire.SNACMessage, error)
}

type AuthService interface {
	BUCPChallenge(ctx context.Context, bodyIn wire.SNAC_0x17_0x06_BUCPChallengeRequest, newUUID func() uuid.UUID) (wire.SNACMessage, error)
	BUCPLogin(ctx context.Context, bodyIn wire.SNAC_0x17_0x02_BUCPLoginRequest, newUserFn func(screenName state.DisplayScreenName) (state.User, error), advertisedHost string) (wire.SNACMessage, error)
	CrackCookie(authCookie []byte) (state.ServerCookie, error)
	FLAPLogin(ctx context.Context, frame wire.FLAPSignonFrame, newUserFn func(screenName state.DisplayScreenName) (state.User, error), here string) (wire.TLVRestBlock, error)
	RegisterBOSSession(ctx context.Context, authCookie state.ServerCookie) (*state.Session, error)
	RegisterChatSession(ctx context.Context, authCookie state.ServerCookie) (*state.Session, error)
	RetrieveBOSSession(ctx context.Context, authCookie state.ServerCookie) (*state.Session, error)
	Signout(ctx context.Context, sess *state.Session)
	SignoutChat(ctx context.Context, sess *state.Session)
}

type LocateService interface {
	SetDirInfo(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x02_0x09_LocateSetDirInfo) (wire.SNACMessage, error)
	SetInfo(ctx context.Context, sess *state.Session, inBody wire.SNAC_0x02_0x04_LocateSetInfo) error
	UserInfoQuery(ctx context.Context, sess *state.Session, inFrame wire.SNACFrame, inBody wire.SNAC_0x02_0x05_LocateUserInfoQuery) (wire.SNACMessage, error)
	DirInfo(ctx context.Context, inFrame wire.SNACFrame, body wire.SNAC_0x02_0x0B_LocateGetDirInfo) (wire.SNACMessage, error)
}

type DirSearchService interface {
	InfoQuery(_ context.Context, inFrame wire.SNACFrame, inBody wire.SNAC_0x0F_0x02_InfoQuery) (wire.SNACMessage, error)
}

type PermitDenyService interface {
	AddDenyListEntries(ctx context.Context, sess *state.Session, body wire.SNAC_0x09_0x07_PermitDenyAddDenyListEntries) error
	AddPermListEntries(ctx context.Context, sess *state.Session, body wire.SNAC_0x09_0x05_PermitDenyAddPermListEntries) error
	DelDenyListEntries(ctx context.Context, sess *state.Session, body wire.SNAC_0x09_0x08_PermitDenyDelDenyListEntries) error
	DelPermListEntries(ctx context.Context, sess *state.Session, body wire.SNAC_0x09_0x06_PermitDenyDelPermListEntries) error
	RightsQuery(_ context.Context, frame wire.SNACFrame) wire.SNACMessage
}

// BuddyListRegistry is the interface for keeping track of users with active
// buddy lists. Once registered, a user becomes visible to other users' buddy
// lists and vice versa.
type BuddyListRegistry interface {
	RegisterBuddyList(ctx context.Context, user state.IdentScreenName) error
	UnregisterBuddyList(ctx context.Context, user state.IdentScreenName) error
}

type TOCConfigStore interface {
	// SetTOCConfig sets the user's TOC config. The TOC config is the server-side
	// buddy list functionality for TOC. This configuration is not available to
	// OSCAR clients.
	SetTOCConfig(ctx context.Context, user state.IdentScreenName, config string) error
	User(ctx context.Context, screenName state.IdentScreenName) (*state.User, error)
}

// CookieBaker defines methods for issuing and verifying AIM authentication tokens ("cookies").
// These tokens are used for authenticating client sessions with AIM services.
type CookieBaker interface {
	// Crack verifies and decodes a previously issued authentication token.
	// Returns the original payload if the token is valid.
	Crack(data []byte) ([]byte, error)

	// Issue creates a new authentication token from the given payload.
	// The resulting token can later be verified using Crack.
	Issue(data []byte) ([]byte, error)
}

type AdminService interface {
	InfoChangeRequest(ctx context.Context, sess *state.Session, frame wire.SNACFrame, body wire.SNAC_0x07_0x04_AdminInfoChangeRequest) (wire.SNACMessage, error)
}

// SessionRetriever provides methods to retrieve OSCAR sessions.
type SessionRetriever interface {
	AllSessions() []*state.Session
	RetrieveSession(screenName state.IdentScreenName) *state.Session
}

// FeedbagRetriever provides methods to retrieve buddy list data.
type FeedbagRetriever interface {
	RetrieveFeedbag(ctx context.Context, screenName state.IdentScreenName) ([]wire.FeedbagItem, error)
	RelationshipsByUser(ctx context.Context, screenName state.IdentScreenName) ([]state.IdentScreenName, error)
}

// FeedbagManager provides methods to manage buddy lists.
type FeedbagManager interface {
	RetrieveFeedbag(ctx context.Context, screenName state.IdentScreenName) ([]wire.FeedbagItem, error)
	InsertItem(ctx context.Context, screenName state.IdentScreenName, item wire.FeedbagItem) error
	UpdateItem(ctx context.Context, screenName state.IdentScreenName, item wire.FeedbagItem) error
	DeleteItem(ctx context.Context, screenName state.IdentScreenName, item wire.FeedbagItem) error
}

// Phase 2: Additional interfaces for messaging and presence

// MessageRelayer relays messages between users
type MessageRelayer interface {
	RelayToScreenName(ctx context.Context, recipient state.IdentScreenName, msg wire.SNACMessage)
}

// OfflineMessageManager manages offline message storage and retrieval
type OfflineMessageManager interface {
	SaveMessage(ctx context.Context, msg state.OfflineMessage) error
	RetrieveMessages(ctx context.Context, recipient state.IdentScreenName) ([]state.OfflineMessage, error)
	DeleteMessages(ctx context.Context, recipient state.IdentScreenName) error
}

// BuddyBroadcaster broadcasts buddy presence updates
type BuddyBroadcaster interface {
	BroadcastBuddyArrived(ctx context.Context, sess *state.Session) error
	BroadcastBuddyDeparted(ctx context.Context, sess *state.Session) error
}

// ProfileManager manages user profiles
type ProfileManager interface {
	SetProfile(ctx context.Context, screenName state.IdentScreenName, profile string) error
	Profile(ctx context.Context, screenName state.IdentScreenName) (string, error)
}

// UserManager defines methods for user authentication.
type UserManager interface {
	// AuthenticateUser verifies username and password
	AuthenticateUser(username, password string) (*state.User, error)
	// FindUserByScreenName finds a user by their screen name
	FindUserByScreenName(screenName state.IdentScreenName) (*state.User, error)
	// InsertUser creates a new user (for DISABLE_AUTH mode)
	InsertUser(ctx context.Context, u state.User) error
}

// TokenStore manages authentication tokens.
type TokenStore interface {
	// StoreToken saves an authentication token for a user
	StoreToken(token string, screenName state.IdentScreenName, expiresAt time.Time) error
	// ValidateToken checks if a token is valid and returns the associated screen name
	ValidateToken(token string) (state.IdentScreenName, error)
	// DeleteToken removes a token
	DeleteToken(token string) error
}

// Phase 3: Preference interfaces

// PreferenceManager provides methods to manage user preferences.
type PreferenceManager interface {
	SetPreferences(ctx context.Context, screenName state.IdentScreenName, prefs map[string]interface{}) error
	GetPreferences(ctx context.Context, screenName state.IdentScreenName) (map[string]interface{}, error)
}

// PermitDenyManager provides methods to manage permit/deny lists.
type PermitDenyManager interface {
	SetPDMode(ctx context.Context, screenName state.IdentScreenName, mode wire.FeedbagPDMode) error
	GetPDMode(ctx context.Context, screenName state.IdentScreenName) (wire.FeedbagPDMode, error)
	GetPermitList(ctx context.Context, screenName state.IdentScreenName) ([]state.IdentScreenName, error)
	GetDenyList(ctx context.Context, screenName state.IdentScreenName) ([]state.IdentScreenName, error)
	AddPermitBuddy(ctx context.Context, me state.IdentScreenName, them state.IdentScreenName) error
	RemovePermitBuddy(ctx context.Context, me state.IdentScreenName, them state.IdentScreenName) error
	AddDenyBuddy(ctx context.Context, me state.IdentScreenName, them state.IdentScreenName) error
	RemoveDenyBuddy(ctx context.Context, me state.IdentScreenName, them state.IdentScreenName) error
}

// Phase 4: OSCAR Bridge interfaces

// OSCARBridgeStore manages the persistence of OSCAR bridge sessions.
type OSCARBridgeStore interface {
	SaveBridgeSession(ctx context.Context, webSessionID string, oscarCookie []byte, bosHost string, bosPort int) error
	SaveBridgeSessionWithDetails(ctx context.Context, session *state.OSCARBridgeSession) error
	GetBridgeSession(ctx context.Context, webSessionID string) (*state.OSCARBridgeSession, error)
	DeleteBridgeSession(ctx context.Context, webSessionID string) error
}

// OSCARConfig provides configuration for OSCAR services.
type OSCARConfig interface {
	GetBOSAddress() (host string, port int)
	GetSSLBOSAddress() (host string, port int)
	IsSSLAvailable() bool
	IsAuthDisabled() bool
}
