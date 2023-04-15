ProgressBar Package
===================

This package provides several types of progress bars that can be used in command-line interfaces.

Usage
-----

Import the package with:

go

```go
import "github.com/mshafiee/progressbar"
```

Call any of the functions in the package to display a progress bar.

For example, to display a simple progress bar:

go

```go
progressbar.SimpleProgressBar(50, 100)
```

This will display a progress bar indicating 50% completion.

Available Progress Bars
-----------------------

The package currently provides the following progress bars:

*   SimpleProgressBar
*   ProgressBarWithPercentage
*   ArrowProgressBar
*   ColorArrowProgressBar
*   BlockProgressBar
*   ColorBlockProgressBar
*   RotatingProgressBar
*   AnimatedArrowProgressBar

Each function takes two arguments: `completed` and `total`. These are the current progress and the total progress, respectively. They should be integers.

Example
-------

go

```go
package main

import (
	"github.com/mshafiee/progressbar"
	"time"
)

func main() {
	total := 100
	for i := 0; i <= total; i++ {
		progressbar.ArrowProgressBar(i, total)
		time.Sleep(50 * time.Millisecond)
	}
}
```

This will display an arrow progress bar that fills up as the loop progresses:

```
[======================>............................] 45.00%
```