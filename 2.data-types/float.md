```md id="f8k2mz"
# 📌 Floating Point Types in Go

## 🔹 float32
- Size: 32 bits (4 bytes)
- Precision: ~6–7 decimal digits
- Range: ~±1.18e−38 to ±3.4e38
- Used when memory matters

    var a float32 = 3.14

## 🔹 float64 (Default)
- Size: 64 bits (8 bytes)
- Precision: ~15–16 decimal digits
- Range: ~±2.23e−308 to ±1.8e308
- Used for most real-world calculations

    var b float64 = 3.141592653589793

## ⚠️ Default Type

    x := 3.14 // float64 by default

## ⚠️ Type Conversion Required

    var a float32 = 3.14
    var b float64 = 2.71

    // ❌ Error
    // c := a + b

    // ✅ Correct
    c := float64(a) + b

## ⚠️ Precision Issue

    fmt.Println(0.1 + 0.2)
    // Output: 0.30000000000000004

Reason: Binary floating-point representation

## 🧠 Float Rule of Thumb
- Use float64 → almost always  
- Use float32 → when memory is critical (large datasets, graphics, etc.)
```
