package storage

func Error(err error) error {
	if err != nil {
		return err
	}
	return nil
}
func (m *IP) Insert(conn *Storage) error {
	err := conn.Create(m)
	if err != nil {
		return err
	}
	return nil
}

func (m *IP) FindOne(conn *Storage, query interface{}) (err error) {
	ses := conn.GetDBSession()
	defer ses.Close()
	err = conn.Find(ses, query).One(m)
	if err != nil {
		return
	}
	return
}
func (m *IP) Remove(conn *Storage, selector interface{}) error {
	ses := conn.GetDBSession()
	defer ses.Close()
	err := conn.Remove(selector)
	if err != nil {
		return err
	}
	return nil
}
func (m *IP) Update(conn *Storage, selector interface{}) error {
	ses := conn.GetDBSession()
	defer ses.Close()
	err := conn.Update(selector, m)
	if err != nil {
		return err
	}
	return nil
}

//Insert if the record does not exist
func (m *IP) Save(conn *Storage, selector interface{}) error {
	ses := conn.GetDBSession()
	defer ses.Close()
	result := IP{}
	err := result.FindOne(conn, selector)
	if err != nil {
		m.Insert(conn)
	}
	return nil
}
func Find(conn *Storage, query interface{}) ([]*IP, error) {
	ses := conn.GetDBSession()
	defer ses.Close()
	var result = []*IP{}
	err := conn.Find(ses, query).All(&result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
