#include <stdio.h>

int main(void){

	long long int k = 1337 ;
	char str[15];

	scanf("%s",str);

	int i = 1 ;
	
	while(str[i-1]){

		k = k*i + str[i-1];
		i++;
	}

	printf("%lld\n",k);
}