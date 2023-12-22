package main

import (
	"net/http"
	"fmt"
	"log"
	"github.com/gin-gonic/gin"
	"encoding/json"
	"encoding/base64"
	"io/ioutil"
	"time"
	"strings"
	) 

	type GetVM struct {
		Boot struct {
			Delay           int    `json:"delay"`
			EfiLegacyBoot   bool   `json:"efi_legacy_boot"`
			EnterSetupMode  bool   `json:"enter_setup_mode"`
			NetworkProtocol string `json:"network_protocol"`
			Retry           bool   `json:"retry"`
			RetryDelay      int    `json:"retry_delay"`
			Type            string `json:"type"`
		} `json:"boot"`
		BootDevices []struct {
			Disks []string `json:"disks"`
			Nic   string   `json:"nic"`
			Type  string   `json:"type"`
		} `json:"boot_devices"`
		Cdroms struct {
			Key struct {
				AllowGuestControl bool `json:"allow_guest_control"`
				Backing           struct {
					AutoDetect       bool   `json:"auto_detect"`
					DeviceAccessType string `json:"device_access_type"`
					HostDevice       string `json:"host_device"`
					IsoFile          string `json:"iso_file"`
					Type             string `json:"type"`
				} `json:"backing"`
				Ide struct {
					Master  bool `json:"master"`
					Primary bool `json:"primary"`
				} `json:"ide"`
				Label string `json:"label"`
				Sata  struct {
					Bus  int `json:"bus"`
					Unit int `json:"unit"`
				} `json:"sata"`
				StartConnected bool   `json:"start_connected"`
				State          string `json:"state"`
				Type           string `json:"type"`
			} `json:"key"`
		} `json:"cdroms"`
		CPU struct {
			CoresPerSocket   int  `json:"cores_per_socket"`
			Count            int  `json:"count"`
			HotAddEnabled    bool `json:"hot_add_enabled"`
			HotRemoveEnabled bool `json:"hot_remove_enabled"`
		} `json:"cpu"`
		Disks struct {
			Key struct {
				Backing struct {
					Type     string `json:"type"`
					VmdkFile string `json:"vmdk_file"`
				} `json:"backing"`
				Capacity int `json:"capacity"`
				Ide      struct {
					Master  bool `json:"master"`
					Primary bool `json:"primary"`
				} `json:"ide"`
				Label string `json:"label"`
				Nvme  struct {
					Bus  int `json:"bus"`
					Unit int `json:"unit"`
				} `json:"nvme"`
				Sata struct {
					Bus  int `json:"bus"`
					Unit int `json:"unit"`
				} `json:"sata"`
				Scsi struct {
					Bus  int `json:"bus"`
					Unit int `json:"unit"`
				} `json:"scsi"`
				Type string `json:"type"`
			} `json:"key"`
		} `json:"disks"`
		Floppies struct {
			Key struct {
				AllowGuestControl bool `json:"allow_guest_control"`
				Backing           struct {
					AutoDetect bool   `json:"auto_detect"`
					HostDevice string `json:"host_device"`
					ImageFile  string `json:"image_file"`
					Type       string `json:"type"`
				} `json:"backing"`
				Label          string `json:"label"`
				StartConnected bool   `json:"start_connected"`
				State          string `json:"state"`
			} `json:"key"`
		} `json:"floppies"`
		GuestOS  string `json:"guest_OS"`
		Hardware struct {
			UpgradeError struct {
			} `json:"upgrade_error"`
			UpgradePolicy  string `json:"upgrade_policy"`
			UpgradeStatus  string `json:"upgrade_status"`
			UpgradeVersion string `json:"upgrade_version"`
			Version        string `json:"version"`
		} `json:"hardware"`
		Identity struct {
			BiosUUID     string `json:"bios_uuid"`
			InstanceUUID string `json:"instance_uuid"`
			Name         string `json:"name"`
		} `json:"identity"`
		InstantCloneFrozen bool `json:"instant_clone_frozen"`
		Memory             struct {
			HotAddEnabled          bool `json:"hot_add_enabled"`
			HotAddIncrementSizeMiB int  `json:"hot_add_increment_size_MiB"`
			HotAddLimitMiB         int  `json:"hot_add_limit_MiB"`
			SizeMiB                int  `json:"size_MiB"`
		} `json:"memory"`
		Name string `json:"name"`
		Nics struct {
			Key struct {
				AllowGuestControl bool `json:"allow_guest_control"`
				Backing           struct {
					ConnectionCookie      int    `json:"connection_cookie"`
					DistributedPort       string `json:"distributed_port"`
					DistributedSwitchUUID string `json:"distributed_switch_uuid"`
					HostDevice            string `json:"host_device"`
					Network               string `json:"network"`
					NetworkName           string `json:"network_name"`
					OpaqueNetworkID       string `json:"opaque_network_id"`
					OpaqueNetworkType     string `json:"opaque_network_type"`
					Type                  string `json:"type"`
				} `json:"backing"`
				Label                   string `json:"label"`
				MacAddress              string `json:"mac_address"`
				MacType                 string `json:"mac_type"`
				PciSlotNumber           int    `json:"pci_slot_number"`
				StartConnected          bool   `json:"start_connected"`
				State                   string `json:"state"`
				Type                    string `json:"type"`
				UptCompatibilityEnabled bool   `json:"upt_compatibility_enabled"`
				WakeOnLanEnabled        bool   `json:"wake_on_lan_enabled"`
			} `json:"key"`
		} `json:"nics"`
		NvmeAdapters struct {
			Key struct {
				Bus           int    `json:"bus"`
				Label         string `json:"label"`
				PciSlotNumber int    `json:"pci_slot_number"`
			} `json:"key"`
		} `json:"nvme_adapters"`
		ParallelPorts struct {
			Key struct {
				AllowGuestControl bool `json:"allow_guest_control"`
				Backing           struct {
					AutoDetect bool   `json:"auto_detect"`
					File       string `json:"file"`
					HostDevice string `json:"host_device"`
					Type       string `json:"type"`
				} `json:"backing"`
				Label          string `json:"label"`
				StartConnected bool   `json:"start_connected"`
				State          string `json:"state"`
			} `json:"key"`
		} `json:"parallel_ports"`
		PowerState   string `json:"power_state"`
		SataAdapters struct {
			Key struct {
				Bus           int    `json:"bus"`
				Label         string `json:"label"`
				PciSlotNumber int    `json:"pci_slot_number"`
				Type          string `json:"type"`
			} `json:"key"`
		} `json:"sata_adapters"`
		ScsiAdapters struct {
			Key struct {
				Label         string `json:"label"`
				PciSlotNumber int    `json:"pci_slot_number"`
				Scsi          struct {
					Bus  int `json:"bus"`
					Unit int `json:"unit"`
				} `json:"scsi"`
				Sharing string `json:"sharing"`
				Type    string `json:"type"`
			} `json:"key"`
		} `json:"scsi_adapters"`
		SerialPorts struct {
			Key struct {
				AllowGuestControl bool `json:"allow_guest_control"`
				Backing           struct {
					AutoDetect      bool   `json:"auto_detect"`
					File            string `json:"file"`
					HostDevice      string `json:"host_device"`
					NetworkLocation string `json:"network_location"`
					NoRxLoss        bool   `json:"no_rx_loss"`
					Pipe            string `json:"pipe"`
					Proxy           string `json:"proxy"`
					Type            string `json:"type"`
				} `json:"backing"`
				Label          string `json:"label"`
				StartConnected bool   `json:"start_connected"`
				State          string `json:"state"`
				YieldOnPoll    bool   `json:"yield_on_poll"`
			} `json:"key"`
		} `json:"serial_ports"`
	}

func main (){
	router := gin.Default()
	vmName := "hello"
	vcenterHost:= "http://vcenter.easlab.co.uk/"
	vcenterUsername:="hello"
	vcenterPassword:= "hello"

	//Repo Exists
	router.GET("/api/vm/exists", func (c *gin.Context) {
		resp, err := http.Get(vcenterHost + "/api/vcenter/vm/" + vmName)
		if err != nil {
			log.Fatalln(err)
		}
		
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		
		var getVMStruct GetVM
		if err := json.Unmarshal(body, &getVMStruct); err != nil {   
    	fmt.Println("Can not unmarshal JSON")
		}

		if (getVMStruct.Name == "-- SOME VALUE HERE?? --"){
			fmt.Println("Good")
        } else {
			fmt.Println("Bad")
        }
	})
	
//Repo Contains
	router.GET("/api/repository/contains", func (c *gin.Context) {
		//http req
		resp, err := http.Get("https://api.github.com/repos/EAS-" + user + "/" + repoName + "/contents")
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(body)
		var result Response
		if err := json.Unmarshal(body, &result); err != nil {   
    	fmt.Println("Can not unmarshal JSON")
		}
		for _, rec := range result {
			fmt.Println(rec.Name)
			if (rec.Name == fileName) {
				fmt.Println("Good")
			} else {
				fmt.Println("Bad")
			}
		}
	})

// File Contains
	router.GET("/api/repository/file/contains", func (c *gin.Context) {
		resp, err := http.Get("https://api.github.com/repos/EAS-" + user + "/" + repoName + "/contents/" + fileName)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		fmt.Println(string(body))
		
		var result3 Response3
		if err := json.Unmarshal(body, &result3); err != nil {   
    	fmt.Println("Can not unmarshal JSON")
		}
		

		decoded, err := base64.StdEncoding.DecodeString(result3.Content)
		ans := strings.Contains(string(decoded), check )
		fmt.Println("hello", ans)
    	if err != nil {
        	fmt.Println("Error decoding string:", err)
		} else {
			fmt.Println("hello")
		}
		fmt.Sprintf("%T", check)
		fmt.Println(check)
    	fmt.Println(string(decoded))
	})
	router.Run(":3000")
}


// func PrettyPrint(i interface{}) string {
//     s, _ := json.MarshalIndent(i, "", "\t")
//     return string(s)
// }

