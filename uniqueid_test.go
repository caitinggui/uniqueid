package uniqueid

import (
	"testing"
)

// 测试生成id的性能
func BenchmarkNextId(b *testing.B) {
	st := Settings{
		WorkerId:  1,
		ReserveId: 1,
	}
	sf := NewUniqueId(st)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sf.NextId()

	}

}

func TestNextIdOnece(t *testing.T) {
	t.Log("测试id参数是否正确")
	st := Settings{
		WorkerId:  1,
		ReserveId: 1,
	}
	sf := NewUniqueId(st)
	uid, _ := sf.NextId()
	result := Prase(uid)
	if result["workerId"] != uint64(st.WorkerId) {
		t.Fatal("workerId error, wanted: ", st.WorkerId, " received: ", result["workerId"])

	}
	if result["reserveId"] != uint64(st.ReserveId) {
		t.Fatal("reserveId error, wanted: ", st.ReserveId, " received: ", result["reserveId"])

	}

}

func TestNextIdIfUnique(t *testing.T) {
	t.Log("测试id是否重复")
	st := Settings{
		WorkerId:  1,
		ReserveId: 1,
	}
	sf := NewUniqueId(st)
	result := map[uint64]bool{}
	for i := 0; i < 5000000; i++ {
		uid, _ := sf.NextId()
		if ok, _ := result[uid]; ok {
			t.Fatal("uid有重复值：", uid, Prase(uid))

		}
		result[uid] = true

	}

}
