package oscar

import (
	"github.com/google/uuid"
	"sync"
)

type Session struct {
	ID         string
	ScreenName string
	MsgChan    chan *XMessage
	Mutex      sync.RWMutex
}

type SessionManager struct {
	store    map[string]*Session
	mapMutex sync.RWMutex
}

func NewSessionManager() *SessionManager {
	return &SessionManager{
		store: make(map[string]*Session),
	}
}

func (s *SessionManager) Retrieve(ID string) (*Session, bool) {
	s.mapMutex.RLock()
	defer s.mapMutex.RUnlock()
	sess, found := s.store[ID]
	return sess, found
}

func (s *SessionManager) RetrieveByScreenName(screenName string) *Session {
	s.mapMutex.RLock()
	defer s.mapMutex.RUnlock()
	for _, sess := range s.store {
		if screenName == sess.ScreenName {
			return sess
		}
	}
	return nil
}

func (s *SessionManager) RetrieveByScreenNames(screenNames []string) []*Session {
	s.mapMutex.RLock()
	defer s.mapMutex.RUnlock()
	var ret []*Session
	for _, sn := range screenNames {
		for _, sess := range s.store {
			if sn == sess.ScreenName {
				ret = append(ret, sess)
			}
		}
	}
	return ret
}

func (s *SessionManager) NewSession() (*Session, error) {
	id, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	sess := &Session{
		ID:      id.String(),
		MsgChan: make(chan *XMessage),
	}
	s.store[sess.ID] = sess
	return sess, nil
}

func (s *SessionManager) Remove(sess *Session) {
	s.mapMutex.Lock()
	defer s.mapMutex.Unlock()
	delete(s.store, sess.ID)
}