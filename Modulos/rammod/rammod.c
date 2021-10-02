#include <linux/module.h>
#include <linux/init.h>
#include <linux/proc_fs.h>
#include <linux/sched.h>
#include <linux/uaccess.h>
#include <linux/fs.h>
#include <linux/sysinfo.h>
#include <linux/seq_file.h>
#include <linux/slab.h>
#include <linux/mm.h>
#include <linux/swap.h>
#include <linux/timekeeping.h>


#include <linux/kernel.h>
#include <linux/hugetlb.h>
#include <linux/mman.h>
#include <linux/mmzone.h>
#include <linux/percpu.h>
#include <linux/vmstat.h>
#include <linux/atomic.h>
#include <linux/vmalloc.h>
#ifdef CONFIG_CMA
#include <linux/cma.h>
#endif
#include <asm/page.h>
#include <asm/pgtable.h>

MODULE_LICENSE("GPL");
MODULE_AUTHOR("PROYECT 1 - GRUPO 12");
MODULE_DESCRIPTION("ESTADO RAM");
MODULE_VERSION("1.0");

extern unsigned long total_swapcache_pages(void);
#define total_swapcache_pages()			0UL

#define nameFile "rammod"

static int my_proc_show(struct seq_file *m, void *v){
    struct sysinfo i;
    long totalRam, free, buffer, totalUse, cached, percentUse;

    si_meminfo(&i);
	
    cached = global_node_page_state(NR_FILE_PAGES) - total_swapcache_pages() - i.bufferram;
            
    if (cached < 0)
		cached = 0;

    //disponible = si_mem_available() << (PAGE_SHIFT - 10);
    totalRam = i.totalram << (PAGE_SHIFT - 10);
    free = i.freeram << (PAGE_SHIFT - 10);
    cached = cached << (PAGE_SHIFT - 10);
    buffer = i.bufferram << (PAGE_SHIFT - 10);

    // CALCULO DE LA MEMORIA USADA
    totalUse = free + buffer + cached; 
    totalUse = totalRam - totalUse;
    percentUse = totalUse * 10000 / totalRam;
	
    // TOTALRAM,LIBRE,CACHED,BUFFER
    seq_printf(m, "{\"ram_total\":%lu,\"ram_uso\":%lu,\"ram_libre\":%lu,\"ram_percent\":%lu}\n",(totalRam / 1000),(totalUse / 1000),(free / 1000), percentUse);

	return 0;
}

static ssize_t my_proc_write(struct file *file, const char __user *buffer, size_t count, loff_t *f_pos){
	return 0;
}

static int my_proc_open(struct inode *inode, struct file *file){
	return single_open(file, my_proc_show, NULL);
}

static struct file_operations my_fops={
	.owner = THIS_MODULE,
	.open = my_proc_open,
	.release = single_release,
	.read = seq_read,
	.llseek = seq_lseek,
	.write = my_proc_write
};

static int __init test_init(void){
	struct proc_dir_entry *entry;
	entry = proc_create(nameFile, 0777, NULL, &my_fops);
	if(!entry){
		return -1;
	}else{
		printk(KERN_INFO "ESTADO MEMORIA RAM BEGIN\n");
	}
	return 0;
}

static void __exit text_exit(void){
	remove_proc_entry(nameFile, NULL);
	printk(KERN_INFO "ESTADO MEMORIA RAM END\n");
}

module_init(test_init);
module_exit(text_exit);