# ngrams

N-grams generation library written in go

## Install

```bash
$ go get -v github.com/f1monkey/ngrams
```

## Usage

```go
    result := ngrams.From("orange", 3)
    fmt.Println(result) // [ora, ran, ang, nge]

    result := ngrams.FromRange("word", 2, 3)
    fmt.Println(result) // [wo, or, rd, wor, ord]
```


## Benchmark

```
goos: linux
goarch: amd64
pkg: github.com/cyradin/ngrams
cpu: Intel(R) Core(TM) i9-8950HK CPU @ 2.90GHz
Benchmark_MakeRange-12        	  493305	      2446 ns/op	    1440 B/op	      31 allocs/op
Benchmark_From_6_3-12         	 2563634	       456.1 ns/op	     224 B/op	       6 allocs/op
Benchmark_FromRunes_6_3-12    	 3208182	       413.3 ns/op	     224 B/op	       6 allocs/op
PASS
coverage: 80.9% of statements
ok  	github.com/cyradin/ngrams	4.593s
```
