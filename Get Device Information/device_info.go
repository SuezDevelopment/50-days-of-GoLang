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
  
  
  // Memory
  fmt.Printf("Total memory: %v, Free memory: %v, Used memory: %v\n", v.Total, v.Free, v.Used)

  // CPU
  fmt.Printf("CPU model: %v, Number of cores: %v\n", c[0].ModelName, c[0].Cores)

  // Uptime
  fmt.Printf("Uptime: %v\n", u.Uptime)
}
