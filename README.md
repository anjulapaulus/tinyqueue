[![Build Status](https://travis-ci.com/anjulapaulus/tinyqueue.svg?branch=master)](https://travis-ci.com/anjulapaulus/tinyqueue)
[![codecov](https://codecov.io/gh/anjulapaulus/tinyqueue/branch/master/graph/badge.svg)](https://codecov.io/gh/anjulapaulus/config)
[![HitCount](http://hits.dwyl.com/anjulapaulus/tinyqueue.svg)](http://hits.dwyl.com/anjulapaulus/tinyqueue)
# tinyqueue

A simple and smallest binary heap priority queue in golang.


#### Installation
````
go get -u github.com/anjulapaulus/tinyqueue
````


#### Acknowledgement
The current implementation is inspired on the javascript library [tinyqueue](https://github.com/mourner/tinyqueue) by Vladimir Agafonkin.

#### Implementation

````
package main

import queue "github.com/anjulapaulus/tinyqueue"

type Val int64

func (a Val) Compare (b queue.Item) bool{
	return a <= b.(Val)
}

func main(){
	q:= queue.Queue{}   //Initialize Queue

	q.Push(Val(2))      //Insert Items to queue
	q.Push(Val(5))
	q.Push(Val(10))

	array := q.All()    //returns array of elemeents in queue

	q.Peek()    //returns the top item without removal

	q.Len()     //length of queue

	q.Pop()     //removes the top item and displays
}
