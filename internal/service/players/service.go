package players

import (
	"github.com/jmoiron/sqlx"
	"github.com/moruezabal/seminario-go/internal/config"
)

//Player ...
type Player struct {
	ID     int64
	name   string
	dorsal int
}

//PlayerService ...
type PlayersService interface {
	AddPlayer(Player) error
	FindByID(int) *Player
	FindAll() []*Player
}

type service struct {
	db   *sqlx.DB
	conf *config.Config
}

// New ...
func New(db *sqlx.DB, c *config.Config) (PlayersService, error) {
	return service{db, c}, nil

}

func (s service) AddPlayer(p Player) error {
	return nil
}

func (s service) FindByID(ID int) *Player {
	return nil
}

func (s service) FindAll() []*Player {
	var list []*Player
	list = append(list, &Player{0, "Juan", 5})
	return list
}
