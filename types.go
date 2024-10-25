package gloop

// SignedInteger represents all signed integer types.
type SignedInteger interface {
	~int |
		~int8 |
		~int16 |
		~int32 |
		~int64
}

// UnsignedInteger represents all unsigned integer types.
type UnsignedInteger interface {
	~uint |
		~uint8 |
		~uint16 |
		~uint32 |
		~uint64 |
		~uintptr
}

// FloatingPoint represents all floating point number types.
type FloatingPoint interface {
	~float32 | ~float64
}

// ComplexNumber represents all complex number types.
type ComplexNumber interface {
	~complex64 | ~complex128
}

// Number represents all numeric types.
type Number interface {
	SignedInteger |
		UnsignedInteger |
		FloatingPoint
}

// Summable represents all types that support summation through the "+"
// operator.
type Summable interface {
	SignedInteger |
		UnsignedInteger |
		FloatingPoint |
		ComplexNumber |
		~string
}

// Productable represents all types that support multiplication through
// the "*" operator.
type Productable interface {
	SignedInteger |
		UnsignedInteger |
		FloatingPoint |
		ComplexNumber
}
