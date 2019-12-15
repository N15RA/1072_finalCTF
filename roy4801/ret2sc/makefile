FLAGS=-m32 -fno-stack-protector -zexecstack

all: ret2sc

ret2sc: main.c
	gcc $(FLAGS) $< -o $@

.PHONY: clean

clean:
	rm -f ./ret2sc