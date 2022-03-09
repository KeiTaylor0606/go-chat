package db

import "testing"

func TestCreate(t *testing.T) {
	testcases := []struct {
		name    string
		user    string
		message string
	}{
		{
			name:    "test_00",
			user:    "A",
			message: "こんにちは",
		},
		{
			name:    "test_01",
			user:    "B",
			message: "Golangで実装しています",
		},
		{
			name:    "test_02",
			user:    "B",
			message: "テストケース書くの面倒",
		},
		{
			name:    "test_03",
			user:    "B",
			message: "北海道に行きたい",
		},
		{
			name:    "test_04",
			user:    "A",
			message: "神戸は最近寒い",
		},
		{
			name:    "test_05",
			user:    "B",
			message: "後期フル単だった",
		},
		{
			name:    "test_06",
			user:    "B",
			message: "卒業できそう",
		},
		{
			name:    "test_07",
			user:    "B",
			message: "どうやら自分は研究者向きじゃなさそう",
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			if err := Create(testcase.user, testcase.message); err != nil {
				t.Errorf("err: %v\n", err)
			}
		})
	}
}
