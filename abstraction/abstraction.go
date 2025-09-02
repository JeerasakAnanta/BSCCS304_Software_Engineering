// abstraction.go

package abstraction

import "fmt"

// สร้าง interface ชื่อ Computer เพื่อกำหนดพฤติกรรมของคอมพิวเตอร์
type Computer interface {
	Start()           // เปิดเครื่อง
	ShutDown()        // ปิดเครื่อง
	GetSpecs() string // แสดงสเปก
}

// Laptop struct implement interface Computer
type Laptop struct {
	Brand string
	RAM   int
	CPU   string
}

func (l Laptop) Start() {
	fmt.Println(l.Brand, "Laptop กำลังเปิดเครื่อง...")
}

func (l Laptop) ShutDown() {
	fmt.Println(l.Brand, "Laptop กำลังปิดเครื่อง...")
}

func (l Laptop) GetSpecs() string {
	return fmt.Sprintf("%s Laptop - CPU: %s, RAM: %dGB", l.Brand, l.CPU, l.RAM)
}

// Desktop struct implement interface Computer
type Desktop struct {
	Brand string
	RAM   int
	CPU   string
}

func (d Desktop) Start() {
	fmt.Println(d.Brand, "Desktop กำลังเปิดเครื่อง...")
}

func (d Desktop) ShutDown() {
	fmt.Println(d.Brand, "Desktop กำลังปิดเครื่อง...")
}

func (d Desktop) GetSpecs() string {
	return fmt.Sprintf("%s Desktop - CPU: %s, RAM: %dGB", d.Brand, d.CPU, d.RAM)
}

func Computer_Abstraction() {
	// สร้างตัวอย่าง Laptop และ Desktop
	var myLaptop Computer = Laptop{Brand: "Dell", CPU: "Intel i7", RAM: 16}
	var myDesktop Computer = Desktop{Brand: "HP", CPU: "AMD Ryzen 5", RAM: 32}

	// เก็บใน slice ของ Computer
	devices := []Computer{myLaptop, myDesktop}

	for _, device := range devices {
		device.Start()
		fmt.Println(device.GetSpecs())
		device.ShutDown()
		fmt.Println("------")
	}
}
