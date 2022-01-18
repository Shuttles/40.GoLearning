package main

import "fmt"

type USB interface {
	Name() string
	PlugIn()
}

type FlashDisk struct {
	name string
}

func (fd FlashDisk) Name() string {
	return fd.name
}

func (fd FlashDisk) PlugIn() {
	fmt.Println(fd.name, "连入电脑中。。。。")
}

// method not in one interface
func (fd FlashDisk) hello() {
	fmt.Println("Hello! I am a method not in one interface")
}

type Mouse struct {
	name string
}

func (m Mouse) Name() string {
	return m.name
}

func (m Mouse) PlugIn() {
	fmt.Println(m.name, "连入电脑中，准备工作。。。")
}

func main() {
	fd := FlashDisk{"U盘"}
	fmt.Println(fd.Name())
	fd.PlugIn()

	m := Mouse{"鼠标"}
	fmt.Println(m.Name())
	m.PlugIn()

	// test
	// proves that a interface variable can not call its method not defined in this interface
	var usb USB = FlashDisk{"U盘"}
	usb.hello()
}
