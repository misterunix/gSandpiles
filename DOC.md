<!-- Code generated by gomarkdoc. DO NOT EDIT -->

# gSandpiles



## Index

- [Variables](<#variables>)
- [func Seconds2Days\(s float64\) string](<#Seconds2Days>)
- [func main\(\)](<#main>)
- [type grid](<#grid>)
  - [func NewGrid\(w, h, b int, rt float64\) grid](<#NewGrid>)
  - [func \(g \*grid\) Get\(x, y int\) int](<#grid.Get>)
  - [func \(g \*grid\) Resize\(w, h int\) grid](<#grid.Resize>)
  - [func \(g \*grid\) SaveImage\(\)](<#grid.SaveImage>)
  - [func \(g \*grid\) Set\(x, y, v int\)](<#grid.Set>)
  - [func \(g \*grid\) StartReport\(\)](<#grid.StartReport>)
  - [func \(g \*grid\) Topple\(\)](<#grid.Topple>)
- [type system](<#system>)


## Variables

<a name="timeImg"></a>

```go
var timeImg = time.Time{}
```

<a name="timeStart"></a>

```go
var timeStart = time.Time{}
```

<a name="Seconds2Days"></a>
## func [Seconds2Days](<https://github.com/misterunix/gSandpiles/blob/main/utilities.go#L6>)

```go
func Seconds2Days(s float64) string
```

Convert seconds to days, hours, minutes and seconds

<a name="main"></a>
## func [main](<https://github.com/misterunix/gSandpiles/blob/main/main.go#L10>)

```go
func main()
```



<a name="grid"></a>
## type [grid](<https://github.com/misterunix/gSandpiles/blob/main/types.go#L22-L28>)

Grid struct

```go
type grid struct {
    Width   int
    Height  int
    Bits    int
    RunTime float64
    Cells   []int
}
```

<a name="NewGrid"></a>
### func [NewGrid](<https://github.com/misterunix/gSandpiles/blob/main/grid.go#L11>)

```go
func NewGrid(w, h, b int, rt float64) grid
```

Create a new grid with the given width, height, bits and runtime

<a name="grid.Get"></a>
### func \(\*grid\) [Get](<https://github.com/misterunix/gSandpiles/blob/main/grid.go#L28>)

```go
func (g *grid) Get(x, y int) int
```

Get the value of the cell at x, y

<a name="grid.Resize"></a>
### func \(\*grid\) [Resize](<https://github.com/misterunix/gSandpiles/blob/main/grid.go#L33>)

```go
func (g *grid) Resize(w, h int) grid
```

Resize the grid to the new size. Copying the old grid to the center of the new grid

<a name="grid.SaveImage"></a>
### func \(\*grid\) [SaveImage](<https://github.com/misterunix/gSandpiles/blob/main/image.go#L13>)

```go
func (g *grid) SaveImage()
```

Save the image to disk

<a name="grid.Set"></a>
### func \(\*grid\) [Set](<https://github.com/misterunix/gSandpiles/blob/main/grid.go#L23>)

```go
func (g *grid) Set(x, y, v int)
```

Set the value of the cell at x, y

<a name="grid.StartReport"></a>
### func \(\*grid\) [StartReport](<https://github.com/misterunix/gSandpiles/blob/main/grid.go#L85>)

```go
func (g *grid) StartReport()
```



<a name="grid.Topple"></a>
### func \(\*grid\) [Topple](<https://github.com/misterunix/gSandpiles/blob/main/grid.go#L101>)

```go
func (g *grid) Topple()
```

Topple the grid. Stopping only to resize the grid if the grid is too small

<a name="system"></a>
## type [system](<https://github.com/misterunix/gSandpiles/blob/main/types.go#L11-L19>)

System struct

```go
type system struct {
    Width     int
    Height    int
    BitsStart int
    BitsEnd   int
    Resume    bool
    Unique    string // Unique part of the filenames
    LastSave  string
}
```

<a name="System"></a>var Grid grid

```go
var System system
```

Generated by [gomarkdoc](<https://github.com/princjef/gomarkdoc>)