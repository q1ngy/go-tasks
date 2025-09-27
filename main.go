package main

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	s := "hello"
	var stack []rune
	for _, v := range s {
		stack = append(stack, v)
	}
	for len(stack) > 0 {
		_ = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
}
