#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <string.h>

char flag[] = "NISRA{NICE_REV_AND_NICETRY_<3<3<3<3<3<3<3<3888888888888888888888@_@}";

void makeflag(int k){

	printf("%d\n",k);	

	srand(k);

	for(int i=0; i < strlen(flag);++i){

		int x = rand() % 255 ;

		printf("\\x%x",flag[i]^x);
	}

	printf("\n");

}

int main(void){

	for (int i = 0; i < 9430678; ++i)
	{
		
		srand(i);
		int x = rand() % 9430678 ;

		if(x == i) makeflag(i);

	}

}