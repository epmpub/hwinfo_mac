package main

import (
	"encoding/json"
	"fmt"
	"os/exec"
)

type HWINFO struct {
	SPNVMeDataType []struct {
		Items []struct {
			Name              string `json:"_name"`
			BsdName           string `json:"bsd_name"`
			DetachableDrive   string `json:"detachable_drive"`
			DeviceModel       string `json:"device_model"`
			DeviceRevision    string `json:"device_revision"`
			DeviceSerial      string `json:"device_serial"`
			PartitionMapType  string `json:"partition_map_type"`
			RemovableMedia    string `json:"removable_media"`
			Size              string `json:"size"`
			SizeInBytes       int64  `json:"size_in_bytes"`
			SmartStatus       string `json:"smart_status"`
			SpnvmeLinkspeed   string `json:"spnvme_linkspeed"`
			SpnvmeLinkwidth   string `json:"spnvme_linkwidth"`
			SpnvmeTrimSupport string `json:"spnvme_trim_support"`
			Volumes           []struct {
				Name             string `json:"_name"`
				BsdName          string `json:"bsd_name"`
				FileSystem       string `json:"file_system,omitempty"`
				Iocontent        string `json:"iocontent"`
				Size             string `json:"size"`
				SizeInBytes      int    `json:"size_in_bytes"`
				VolumeUUID       string `json:"volume_uuid,omitempty"`
				FreeSpace        string `json:"free_space,omitempty"`
				FreeSpaceInBytes int64  `json:"free_space_in_bytes,omitempty"`
				MountPoint       string `json:"mount_point,omitempty"`
				Writable         string `json:"writable,omitempty"`
			} `json:"volumes"`
		} `json:"_items"`
		Name string `json:"_name"`
	} `json:"SPNVMeDataType"`
}

func main() {
	//system_profiler -json SPNVMeDataType
	out, err := exec.Command("system_profiler", "-json", "SPNVMeDataType").Output()
	if err != nil {
		return
	}

	data := HWINFO{}

	err = json.Unmarshal([]byte(out), &data)
	if err != nil {
		return
	}

	//  useful website , find you want information JSON PATH and print
	//	JSON to Struct:			https://jsonpathfinder.com/
	// 	JSON PATH finder:		https://transform.tools/json-to-go
	//  Author : Andy Hu
	//  Date: 2023 / 2 / 15 07:59
	//  email:327656021@qq.com

	// some test case , use JSON path print
	fmt.Printf("%T\n", data)
	fmt.Println(data.SPNVMeDataType[0].Items[0].DeviceModel)

}
