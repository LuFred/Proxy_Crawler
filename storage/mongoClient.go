package storage

import (
	"log"

	"gopkg.in/mgo.v2"
)

type Storage struct {
	conn       string
	database   string
	collection string
	session    *mgo.Session
}

func NewStorage(conn, db, coll string) *Storage {

	session, err := mgo.Dial(conn)
	if err != nil {
		log.Println(err)
	}
	return &Storage{conn: conn, database: db, collection: coll, session: session}
}

// GetDBSession returns a new connection from the pool
func (s *Storage) GetDBSession() *mgo.Session {
	return s.session.Copy()
}

// Create insert new item
func (s *Storage) Create(item interface{}) error {
	ses := s.GetDBSession()
	defer ses.Close()
	err := ses.DB(s.database).C(s.collection).Insert(item)
	if err != nil {
		return err
	}
	return nil
}

// update  item
func (s *Storage) Update(selector interface{}, update interface{}) error {
	ses := s.GetDBSession()
	defer ses.Close()
	err := ses.DB(s.database).C(s.collection).Update(selector, update)
	if err != nil {
		return err
	}
	return nil
}

// remove  item
func (s *Storage) Remove(selector interface{}) error {
	ses := s.GetDBSession()
	defer ses.Close()
	err := ses.DB(s.database).C(s.collection).Remove(selector)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) Find(ses *mgo.Session, q interface{}) (query *mgo.Query) {
	query = ses.DB(s.database).C(s.collection).Find(q)
	return
}
