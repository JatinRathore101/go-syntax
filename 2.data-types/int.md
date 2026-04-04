```md
# 📌 Integer Types in Go

## 🔹 int
- Platform-dependent:
  - 32-bit system → 32 bits
  - 64-bit system → 64 bits
- Range:
  - 32-bit → −2,147,483,648 to 2,147,483,647
  - 64-bit → very large (~±9e18)
- Default integer type
- Best for general-purpose usage

## 🔹 int8
- Size: 8 bits (1 byte)
- Range: −128 to 127
- Used when memory optimization is important

    var x int8 = 100

## 🔹 int64
- Size: 64 bits (8 bytes)
- Range: very large (~±9e18)
- Used for large numbers (timestamps, IDs, etc.)

    var y int64 = 10000000000

## ⚠️ Type Conversion Required
Go does NOT auto-convert between integer types:

    var a int8 = 10
    var b int64 = 20

    // ❌ Error
    // c := a + b

    // ✅ Correct
    c := int64(a) + b

## 🧠 Integer Rule of Thumb
- Use int → general use  
- Use int8 → memory-critical cases  
- Use int64 → large values / consistency across systems  
```
