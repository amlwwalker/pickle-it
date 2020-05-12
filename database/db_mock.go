// +build mock

package database

func (db *DB) ConfigureDB() {
	db.Notice("Fake Database Configured")
}

func (db *DB) CheckIfFileCurrentlyMonitored(src string) (bool, error) {
	return true, nil
}

func (db *DB) InitialiseFileInDatabase(src string) (bool, error) {
	return true, nil
}
