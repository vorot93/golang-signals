# golang-signals
Very simple signals library.

## Example usage
```
func main() {
	type MyType struct {
		MySignal Signal
	}

	var v MyType
	var result string

	var id = v.MySignal.Connect(func(data interface{}) {
		var str, _ = data.(string)
		result = str
	})
	v.MySignal.Emit("edited")
	
	fmt.Printf("Callback %d printed: %s\n", id, result)
}
```
