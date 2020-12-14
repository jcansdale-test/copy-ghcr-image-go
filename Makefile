run:
	GO111MODULE=on go run -mod=vendor .

clean:
	rm -f tmp.img
