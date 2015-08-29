all:
	gcc -Wall lexer.c -o build/lexer && cp lib/reserved_words.txt build/

clean:
	rm build/*
