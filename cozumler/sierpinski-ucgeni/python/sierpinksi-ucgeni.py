# https://rosettacode.org/wiki/Sierpinski_triangle#Python

def sierpinski(n):
    d = ["*"]
    for n in range(n):
        sp = " " * (2 ** n)
        d = [sp + x + sp for x in d] + [ x + " " + x for x in d]
    return d

print("\n".join(sierpinski(4)))
