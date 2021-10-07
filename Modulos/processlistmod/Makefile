obj-m += processlistmod.o
processlistmod-objs += main.o
processlistmod-objs += cpu.o

all:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) modules
clean:
	make -C /lib/modules/$(shell uname -r)/build M=$(PWD) clean