package main

import (
	"sync"
)

type Song struct {
	name string
	artist string
	prev *Song
	next *Song
}

type Playlist struct {
	name string
	head *Song
	tail *Song
	nowPlaing *Song
	lock sync.RWMutex
}

func (p *Playlist) CreatePlaylist(name string) *Playlist {
	return &Playlist{
		name: name,
		lock: sync.RWMutex{},
	}
}

func (p *Playlist) AddSong(name, author string) {
	p.lock.Lock()
	defer p.lock.Unlock()
	
	s := &Song{
		name:   name,
		artist: author,
	}
	if p.head == nil {
		p.head = s
	} else {
		currentNode := p.tail
		currentNode.next = s
		s.prev = p.tail
	}
	p.tail = s
}


func main() {
	playlist := &Playlist{}
	playlist.CreatePlaylist("My songs")
	playlist.AddSong("New Born", "Muse")
}
