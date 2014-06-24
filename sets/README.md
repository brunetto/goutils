## String set for golang

From

* [http://programmers.stackexchange.com/questions/177428/sets-data-structure-in-golang](http://programmers.stackexchange.com/questions/177428/sets-data-structure-in-golang)
* [http://play.golang.org/p/_FvECoFvhq](http://play.golang.org/p/_FvECoFvhq)
* [https://groups.google.com/forum/#!topic/golang-nuts/lb4xLHq7wug](https://groups.google.com/forum/#!topic/golang-nuts/lb4xLHq7wug)
* [https://github.com/deckarep/golang-set/blob/master/](https://github.com/deckarep/golang-set/blob/master/)

Use like:

````go
set := NewIntSet()
set.Add(1)
set.Add(2)
set.Add(3)
fmt.Println(set.Get(2))
set.Remove(2)
fmt.Println(set.Get(2))
````
