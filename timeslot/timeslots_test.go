package utils

import (
	"fmt"
	"testing"
	"time"

	"git.libmed.fr/postier/v2019/proto"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func TestTimeslotsSum(t *testing.T) {
	t0 := time.Now()
	t1 := t0.Add(time.Hour * 2)
	t2 := t0.Add(time.Hour * 24)
	t3 := t2.Add(time.Hour * 3)
	sum := TimeSlotSum([]*proto.Mission_TimeSlot{
		&proto.Mission_TimeSlot{
			Start: &timestamp.Timestamp{
				Seconds: t0.Unix(),
			},
			Stop: &timestamp.Timestamp{
				Seconds: t1.Unix(),
			},
		},
		&proto.Mission_TimeSlot{
			Start: &timestamp.Timestamp{
				Seconds: t2.Unix(),
			},
			Stop: &timestamp.Timestamp{
				Seconds: t3.Unix(),
			},
		},
	})
	fmt.Println(sum)
	if sum != time.Hour*5 {
		t.Fail()
	}
}
