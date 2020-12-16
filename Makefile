setup:
	sudo apt-get update
	sudo apt-get install -y libgpgme-dev libdevmapper-dev

run:
	GO111MODULE=on go run -mod=vendor .

clean:
	rm -f tmp.img
