# In The Name Of God
# ========================================
# [] File Name : Makefile
#
# [] Creation Date : 24-03-2015
#
# [] Created By : Parham Alvani (parham.alvani@gmail.com)
# =======================================

P=GSource
OBJ=GSource.o main.o
LIB=glib-2.0 gobject-2.0 gio-2.0
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
	$(CC) -o $(P) $(OBJ) $(LDLIBS)
GSource.o:
main.o:

.PHONY: clean
clean:
	rm $(OBJ) $(P)
