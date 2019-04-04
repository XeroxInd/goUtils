package timeslots

import (
	"fmt"
	"testing"
	"time"

	"git.libmed.fr/LibMed/goUtil/v2019/lm_object"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func TestTimeslotsSum(t *testing.T) {
	t0 := time.Now()
	t1 := t0.Add(time.Hour * 2)
	t2 := t0.Add(time.Hour * 24)
	t3 := t2.Add(time.Hour * 3)
	sum := TimeSlotSum([]*lm_object.Mission_TimeSlot{
		&lm_object.Mission_TimeSlot{
			Start: &timestamp.Timestamp{
				Seconds: t0.Unix(),
			},
			Stop: &timestamp.Timestamp{
				Seconds: t1.Unix(),
			},
		},
		&lm_object.Mission_TimeSlot{
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
