package main

import "fmt"

func main() {
    // Khai báo biến
    var name string = "Tân"
    age := 22

    fmt.Println("Tên:", name)
    fmt.Println("Tuổi:", age)

    // Nhập số lần lặp từ người dùng
    var n int
    fmt.Print("Nhập số lần lặp: ")
    fmt.Scan(&n)

    // Khai báo slice để lưu số chẵn và lẻ
    evenNumbers := []int{}
    oddNumbers := []int{}

    // Vòng lặp for từ 1 đến n
    for i := 1; i <= n; i++ {
        fmt.Printf("Lần lặp thứ %d\n", i)

        if i%2 == 0 {
            fmt.Println("=> Đây là số chẵn")
            evenNumbers = append(evenNumbers, i) // Lưu vào danh sách số chẵn
        } else {
            fmt.Println("=> Đây là số lẻ")
            oddNumbers = append(oddNumbers, i) // Lưu vào danh sách số lẻ
        }
    }

    // Hiển thị kết quả
    fmt.Println("\n🔹 Tổng số lần lặp:", n)
    fmt.Println("🔹 Danh sách số chẵn:", evenNumbers)
    fmt.Println("🔹 Danh sách số lẻ:", oddNumbers)
    fmt.Println("🔹 Tổng số chẵn:", len(evenNumbers))
    fmt.Println("🔹 Tổng số lẻ:", len(oddNumbers))
}
