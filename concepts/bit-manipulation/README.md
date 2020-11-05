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
