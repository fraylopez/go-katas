package main

type Note struct {
	value int
}

func InvalidNote() Note {
	return Note{-1}
}
