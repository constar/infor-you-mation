package main

type Message interface {
	GetTitle() string
	GetContent() string
	GetUrl() string
}
