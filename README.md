[![GoDoc](https://pkg.go.dev/badge/github.com/symonk/set)](https://pkg.go.dev/github.com/symonk/set)
[![Build Status](https://github.com/symonk/set/actions/workflows/go_test.yml/badge.svg)](https://github.com/symonk/set/actions/workflows/go_test.yml)
[![codecov](https://codecov.io/gh/symonk/set/branch/main/graph/badge.svg)](https://codecov.io/gh/symonk/set)
[![Go Report Card](https://goreportcard.com/badge/github.com/symonk/set)](https://goreportcard.com/report/github.com/symonk/set)
[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://github.com/symonk/set/blob/master/LICENSE)


> [!CAUTION]
> set is currently in alpha and not fit for production level use.

# set 

`set` is a basic implementation of a `Hashset` in golang.  `set` is underpinned by a hashmap
(`map[comparable]struct{}`) to be exact and shares the same concurrency traits of a map.

-----