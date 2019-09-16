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

## Output image

![testdata/isucon.png](testdata/isucon.png)
