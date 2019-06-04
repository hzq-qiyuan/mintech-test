package service

import (
	"db"
	"fmt"
	"model"
	"testing"
)

func TestInsertOrder(t *testing.T) {
	db.OpenDB()

	id, err := CreateOrder(&model.Order{OrderId:"12343", Amount:11, FileUrl:"dfd", UserName:"zhangsan", Status:"good"})
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println("id:",id)
}


func TestUpdateOrder(t *testing.T) {
	db.OpenDB()

	err := UpdateOrder(&model.Order{OrderId:"11111", Amount:1, FileUrl:"1111334", UserName:"11114", Status:"11114"})
	if err != nil {
		t.Fatal(err)
	}
}


func TestQueryByOrderId(t *testing.T) {
	db.OpenDB()

	ret, err := QueryByOrderId("11111")
	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(fmt.Printf("%+v",ret))

}


func TestQueryByCondition(t *testing.T) {
	db.OpenDB()

	ret, err := QueryByCondition(&model.QueryCondition{Key:"", LikeStr:"", Desc:true})
	if err != nil {
		t.Fatal(err)
	}

	for i := 0; i < len(ret); i++ {
		fmt.Println(fmt.Printf("%+v",ret[i]))
	}
}