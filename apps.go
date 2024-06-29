package main

import (
	"fmt"
	"log"
	"runtime"
	"sync"
	"sync/atomic"
)

//user type
type user struct {
	name string
	email string
}


// changeEmail implements a method with a pointer receiver.
func (u *user) changeEmail(email string) {
	u.email=email
}

type notifier interface {
	notify()
}

func (u *admin) notify(){
	fmt.Printf("Sending email to user %s<%s>\n", u.name, u.email)

}

type admin struct {
	user
	level string
}

// init is called before main.
// func init() {
// 	if len(os.Args) !=2 {
// 		fmt.Println("Usage: ./example <url>")
// 		os.Exit(-1)
// 	}
// }

func sendNotification(n notifier){
	n.notify()
}

var (
	counter int64
	wg sync.WaitGroup
)

func incCounter(id int64){
	defer wg.Done()
	for count:=0; count<2; count++ {
		atomic.AddInt64(&counter, id)
		runtime.Gosched()
	}
}

func init(){
	//initialize log settings
	log.SetPrefix("TRACE: ")
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
}

var (
	Trace *log.Logger
	Info *log.Logger
	Warning *log.Logger
	Error *log.Logger
)
// main is the entry point for the application.
func oldmain() {

	log.Println("message")
	//log.Fatalln("fatal message")
	defer log.Panicln("panic message")
	// wg.Add(2)
	// go incCounter(1)
	// go incCounter(2)
	// wg.Wait()
	// fmt.Println("Final Counter value ", counter)
	// fmt.Printf("Number of processor %d\n", runtime.GOMAXPROCS(runtime.NumCPU()))
	//create an admin user
	// ad := admin {
	// 	user: user{
	// 		name: "Ravi",
	// 		email:  "rs@email.com",
	// 	},
	// 	level: "supervisor",
	// }
	// sendNotification(&ad)
	// ad.notify()
	
	

	// fmt.Println("Hello, Go")

	// var b bytes.Buffer
	
	//      // Write a string to the buffer.
	//      b.Write([]byte("Hello"))
	
	//  // Use Fprintf to concatenate a string to the Buffer.
	//     fmt.Fprintf(&b, "World!")
	
	//      // Write the content of the Buffer to stdout.
	//      io.Copy(os.Stdout, &b)

	// 	 // Get a response from the web server. 
// r, err := http.Get(os.Args[1]) 
// if err!=nil {
// 	   fmt.Println(err)
// 	   return
// }
// // Copies from the Body to Stdout.
//  //io.Copy(os.Stdout, r.Body)
//  if err := r.Body.Close(); err != nil {
// 		fmt.Println(err)
// 	 }
}
