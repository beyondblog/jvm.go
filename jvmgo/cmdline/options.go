package cmdline

import (
	"strconv"
	"strings"
)

const (
	_1k = 1024
	_1m = _1k * _1k
	_1g = _1k * _1m
)

type Options struct {
	argReader    *ArgReader
	classpath    string
	verboseClass bool
	xss          int
	Xcpuprofile  string
	XuseJavaHome bool
}

func newOptions(argReader *ArgReader) *Options {
	return &Options{
		argReader: argReader,
		xss:       16 * _1k,
	}
}

// getters
func (self *Options) Classpath() string {
	return self.classpath
}
func (self *Options) VerboseClass() bool {
	return self.verboseClass
}
func (self *Options) Xss() int {
	return self.xss
}

func parseOptions(argReader *ArgReader) *Options {
	options := newOptions(argReader)

	for argReader.hasMoreOptions() {
		optionName := argReader.removeFirst()
		_ = options.parseClassPathOption(optionName) ||
			options.parseVerboseOption(optionName) ||
			options.parseXssOption(optionName) ||
			options.parseXcpuprofile(optionName) ||
			options.parseXuseJavaHome(optionName)
		// todo
	}

	return options
}

func (self *Options) parseClassPathOption(optionName string) bool {
	if optionName == "-classpath" || optionName == "-cp" {
		self.classpath = self.argReader.removeFirst()
		return true
	}
	return false
}

func (self *Options) parseVerboseOption(optionName string) bool {
	if optionName == "-verbose" || optionName == "-verbose:class" {
		self.verboseClass = true
		return true
	}
	return false
}

// -Xss<size>[g|G|m|M|k|K]
func (self *Options) parseXssOption(optionName string) bool {
	if strings.HasPrefix(optionName, "-Xss") {
		size := optionName[4:]
		switch size[len(size)-1] {
		case 'g', 'G':
			self.xss = _1g * parseInt(size[:len(size)-1])
		case 'm', 'M':
			self.xss = _1m * parseInt(size[:len(size)-1])
		case 'k', 'K':
			self.xss = _1k * parseInt(size[:len(size)-1])
		default:
			self.xss = parseInt(size)
		}
		return true
	}
	return false
}

func parseInt(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		panic(err.Error())
	}
	return i
}

func (self *Options) parseXcpuprofile(optionName string) bool {
	if optionName == "-Xcpuprofile" {
		self.Xcpuprofile = self.argReader.removeFirst()
		return true
	}
	return false
}

func (self *Options) parseXuseJavaHome(optionName string) bool {
	if optionName == "-XuseJavaHome" {
		self.XuseJavaHome = true
		return true
	}
	return false
}
