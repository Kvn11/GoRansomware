# GoRansomware
Ransomware written in Go for Windows. Done as a learning project, no harm is intended. I am not responsible if you download and break your computer.

This project is nowhere near complete yet, only the encryption methods are done. Still need to add monero address generation, as well as the features listed below.

I am a noob who is just trying to learn more about malware, and since there were no solid guides online, I am making this to hopefully help others who may try to write their own. My plan is to hardcode a public RSA key that will be used to encrypt an AES key. The unencrypted AES key is what will be used to encrypt files. The AES key is randomly generated the first time the program is run, and will remain unencrypted until the last file has been encrypted. At that point, the only way to unencrypt the key, and thus unencrypt the filesystem, is to pay the ransom.

## Features I plan to implement:
- A GUI to distract the user while encryption takes place
- Threads, so that encryption occurs faster
- Anti-Debug and Obfuscation
- AV Evasion