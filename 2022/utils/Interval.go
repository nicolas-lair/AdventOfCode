package utils

type Interval struct {
	Lower int
	Upper int
}

func (i *Interval) Contains(other Interval) bool {
	return i.Lower <= other.Lower && i.Upper >= other.Upper
}

func (i *Interval) Overlaps(other Interval) bool {
	return (i.Lower <= other.Lower && i.Upper >= other.Lower) || (i.Lower <= other.Upper && i.Upper >= other.Upper) || other.Contains(*i)
}
