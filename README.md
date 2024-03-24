# Go package for Tailwind CSS heroicons set

## Usage

```sh
go get github.com/bilal-bhatti/goheroicons
```

``` go
package main

import "github.com/bilal-bhatti/goheroicons/pkg/s24/outline"

func main() {
    fmt.Println(outline.AcademicCap())
}
```

## Upstream Repository

To pull a Git submodule, use the “git submodule update” command with the “–init” and the “–recursive” options.

``` sh
git submodule update --init --recursive
```

In order to update an existing Git submodule, you need to execute the “git submodule update” with the “–remote” and the “–merge” option.

``` sh
git submodule update --remote --merge
```

To fetch new commits done in the submodule repository, head into your submodule folder and run the “git fetch” command first (you will get the new submodule commits)

``` sh
cd repository/submodule 
git fetch
```
