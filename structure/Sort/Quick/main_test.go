/**
* $ go test -bench=.
*
* goos: darwin
* goarch: arm64
* pkg: github.com/pure-juan/baekjoon/structure/Sort/Quick
* BenchmarkWithoutGoRoutine-8     2 6 7 10 12 22 26 32 34 38 48 50 58 65 68 76 87 94 96 97
* 2 6 7 10 12 22 26 32 34 38 48 50 58 65 68 76 87 94 96 97
* 2 6 7 10 12 22 26 32 34 38 48 50 58 65 68 76 87 94 96 97
* 2 6 7 10 12 22 26 32 34 38 48 50 58 65 68 76 87 94 96 97
* 2 6 7 10 12 22 26 32 34 38 48 50 58 65 68 76 87 94 96 97
* 1000000000               0.0000131 ns/op
* 2 6 7 10 12 22 26 32 34 38 48 50 58 65 68 76 87 94 96 97
* BenchmarkGoRoutine-8            2 6 7 10 12 22 26 32 34 38 48 50 58 65 68 76 87 94 96 97
* 2 6 7 10 12 22 26 32 34 38 48 50 58 65 68 76 87 94 96 97
* 2 6 7 10 12 22 26 32 34 38 48 50 58 65 68 76 87 94 96 97
* 2 6 7 10 12 22 26 32 34 38 48 50 58 65 68 76 87 94 96 97
* 2 6 7 10 12 22 26 32 34 38 48 50 58 65 68 76 87 94 96 97
* 1000000000               0.0000276 ns/op
* PASS
* ok      github.com/pure-juan/baekjoon/structure/Sort/Quick      0.157s
 */

package main

import "testing"

func BenchmarkWithoutGoRoutine(b *testing.B) {
	main()
}

func BenchmarkGoRoutine(b *testing.B) {
	main2()
}
