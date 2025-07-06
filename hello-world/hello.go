//  1. 宣告這個檔案屬於 main 套件 (package)
//     只有 main 套件才能被編譯成可執行檔
package main

// 2. 導入 (import) "fmt" 函式庫
//    "fmt" 包含了格式化 I/O (輸入/輸出) 的函式，例如列印到終端機
import "fmt"

// 3. 程式執行的入口：main 函式 (function)
func main() {
	// 使用 fmt 套件中的 Println 函式，它會將文字印出並換行
	fmt.Println("Hello, World!")
}
