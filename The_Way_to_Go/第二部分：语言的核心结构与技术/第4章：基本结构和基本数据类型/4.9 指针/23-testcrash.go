package main

func main() {
	var p *int = nil
	*p = 0
}

/*
	panic: runtime error: invalid memory address or nil pointer dereference
	[signal 0xc0000005 code=0x1 addr=0x0 pc=0x44ddb2]
*/
