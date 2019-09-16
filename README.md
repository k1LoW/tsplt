# tsplt

plot time series data

## Usage

``` console
$ cat ts.tsv | tsplt -o ts.png
```

``` console
$ tsplt -i ts.tsv -o ts.png
```

## Input time series data

Default delimiter : `\t`

| timestamp | d1 | d2 | ... |
| --- | --- | --- | --- | 
| `[timestamp format]` | number | number | ... |
| `[timestamp format]` | number | number | ... |
| `[timestamp format]` | number | number | ... |
| ... | ... | ... | ... |

### Timestamp format

tsplt parse timestamp using [araddon/dateparse](https://github.com/araddon/dateparse).

## Output plot image

Output plot image of [sample time series tsv](testdata/isucon.tsv).

![testdata/isucon.png](testdata/isucon.png)

