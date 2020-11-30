package main

import (
  "fmt"
  "sync"
)

func processInput(ip int, c chan int, currEating *int, totalDone *int,
                  maxEating int) {
  if ip == 1 {
    if *currEating < maxEating {
      (*currEating)++
        //signal the philosopher to eat
        c <- 1    
    } else {
        c <- -1    
    }
  } else if ip == 0 {
    (*currEating)--
  } else {
    (*totalDone)++ 
  }
}

func runHost(chans [5]chan int, wg *sync.WaitGroup) {
 maxEating := 2
 currEating := 0
 totalDone := 0
 for {

  if totalDone == 5 {
    break
  }
   
   select {
     case p1 := <- chans[0]:
       processInput(p1, chans[0], &currEating, &totalDone, maxEating)
     case p2 := <- chans[1]:
       processInput(p2, chans[1], &currEating, &totalDone, maxEating)
     case p3 := <- chans[2]:
       processInput(p3, chans[2], &currEating, &totalDone, maxEating)
     case p4 := <- chans[3]:
       processInput(p4, chans[3], &currEating, &totalDone, maxEating)
     case p5 := <- chans[4]:
       processInput(p5, chans[4], &currEating, &totalDone, maxEating)
   }
 }
  
  wg.Done()
} 

func eat(id int, c chan int, wg *sync.WaitGroup) {
  count := 0
  for {
    if count == 3 {
      break
    }
    //fmt.Println("Inside ", id)
    //first ask from the host on the channel
    c <- 1
    resp := <- c //block until response is given by host
    if resp < 0 {
      continue
    }

    fmt.Println("Starting to eat ", id)
     
    //finished eating once
    fmt.Println("Finishing to eat ", id)
    //write 0 on the channel which means finished eating once
    c <- 0
    count++
  }
  //2 means fully finished
  c <- 2
  //fmt.Println("Done ", id)
  wg.Done()
}

func main() {
  var wg sync.WaitGroup
  var chans [5]chan int

  for i := range chans {
    chans[i] = make(chan int)
  }

  wg.Add(6)
  go runHost(chans, &wg)
  for i := 0; i < 5; i++ {
    go eat(i+1, chans[i], &wg)
  
  }
  wg.Wait()
}

