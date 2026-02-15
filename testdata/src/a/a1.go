package a

import (
	"a/b"

	k8sptr "k8s.io/utils/ptr"
)

func complexScenarios() {
	_ = &b.TestType{
		A: k8sptr.To(string(b.H1)),   // want `use the "new" operator instead of k8sptr\.To; new\(string\(b\.H1\)\)`
		B: k8sptr.To(b.IntFunc()),    // want `use the "new" operator instead of k8sptr\.To; new\(b.IntFunc\(\)\)`
		C: k8sptr.To(true),           // want `use the "new" operator instead of k8sptr\.To; new\(true\)`
		D: k8sptr.To[float32](12.34), // want `use the "new" operator instead of k8sptr\.To; new\(float32\(12\.34\)\)`
		E: &b.Nested{
			F: k8sptr.To(uint64(56)), // want `use the "new" operator instead of k8sptr\.To; new\(uint64\(56\)\)`
		},
		I: k8sptr.To(b.H2), // want `use the "new" operator instead of k8sptr\.To; new\(b\.H2\)`
	}

	f1(k8sptr.To(42), k8sptr.To(false)) // want `use the "new" operator instead of k8sptr\.To; new\(42\)` `use the "new" operator instead of k8sptr\.To; new\(false\)`
}

func f1(_ *int, _ *bool) {
	// no implementation
}
