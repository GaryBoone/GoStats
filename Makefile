include $(GOROOT)/src/Make.inc

TARG=stats
GOFILES=\
	stats.go\
	regression.go\

include $(GOROOT)/src/Make.pkg
