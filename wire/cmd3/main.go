package main

type Foo int 
type Bar int 
 
func ProvideFoo() Foo {
	return Foo(123)
} 
 
func ProvideBar() Bar {
	return Bar(456)
} 
 
type FooBar struct { 
    MyFoo Foo 
    MyBar Bar 
} 
 
func main() {

}