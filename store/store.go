package store

import "github.com/cloudflare/gokabinet/kc"

type Store struct {
	db *kc.DB;
}

func NewStore() (*Store, error) {
	s := new(Store)
	err := s.Open()

	return s, err
}

func (s Store) Open() error {
	db, err := kc.Open("/tmp/cache.kch", kc.WRITE)
	s.db = db
	return err;
}

func (s Store) Close(){
	s.db.Close()
}

func (s Store) GetSubscribers(key string) (string, error){
	return s.db.Get(key);
}

func (s Store) AddSubsriber(key, value string) error{
	return s.db.Append(key, value);
}
