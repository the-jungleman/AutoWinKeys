package	main
import(
	"log"
	"os/exec"
	"./winStatus"
)

func main(){
	isWinActivated()
	switch{
	case winActivated == false:
		selectKey()
		
	case winActivated == true:
	}
	
}

func Cmd(command string){
	cmd := exec.Command(command)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

func winActivation(){
	homeSingleLangKeys := [3]string{"NTRHT-XTHTG-GBWCG-4MTMP-HH64C", "7HNRX-D7KGG-3K4RQ-4WPJ4-YTDFH", "BT79Q-G7N6G-PGBYW-4YWX6-6F4BT"}
	homeKeys := [3]string{"TX9XD-98N7V-6WMQ6-BX7FG-H8Q99", "TX9XD-98N7V-6WMQ6-BX7FG-H8Q99", "YTMG3-N6DKC-DKB77-7M9GH-8HVX7"}
	proKeys := [3]string{"VK7JG-NPHTM-C97JM-9MPGT-3V66T", "W269N-WFGWX-YVC9B-4J6C9-T83GX", "NRG8B-VKK3Q-CXVCJ-9G2XF-6Q84J"}
	workKeys := [3]string{"NRG8B-VKK3Q-CXVCJ-9G2XF-6Q84J", "NRG8B-VKK3Q-CXVCJ-9G2XF-6Q84J", "9FNHH-K3HBT-3W4TD-6383H-6XYWF"}
	educationKeys := [2]string{"NW6C2-QMPVW-D7KKK-3GKT6-VCFB2", "2WH4N-8QGBV-H22JP-CT43Q-MDWWJ"}
	proEducationKeys := [2]string{"6TP4R-GNPTD-KYYHQ-7B7DP-J447Y", "YVWGF-BXNMC-HTQYQ-CPQ99-66QFC"}
	
	Cmd("Slui")
	winKey := selectKey()
	Cmd("slmgr.vbs /ipk %v", winKey)
	Cmd("slmgr.vbs /skms kms.lotro.cc")
	Cmd("slmgr.vbs /ato")
	//Cmd("slmgr.vbs /dli")
}

func selectKey()string{
	ascii := `

	 /$$$$$$              /$$               /$$      /$$ /$$           /$$   /$$                              
	/$$__  $$            | $$              | $$  /$ | $$|__/          | $$  /$$/                              
   | $$  \ $$ /$$   /$$ /$$$$$$    /$$$$$$ | $$ /$$$| $$ /$$ /$$$$$$$ | $$ /$$/   /$$$$$$  /$$   /$$  /$$$$$$$
   | $$$$$$$$| $$  | $$|_  $$_/   /$$__  $$| $$/$$ $$ $$| $$| $$__  $$| $$$$$/   /$$__  $$| $$  | $$ /$$_____/
   | $$__  $$| $$  | $$  | $$    | $$  \ $$| $$$$_  $$$$| $$| $$  \ $$| $$  $$  | $$$$$$$$| $$  | $$|  $$$$$$ 
   | $$  | $$| $$  | $$  | $$ /$$| $$  | $$| $$$/ \  $$$| $$| $$  | $$| $$\  $$ | $$_____/| $$  | $$ \____  $$
   | $$  | $$|  $$$$$$/  |  $$$$/|  $$$$$$/| $$/   \  $$| $$| $$  | $$| $$ \  $$|  $$$$$$$|  $$$$$$$ /$$$$$$$/
   |__/  |__/ \______/    \___/   \______/ |__/     \__/|__/|__/  |__/|__/  \__/ \_______/ \____  $$|_______/ 
																						   /$$  | $$          
																						  |  $$$$$$/          
																						   \______/           
	`

	var version int
	home := `

			SELECT YOUR WINDOWS 10 VERSION

			[1] Home Single Language
			[2]	Home
			[3] Pro
			[4] Workstations
			[5] Education
			[6] Pro Education

	`
	fmt.Println("%v %v	>>", ascii, home)
	fmt.Scanln(&version)

	var key string
	switch{
	case version == 1:

		Cmd("slmgr.vbs /ipk %v", keys)
		Cmd("slmgr.vbs /skms kms.lotro.cc")
		Cmd("slmgr.vbs /ato")
	}
	return key
}


