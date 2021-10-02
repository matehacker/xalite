package main
import _ "net/http/pprof"
 
import (
    "fmt"
    "io"
    "log"
    "os"
)

var _ = fmt.Printf // For debugging; delete when done.
var _ io.Reader    // For debugging; delete when done.

func main1() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
    _ = fd
}
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
type ReadWriter struct {
    reader *Reader
    writer *Writer
}
type Job struct {
    Command string
    *log.Logger
}
func NewJob(command string, logger *log.Logger) *Job {
    return &Job{command, logger}
}
func (job *Job) Printf(format string, args ...interface{}) {
    job.Logger.Printf("%q: %s", job.Command, fmt.Sprintf(format, args...))
}
type Request struct {
    args        []int
    f           func([]int) int
    resultChan  chan int
}
func sum(a []int) (s int) {
    for _, v := range a {
        s += v
    }
    return
}
func handle(queue chan *Request) {
    for req := range queue {
        req.resultChan <- req.f(req.args)
    }
}
type Vector []float64
const numCPU = 4 // number of CPU cores

func (v Vector) DoAll(u Vector) {
    c := make(chan int, numCPU)  // Buffering optional but sensible.
    for i := 0; i < numCPU; i++ {
       type Request struct {
    args        []int
    f           func([]int) int
    resultChan  chan int
}
    }
    // Drain the channel.
    for i := 0; i < numCPU; i++ {
        <-c    // wait for one task to complete
    }
    // All done.
}
