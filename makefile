all:
	gcc -Wall lexer.c -o build/lexer && cp reserved_words.txt build/

clean:
	rm build/*

