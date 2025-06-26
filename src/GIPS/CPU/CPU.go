package CPU

type Word struct {
	value  int32
	Signed bool
}

func (w Word) getValue() (int32, uint32) {
	if w.Signed {
		return w.value, nil
	}
}

type Register Word

type CPU struct {
	ZERO               Register
	AT                 Register
	V0, V1             Register
	A0, A1, A2, A3     Register
	T0, T1, T2, T3, T4 Register
	T5, T6, T7, T8, T9 Register
	K0, K1             Register
	GP                 Register
	SP                 Register
	FP                 Register
	RA                 Register
}

func NewCPU() *CPU {
	return &CPU{}
}

/*
	INSTRUCTIONS
	add
	addi
	addiu

*/
