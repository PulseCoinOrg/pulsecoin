build:
	go build -o ./pulse ./cmd/gpulse

run:
	./pulse

clean:
	rm ./pulse
