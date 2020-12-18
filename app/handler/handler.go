package handler

import (
	"github.com/memochou1993/github-rankings/logger"
	"time"
)

type Worker struct {
	starter chan struct{}
}

func NewWorker() *Worker {
	return &Worker{
		starter: make(chan struct{}, 1),
	}
}

func (w *Worker) BuildUserModel() {
	u := NewUserHandler()
	u.Init(w.starter)
	<-w.starter
	go w.collectUsers()
	go w.updateUsers()
	go w.rankUsers()
}

func (w *Worker) collectUsers() {
	u := NewUserHandler()
	t := time.NewTicker(10 * time.Minute) // FIXME
	for ; true; <-t.C {
		if err := u.Collect(); err != nil {
			logger.Error(err.Error())
		}
	}
}

func (w *Worker) updateUsers() {
	u := NewUserHandler()
	t := time.NewTicker(10 * time.Minute) // FIXME
	for ; true; <-t.C {
		if err := u.Update(); err != nil {
			logger.Error(err.Error())
		}
	}
}

func (w *Worker) rankUsers() {
	u := NewUserHandler()
	t := time.NewTicker(10 * time.Minute) // FIXME
	for ; true; <-t.C {
		u.RankFollowers()
		u.RankGistStars()
		u.RankRepositoryStars()
		u.RankRepositoryStarsByLanguage()
	}
}