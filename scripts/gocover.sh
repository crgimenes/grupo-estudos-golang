

gocover () {
	echo "Testing..."
	t="/tmp/go-cover.$$.tmp"
	go test ./... -coverprofile=$t && go tool cover -html=$t -o /tmp/cover.html && rm -f $t || return 1
	echo "Open /tmp/cover.html"
	open /tmp/cover.html
}

