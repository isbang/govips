package vips

import "math"

type Scalar struct {
	Value    float64
	Relative bool
}

func ValueOf(value float64) Scalar {
	return Scalar{value, false}
}

func (s *Scalar) IsZero() bool {
	return s.Value == 0 && !s.Relative
}

func (s *Scalar) SetInt(value int) {
	s.Set(float64(value))
}

func (s *Scalar) Set(value float64) {
	s.Value = value
	s.Relative = false
}

func (s *Scalar) SetScale(f float64) {
	s.Value = f
	s.Relative = true
}

func (s *Scalar) Get(base int) float64 {
	if s.Relative {
		return s.Value * float64(base)
	}
	return s.Value
}

func (s *Scalar) GetRounded(base int) int {
	return roundFloat(s.Get(base))
}

func roundFloat(f float64) int {
	if f < 0 {
		return int(math.Ceil(f - 0.5))
	}
	return int(math.Floor(f + 0.5))
}
