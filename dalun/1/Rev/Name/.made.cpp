#include <cstdio>
#include <cstring>

int main(void){

	char name[] = "h4nds0meda1un" ;

	char c ;
	int i = 0;
	while(scanf("%c",&c) != EOF){

		printf("%c",c ^ name[i%strlen(name)]);
		i++;
	}

}