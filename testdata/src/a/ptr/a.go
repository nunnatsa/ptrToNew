package ptr

import (
	"k8s.io/utils/ptr"
)

func f() uint64 {
	return 42
}

const int42 = 42

var (
	_ *bool   = ptr.To(true)       // want `replace ptr\.To\(true\) with the "new\(\)" built-in function: new\(true\)`
	_ *string = ptr.To("a string") // want `replace ptr\.To\("a string"\) with the "new\(\)" built-in function: new\("a string"\)`
	_ *int    = ptr.To(int42)      // want `replace ptr\.To\(int42\) with the "new\(\)" built-in function: new\(int42\)`
	_ *int32  = ptr.To(int32(10))  // want `replace ptr\.To\(int32\(10\)\) with the "new\(\)" built-in function: new\(int32\(10\)\)`
	_ *int32  = ptr.To[int32](10)  // want `replace ptr\.To\[int32\]\(10\) with the "new\(\)" built-in function: new\(int32\(10\)\)`
	_ *uint64 = ptr.To(f())        // want `replace ptr\.To\(f\(\)\) with the "new\(\)" built-in function: new\(f\(\)\)`
)
