# GoRansomware
Ransomware written in Go for Windows. Done as a learning project, no harm is intended. I am not responsible if you download and break your computer.

I am a cybersecurity enthusiast who is just trying to learn more about malware, and since there were no solid guides online, I am making this to hopefully help others who may try to write their own. My plan is to hardcode a public RSA key that will be used to encrypt an AES key. The unencrypted AES key is what will be used to encrypt files. The AES key is randomly generated the first time the program is run, and will remain unencrypted until the last file has been encrypted. At that point, the only way to unencrypt the key, and thus unencrypt the filesystem, is to pay the ransom.

Most of the frontend is done, but there is a slight bug where the KeyID doesn't show on the GUI window. The encryption features are all done. I still need to test whether inputting the base64 decrypted key will actually decrypt the file system.

## Features I plan to implement:
- Threading, so that encryption occurs faster
- Anti-Debug and Obfuscation
- AV Evasion
