package db

//InsertOne : insert one record
func InsertOne(tableName string, obj interface{}) (err error) {
	sess := NewSQLSession()
	defer sess.Close()

	sess = sess.Table(tableName)
	_, err = sess.InsertOne(obj)
	return
}

//InsertMultiple : insert multiple item into table
func InsertMultiple(tableName string, objs []interface{}) (err error) {
	sess := NewSQLSession()
	defer sess.Close()

	sess = sess.Table(tableName)
	_, err = sess.InsertMulti(objs)
	return
}
