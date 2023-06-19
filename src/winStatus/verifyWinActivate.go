package winStatus

import(
	"fmt",
	"github.com/StackExchange/wmi"
)

type Win32_OperatingSystem struct {
	OperatingSystemSKU uint32
}

func isWinActivated() {
	winActivated := false
	var os Win32_OperatingSystem
	err := wmi.Query("SELECT OperatingSystemSKU FROM Win32_OperatingSystem", &os)
	if err != nil {
		fmt.Println("Erro ao obter informações do sistema operacional:", err)
		return
	}

	if os.OperatingSystemSKU == 5 || os.OperatingSystemSKU == 6 || os.OperatingSystemSKU == 10 {
		fmt.Println("Windows is activated")
		winActivated := true
	} else {
		fmt.Println("Windows is not activated")
		winActivated := false
	}
}
