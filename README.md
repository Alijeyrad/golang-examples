# Topics Explained in this repo

### 1. Context Package: Cancellation, Timeout, and Background

The context package in Go provides a way to manage deadlines, cancellation signals, and other request-scoped values across API boundaries and between goroutines. It is essential for controlling the lifecycle of operations, especially in concurrent programming.

	•	Cancellation: Allows you to cancel operations when they are no longer needed.
	•	Timeout: Automatically cancels operations that exceed a specified duration.
	•	Background: context.Background() returns an empty context, often used as the root context.


```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Create a context with a timeout of 2 seconds
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Ensure resources are cleaned up

	// Simulate a long-running operation
	ch := make(chan string)
	go func() {
		time.Sleep(3 * time.Second) // Simulate work
		ch <- "Operation Complete"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
	case <-ctx.Done():
		fmt.Println("Operation timed out:", ctx.Err())
	}
}
```

### 2. Net Package: Creating a TCP Server

The net package allows you to create network applications. You can build TCP servers that handle client connections either by responding to each request and closing the connection or by keeping the connection open for continuous communication.

Request-Response Based TCP Server:

```go
package main

import (
	"fmt"
	"net"
)

func main() {
	// Listen on TCP port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	fmt.Println("TCP server listening on port 8080")

	for {
		// Accept a connection
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Handle the connection in a new goroutine
		go func(c net.Conn) {
			// Read data from the connection
			buf := make([]byte, 1024)
			n, _ := c.Read(buf)
			fmt.Printf("Received: %s\n", string(buf[:n]))
			// Write response and close the connection
			c.Write([]byte("Hello from server\n"))
			c.Close()
		}(conn)
	}
}
```

Socket-Based TCP Server:

```go
package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	// Listen on TCP port 8080
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()
	fmt.Println("Socket server listening on port 8080")

	for {
		// Accept a connection
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		// Handle the connection in a new goroutine
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	defer c.Close()
	reader := bufio.NewReader(c)
	for {
		// Read data from the connection
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected")
			return
		}
		fmt.Printf("Received: %s", message)
		// Echo the message back to the client
		c.Write([]byte("Echo: " + message))
	}
}
```

### 3. Writing Unit Tests and Integration Tests

	•	Unit Tests: Test individual functions for correctness.
	•	Integration Tests: Test the interaction between components, such as client-server communication.

Unit Test:

```go
Function (main.go):

package main

import "fmt"

// Add sums two integers
func Add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("2 + 3 =", Add(2, 3))
}

Unit Test (main_test.go):

package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5

	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}
```

Integration Test for TCP Server:

```go
package main

import (
	"bufio"
	"net"
	"testing"
	"time"
)

func TestTCPServer(t *testing.T) {
	go main() // Start the server
	time.Sleep(time.Second) // Wait for the server to start

	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer conn.Close()

	// Send a message to the server
	conn.Write([]byte("Ping\n"))

	// Read the response
	message, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		t.Fatal(err)
	}

	expected := "Hello from server\n"
	if message != expected {
		t.Errorf("Expected %q, got %q", expected, message)
	}
}
```

### 4. container/heap Package and Benchmarking

The container/heap package provides heap operations for any type that implements heap.Interface. Heaps are efficient for priority queue implementations. We can compare its performance with slices.

Heap Implementation:

```go
package main

import (
	"container/heap"
	"fmt"
)

// An IntHeap is a min-heap of integers.
type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] } // For min-heap
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }

// Pop removes and returns the minimum element (according to Less).
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[0]
	*h = old[1:n]
	return x
}

func main() {
	h := &IntHeap{5, 2, 3}
	heap.Init(h)
	heap.Push(h, 1)
	fmt.Printf("Minimum: %d\n", (*h)[0])

	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
}
```

Slice Implementation:

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{5, 2, 3}
	s = append(s, 1)
	sort.Ints(s)
	fmt.Printf("Minimum: %d\n", s[0])

	for _, val := range s {
		fmt.Printf("%d ", val)
	}
}

Benchmark Tests:

package main

import (
	"container/heap"
	"sort"
	"testing"
)

func BenchmarkHeap(b *testing.B) {
	h := &IntHeap{}
	heap.Init(h)
	for i := 0; i < b.N; i++ {
		heap.Push(h, i)
		heap.Pop(h)
	}
}

func BenchmarkSlice(b *testing.B) {
	s := []int{}
	for i := 0; i < b.N; i++ {
		s = append(s, i)
		sort.Ints(s)
		s = s[1:]
	}
}
```

### 5. JSON Marshalling and Unmarshalling, Encoders and Decoders

The encoding/json package allows you to encode (marshal) Go structs into JSON and decode (unmarshal) JSON into Go structs. Encoders and decoders are useful for streaming JSON data over connections.


```go
package main

import (
	"encoding/json"
	"fmt"
	"net"
)

// Message represents a simple message structure
type Message struct {
	Text string `json:"text"`
}

func main() {
	// Marshalling a struct to JSON
	msg := Message{Text: "Hello, World!"}
	jsonData, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}
	fmt.Println("Marshalled JSON:", string(jsonData))

	// Unmarshalling JSON to a struct
	var msg2 Message
	err = json.Unmarshal(jsonData, &msg2)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}
	fmt.Println("Unmarshalled Struct:", msg2)

	// Start a TCP server to demonstrate encoders and decoders
	go startServer()

	// Give the server a moment to start
	conn, err := net.Dial("tcp", ":8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	encoder := json.NewEncoder(conn)
	err = encoder.Encode(msg)
	if err != nil {
		fmt.Println("Error encoding:", err)
		return
	}
}

func startServer() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()
	fmt.Println("JSON server listening on port 8080")

	conn, err := ln.Accept()
	if err != nil {
		fmt.Println("Error accepting:", err)
		return
	}
	defer conn.Close()

	decoder := json.NewDecoder(conn)
	var msg Message
	err = decoder.Decode(&msg)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}
	fmt.Println("Received message:", msg)
}
```
