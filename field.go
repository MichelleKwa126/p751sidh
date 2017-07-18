package cln16sidh

// Represents an element of the extension field F_{p^2}.
type FieldElement struct {
	A Fp751Element
	B Fp751Element
}

func (dest *FieldElement) Mul(lhs, rhs *FieldElement) {
	// Let (a,b,c,d) = (lhs.A,lhs.B,rhs.A,rhs.B).

	a := &lhs.A
	b := &lhs.B
	c := &rhs.A
	d := &rhs.B

	// We want to compute
	//
	// (a + bi)*(c + di) = (a*c - b*d) + (a*d + b*c)i
	//
	// Use Karatsuba's trick: note that
	//
	// (b - a)*(c - d) = (b*c + a*d) - a*c - b*d
	//
	// so (a*d + b*c) = (b-a)*(c-d) + a*c + b*d.

	var ac, bd Fp751X2
	Fp751Mul(&ac, a, c)				// = a*c*R*R
	Fp751Mul(&bd, b, d)				// = b*d*R*R

	var b_minus_a, c_minus_d Fp751Element
	Fp751SubReduced(&b_minus_a, b, a)		// = (b-a)*R
	Fp751SubReduced(&c_minus_d, c, d)		// = (c-d)*R

	var ad_plus_bc Fp751X2
	Fp751Mul(&ad_plus_bc, &b_minus_a, &c_minus_d)	// = (b-a)*(c-d)*R*R
	Fp751X2AddLazy(&ad_plus_bc, &ad_plus_bc, &ac)	// = ((b-a)*(c-d) - a*c)*R*R
	Fp751X2AddLazy(&ad_plus_bc, &ad_plus_bc, &bd)	// = ((b-a)*(c-d) - a*c - b*d)*R*R

	Fp751MontgomeryReduce(&dest.B, &ad_plus_bc)	// = (a*d + b*c)*R mod p

	Fp751X2AddLazy(&ac, &ac, &bd)			// = (a*c + b*d)*R*R
	Fp751MontgomeryReduce(&dest.A, &ac)		// = (a*c + b*d)*R mod p
}

func (dest *FieldElement) Add(lhs, rhs *FieldElement) {
	Fp751AddReduced(&dest.A, &lhs.A, &rhs.A)
	Fp751AddReduced(&dest.B, &lhs.B, &rhs.B)
}

func (dest *FieldElement) Sub(lhs, rhs *FieldElement) {
	Fp751SubReduced(&dest.A, &lhs.A, &rhs.A)
	Fp751SubReduced(&dest.B, &lhs.B, &rhs.B)
}

func (lhs *FieldElement) VartimeEq(rhs *FieldElement) bool {
	return lhs.A.VartimeEq(rhs.A) && lhs.B.VartimeEq(rhs.B)
}

const Fp751NumWords = 12

// Represents an element of the base field F_p, in Montgomery form.
type Fp751Element [Fp751NumWords]uint64

// Represents an intermediate product of two elements of the base field F_p.
type Fp751X2 [2 * Fp751NumWords]uint64

// Compute z = x + y (mod p).
//go:noescape
func Fp751AddReduced(z, x, y *Fp751Element)

// Compute z = x - y (mod p).
//go:noescape
func Fp751SubReduced(z, x, y *Fp751Element)

// Compute z = x + y, without reducing mod p.
//go:noescape
func Fp751AddLazy(z, x, y *Fp751Element)

// Compute z = x + y, without reducing mod p.
//go:noescape
func Fp751X2AddLazy(z, x, y *Fp751X2)

// Compute z = x * y.
//go:noescape
func Fp751Mul(z *Fp751X2, x, y *Fp751Element)

// Perform Montgomery reduction: set z = x R^{-1} (mod p).
// Destroys the input value.
//go:noescape
func Fp751MontgomeryReduce(z *Fp751Element, x *Fp751X2)

// Reduce a field element in [0, 2*p) to one in [0,p).
//go:noescape
func Fp751StrongReduce(x *Fp751Element)

func (x Fp751Element) VartimeEq(y Fp751Element) bool {
	Fp751StrongReduce(&x)
	Fp751StrongReduce(&y)
	eq := true
	for i := 0; i < Fp751NumWords; i++ {
		eq = (x[i] == y[i]) && eq
	}

	return eq
}
