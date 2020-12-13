def factorial(n):
    if n == 0:
        return 1
    else:
        return n * factorial(n-1)

a = [0, 2, 5, 10]
for x in a:
    print(factorial(x))

