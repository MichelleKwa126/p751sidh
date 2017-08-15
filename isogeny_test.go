package cln16sidh

import (
	"testing"
)

// Test the first four-isogeny from the base curve E_0(F_{p^2})
func TestFirstFourIsogenyVersusSage(t *testing.T) {
	var xR, isogenized_xR, sageIsogenized_xR ProjectivePoint

	// sage: p = 2^372 * 3^239 - 1; Fp = GF(p)
	//   ***   Warning: increasing stack size to 2000000.
	//   ***   Warning: increasing stack size to 4000000.
	// sage: R.<x> = Fp[]
	// sage: Fp2 = Fp.extension(x^2 + 1, 'i')
	// sage: i = Fp2.gen()
	// sage: E0Fp = EllipticCurve(Fp, [0,0,0,1,0])
	// sage: E0Fp2 = EllipticCurve(Fp2, [0,0,0,1,0])
	// sage: x_PA = 11
	// sage: y_PA = -Fp(11^3 + 11).sqrt()
	// sage: x_PB = 6
	// sage: y_PB = -Fp(6^3 + 6).sqrt()
	// sage: P_A = 3^239 * E0Fp((x_PA,y_PA))
	// sage: P_B = 2^372 * E0Fp((x_PB,y_PB))
	// sage: def tau(P):
	// ....:     return E0Fp2( (-P.xy()[0], i*P.xy()[1]))
	// ....:
	// sage: m_B = 3*randint(0,3^238)
	// sage: m_A = 2*randint(0,2^371)
	// sage: R_A = E0Fp2(P_A) + m_A*tau(P_A)
	// sage: def y_recover(x, a):
	// ....:     return (x**3 + a*x**2 + x).sqrt()
	// ....:
	// sage: first_4_torsion_point = E0Fp2(1, y_recover(Fp2(1),0))
	// sage: sage_first_4_isogeny = E0Fp2.isogeny(first_4_torsion_point)
	// sage: a = Fp2(0)
	// sage: sage_isomorphism = sage_first_4_isogeny.codomain().isomorphism_to(EllipticCurve(Fp2, [0,(2*(a+6))/(a-2),0,1,0]))
	// sage: isogenized_R_A = sage_isomorphism(sage_first_4_isogeny(R_A))

	xR.fromAffine(&ExtensionFieldElement{a: fp751Element{0xa179cb7e2a95fce9, 0xbfd6a0f3a0a892c0, 0x8b2f0aa4250ab3f3, 0x2e7aa4dd4118732d, 0x627969e493acbc2a, 0x21a5b852c7b8cc83, 0x26084278586324f2, 0x383be1aa5aa947c0, 0xc6558ecbb5c0183e, 0xf1f192086a52b035, 0x4c58b755b865c1b, 0x67b4ceea2d2c}, b: fp751Element{0xfceb02a2797fecbf, 0x3fee9e1d21f95e99, 0xa1c4ce896024e166, 0xc09c024254517358, 0xf0255994b17b94e7, 0xa4834359b41ee894, 0x9487f7db7ebefbe, 0x3bbeeb34a0bf1f24, 0xfa7e5533514c6a05, 0x92b0328146450a9a, 0xfde71ca3fada4c06, 0x3610f995c2bd}})

	sageIsogenized_xR.fromAffine(&ExtensionFieldElement{a: fp751Element{0xff99e76f78da1e05, 0xdaa36bd2bb8d97c4, 0xb4328cee0a409daf, 0xc28b099980c5da3f, 0xf2d7cd15cfebb852, 0x1935103dded6cdef, 0xade81528de1429c3, 0x6775b0fa90a64319, 0x25f89817ee52485d, 0x706e2d00848e697, 0xc4958ec4216d65c0, 0xc519681417f}, b: fp751Element{0x742fe7dde60e1fb9, 0x801a3c78466a456b, 0xa9f945b786f48c35, 0x20ce89e1b144348f, 0xf633970b7776217e, 0x4c6077a9b38976e5, 0x34a513fc766c7825, 0xacccba359b9cd65, 0xd0ca8383f0fd0125, 0x77350437196287a, 0x9fe1ad7706d4ea21, 0x4d26129ee42d}})

	var a ExtensionFieldElement
	a.Zero()

	_, phi := ComputeFirstFourIsogeny(&a)

	isogenized_xR = phi.Eval(&xR)

	if !sageIsogenized_xR.VartimeEq(&isogenized_xR) {
		t.Error("\nExpected\n", sageIsogenized_xR.toAffine(), "\nfound\n", isogenized_xR.toAffine())
	}
}

func TestThreeIsogenyVersusSage(t *testing.T) {
	var xR, xP3, isogenized_xR, sageIsogenized_xR ProjectivePoint

	// sage: %colors Linux
	// sage: p = 2^372 * 3^239 - 1; Fp = GF(p)
	//   ***   Warning: increasing stack size to 2000000.
	//   ***   Warning: increasing stack size to 4000000.
	// sage: R.<x> = Fp[]
	// sage: Fp2 = Fp.extension(x^2 + 1, 'i')
	// sage: i = Fp2.gen()
	// sage: E0Fp = EllipticCurve(Fp, [0,0,0,1,0])
	// sage: E0Fp2 = EllipticCurve(Fp2, [0,0,0,1,0])
	// sage: x_PA = 11
	// sage: y_PA = -Fp(11^3 + 11).sqrt()
	// sage: x_PB = 6
	// sage: y_PB = -Fp(6^3 + 6).sqrt()
	// sage: P_A = 3^239 * E0Fp((x_PA,y_PA))
	// sage: P_B = 2^372 * E0Fp((x_PB,y_PB))
	// sage: def tau(P):
	// ....:     return E0Fp2( (-P.xy()[0], i*P.xy()[1]))
	// ....:
	// sage: m_B = 3*randint(0,3^238)
	// sage: R_B = E0Fp2(P_B) + m_B*tau(P_B)
	// sage: P_3 = (3^238)*R_B
	// sage: def three_isog(P_3, P):
	// ....:     X3, Z3 = P_3.xy()[0], 1
	// ....:     XP, ZP = P.xy()[0], 1
	// ....:     x = (XP*(X3*XP - Z3*ZP)^2)/(ZP*(Z3*XP - X3*ZP)^2)
	// ....:     A3, C3 = (Z3^4 + 9*X3^2*(2*Z3^2 - 3*X3^2)), 4*X3*Z3^3
	// ....:     cod = EllipticCurve(Fp2, [0,A3/C3,0,1,0])
	// ....:     return cod.lift_x(x)
	// ....:
	// sage: isogenized_R_B = three_isog(P_3, R_B)

	xR.fromAffine(&ExtensionFieldElement{a: fp751Element{0xbd0737ed5cc9a3d7, 0x45ae6d476517c101, 0x6f228e9e7364fdb2, 0xbba4871225b3dbd, 0x6299ccd2e5da1a07, 0x38488fe4af5f2d0e, 0xec23cae5a86e980c, 0x26c804ba3f1edffa, 0xfbbed81932df60e5, 0x7e00e9d182ae9187, 0xc7654abb66d05f4b, 0x262d0567237b}, b: fp751Element{0x3a3b5b6ad0b2ac33, 0x246602b5179127d3, 0x502ae0e9ad65077d, 0x10a3a37237e1bf70, 0x4a1ab9294dd05610, 0xb0f3adac30fe1fa6, 0x341995267faf70cb, 0xa14dd94d39cf4ec1, 0xce4b7527d1bf5568, 0xe0410423ed45c7e4, 0x38011809b6425686, 0x28f52472ebed}})

	xP3.fromAffine(&ExtensionFieldElement{a: fp751Element{0x7bb7a4a07b0788dc, 0xdc36a3f6607b21b0, 0x4750e18ee74cf2f0, 0x464e319d0b7ab806, 0xc25aa44c04f758ff, 0x392e8521a46e0a68, 0xfc4e76b63eff37df, 0x1f3566d892e67dd8, 0xf8d2eb0f73295e65, 0x457b13ebc470bccb, 0xfda1cc9efef5be33, 0x5dbf3d92cc02}, b: fp751Element{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0}})

	sageIsogenized_xR.fromAffine(&ExtensionFieldElement{a: fp751Element{0x286db7d75913c5b1, 0xcb2049ad50189220, 0xccee90ef765fa9f4, 0x65e52ce2730e7d88, 0xa6b6b553bd0d06e7, 0xb561ecec14591590, 0x17b7a66d8c64d959, 0x77778cecbe1461e, 0x9405c9c0c41a57ce, 0x8f6b4847e8ca7d3d, 0xf625eb987b366937, 0x421b3590e345}, b: fp751Element{0x566b893803e7d8d6, 0xe8c71a04d527e696, 0x5a1d8f87bf5eb51, 0x42ae08ae098724f, 0x4ee3d7c7af40ca2e, 0xd9f9ab9067bb10a7, 0xecd53d69edd6328c, 0xa581e9202dea107d, 0x8bcdfb6c8ecf9257, 0xe7cbbc2e5cbcf2af, 0x5f031a8701f0e53e, 0x18312d93e3cb}})

	_, phi := ComputeThreeIsogeny(&xP3)

	isogenized_xR = phi.Eval(&xR)

	if !sageIsogenized_xR.VartimeEq(&isogenized_xR) {
		t.Error("\nExpected\n", sageIsogenized_xR.toAffine(), "\nfound\n", isogenized_xR.toAffine())
	}
}
