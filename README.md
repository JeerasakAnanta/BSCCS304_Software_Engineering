# BSCCS304 วิศวกรรมซอฟต์แวร์ (Software Engineering)

**OOP Concepts (Object-Oriented Programming)**
- Interface 

---

## Features

- ใช้ **Interface** `Computer` เพื่อกำหนดพฤติกรรมพื้นฐานของคอมพิวเตอร์
- สร้าง Struct `Laptop` และ `Desktop` ที่ implement Interface `Computer`
- เรียกใช้ฟังก์ชัน **Start**, **GetSpecs**, และ **ShutDown** ของคอมพิวเตอร์แต่ละประเภท

---

## การติดตั้ง

1. Clone โปรเจกต์จาก GitHub
    
```bash
git clone https://github.com/JeerasakAnanta/BSCCS304_Software_Engineering
```

2. เข้าไปยังโฟลเดอร์โปรเจกต์
```bash
cd BSCCS304_Software_Engineering
```

3. สร้าง Go Module ใหม่

```bash
go mod init BSCCS304_Software_Engineering
```

---

## การรันโปรเจกต์

```bash
go run main.go
```

---

## ตัวอย่างผลลัพธ์

```bash
Dell Laptop กำลังเปิดเครื่อง...
Dell Laptop - CPU: Intel i7, RAM: 16GB
Dell Laptop กำลังปิดเครื่อง...
------
HP Desktop กำลังเปิดเครื่อง...
HP Desktop - CPU: AMD Ryzen 5, RAM: 32GB
HP Desktop กำลังปิดเครื่อง...
------
```

---

## แนวคิดที่ใช้

* **Abstraction (การซ่อนรายละเอียด):** Interface `Computer` ซ่อนการทำงานภายในของ Laptop และ Desktop
---

## การขยายเพิ่มเติม

คุณสามารถเพิ่มฟีเจอร์เพิ่มเติม เช่น:

* อัปเกรด RAM (`UpgradeRAM`)
* ติดตั้งโปรแกรม (`InstallSoftware`)
* เชื่อมต่อเครือข่าย (`ConnectNetwork`)
* หรือเพิ่มประเภทคอมพิวเตอร์อื่น ๆ


