package models

import (
	list "github.com/migege/go-bidir-list"
)

type RoomModel struct {
	rooms map[string]*list.List
}

func NewRoomModel() *RoomModel {
	r := new(RoomModel)
	r.rooms = make(map[string]*list.List)
	return r
}

func (r *RoomModel) getRoom(room string) *list.List {
	l := r.rooms[room]
	if l == nil {
		l = list.NewList()
		r.rooms[room] = l
	}
	return l
}

func (r *RoomModel) AddMoney(room, user string, money uint) bool {
	if l := r.getRoom(room); l != nil {
		return l.AddMoney(user, money)
	}
	return false
}

func (r *RoomModel) Top(room string, n int) []*list.Item {
	if l := r.getRoom(room); l != nil {
		return l.Top(n)
	}
	return nil
}

func (r *RoomModel) CheckList(room string) string {
	if l := r.getRoom(room); l != nil {
		return l.CheckList()
	}
	return "No Room"
}

func (r *RoomModel) Reset(room string) bool {
	delete(r.rooms, room)
	return true
}
