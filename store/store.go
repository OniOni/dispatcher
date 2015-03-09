package store

import (
	"strings"
	"github.com/cloudflare/gokabinet/kc"
)

type Store struct {
	db *kc.DB;
}

func (s Store) Open() (error, Store) {
	db, err := kc.Open("/tmp/cache.kch", kc.WRITE)
	s.db = db
	return err, s;
}

func (s Store) Close(){
	s.db.Close()
}

func (s Store) GetSubscribers(key string) ([]string, error){
	str, error := s.db.Get(key);
	if error != nil {
		return nil, error
	}

	return strings.Split(str, ";")[1:], nil
}

func (s Store) AddSubsriber(key, value string) error{
	return s.db.Append(key, ";" + value);
}

func (s Store) IsSubscribed(key, value string) (bool, error){
	str, error := s.db.Get(key)
	if  error != nil {
		return false, error
	}

	return strings.Contains(str, value), nil
}

func NewStore() (Store, error) {
	s_ptr := new(Store)
	err, s := s_ptr.Open()

	return s, err
}
