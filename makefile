all:
	gcc -Wall src/compiler.c -o build/compiler.exe && cp resources/reserved_words.txt build/

clean:
	rm build/*
