package device_info

import (
    "fmt"
    "github.com/shirou/gopsutil/cpu"
    "github.com/shirou/gopsutil/mem"
    "github.com/shirou/gopsutil/host"
)

func GetDeviceInfo(){
   v, _ := mem.VirtualMemory()
   c, _ := cpu.Info()
   u, _ := host.Info()
  
    
  // Unique HostID
  fmt.Printf("Host ID: %v\n", u.HostID)
    
  // Platform 
  fmt.Printf("Platform: %v\n", u.Platform)
    
  // Memory
  fmt.Printf("Total memory: %v, Free memory: %v, Used memory: %v", Used Percent:%f%%\n", v.Total, v.Free, v.Used, v.UsedPercent)

  // CPU
  fmt.Printf("CPU model: %v, Number of cores: %v\n", c[0].ModelName, c[0].Cores)

  // Uptime
  fmt.Printf("Uptime: %v\n", u.Uptime)
}
