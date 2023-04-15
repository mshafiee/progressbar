package progressbar

import (
	"fmt"
	"strings"
	"time"
)

const (
	progressWidth = 50
	progressChars = `|/-\`
)

func ColorArrowProgressBar(completed, total int) {
	progress := float64(completed) / float64(total)
	blocks := int(progress * progressWidth)

	color := "\033[1;31m"
	if progress >= 0.5 {
		color = "\033[1;33m"
	}
	if progress >= 0.9 {
		color = "\033[1;32m"
	}

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%s", color))
	sb.WriteString(strings.Repeat("=", blocks))
	if progressWidth-blocks != 0 {
		sb.WriteString(">")
	} else {
		sb.WriteString("=")
	}
	sb.WriteString("\u001B[0m")
	sb.WriteString(strings.Repeat(".", progressWidth-blocks))
	sb.WriteString(fmt.Sprintf("] %.2f%%", progress*100))
	fmt.Printf("\r%s\033[0m", sb.String())
}

func RotatingProgressBar(completed, total int) {
	progress := float64(completed) / float64(total)
	blocks := int(progress * progressWidth)

	index := (completed / 2) % len(progressChars)
	char := progressChars[index : index+1]

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%s", strings.Repeat("=", blocks)))
	sb.WriteString(char)
	sb.WriteString(strings.Repeat(".", progressWidth-blocks))
	sb.WriteString(fmt.Sprintf("] %.2f%%", progress*100))
	fmt.Printf("\r%s", sb.String())
	time.Sleep(100 * time.Millisecond)
}

func BlockProgressBar(completed, total int) {
	progress := float64(completed) / float64(total)
	blocks := int(progress * progressWidth)

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%s%s] %.2f%%", strings.Repeat("█", blocks), strings.Repeat("░", progressWidth-blocks), progress*100))
	fmt.Printf("\r%s", sb.String())
}

func ArrowProgressBar(completed, total int) {
	progress := float64(completed) / float64(total)
	blocks := int(progress * progressWidth)

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%s>%s] %.2f%%", strings.Repeat("=", blocks), strings.Repeat(".", progressWidth-blocks), progress*100))
	fmt.Printf("\r%s", sb.String())
}

func ColorBlockProgressBar(completed, total int) {
	progress := float64(completed) / float64(total)
	blocks := int(progress * progressWidth)

	var sb strings.Builder
	sb.WriteString("\r[")
	sb.WriteString(strings.Repeat("\033[32m█\033[0m", blocks))
	sb.WriteString(strings.Repeat("\033[31m░\033[0m", progressWidth-blocks))
	sb.WriteString(fmt.Sprintf("] %.2f%%", progress*100))
	fmt.Printf("%s", sb.String())
}

func AnimatedArrowProgressBar(completed, total int) {
	progress := float64(completed) / float64(total)
	blocks := int(progress * progressWidth)

	index := int(time.Now().UnixNano()/100000000) % len(progressChars)
	char := progressChars[index : index+1]

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%s%s>%s] %.2f%%", strings.Repeat("=", blocks), char, strings.Repeat(".", progressWidth-blocks), progress*100))
	fmt.Printf("\r%s", sb.String())
}

func SimpleProgressBar(completed, total int) {
	progress := float64(completed) / float64(total)
	blocks := int(progress * progressWidth)

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%s%s] %.2f%%", strings.Repeat("=", blocks), strings.Repeat(" ", progressWidth-blocks), progress*100))
	fmt.Printf("\r%s", sb.String())
}

func ProgressBarWithPercentage(completed, total int) {
	progress := float64(completed) / float64(total)
	blocks := int(progress * progressWidth)

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("[%s%s] %.2f%% ", strings.Repeat("=", blocks), strings.Repeat(" ", progressWidth-blocks), progress*100))
	sb.WriteString(fmt.Sprintf("  %d/%d", completed, total))
	fmt.Printf("\r%s", sb.String())
}
