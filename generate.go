package maestro

//go:generate echo "* installing govendor..."
//go:generate go get -u github.com/kardianos/govendor
//go:generate echo "* syncing dependencies. Go get a coffee..."
//go:generate $GOPATH/bin/govendor sync
