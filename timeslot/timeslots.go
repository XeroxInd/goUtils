package timeslots

import (
	"errors"
	"log"
	"sort"
	"time"

	. "git.libmed.fr/LibMed/ProtoObjects"
	"github.com/getsentry/raven-go"
)

type MultiMissions map[string]map[string]struct { //map[institutionName]map[ServiceName]struct{}
	Start time.Time
	Stop  time.Time
	NbTs  int
}

func MissionTimeSlotSum(mission *Mission) (dur time.Duration, err error) {
	tsw := mission.GetTimeSlotsWithSelected()
	if tsw == nil {
		err = errors.New("timeslots with selected is nil")
		return
	}
	dur = TimeSlotSum(tsw.GetTimeSlots())
	return
}

func TimeSlotSum(ts []*Mission_TimeSlot) (total time.Duration) {
	for _, slot := range ts {
		start := time.Unix(slot.Start.GetSeconds(), int64(slot.Start.GetNanos()))
		stop := time.Unix(slot.Stop.GetSeconds(), int64(slot.Stop.GetNanos()))
		total += stop.Sub(start)
	}
	return
}

func Sort(ts []*Mission_TimeSlot) (sorted []*Mission_TimeSlot) {
	sort.Slice(ts, func(i, j int) bool {
		si := time.Unix(ts[i].Start.Seconds, int64(ts[i].Start.Nanos))
		sj := time.Unix(ts[j].Start.Seconds, int64(ts[j].Start.Nanos))
		return si.Before(sj)
	})
	return ts
}

func GetBounds(ts []*Mission_TimeSlot) (start, stop time.Time) {
	if len(ts) != 0 {
		sorted := Sort(ts)
		first := sorted[0].Start
		last := sorted[len(sorted)-1].Stop
		start = time.Unix(first.GetSeconds(), int64(first.GetNanos()))
		stop = time.Unix(last.GetSeconds(), int64(last.GetNanos()))
	}
	return
}

func GetBoundsWithoutPast(ts []*Mission_TimeSlot) (start, stop time.Time, length int) {
	sorted := Sort(ts)
	first := sorted[0].Start
	last := sorted[len(sorted)-1].Stop
	length = len(sorted)
	current := time.Now()
	for _, v := range sorted {
		if time.Unix(v.Start.GetSeconds(), int64(v.Start.GetNanos())).After(current) {
			first = v.Start
			break
		}
		length--
	}
	start = time.Unix(first.GetSeconds(), int64(first.GetNanos()))
	stop = time.Unix(last.GetSeconds(), int64(last.GetNanos()))
	return
}

func GetMissionBounds(mission *Mission) (start, stop time.Time, length int, err error) {
	tsw := mission.GetTimeSlotsWithSelected()
	if tsw == nil {
		err = errors.New("timeslots with selected is nil")
		return
	} else if len(tsw.GetTimeSlots()) == 0 {
		err = errors.New("timeslots array is empty")
		return
	}
	length = len(tsw.TimeSlots)
	start, stop = GetBounds(tsw.GetTimeSlots())
	return
}

func GetMissionBoundsWithoutPastTS(mission *Mission) (start, stop time.Time, length int, err error) {
	tsw := mission.GetTimeSlotsWithSelected()

	if tsw == nil {
		err = errors.New("timeslots with selected is nil")
		return
	} else if len(tsw.GetTimeSlots()) == 0 {
		err = errors.New("timeslots array is empty")
		return
	}
	start, stop, length = GetBoundsWithoutPast(tsw.GetTimeSlots())
	return
}

func GroupMissions(m *Missions) (all MultiMissions) {
	all = make(MultiMissions)
	for _, mi := range m.Missions {
		start, stop, length, err := GetMissionBoundsWithoutPastTS(mi)
		if err != nil {
			raven.CaptureError(err, map[string]string{"category": "new_missions_mail"})
			log.Printf(err.Error())
		}

		if length > 0 {
			if _, ok := all[mi.Institution.Name]; !ok { //Empty value for this institution
				all[mi.Institution.Name] = make(map[string]struct {
					Start time.Time
					Stop  time.Time
					NbTs  int
				})
			}

			if s, ok := all[mi.Institution.Name][mi.Service]; !ok { //Empty value for this service in this institution
				all[mi.Institution.Name][mi.Service] = struct {
					Start time.Time
					Stop  time.Time
					NbTs  int
				}{Start: start, Stop: stop, NbTs: length}
			} else {
				if s.Start.After(start) {
					s.Start = start
				}
				if s.Stop.Before(stop) {
					s.Stop = stop
				}
				s.NbTs = s.NbTs + length
				all[mi.Institution.Name][mi.Service] = s
			}
		}
	}
	return
}

func GetMissionsBounds(missions *Missions) (start, stop time.Time) {
	for _, m := range missions.Missions {
		for _, t := range m.GetTimeSlotsWithSelected().TimeSlots {
			tStart := time.Unix(t.Start.Seconds, int64(t.Start.Nanos))
			if start.IsZero() || tStart.Before(start) {
				start = tStart
			}
			tStop := time.Unix(t.Stop.Seconds, int64(t.Stop.Nanos))
			if stop.IsZero() || tStop.After(stop) {
				stop = tStop
			}
		}
	}
	return
}

func GetMultiMissionsBounds(missions MultiMissions) (start, stop time.Time, length int) {
	for _, inst := range missions {
		for _, service := range inst {
			length = length + service.NbTs
			if start.IsZero() || start.After(service.Start) {
				start = service.Start
			}
			if stop.IsZero() || stop.Before(service.Stop) {
				stop = service.Stop
			}
		}
	}
	return
}
