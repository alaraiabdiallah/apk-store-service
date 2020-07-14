package tests

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"reflect"
	"strconv"
	"testing"
)

func TestReflect(t *testing.T) {
	var number interface{} = 23
	var n interface{}
	ref := reflect.ValueOf(number)

	switch ref.Type().Name() {
	case "int":
		fmt.Println(ref.Type().Name())
		n = strconv.Itoa(number.(int))
	case "string":
		fmt.Println(ref.Type().Name())
		n, _ = strconv.Atoi(number.(string))
	}
	t.Log(n)
}

func TestBson(t *testing.T) {
	someData := bson.M{"akey": "a", "bkey":"b"}
	var someInter map[string]interface{}
	someInter = someData
	for k, v := range someInter{
		t.Logf("%v:%v",k,v)
	}
}
