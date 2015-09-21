all:
	gcc -Wall src/lexer.c -o build/lexer && cp resources/reserved_words.txt build/

clean:
	rm build/*
