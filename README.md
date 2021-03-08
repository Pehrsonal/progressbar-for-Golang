# Progressbar-For-Golang
Develop of a simple progressbar in Go. School project in the course PA1456 at BTH

works on all operating system with go installed!!

## Install
```
go get -u github.com/Pehrsonal/Progressbar-For-Golang/progressbar
```
## HOW TO USE
```
package main

import (
  "github.com/Pehrsonal/Progressbar-For-Golang/progressbar"
  "time"
)

func main() {
  bar := progressbar.StartNew(50, ...Custom features)
  for i := 0; i < bar.GetMaxvalue(); i++ {
		bar.Increment()
		time.Sleep(60 * time.Millisecond) // Sleep represent you doing something inside loop
  }
  bar.Finish()
}
```
Custom features : 
```
SetWidth(int) , ShowPercent(bool) , ShowTime(bool), Newdescription(string), Setstyle(struct Style)
```
Style and color :
Available colors atm = White, Black, Red, Blue, Green and Yellow!
```
//You can change style of the bar with the style struct. 
do so with function setStyle. Here is two examples
barstyle := progressbar.Style{
		StartChar:     "!",
		EndChar:       "!",
		ProgressChar:  "C",
		StartEndColor: "Red",
		ProgressColor: "Yellow",
	}

	barstyle2 := progressbar.Style{
		StartChar:     "{",
		EndChar:       "}",
		ProgressChar:  "%",
		StartEndColor: "Blue",
		ProgressColor: "Red",
	}
  Then you do it like this: 
  testbarwithStyle := progressbar.StartNew(50, progressbar.SetWidth(10), progressbar.SetStyle(barstyle))
```

## Time Limits 
125 hours for budget on this project.
for more info see -> [Time_Estimate_File](https://github.com/Pehrsonal/progressbar-for-Golang/blob/main/TIME_ESTIMATE.md)

## Todo List
[Link TODO LIST](https://github.com/Pehrsonal/progressbar-for-Golang/projects/1)

## Weekly Diary-blog
[Link Weekly Diary](https://docs.google.com/document/d/1dffjBnzQhBu6OpY11p0MNKGw9TnxKaFTgpuDHYp3rK4/edit?usp=sharing)

## License
See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
