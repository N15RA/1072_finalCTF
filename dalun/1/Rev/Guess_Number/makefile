CXX=g++
FLAG=-fstack-protector-all -fpie -pie -z now

all: guess

guess: guessnum.cpp
	$(CXX) $(FLAG) $< -o $@