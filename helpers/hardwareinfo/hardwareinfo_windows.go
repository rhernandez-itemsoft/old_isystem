// +build windows
package hardwareinfo

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

//GetUID  retornamos el identificador único
func GetUID() string {
	hostName := getHostName()
	disk, _ := getDiskPartition(strings.ToUpper(hostName))
	return serialDisk(disk)
}

//getHostName obtenemos el hostname
func getHostName() string {
	cmd := exec.Command("hostname")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	stringInfo := fmt.Sprintf("%s\n", stdoutStderr)
	stringInfo = strings.TrimSpace(stringInfo)
	return stringInfo
}

//getDiskPartition Obtenemos el numero de partition de C:\ (instalacion de windows)
func getDiskPartition(hostName string) (string, string) {
	// wmic logicaldisk where (DeviceID="C:") assoc /assocclass:Win32_LogicalDiskToPartition
	// where (DeviceID=\"C:\")
	cmd := exec.Command("wmic", "logicaldisk", "where", "(DeviceID='C:')", "assoc", "/assocclass:Win32_LogicalDiskToPartition")
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	stringInfo := fmt.Sprintf("%s\n", stdoutStderr)

	//	arrInfoClear := strings.Fields(stringInfo)
	//fmt.Println(arrInfoClear)
	stringInfo = strings.TrimSpace(stringInfo)
	stringInfo = strings.Replace(stringInfo, "\n", "", -1)
	stringInfo = strings.Replace(stringInfo, "\r", "", -1)
	stringInfo = strings.Replace(stringInfo, "  ", "", -1)
	stringInfo = strings.Replace(stringInfo, "\t", "", -1)

	var arrInfo []string
	arrInfo = strings.Split(stringInfo, hostName)
	//fmt.Println(arrInfo)

	stringInfo = ""
	for _, val := range arrInfo {
		value := strings.Replace(val, "\n", "", 1)
		//	fmt.Println("[" + value + "]")
		if strings.Index(value, "DiskPartition") > -1 {
			stringInfo = value
			break
		}
	}

	arrInfo = strings.Split(stringInfo, "DeviceID=\"")
	arrInfo = strings.Split(arrInfo[1], "\"")
	arrInfo = strings.Split(arrInfo[0], ",")

	disk := strings.Replace(arrInfo[0], "Disk #", "", 1)
	partition := strings.Replace(strings.TrimSpace(arrInfo[1]), "Partition #", "", 1)

	return disk, partition

}

//serialDisk Obtiene el numero de serie del disco duro
func serialDisk(disk string) string {
	//wmic path win32_diskdrive where deviceid='\\\\.\\PHYSICALDRIVE0' get serialnumber
	param := "deviceid='\\\\\\\\.\\\\PHYSICALDRIVE" + disk + "'"
	cmd := exec.Command("wmic", "path", "win32_diskdrive", "where", param, "get", "caption,serialnumber")
	/*
		additionalEnv := param
		newEnv := append(os.Environ(), additionalEnv)
		cmd.Env = newEnv
	*/
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	stringInfo := fmt.Sprintf("%s\n", stdoutStderr)
	stringInfo = strings.TrimSpace(strings.Replace(stringInfo, "CaptionSerialNumber", "", -1))

	arrString := strings.Fields(stringInfo)
	//fmt.Println(arrString)

	serialNumber := ""
	for _, val := range arrString {
		value := strings.Replace(val, "\n", "", 1)
		//	fmt.Println("[" + value + "]")
		if strings.TrimSpace(value) != "Caption" && strings.TrimSpace(value) != "SerialNumber" {
			if serialNumber != "" {
				serialNumber += "_"
			}
			serialNumber += value
		}
	}
	return serialNumber
}
