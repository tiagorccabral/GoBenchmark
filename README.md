## CPU benchmark writen in Go

---

### This software aims to test the performance of a computional system, with focus on the CPU performance

---

####Installation

```
git clone https://github.com/tiagorccabral/GoBenchmark.git
```

---

##### To run the benchmarks, you can run one of the following commands:

Inside the main folder:
```
cd GoBenchmark/
```

To run All benchmarks:
```
$ go test -bench=. -benchmem
```

To run string concatenation benchmarks:
```
$ go test -bench=String -benchmem
```

To run floating point benchmarks:
```
$ go test -bench=Fp -benchmem
```

To run slice sorting benchmarks:
```
$ go test -bench=Sort -benchmem
```