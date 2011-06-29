include $(GOROOT)/src/Make.inc

TARG=stats
GOFILES=\
	stats.go\
	regression.go\
	rand_normal.go

include $(GOROOT)/src/Make.pkg
