all:
	gcc -Wall lexer.c -o lexer

clean:
	rm *.o lexer
