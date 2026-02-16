package pointer

import (
	"a/b"
	"time"

	"k8s.io/utils/pointer"
)

const (
	int42    = 42
	uint42   = uint(42)
	float42  = 42.0
	fortyTwo = "42"

	enabled = true
)

var (
	vuint42 = uint(42)
)

var (
	_ *int = pointer.Int(42)          // want `replace pointer\.Int\(42\) with the "new\(\)" built-in function: new\(42\)`
	_ *int = pointer.Int(b.IntFunc()) // want `replace pointer\.Int\(b\.IntFunc\(\)\) with the "new\(\)" built-in function: new\(b\.IntFunc\(\)\)`
	_ *int = pointer.Int(int42)       // want `replace pointer\.Int\(int42\) with the "new\(\)" built-in function: new\(int42\)`
	_ *int = pointer.Int(42.0)        // want `replace pointer\.Int\(42\.0\) with the "new\(\)" built-in function: new\(int\(42.0\)\)`
	_ *int = pointer.Int(float42)     // want `replace pointer\.Int\(float42\) with the "new\(\)" built-in function: new\(int\(float42\)\)`
	_ *int = pointer.Int(b.FG)        // want `replace pointer\.Int\(b\.FG\) with the "new\(\)" built-in function: new\(int\(b\.FG\)\)`

	_ *uint = pointer.Uint(42)      // want `replace pointer\.Uint\(42\) with the "new\(\)" built-in function: new\(uint\(42\)\)`
	_ *uint = pointer.Uint(b.G)     // want `replace pointer\.Uint\(b\.G\) with the "new\(\)" built-in function: new\(uint\(b\.G\)\)`
	_ *uint = pointer.Uint(b.UG)    // want `replace pointer\.Uint\(b\.UG\) with the "new\(\)" built-in function: new\(b\.UG\)`
	_ *uint = pointer.Uint(int42)   // want `replace pointer\.Uint\(int42\) with the "new\(\)" built-in function: new\(uint\(int42\)\)`
	_ *uint = pointer.Uint(uint42)  // want `replace pointer\.Uint\(uint42\) with the "new\(\)" built-in function: new\(uint42\)`
	_ *uint = pointer.Uint(vuint42) // want `replace pointer\.Uint\(vuint42\) with the "new\(\)" built-in function: new\(vuint42\)`

	_ *string = pointer.String("forty two")  // want `replace pointer\.String\("forty two"\) with the "new\(\)" built-in function: new\("forty two"\)`
	_ *string = pointer.String(string(b.H1)) // want `replace pointer\.String\(string\(b\.H1\)\) with the "new\(\)" built-in function: new\(string\(b\.H1\)\)`
	_ *string = pointer.String(fortyTwo)     // want `replace pointer\.String\(fortyTwo\) with the "new\(\)" built-in function: new\(fortyTwo\)`

	_ *bool = pointer.Bool(true)    // want `replace pointer\.Bool\(true\) with the "new\(\)" built-in function: new\(true\)`
	_ *bool = pointer.Bool(false)   // want `replace pointer\.Bool\(false\) with the "new\(\)" built-in function: new\(false\)`
	_ *bool = pointer.Bool(enabled) // want `replace pointer\.Bool\(enabled\) with the "new\(\)" built-in function: new\(enabled\)`

	_ *int32 = pointer.Int32(42)    // want `replace pointer\.Int32\(42\) with the "new\(\)" built-in function: new\(int32\(42\)\)`
	_ *int32 = pointer.Int32(int42) // want `replace pointer\.Int32\(int42\) with the "new\(\)" built-in function: new\(int32\(int42\)\)`

	_ *uint32 = pointer.Uint32(42)    // want `replace pointer\.Uint32\(42\) with the "new\(\)" built-in function: new\(uint32\(42\)\)`
	_ *uint32 = pointer.Uint32(int42) // want `replace pointer\.Uint32\(int42\) with the "new\(\)" built-in function: new\(uint32\(int42\)\)`

	_ *int64 = pointer.Int64(42)    // want `replace pointer\.Int64\(42\) with the "new\(\)" built-in function: new\(int64\(42\)\)`
	_ *int64 = pointer.Int64(int42) // want `replace pointer\.Int64\(int42\) with the "new\(\)" built-in function: new\(int64\(int42\)\)`

	_ *uint64 = pointer.Uint64(42)    // want `replace pointer\.Uint64\(42\) with the "new\(\)" built-in function: new\(uint64\(42\)\)`
	_ *uint64 = pointer.Uint64(int42) // want `replace pointer\.Uint64\(int42\) with the "new\(\)" built-in function: new\(uint64\(int42\)\)`

	_ *float32 = pointer.Float32(42.42)   // want `replace pointer\.Float32\(42\.42\) with the "new\(\)" built-in function: new\(float32\(42\.42\)\)`
	_ *float32 = pointer.Float32(float42) // want `replace pointer\.Float32\(float42\) with the "new\(\)" built-in function: new\(float32\(float42\)\)`

	_ *float64 = pointer.Float64(42.42)   // want `replace pointer\.Float64\(42\.42\) with the "new\(\)" built-in function: new\(42\.42\)`
	_ *float64 = pointer.Float64(float42) // want `replace pointer\.Float64\(float42\) with the "new\(\)" built-in function: new\(float42\)`
	_ *float64 = pointer.Float64(42)      // want `replace pointer\.Float64\(42\) with the "new\(\)" built-in function: new\(float64\(42\)\)`
	_ *float64 = pointer.Float64(int42)   // want `replace pointer\.Float64\(int42\) with the "new\(\)" built-in function: new\(float64\(int42\)\)`
	_ *float64 = pointer.Float64(b.G)     // want `replace pointer\.Float64\(b\.G\) with the "new\(\)" built-in function: new\(float64\(b\.G\)\)`

	_ *time.Duration = pointer.Duration(time.Second * 42) // want `replace pointer\.Duration\(time\.Second \* 42\) with the "new\(\)" built-in function: new\(time\.Second \* 42\)`
)
