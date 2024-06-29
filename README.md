# Devlopment Setup in Macos

## System

1. `brew install go`
2. Edit ~/.profile and below lines
<pre>
export GOPATH=~/go/
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
</pre>
3. Update the ~/.profile file using `source ~/.profile`
4. Edit `code ~/Library/ApplicationSupport/Code/User/settings.json` with the following values at the top inside {}
<pre>
 "go.toolsGopath": "/Users/rssingh/go/",
 "go.gopath": "/Users/rssingh/go/",
</pre>

## VS Code

5. Setup VS Code for Go Development, install plugin Go, Rich Go language support for Visual Studio code by Go team at Google
6. Setup Debuuger for Go in VS Code, Select Run and Debug icon and click on run by selecting a go program and allow to install the devl tool for debugging
and setup the .vscode/launch.json file with below content
<pre>
{
"name": "Launch Package",
"type": "go",
"request": "launch",
"mode": "debug",
"program": "${workspaceFolder}",
"args": ["serve"]
}

</pre>

# Methods to read the go documentation

We can refer the source code of the library/module using either one of below three approaches.

1. from command line `godoc <module_name>` or
2. from command line `godoc -http=:6060` and then open `http://localhost:6060` in browser
3. Refer the standard library from website `https://pkg.go.dev/std`

Note: The standard library of Go source code lies in the location `$GOROOT/src`

# Alernative process to build go

Refer [gb](http://getgb.io/) as go alternate build tool. It is vendoring tool to manage dependencies.
Ex: godep, vendor, gb

# Go Data Structure

Go has three different data structures that allow you to manage collections of data. Command on these structure will makes Go fun, fast, and flexible

- arrays: type of array and total no. of elements, var data [5] int , data type and lenth cannot be changed once declared, data := [5]int{4,2,6,7,10}, data :=[...]int{10,20,32,40,50, 45,50,65}, here the lenght is determined based on the initial data passed. data := [5]int{1:5, 2:15}, here index 1 and 2 are initaialized to 5 and 15 respectively while rest are initialized to zero. data := [5]\*int{0:new(int), 1:new(int)}; *data[0]=10; *data[1]=20

- slices: on 64 bit architecture, slice size is 24 bytes. It has pointer(8 bytes), length(8 bytes) and capacity(8 bytes) in its data structure. dynamic length array but same data type. slice := make([]string, 5), slice := []string{"Red","Green", "Blue", "Red"}, nil slice > var slice [] int, pointter nil, length 0 capacity 0, slice := make([] int, 0), slice := []int{}. It has function - append, len and cap. Calculate the slice length and capacity for slick[i:j] for capacity k
  length j-i
  capacity k-i
  Using range over slices with for loop we can iterate each element of slices.

- maps: unordered collection of key/value pairs. The strength of a map is its ability to retrieve data quickly based on the key. A key works like an index, pointing to the value you associate with that key. Map is implementted using hashtable. There are two methods to declare map

1. dict := make(map[string]int) key is type string and value is type int
2. dict := map[string]string{} The initial length will be based on the number of key/value pairs you specify during initialization. A map can be created for a slice of values dict := map[int]int[]string{}
   If you want to remove a key/value pair from the map, you use the built-in function delete.
   Example: Remove the key/value pair for the key "Coral". delete(colors, "Coral")

# Type in Go

Go is statically typed language. It means compiler wants to know type of every value in the program. And this ensure the safety as well. Reduce potential memory corruption and bugs and provide opportunity to compiler for more performant code. Each type has size and representation in memory.

## User Define Types

1. struct type
      <pre>
        type user struct {
         name string
         email string
         ext int
         privileged bool
         }
         </pre>

   declare a variable bill of type user as below. And once declared, the values are intiailzied to zeros of types.
   var bill user
   user {
   name: "Ravi"
   email: "rr@gmail.com",
   ext: 567,
   privileged: true,
   }
   lisa := user{"Lisa", "lisa@gmail.com", 345, false}
   type admin struct {
   person user
   level
   string
   }

   Value receivers operate on a copy of the value used to make the method call and pointer receivers operate on the actual value.
   After declaring a new type, try to answer this question before declaring methods for the type. Does adding or removing something from a value of this type need to create a new value or mutate the existing one?
   If the answer is create a new value, then use value receivers for your methods. If the answer is mutate the value, then use pointer receivers.

2. Reference type
   Reference types in Go are the set of slice, map, channel, interface, and function types. When you declare a variable from one of these types, the value that’s created is called a <i>header</i> value. The header value contains a pointer; there- fore, you can pass a copy of any reference type value and share the underlying data structure intrinsically.

## Interfaces

Polymorphism is the ability to write code that can take on different behavior through the implementation of types. Once a type implements an interface, an entire world of functionality can be opened up to values of that type.
Interfaces are types that just declare behavior. This behavior is never implemented by the interface type directly but instead by user-defined types via methods. When a user- defined type implements the set of methods declared by an interface type, values of the user-defined type can be assigned to values of the interface type. This assignment stores the value of the user-defined type into the interface value.
Interface values are two-word data structures. The first word contains a pointer to an internal table called an iTable, which contains type information about the stored value. The iTable contains the type of value that has been stored and a list of methods associated with the value. The second word is a pointer to the stored value. The combination of type information and pointer binds the relationship between the two values.
// notifier is an interface that defined notification
// type behavior.
type notifier interface {
notify()
}

## Exported or Unexported Identifiers in Go

Identifier name starting Captial letter is exported while identifier starting with lower is unexported

# Concurrency in Go

Multiple request for data on individual sockets at the same time.
Parallelism is about doing a lot of things at once. Concurrency is about managing a lot of things at once.

## Goroutines

When a function is created as a goroutine, it’s treated as an independent unit of work that gets scheduled and then executed on an available logical processor. Concurrency synchronization comes from a paradigm called communicating sequen- tial processes or CSP. CSP is a message-passing model that works by communicating data between goroutines instead of locking data to synchronize access.
The key data type for synchronizing and passing messages between goroutines is called a channel. Using channels makes it easier to write concurrent programs and makes them less prone to errors.

## Detecting and fixing race conditions

## Sharing data with channels

# Testing Go

`go test -v`

## Convention for wrting test case

<pre>
func TestFunction(t \*testing.T){
t.Log("description of test"){
   t.Logf()
   t.Errorf()
   }
}
</pre>

A test function must be an exported function that begins with the word Test. Not only must the function start with the word Test, it must have a signature that accepts a pointer of type testing.T and returns no value. If we don’t follow these conventions, the testing framework won’t recognize the function as a test function and none of the tooling will work against it.

The pointer of type testing.T is important. It provides the mechanism for report- ing the output and status of each test. There’s no one standard for formatting the out- put of your tests. I like the test output to read well, which does follow the Go idioms for writing documentation. For me, the testing output is documentation for the code. The test output should document why the test exists, what’s being tested, and the result of the test in clear complete sentences that are easy to read. Let’s see how I accomplish this as we review more of the code.

## Mocking

When webserver is not available, use "net/http/httptest" module

<pre>
func mockServer() *httptest.Server {
   f := func(w http.ResponseWriter, r *http.Request) {
      w.WriteHeader(200)
    w.Header().Set("Content-Type", "application/xml")
    fmt.Fprintln(w, feed)
   }
   return httptest.NewServer(http.HandlerFunc(f))
}
</pre>

## Benchmarking

run a go module BenchmarkSprintf
`go test -v -run="none" -bench="BenchmarkSprintf"`
