package lenconv

import "fmt"

type Meter float64
type Foot float64

func (m Meter) String() string { return fmt.Sprintf("%gm", m) }
func (f Foot) String() string  { return fmt.Sprintf("%gft", f) }

func MTof(m Meter) Foot { return Foot(m * 3.2808) }
func FToM(f Foot) Meter { return Meter(f / 3.2808) }
