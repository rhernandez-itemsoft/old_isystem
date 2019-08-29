// +build linux
package hardwareinfo

import (
	"github.com/shirou/gopsutil/disk"
)

//GetUID  retornamos el identificador único
func GetUID() string {
	//return "UID001-002-003"
	return serialDisk()
}

//GetserialDiskUID  obtenemos el número de serie del disco duro
func serialDisk() string {
	serial := disk.GetDiskSerialNumber("/dev/sda")
	//	serial := fmt.Sprintf("%v", asd)
	return serial
}
