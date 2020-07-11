# Examples

- YAML output

```sh
$ cd _examples
$ go run . yaml /Volumes/ZOIA/from_zoia/011_zoia_Crunch_Time.bin
```

- Graphviz output

```sh
$ cd _examples
$ go run . graph /Volumes/ZOIA/from_zoia/011_zoia_Crunch_Time.bin | dot -Tsvg > /Downloads/patch.svg
```