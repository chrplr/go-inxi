BINARY  = inxi-go
LDFLAGS = -s -w

.PHONY: all clean

all: $(BINARY)

$(BINARY): *.go
	go build -ldflags="$(LDFLAGS)" -o $(BINARY) .

clean:
	rm -f $(BINARY)
