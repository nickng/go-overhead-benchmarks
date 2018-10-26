# go-overhead-benchmarks

Various overhead comparison benchmarks, for guiding
design decisions when using Go as implementation language.

## Interface overhead

Tests the overhead of using interface instead of a native type as a map key.
The interface would be used to abstract both a primitive type and a struct.

    BenchmarkIfaceOverheadNativePrimitive
    BenchmarkIfaceOverheadNativeStruct
    BenchmarkIfaceOverheadIfacePrimitive
    BenchmarkIfaceOverheadIfaceStruct

