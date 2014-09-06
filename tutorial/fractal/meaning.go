//
// Run:
//	escher -x tutorial/fractal -y acid/escher
//
main {
	?? pause!!
	explore model.Explore
	explore.When = 0
	explore.Charge = charge._
	explore.Sequence = sequence._

	charge model.ForkCharge
	charge.Circuit = ?
	charge.Peer = "b"
	charge.Valve = ?

	sequence model.ForkSequence
	sequence.When = ?
	sequence.Index = ?
	sequence.Charge = ?
}

// Circuits A, B, C and D describe a complete system.
A {
	b B
	b.X = b.Y
}

B {
	c1 C
	c2 C
	X = c1.X
	c1.Y = c2.X
	c2.Y = Y
}

C {
	d1 D
	d2 D
	X = d1.X
	d1.Y = d2.X
	d2.Y = Y
}

D {
	X = Y
}
