package utilities

import (
	"io"
	"time"

	"github.com/schollz/progressbar/v3"
)

// ProgressBar is a wrapper around the progressbar library
// so that we can add our own functions
type ProgressBar struct {
	io.WriteCloser
	*progressbar.ProgressBar
}

// NewProgressbarTime returns a new progress bar that expects to be time/absolute value controlled
func NewProgressbarTime(maximumInt int, buffer io.WriteCloser) (ProgressBar, error) {
	bar := progressbar.NewOptions(
		maximumInt,
		progressbar.OptionSetTheme(progressbar.Theme{Saucer: "█", SaucerPadding: "_", BarEnd: "< "}),
		progressbar.OptionSetWidth(100),
		progressbar.OptionSetWriter(buffer),
		progressbar.OptionSetRenderBlankState(true),
	)
	p := ProgressBar{
		buffer,
		bar,
	}
	return p, nil
}

// DemoBarTime - just a simple demo function
func (p *ProgressBar) DemoBarTime(maximumInt int) {
	defer p.Close()
	for i := 0; i < maximumInt; i++ {
		p.Add(1)
		time.Sleep(10 * time.Millisecond)
	}
}
func (p *ProgressBar) Incremement() {
	p.Add(1)
}

func (p *ProgressBar) Close() {

}
