#include <linux/mm.h> 
#include <linux/mmzone.h> 
#include <linux/blkdev.h> 
#include <linux/list.h> 
#include <linux/swap.h> 
#include <linux/cpumask.h> 
#include <linux/kernel_stat.h>

unsigned long cpu_user_time(int c)
{
	struct kernel_cpustat *base = (struct kernel_cpustat *)((unsigned long)__per_cpu_offset[c]+(unsigned long)&kernel_cpustat);
	return base->cpustat[CPUTIME_USER];
}

unsigned long cpu_nice_time(int c)
{
	struct kernel_cpustat *base = (struct kernel_cpustat *)((unsigned long)__per_cpu_offset[c]+(unsigned long)&kernel_cpustat);
	return base->cpustat[CPUTIME_NICE];
}

unsigned long cpu_sys_time(int c)
{
	struct kernel_cpustat *base = (struct kernel_cpustat *)((unsigned long)__per_cpu_offset[c]+(unsigned long)&kernel_cpustat);
	return base->cpustat[CPUTIME_SYSTEM];
}

unsigned long cpu_idle_time(int c)
{
	struct kernel_cpustat *base = (struct kernel_cpustat *)((unsigned long)__per_cpu_offset[c]+(unsigned long)&kernel_cpustat);
	return base->cpustat[CPUTIME_IDLE];
}

unsigned long cpu_iowait_time(int c)
{
	struct kernel_cpustat *base = (struct kernel_cpustat *)((unsigned long)__per_cpu_offset[c]+(unsigned long)&kernel_cpustat);
	return base->cpustat[CPUTIME_IOWAIT];
}

unsigned long cpu_hirq_time(int c)
{
	struct kernel_cpustat *base = (struct kernel_cpustat *)((unsigned long)__per_cpu_offset[c]+(unsigned long)&kernel_cpustat);
	return base->cpustat[CPUTIME_IRQ];
}

unsigned long cpu_sirq_time(int c)
{
	struct kernel_cpustat *base = (struct kernel_cpustat *)((unsigned long)__per_cpu_offset[c]+(unsigned long)&kernel_cpustat);
	return base->cpustat[CPUTIME_SOFTIRQ];
}

unsigned long cpu_steal_time(int c)
{
	struct kernel_cpustat *base = (struct kernel_cpustat *)((unsigned long)__per_cpu_offset[c]+(unsigned long)&kernel_cpustat);
	return base->cpustat[CPUTIME_STEAL];
}