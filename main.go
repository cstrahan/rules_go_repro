package main

// extern void cc_main();
import "C"

func main() {
  C.cc_main()
}