#include <stdio.h>
#include <unistd.h>

void disable_buffer()
{
    setvbuf(stdin, NULL, _IONBF, 0);
    setvbuf(stdout, NULL, _IONBF, 0);
}

void input()
{
    char buf[20];
    printf("\n%p", buf);
    printf("\nthis might be helpful to you.");
    gets(buf);
}

int main()
{
    disable_buffer();
    printf("i put all the Sh3LL inside the Cod3."
           " go for finding it!");
    input();
    return 0;
}