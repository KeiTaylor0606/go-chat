package db

import "testing"

func TestCreate(t *testing.T) {
	testcases := []struct {
		name      string
		teacherId uint
		studentId uint
		message   string
	}{
		{
			name:      "test_00",
			teacherId: 1,
			studentId: 2,
			message:   "こんにちは",
		},
		{
			name:      "test_01",
			teacherId: 3,
			studentId: 4,
			message:   "Golangで実装しています",
		},
		{
			name:      "test_02",
			teacherId: 5,
			studentId: 6,
			message:   "テストケース書くの面倒",
		},
		{
			name:      "test_03",
			teacherId: 7,
			studentId: 8,
			message:   "北海道に行きたい",
		},
		{
			name:      "test_04",
			teacherId: 9,
			studentId: 10,
			message:   "神戸は最近寒い",
		},
		{
			name:      "test_05",
			teacherId: 11,
			studentId: 12,
			message:   "後期フル単だった",
		},
		{
			name:      "test_06",
			teacherId: 13,
			studentId: 14,
			message:   "卒業できそう",
		},
		{
			name:      "test_07",
			teacherId: 15,
			studentId: 16,
			message:   "どうやら自分は研究者向きじゃなさそう",
		},
	}
	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			if err := Create(testcase.teacherId, testcase.studentId, testcase.message); err != nil {
				t.Errorf("err: %v\n", err)
			}
		})
	}
}
