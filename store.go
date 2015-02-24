package store

import "github.com/cloudflare/gokabinet/kc"

type Store struct {
	db *kc.DB;
}

func (s Store) Open() error {
	db, err := kc.Open("/tmp/cache.kch", kc.WRITE)
	s.db = db
	return err;
}

func (s Store) Close(){
	s.db.Close()
}
