#include<stdio.h>

// ocozalp
int main()
{
    for(int i = 1; i<=110; i++) printf("%c", "*\n"[i%11==0]);
    return 0;
}
