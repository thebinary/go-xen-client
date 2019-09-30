generate:
	$(MAKE) -C builder generate
	go generate .

clean:
	rm -vf *.gen.go *_string.go
