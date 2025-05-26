package main

func main() {
	console := ConsoleLog{}
	file := FileLog{
		ConsoleLog: ConsoleLog{},
		FilePath:   "/tmp",
	}
	Process(console)
	Process(file)
}
