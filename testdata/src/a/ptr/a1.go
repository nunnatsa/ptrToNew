package ptr

import (
	"a/b"

	k8sptr "k8s.io/utils/ptr"
)

func complexScenarios() {
	_ = &b.TestType{
		A: k8sptr.To(string(b.H1)),   // want `replace k8sptr\.To\(string\(b\.H1\)\) with the "new\(\)" built-in function: new\(string\(b\.H1\)\)`
		B: k8sptr.To(b.IntFunc()),    // want `replace k8sptr\.To\(b\.IntFunc\(\)\) with the "new\(\)" built-in function: new\(b\.IntFunc\(\)\)`
		C: k8sptr.To(true),           // want `replace k8sptr\.To\(true\) with the "new\(\)" built-in function: new\(true\)`
		D: k8sptr.To[float32](12.34), // want `replace k8sptr\.To\[float32\]\(12\.34\) with the "new\(\)" built-in function: new\(float32\(12\.34\)\)`
		E: &b.Nested{
			F: k8sptr.To(uint64(56)), // want `replace k8sptr\.To\(uint64\(56\)\) with the "new\(\)" built-in function: new\(uint64\(56\)\)`
		},
		I: k8sptr.To(b.H2), // want `replace k8sptr\.To\(b\.H2\) with the "new\(\)" built-in function: new\(b\.H2\)`
	}

	f1(k8sptr.To(42), k8sptr.To(false)) // want `replace k8sptr\.To\(42\) with the "new\(\)" built-in function: new\(42\)` `replace k8sptr\.To\(false\) with the "new\(\)" built-in function: new\(false\)`
}

func f1(_ *int, _ *bool) {
	// no implementation
}
