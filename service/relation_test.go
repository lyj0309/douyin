package service

import (
	"testing"
)

var testData1 = []int{1, 1, 2, 2, 3, 3, 4, 4}
var testData2 = []int{2, 3, 1, 3, 1, 2, 1, 2}

func TestRelation(t *testing.T) {
	//addData()
	RelationAction("1", "2", "2")

	//fmt.Println(FollowerList("1"))
	//fmt.Println(IsFollow("1", "2"))
	//fmt.Println(FollowerList("2"))
	//fmt.Println(FollowerList("3"))
	//
	//fmt.Println(FollowList("1"))
	//fmt.Println(FollowList("2"))
	//fmt.Println(FollowList("3"))

}
func addData() {
	for i := 0; i < len(testData1); i++ {
		RelationAction(Itoa(testData1[i]), Itoa(testData2[i]), "1")
	}
}
func delData() {
	for i := 0; i < len(testData1); i++ {
		RelationAction(Itoa(testData1[i]), Itoa(testData2[i]), "2")
	}
}
