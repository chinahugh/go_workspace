package main

import (
	"fmt"
	"mypkg"
)

func main() {
	fmt.Println("Hello Go !!!")
	mypkg.Myfunc()
	/*ff :=1 ;
		fmt.Println(ff);

	    const (
	            a = iota   //0
	            b          //1
	            c          //2
	            d = "ha"   //独立值，iota += 1
	            e          //"ha"   iota += 1
	            f = 100    //iota +=1
	            g          //100  iota +=1
	            h = iota   //7,恢复计数
	            i          //8
	    )
	    fmt.Println(a,b,c,d,e,f,g,h,i)*/

	/* 定义局部变量 */
	/*var a int = 100
	  var b int = 200

	  /* 判断条件 */
	/*if (a == 100) {
	      /* if 条件语句为 true 执行 /
	      if (b == 200) {
	         /* if 条件语句为 true 执行 /
	         fmt.Printf("a 的值为 100 ， b 的值为 200\n" );
	      }
	  }
	  fmt.Printf("a 值为 : %d\n", a );
	  fmt.Printf("b 值为 : %d\n", b );*/

	/*var c1, c2, c3 chan int
	  var i1, i2 int
	  select {
	     case i1 = <-c1:
	        fmt.Printf("received ", i1, " from c1\n")
	     case c2 <- i2:
	        fmt.Printf("sent ", i2, " to c2\n")
	     case i3, ok := (<-c3):  // same as: i3, ok := <-c3
	        if ok {
	           fmt.Printf("received ", i3, " from c3\n")
	        } else {
	           fmt.Printf("c3 is closed\n")
	        }
	     default:
	        fmt.Printf("no communication\n")
	  }    */
	// for m := 1; m < 10; m++ {
	// 	for n := 1; n <= m; n++ {
	// 		fmt.Printf("%dx%d=%d ", n, m, m*n)
	// 	}
	// 	fmt.Println("")
	// }
}
