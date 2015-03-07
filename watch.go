package postman

import "github.com/clipperhouse/fsnotify"

func (p Postman) watch() error {
	go func() {
		// file notify
		for {
			select {
			case event := <-p.watcher.Events:
				switch {
				case event.Op&fsnotify.Create == fsnotify.Create:
					p.Pub(event.Name)
				}
			case err := <-p.watcher.Errors:
				p.done <- err
			}
		}
	}()

	err := p.watcher.Add(p.dir)
	if err != nil {
		return err
	}
	return nil
}
