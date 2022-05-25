package sd

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"net/http"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

// HealthCheck 检查是否健康
func HealthCheck(ctx *gin.Context) {
	message := "OK"
	ctx.String(http.StatusOK, "\n"+message)
	log.Info("路由正常")
}

// DiskCheck 检查磁盘
func DiskCheck(ctx *gin.Context) {
	// 返回文件系统使用的情况，返回的是一个UsageStat类型的
	usage, _ := disk.Usage("/")
	// usedMB已使用多少MB
	usedMB := int(usage.Used) / MB
	// usedGB已使用多少GB
	usedGB := int(usage.Used) / GB
	// totalMB一共有多少MB
	totalMB := int(usage.Total) / MB
	// totalGB一共有多少GB
	totalGB := int(usage.Total) / GB
	// 硬盘使用的百分比
	usedPercent := int(usage.UsedPercent)

	status := http.StatusOK
	text := "OK"
	if usedPercent >= 95 {
		status = http.StatusOK
		text = "危急"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "警告"
	}
	message := fmt.Sprintf("%s - 可用空间：%dMB(%dGB) / %dMB(%dGB) | 已使用：%d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	log.Info(message)
	ctx.String(status, "\n"+message)
}

// CPUCheck 检查CPU
func CPUCheck(ctx *gin.Context) {
	// 返回系统中物理或逻辑内核的数量 传入false，返回物理核数，传入true，返回逻辑核数
	physicsKernelCounts, _ := cpu.Counts(false)
	avg, _ := load.Avg()

	load1 := avg.Load1
	load5 := avg.Load5
	load15 := avg.Load15
	//log.Infof(fmt.Sprintf("avg:1min = %v ,5min = %v,15min = %v", avg.Load1, avg.Load5, avg.Load15))
	status := http.StatusOK
	text := "OK"
	if load5 >= float64(physicsKernelCounts-1) {
		status = http.StatusInternalServerError
		text = "危急"
	} else if load5 >= float64(physicsKernelCounts-2) {
		status = http.StatusTooManyRequests
		text = "警告"
	}
	message := fmt.Sprintf("%s - 平均负荷: %.2f, %.2f, %.2f | 物理内核: %d", text, load1, load5, load15, physicsKernelCounts)
	log.Info(message)
	ctx.String(status, "\n"+message)
}

// RAMCheck men检查磁盘使用情况
func RAMCheck(ctx *gin.Context) {
	// 内存使用统计
	usage, _ := mem.VirtualMemory()
	// 已使用多少MB
	usedMB := int(usage.Used) / MB
	// 已使用多少GB
	usedGB := int(usage.Used) / GB
	// 一共有多少MB
	totalMB := int(usage.Total) / MB
	// 一共有多少GB
	totalGB := int(usage.Total) / GB
	// 硬盘使用的百分比
	usedPercent := int(usage.UsedPercent)

	status := http.StatusOK
	text := "OK"
	if usedPercent >= 95 {
		status = http.StatusOK
		text = "危急"
	} else if usedPercent >= 90 {
		status = http.StatusTooManyRequests
		text = "警告"
	}
	message := fmt.Sprintf("%s - 可用空间：%dMB(%dGB) / %dMB(%dGB) | 已使用：%d%%", text, usedMB, usedGB, totalMB, totalGB, usedPercent)
	log.Info(message)
	ctx.String(status, "\n"+message)
}
