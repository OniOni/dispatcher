package store

import (
	"strings"
	"errors"
	"github.com/cloudflare/gokabinet/kc"
)

type Store struct {
	db *kc.DB;
	keys []string;
}

func (s *Store) Open() (error, *Store) {
	db, err := kc.Open("/tmp/cache.kch", kc.WRITE)
	s.db = db
	return err, s;
}

func (s *Store) Close(){
	s.db.Close()
}

func (s *Store) GetKeys() []string {
	return s.keys
}

func (s *Store) HasKey(key string) (bool) {
	for _, str := range s.keys {
		if str == key {
			return true;
		}
	}

	return false;
}

func (s *Store) GetSubscribers(key string) ([]string, error){
	str, error := s.db.Get(key);
	if error != nil {
		return nil, error
	}

	return strings.Split(str, ";")[1:], nil
}

func (s *Store) AddSubsriber(key, value string) error{
	subscribed, _ := s.IsSubscribed(key, value)

	if subscribed {
		return errors.New("Already subscribed")
	}

	s.keys = append(s.keys, key)
	return s.db.Append(key, ";" + value)
}

func (s *Store) IsSubscribed(key, value string) (bool, error){
	str, error := s.db.Get(key)
	if  error != nil {
		return false, error
	}

	return strings.Contains(str, value), nil
}

func NewStore() (*Store, error) {
	s_ptr := new(Store)
	err, s := s_ptr.Open()

	s.keys = make([]string, 0)

	return s, err
}
