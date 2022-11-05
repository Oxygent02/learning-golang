package database

var connection string

/*
func init() akan otomatis tereksekusi ketika package ini dipanggil
*/
func init(){
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}