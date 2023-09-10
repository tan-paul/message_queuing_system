package model

import (
	database "message_queuing_system/database"
	"testing"
)

var RequestDatas = []RequestData{
	{
		1,
		"test_product",
		"demo product for test",
		[]string{"https://image.similarpng.com/very-thumbnail/2020/09/Raw-beef-meat-pieces-on-transparent-background-PNG.png", "https://cdn.pixabay.com/photo/2012/06/19/10/32/owl-50267_1280.jpg"},
		102.454,
	},
	{
		2,
		"test_product_",
		"demo product for test_",
		[]string{"https://image.similarpng.com/very-thumbnail/2020/05/Cartoon-chef-holding-hamburger-transparent-background-PNG.png", "https://cdn.pixabay.com/photo/2012/06/19/10/32/owl-50267_1280.jpg"},
		102.454,
	},
}

/*
it will throw Log for all the invalid client and error if
any internal error found during insertion.

Here i've inserted 1 & 2 as my client in DB and
tested for 1-10.
*/
func TestIsAuthenticClient(t *testing.T) {
	err := database.ConnectMysqlDB()
	if err != nil {
		panic(err)
	}
	defer database.DisConnectMysqlDB(database.Database_Conn)
	m := new(Models)
	for i := 1; i <= 10; i++ {
		if ok, err := m.IsAuthenticClient(i); !ok {
			if err != nil {
				t.Errorf("internal error ::  %q", i)
			} else {
				t.Log("invalid client :: ", i)
			}
		}
	}
}

func TestInsertIntoDatabase(t *testing.T) {
	err := database.ConnectMysqlDB()
	if err != nil {
		panic(err)
	}
	defer database.DisConnectMysqlDB(database.Database_Conn)
	m := new(Models)
	for _, req := range RequestDatas {
		_, err := m.InsertIntoDatabase(req)
		if err != nil {
			t.Errorf("Error in inserting data to DB :: %q", err)
		}
	}
}
func BenchmarkInsertIntoDatabase(b *testing.B) {
	err := database.ConnectMysqlDB()
	if err != nil {
		panic(err)
	}
	defer database.DisConnectMysqlDB(database.Database_Conn)
	m := new(Models)
	for i := 0; i < b.N; i++ {
		m.InsertIntoDatabase(RequestData{
			2,
			"test_product",
			"demo product for test",
			[]string{"https://image.similarpng.com/very-thumbnail/2020/09/Raw-beef-meat-pieces-on-transparent-background-PNG.png", "https://cdn.pixabay.com/photo/2012/06/19/10/32/owl-50267_1280.jpg"},
			102.454,
		})
	}
}
