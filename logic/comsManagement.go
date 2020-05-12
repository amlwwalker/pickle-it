package logic

//FileManager is a series of channels to communicate the state of a new file
// TODO: its unutilised at the moment
type FileCommsManager struct {
	errorsChannel   chan error   //if you want to error, send a message on this channel and it will be handled
	progressChannel chan float64 //whenever you want to update progress, use this channel
	resultChannel   chan string  //if and when a result is reached, send it over this channel
}

func (f *FileCommsManager) Close() {
	close(f.errorsChannel)
	close(f.progressChannel)
	close(f.resultChannel)
}

func NewFileManager() FileCommsManager {
	f := FileCommsManager{
		errorsChannel:   make(chan error),
		progressChannel: make(chan float64),
		resultChannel:   make(chan string),
	}
	return f
}
