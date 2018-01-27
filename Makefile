ifdef SLACK_TOKEN
else
	# empty token
	echo "Empty SLACK_TOKEN"
endif

all: main.go
	go build -o notifier -ldflags "-X main.token=$(SLACK_TOKEN)" main.go

clean:
	rm notifier
