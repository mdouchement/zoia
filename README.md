# Zoia

[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/mdouchement/zoia)

Patch parser for the Empress Effects ZOIA. It takes binary patch and parses its content as structured Golang struct.

Lots of module are still not implemented, any help is appreciated to create module parsers.

## Example

This graph has been generated with the script in `_example` and the patch `005_zoia_Clown_Shoes.bin`:

![005_zoia_Clown_Shoes.bin](https://user-images.githubusercontent.com/6150317/86532522-01705900-becb-11ea-8faf-5eeacd5a211e.png)

## Usage

```go
package main

import (
	"fmt"
	"io/ioutil"

	"github.com/mdouchement/zoia"
	"gopkg.in/yaml.v2"
)

func main() {
	data, err := ioutil.ReadFile("011_zoia_Crunch_Time.bin")
	check(err)

	patch := zoia.Parse(data)

	payload, err := yaml.Marshal(patch)
	check(err)
	fmt.Println(string(payload))

	// fmt.Println(patch[:patch.Header.PatchLength])
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}
```

## Module position

```
+----+----+----+----+----+----+----+----+
| 0  | 1  | 2  | 3  | 4  | 5  | 6  | 7  |
| 8  | 9  | 10 | 11 | 12 | 13 | 14 | 15 |
| 16 | 17 | 18 | 19 | 20 | 21 | 22 | 23 |
| 24 | 25 | 26 | 27 | 28 | 29 | 30 | 31 |
| 32 | 33 | 34 | 35 | 36 | 37 | 38 | 39 |
+----+----+----+----+----+----+----+----+
```

## Patch structure format

### Firmwares `< 1.10`

```md
[header][module#0][module#1][#patch_cables][patch_cables][#pages][page#1][page#2][padding]
^-------------------------------- Patch length ---------------------------------^
```

### Firmwares `>= 1.10`

- Extended Colours has been added

```md
[header][module#0][module#1][#patch_cables][patch_cables][#pages][page#1][page#2][module#0_extra][module#1_extra][padding]
^--------------------------------------------------- Patch length ----------------------------------------------^
```

## License

**GPLv3**


## Contributing

All PRs are welcome.

1. Fork it
2. Create your feature branch (git checkout -b my-new-feature)
3. Commit your changes (git commit -am 'Add some feature')
5. Push to the branch (git push origin my-new-feature)
6. Create new Pull Request