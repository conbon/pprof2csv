# pprof2csv

Converts [pprof](https://github.com/google/pprof) profile format to CSV.

Inspired by Google's [csv2pprof](https://github.com/google/csv2pprof).

## Installation

```bash
go install github.com/conbon/pprof2csv@latest
```

## Usage

Pprof text output using `-text` or `-top` flags supported.

```bash
pprof2csv -input=profile.txt -output=profile.csv
```
