package nocmd_test

// example_test.go is a black-box test file. The driver delivers it as an
// external test package (clause nocmd_test) plus a synthesized test-main
// package; neither defines Command(), so without primaryBuild's test-variant
// skip they would each falsely report a missing entry point.
func ExampleCommand() {}
