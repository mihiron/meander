package meander

import "strings"

type Cost int8

const (
	_ Cost = iota
	Cost1
	Cost2
	Cost3
	Cost4
	Cost5
)

var costStrings = map[string]Cost{
	"$":     Cost1,
	"$$":    Cost2,
	"$$$":   Cost3,
	"$$$$":  Cost4,
	"$$$$$": Cost5,
}

func (l Cost) String() string {
	for s, v := range costStrings {
		if l == v {
			return s
		}
	}
	return "不正な値です"
}

func ParceCost(s string) Cost {
	return costStrings[s]
}

type CostRange struct {
	From Cost
	To   Cost
}

func (r CostRange) String() string {
	return r.From.String() + "..." + r.To.String()
}

func ParceCostRange(s string) *CostRange {
	segs := strings.Split(s, "...")
	return &CostRange{
		From: ParceCost(segs[0]),
		To:   ParceCost(segs[1]),
	}
}
