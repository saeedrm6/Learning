# In The Name Of God
# ========================================
# [] File Name : Makefile
#
# [] Creation Date : 31-01-2015
#
# [] Created By : Parham Alvani (parham.alvani@gmail.com)
# =======================================

P=preproc
OBJ=preproc.o
LIB=
MLIB=

ifdef LIB
CFLAGS=-Wall -O0 -g -std=gnu11 \
       $(shell pkg-config --cflags $(LIB))
LDLIBS=$(shell pkg-config --libs $(LIB)) $(MLIB)
else
CFLAGS=-Wall -O0 -g -std=gnu11
LDLIBS=$(MLIB)
endif
ASFLAGS=-Wall -O0 -g

$(P): $(OBJ)
preproc.o:
	$(CC) -E preproc.c > preproc-E.c
	$(CC) $(CFLAGS) -c preproc.c

.PHONY: clean
clean:
	rm $(OBJ) $(P)
