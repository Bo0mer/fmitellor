

# logger
`import "github.com/Bo0mer/logger"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)

## <a name="pkg-overview">Overview</a>
Package logger implements simple logger library.




## <a name="pkg-index">Index</a>
* [type Logger](#Logger)
  * [func NewLogger(output io.Writer) *Logger](#NewLogger)
  * [func (l *Logger) Fatalf(format string, v ...interface{})](#Logger.Fatalf)
  * [func (l *Logger) Printf(format string, v ...interface{})](#Logger.Printf)


#### <a name="pkg-files">Package files</a>
[doc.go](/src/github.com/Bo0mer/logger/doc.go) [logger.go](/src/github.com/Bo0mer/logger/logger.go) 






## <a name="Logger">type</a> [Logger](/src/target/logger.go?s=77:116#L9)
``` go
type Logger struct {
    // contains filtered or unexported fields
}

```
Logger can be used to log stuff.







### <a name="NewLogger">func</a> [NewLogger](/src/target/logger.go?s=183:223#L14)
``` go
func NewLogger(output io.Writer) *Logger
```
NewLogger returns new Logger writing to the specified output.





### <a name="Logger.Fatalf">func</a> (\*Logger) [Fatalf](/src/target/logger.go?s=537:593#L27)
``` go
func (l *Logger) Fatalf(format string, v ...interface{})
```
Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).




### <a name="Logger.Printf">func</a> (\*Logger) [Printf](/src/target/logger.go?s=375:431#L22)
``` go
func (l *Logger) Printf(format string, v ...interface{})
```
Printf writes to the logger's output. Arguments are handled in the manner of
fmt.Printf.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)