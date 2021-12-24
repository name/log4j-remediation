sakuraji_log4j

This tool is used to discover and remedidate the Log4Shell vulnerability [CVE-2021-45105] by removing the 'JndiLookup.class' file from '.ear|.jar|.war' files on Windows, Linux & Mac OS.

Usage:
    Scan OS for Log4Shell vulnerability.
    - sakuraji_log4j discover
    - Example: https://i.imgur.com/VgXCrH1.png

    Scan OS and remediate Log4Shell vulnerability.
    - sakuraji_log4j remediate
    - Example: https://i.imgur.com/JijGWj7.png

Build:
    Clone & go build .