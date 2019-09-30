generate:
	$(MAKE) -C builder generate
	go generate .

clean:
	rm -vf *.go
