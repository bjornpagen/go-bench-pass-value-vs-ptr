### Verdict: Just pass by pointer in go

```
Running tool: /opt/homebrew/bin/go test -benchmem -run=^$ -bench ^(Benchmark1B|Benchmark4B|Benchmark16B|Benchmark64B|Benchmark128B|Benchmark256B|Benchmark1K|Benchmark4K|Benchmark16K|Benchmark64K|Benchmark256K|Benchmark1M)$ github.com/bjornpagen/go-bench-pass-value-vs-ptr

goos: darwin
goarch: arm64
pkg: github.com/bjornpagen/go-bench-pass-value-vs-ptr
=== RUN   Benchmark1B
Benchmark1B
=== RUN   Benchmark1B/passByPtr
Benchmark1B/passByPtr
Benchmark1B/passByPtr-8                 751308920                1.535 ns/op          0 B/op           0 allocs/op
=== RUN   Benchmark1B/passByValue
Benchmark1B/passByValue
Benchmark1B/passByValue-8               779551708                1.539 ns/op          0 B/op           0 allocs/op
=== RUN   Benchmark4B
Benchmark4B
=== RUN   Benchmark4B/passByPtr
Benchmark4B/passByPtr
Benchmark4B/passByPtr-8                 292158169                4.116 ns/op          0 B/op           0 allocs/op
=== RUN   Benchmark4B/passByValue
Benchmark4B/passByValue
Benchmark4B/passByValue-8               293396619                4.098 ns/op          0 B/op           0 allocs/op
=== RUN   Benchmark16B
Benchmark16B
=== RUN   Benchmark16B/passByPtr
Benchmark16B/passByPtr
Benchmark16B/passByPtr-8                100000000               10.65 ns/op           0 B/op           0 allocs/op
=== RUN   Benchmark16B/passByValue
Benchmark16B/passByValue
Benchmark16B/passByValue-8              100000000               10.37 ns/op           0 B/op           0 allocs/op
=== RUN   Benchmark64B
Benchmark64B
=== RUN   Benchmark64B/passByPtr
Benchmark64B/passByPtr
Benchmark64B/passByPtr-8                27765511                42.90 ns/op           0 B/op           0 allocs/op
=== RUN   Benchmark64B/passByValue
Benchmark64B/passByValue
Benchmark64B/passByValue-8              33241956                35.76 ns/op           0 B/op           0 allocs/op
=== RUN   Benchmark128B
Benchmark128B
=== RUN   Benchmark128B/passByPtr
Benchmark128B/passByPtr
Benchmark128B/passByPtr-8               15134539                79.06 ns/op           0 B/op           0 allocs/op
=== RUN   Benchmark128B/passByValue
Benchmark128B/passByValue
Benchmark128B/passByValue-8             13177701                90.17 ns/op           0 B/op           0 allocs/op
=== RUN   Benchmark256B
Benchmark256B
=== RUN   Benchmark256B/passByPtr
Benchmark256B/passByPtr
Benchmark256B/passByPtr-8                8183865               146.4 ns/op            0 B/op           0 allocs/op
=== RUN   Benchmark256B/passByValue
Benchmark256B/passByValue
Benchmark256B/passByValue-8              7392670               161.6 ns/op            0 B/op           0 allocs/op
=== RUN   Benchmark1K
Benchmark1K
=== RUN   Benchmark1K/passByPtr
Benchmark1K/passByPtr
Benchmark1K/passByPtr-8                  2220846               539.9 ns/op            0 B/op           0 allocs/op
=== RUN   Benchmark1K/passByValue
Benchmark1K/passByValue
Benchmark1K/passByValue-8                1778340               674.7 ns/op            0 B/op           0 allocs/op
=== RUN   Benchmark4K
Benchmark4K
=== RUN   Benchmark4K/passByPtr
Benchmark4K/passByPtr
Benchmark4K/passByPtr-8                   562999              2110 ns/op              0 B/op           0 allocs/op
=== RUN   Benchmark4K/passByValue
Benchmark4K/passByValue
Benchmark4K/passByValue-8                 427296              2773 ns/op              0 B/op           0 allocs/op
=== RUN   Benchmark16K
Benchmark16K
=== RUN   Benchmark16K/passByPtr
Benchmark16K/passByPtr
Benchmark16K/passByPtr-8                  142136              8392 ns/op              0 B/op           0 allocs/op
=== RUN   Benchmark16K/passByValue
Benchmark16K/passByValue
Benchmark16K/passByValue-8                107124             11161 ns/op              0 B/op           0 allocs/op
=== RUN   Benchmark64K
Benchmark64K
=== RUN   Benchmark64K/passByPtr
Benchmark64K/passByPtr
Benchmark64K/passByPtr-8                   35250             33542 ns/op              0 B/op           0 allocs/op
=== RUN   Benchmark64K/passByValue
Benchmark64K/passByValue
Benchmark64K/passByValue-8                 27075             44313 ns/op              0 B/op           0 allocs/op
=== RUN   Benchmark256K
Benchmark256K
=== RUN   Benchmark256K/passByPtr
Benchmark256K/passByPtr
Benchmark256K/passByPtr-8                   7628            154485 ns/op              0 B/op           0 allocs/op
=== RUN   Benchmark256K/passByValue
Benchmark256K/passByValue
Benchmark256K/passByValue-8                 6699            177457 ns/op              0 B/op           0 allocs/op
=== RUN   Benchmark1M
Benchmark1M
=== RUN   Benchmark1M/passByPtr
Benchmark1M/passByPtr
Benchmark1M/passByPtr-8                     2053            555195 ns/op              0 B/op           0 allocs/op
=== RUN   Benchmark1M/passByValue
Benchmark1M/passByValue
Benchmark1M/passByValue-8                   1680            711717 ns/op              0 B/op           0 allocs/op
PASS
ok      github.com/bjornpagen/go-bench-pass-value-vs-ptr        32.754s


> Test run finished at 5/19/2023, 9:01:52 PM <
```