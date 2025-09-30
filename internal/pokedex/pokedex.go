package pokedex

import (
    "sync"
    "time"
)

type Pokedex struct {
    myPokemons	map[string]Pokemon
    mux		*sync.RWMutex
}

type Pokemon struct {
    caughtOn	time.Time
    val		[]byte
}

func NewPokedex() Pokedex {
    p := Pokedex{
	myPokemons: make(map[string]Pokemon),
	mux:	    &sync.RWMutex{},
    }
    return p
}

func (p *Pokedex) Add(name string, value []byte) error {
    p.mux.Lock()
    defer p.mux.Unlock()
    p.myPokemons[name] = Pokemon{
	caughtOn:   time.Now().UTC(),
	val:	    value,
    }
    return nil
}

func (p *Pokedex) Get(name string) ([]byte, bool) {
    p.mux.RLock()
    defer p.mux.RUnlock()
    pokemon, ok := p.myPokemons[name]
    return pokemon.val, ok
}
