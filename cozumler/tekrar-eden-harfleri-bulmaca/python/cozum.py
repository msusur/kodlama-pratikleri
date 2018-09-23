import collections

text = input("Bir cümle yazınız:")

def repeatStringCount(text):
    d = collections.defaultdict(int)
    for c in text:
        d[c] +=1
    for c in sorted(d,key=d.get,reverse=True):
        if d[c] > 1:
            print("%s : %d" %(c, d[c]))


repeatStringCount(text)

