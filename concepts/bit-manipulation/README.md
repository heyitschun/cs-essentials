**Bit manipulation** is the act of manipulating bits or other pieces of data shorter than a word. It can be used for tasks such as data compression, encryption and optimization. Bit manipulation is done with bitwise operators: `AND`, `OR`, `XOR`, `NOT` and possibly others.

# Operators

Indications of a bit's position are counted from right to left. 

## `NOT`

The `NOT` or **complement** is a unary operation that performs logical negation on each bit. Put simply, ones become zeros and zeros become ones.

```
NOT 0101 (dec 5)
  = 1010 (dec 10)
```

## `AND`

The `AND` operator takes two equal length binary representations and performs the AND operation on each pair of bits. This is the same as multiplying bits that are in the same index.

```
    1010 (dec 10)
AND 1100 (dec 12)
  = 1000 (dec 8)
```

This operation can be used to determine whether a particular bit is *set* or *clear*. It can also be used to clear selected bits of a register in which each bit represents an individual Boolean state.

## `OR`

The `OR` operator takes two equal length binary representations and performs the *inclusive* OR operation on each pair of bits. The result in each position is `0` if both bits are `0`, otherwise the result in that position is `1`.

```
    0101 (dec 5)
OR  0011 (dec 3)
  = 0111 (dec 7)
```

## `XOR`

The `XOR` operator takes two equal length bit representations and performs the *exclusive* OR operation on each pair of bits. The result in each position is `1` only if one of the bits are `1`, other wise it will be `0`.

```
    0101 (dec 5)
XOR 0011 (dec 3)
  = 0110 (dec 6)
```

The `XOR` operator may be used to invert the selected bits in a register. It is a technique that can be used to manipulate bit patterns that represent Boolean states.

# Bit shifts

Operations in which digits are moved, or *shifted*, are called **bit shifts**. This can happen to the left or right. Bit registers in a computer have a fixed width, so bits can sometimes be shifted out of the register on one end.

## Bit addressing

If the width of the register (usually 32 or 64) is larger than the number of bits (usually 8) of the smallest addressable unit, usually called a *byte*, the shift operation induces an addressing scheme on the bits. Arithmetic and logical shift operations behave the same and a shift by 8 bit positions transport the bit pattern by 1 byte in the following way:

- Little-endian ordering: a left shift by 8 positions increases the byte address by 1
- Little-endian ordering: a right shift by 8 positions decreases the byte address by 1
- Big-endian ordering: a left shift by 8 positions decreases the byte address by 1
- Big-endian ordering: a right shift by 8 positions decreases the byte address by 1

## Arithmetic shift

In an **arithmetic shift**, the bits that are shifted out of either end are discarded. In a left arithmetic shift, zeros are shifted in on the right; in a right arithmetic shift, the sign bit is shifted on the left, thus preserving the sign of the operand.

```
    00010111 (dec 23) LEFT-SHIFT
  = 00101110 (dec 46)
```

In the above example, the left-most digit was shifted past the end of the register. A new `0` was shifted into the right-most position.

```
    10010111 (dec -23) RIGHT-SHIFT
  = 11001011 (dec -11)
```

In the above example, the right-most `1` was shifted out. A new `1` was copied into the left-most position, preserving the sign of the number.

A left arithmetic shift by `n` is equivalent to multiplying by `2^n` (assuming the number does not overflow). A right arithmetic shift by `n` of a [two's complement](https://en.wikipedia.org/wiki/Two%27s_complement) value is the equivalent of dividing by `2^n` and rounding toward negative infinity. If the binary number is treated as [one's complement](https://en.wikipedia.org/wiki/Ones%27_complement), then the same right-shift operation results in division by `2^n` and rounding towards zero.

## Logical shift

In a **logical shift**, the zeros are shifted in to replace the discarded bits. So logical and arithmetic left-shifts lead to the same results.

A logical right-shift is ideal for unsigned binary numbers, because it inserts zeros in to the most significant bits. Arithmetic right-shifts are better for signed two's complement binary representations.

## Circular shift

In a *rotate* operation, the bits are rotated as if the left and right ends of the register were joined. This is useful for retaining all existing bits. Its usage is found most often in cryptography.
