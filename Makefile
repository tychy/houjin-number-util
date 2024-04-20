# get toukijo code
toukijo:
	curl -o toukijo.csv https://raw.githubusercontent.com/tychy/toukijo_code/main/toukijo.csv
	./gen_toukijo.sh

test:
	go test
