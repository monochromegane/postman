package postman

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/clipperhouse/fsnotify"
	"github.com/mattn/go-pubsub"
	"github.com/otiai10/gosseract"
)

type Postman struct {
	*pubsub.PubSub
	done    chan error
	dir     string
	watcher *fsnotify.Watcher
}

func NewPostman(dir string) *Postman {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		panic(err)
	}
	return &Postman{
		PubSub:  pubsub.New(),
		done:    make(chan error),
		dir:     dir,
		watcher: watcher,
	}
}

func (p Postman) Run() {
	defer func() {
		p.Leave(nil)
		p.watcher.Close()
	}()

	// subscribe
	p.Sub(p.onCreate)
	p.Sub(p.onScan)

	// scan already exist files
	files, err := ioutil.ReadDir(p.dir)
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	for _, file := range files {
		p.Pub(filepath.Join(p.dir, file.Name()))
	}

	// watch directory
	err = p.watch()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	// wait
	err = <-p.done
	fmt.Printf("%v\n", err)
}

type onCreate func(string)
type file struct {
	name    string
	content string
}
type onScan func(file)

func (p Postman) onCreate(name string) {
	out := gosseract.Must(map[string]string{"src": name})
	p.Pub(file{name, out})
}

func (p Postman) onScan(file file) {
	// post to api
	fmt.Printf("%v\n", file)
	// remove file if successfuly
}
