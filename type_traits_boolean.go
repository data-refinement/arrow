package arrow

type BooleanTraits struct{}

// BytesRequired returns the number of bytes required to store the requested number of elements.
func (BooleanTraits) BytesRequired(elements int) int { return ceilByte(elements) / 8 }
func (BooleanTraits) CastFromBytes(b []byte) []byte  { return b }
func (BooleanTraits) CastToBytes(b []byte) []byte    { return b }
func (t BooleanTraits) Copy(dst, src []byte)         { copy(dst, src) }