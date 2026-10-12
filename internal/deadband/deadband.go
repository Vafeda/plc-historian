package deadband

import (
	"errors"
	"math"

	"github.com/Vafeda/plc-historian/internal/models"
)

type Filter struct {
	defaultThreshold float64
	thresholds       map[uint32]float64
	last             map[uint32]float64
}

func New(defaultThreshold float64, thresholds map[uint32]float64) *Filter {
	return &Filter{
		defaultThreshold: defaultThreshold,
		thresholds:       thresholds,
		last:             make(map[uint32]float64),
	}
}

func (f *Filter) Pass(smpl models.Sample) (bool, error) {
	if nil != smpl.ValBool || nil != smpl.ValStr {
		return true, nil
	}

	threshold := f.defaultThreshold
	if t, ok := f.thresholds[smpl.ItemId]; ok {
		threshold = t
	}

	var value float64
	if nil != smpl.ValInt {
		value = float64(*smpl.ValInt)
	} else if nil != smpl.ValDouble {
		value = *smpl.ValDouble
	} else {
		return false, errors.New("invalid item id")
	}

	lastValue, ok := f.last[smpl.ItemId]
	if !ok || math.IsNaN(value) || math.IsNaN(lastValue) {
		f.last[smpl.ItemId] = value
		return true, nil
	}

	if math.Abs(value-lastValue) > threshold {
		f.last[smpl.ItemId] = value
		return true, nil
	}

	return false, nil
}

func (f *Filter) Reset() {
	clear(f.last)
}
