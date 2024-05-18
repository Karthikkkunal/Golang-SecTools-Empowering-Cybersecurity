package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func checkSources() bool {
	sources, err := os.Open("/etc/apt/sources.list")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer sources.Close()

	scanner := bufio.NewScanner(sources)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, "kali-rolling") {
			return true
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
		return false
	}

	return false
}

func addKaliRepository() {
	if checkSources() {
		fmt.Println("Kali sources already exist. Run the script again and install the tools")
	} else {
		fmt.Println("Adding Kali sources")
		sources, err := os.OpenFile("/etc/apt/sources.list", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer sources.Close()

		_, err = sources.WriteString("deb http://http.kali.org/kali kali-rolling main non-free contrib\n")
		if err != nil {
			fmt.Println(err)
			return
		}

		cmd := exec.Command("apt-key", "adv", "--keyserver", "keyserver.ubuntu.com", "--recv-keys", "ED444FF07D8D0BF6")
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}

		cmd = exec.Command("apt", "update")
		err = cmd.Run()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Kali sources added")
	}
}

func installCommand(category string, tool int) {

	switch category {
	case "InformationGathering":
		cmd := exec.Command("apt", "install", "-y", InformationGathering[tool-1])
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "VulnerabilityAnalysis":
		cmd := exec.Command("apt", "install", "-y", VulnerabilityAnalysis[tool-1])
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "PasswordAttacks":
		cmd := exec.Command("apt", "install", "-y", PasswordAttacks[tool-1])
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "Forensics":
		cmd := exec.Command("apt", "install", "-y", Forensics[tool-1])
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "WebApplicationAnalysis":
		cmd := exec.Command("apt", "install", "-y", WebApplicationAnalysis[tool-1])
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "WirelessAttacks":
		cmd := exec.Command("apt", "install", "-y", WirelessAttacks[tool-1])
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "ReverseEngineering":
		cmd := exec.Command("apt", "install", "-y", ReverseEngineering[tool-1])
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "SniffingSpoofing":
		cmd := exec.Command("apt", "install", "-y", SniffingSpoofing[tool-1])
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}
}

func installAll(category string) {
	fmt.Println("Installing all tools")
	for i := 1; i <= 58; i++ {
		installCommand(category, i)
	}
	fmt.Printf("All tools of %s installed\n", category)
}

func installTools() {
	fmt.Println("Select Category:")
	fmt.Println("1 - Information Gathering")
	fmt.Println("2 - Vulnerability Analysis")
	fmt.Println("3 - Password Attacks")
	fmt.Println("4 - Database Assessment")
	fmt.Println("5 - Web Applications Analysis")
	fmt.Println("6 - Wireless Attacks")
	fmt.Println("7 - Reverse Engineering")
	fmt.Println("8 - Exploitation Tools")
	fmt.Println("9 - Sniffing & Spoofing")
	fmt.Println("10 - Post Exploitation")
	fmt.Println("11 - Forensics")
	fmt.Println("12 - Social Engineering Tools")
	fmt.Println("13 - Most Used Tools")

	var option int
	fmt.Print("Enter your option: ")
	_, err := fmt.Scan(&option)
	if err != nil {
		fmt.Println(err)
		return
	}

	var category string
	switch option {
	case 1:
		category = "InformationGathering"
		fmt.Println("1) acccheck")
		fmt.Println("2) amap")
		fmt.Println("3) nmap")
		fmt.Println("4) dnsenum")
		fmt.Println("5) dnsmap")
		fmt.Println("6) dnsrecon")
		fmt.Println("7) dnstracer")
		fmt.Println("8) dnswalk")
		fmt.Println("9) enum4linux")
		fmt.Println("10) fierce")
		fmt.Println("11) fragroute")
		fmt.Println("12) fragrouter")
		fmt.Println("13) goofile")
		fmt.Println("14) lbd")
		fmt.Println("15) masscan")
		fmt.Println("16) metagoofil")
		fmt.Println("17) ntop")
		fmt.Println("18) p0f")
		fmt.Println("19) recon-ng")
		fmt.Println("20) set")
		fmt.Println("21) smtp-user-enum")
		fmt.Println("22) snmpcheck")
		fmt.Println("23) sslcaudit")
		fmt.Println("24) sslsplit")
		fmt.Println("25) sslstrip")
		fmt.Println("26) sslyze")
		fmt.Println("27) theharvester")
		fmt.Println("28) urlcrazy")
		fmt.Println("29) wireshark")
		fmt.Println("30) xplico")
		fmt.Println("31) intrace")
		fmt.Println("32) hping3")
		fmt.Println("0) INSTALL ALL TOOLS")

		var tool int
		fmt.Print("Enter a option: ")
		_, err := fmt.Scan(&tool)
		if err != nil {
			fmt.Println(err)
			return
		}

		if tool != 0 {
			installCommand(category, tool)
		} else {
			installAll(category)
		}
	case 2:
		category = "VulnerabilityAnalysis"
		fmt.Println("1) cisco-auditing-tool")
		fmt.Println("2) commix")
		fmt.Println("3) lynis")
		fmt.Println("4) nmap")
		fmt.Println("5) sqlmap")
		fmt.Println("6) nikto")
		fmt.Println("7) unix-privesc-check")
		fmt.Println("8) legion")
		fmt.Println("9) spike")
		fmt.Println("0) INSTALL ALL TOOLS")

		var tool int
		fmt.Print("Enter a option: ")
		_, err := fmt.Scan(&tool)
		if err != nil {
			fmt.Println(err)
			return
		}

		if tool != 0 {
			installCommand(category, tool)
		} else {
			installAll(category)
		}
	case 3:
		category = "PasswordAttacks"
		fmt.Println("1) cewl")
		fmt.Println("2) cisco-auditing-tool")
		fmt.Println("3) chntpw")
		fmt.Println("4) cmospwd")
		fmt.Println("5) creddump")
		fmt.Println("6) crunch")
		fmt.Println("7) findmyhash")
		fmt.Println("8) impacket")
		fmt.Println("9) hydra")
		fmt.Println("10) john")
		fmt.Println("11) johnny")
		fmt.Println("12) ncrack")
		fmt.Println("13) patator")
		fmt.Println("14) rainbowcrack")
		fmt.Println("15) rsmangler")
		fmt.Println("16) wordlists")
		fmt.Println("17) zaproxy")
		fmt.Println("0) INSTALL ALL TOOLS")

		var tool int
		fmt.Print("Enter a option: ")
		_, err := fmt.Scan(&tool)
		if err != nil {
			fmt.Println(err)
			return
		}

		if tool != 0 {
			installCommand(category, tool)
		} else {
			installAll(category)
		}
	case 4:
		category = "DatabaseAssessment"
		fmt.Println("1) sqlite")
		fmt.Println("2) sqlmap")

		var tool int
		fmt.Print("Enter a option: ")
		_, err := fmt.Scan(&tool)
		if err != nil {
			fmt.Println(err)
			return
		}

		if tool != 0 {
			if tool == 1 {
				cmd := exec.Command("apt", "install", "-y", "sqlite")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else if tool == 2 {
				cmd := exec.Command("apt", "install", "-y", "sqlmap")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else {
				fmt.Println("Invalid option")
			}
		} else {
			cmd := exec.Command("apt", "install", "-y", "sqlite", "sqlmap")
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("All tools of %s installed\n", category)
			}
		}
	case 5:
		category = "WebApplicationsAnalysis"
		fmt.Println("1) burpsuite")
		fmt.Println("2) wpscan")
		fmt.Println("3) dirbuster")
		fmt.Println("4) dirb")
		fmt.Println("5) gobuster")
		fmt.Println("6) cadaver")
		fmt.Println("7) nikto")
		fmt.Println("8) davtest")
		fmt.Println("9) wafw00f")
		fmt.Println("10) wapiti")
		fmt.Println("11) whatweb")
		fmt.Println("12) commix")
		fmt.Println("0) INSTALL ALL TOOLS")

		var tool int
		fmt.Print("Enter a option: ")
		_, err := fmt.Scan(&tool)
		if err != nil {
			fmt.Println(err)
			return
		}

		if tool != 0 {
			installCommand(category, tool)
		} else {
			installAll(category)
		}
	case 6:
		category = "WirelessAttacks"
		fmt.Println("1) bully")
		fmt.Println("2) fern-wifi-cracker")
		fmt.Println("3) spooftooph")
		fmt.Println("4) aircrack-ng")
		fmt.Println("5) kismet")
		fmt.Println("6) wifite")
		fmt.Println("7) pixiewps")
		fmt.Println("8) reaver")
		fmt.Println("0) INSTALL ALL TOOLS")

		var tool int
		fmt.Print("Enter a option: ")
		_, err := fmt.Scan(&tool)
		if err != nil {
			fmt.Println(err)
			return
		}

		if tool != 0 {
			installCommand(category, tool)
		} else {
			installAll(category)
		}
	case 7:
		category = "ReverseEngineering"
		fmt.Println("1) clang")
		fmt.Println("2) clang++")
		fmt.Println("3) ghidra")
		fmt.Println("4) NASM shell")
		fmt.Println("5) radare2")
		fmt.Println("0) INSTALL ALL TOOLS")

		var tool int
		fmt.Print("Enter a option: ")
		_, err := fmt.Scan(&tool)
		if err != nil {
			fmt.Println(err)
			return
		}

		if tool != 0 {
			installCommand(category, tool)
		} else {
			installAll(category)
		}
	case 8:
		category = "ExploitTools"
		fmt.Println("1) exploitdb")
		fmt.Println("2) metasploit-framework")
		fmt.Println("3) searchsploit")
		fmt.Println("4) sqlmap")
		fmt.Println("5) crackmapexec")
		fmt.Println("6) social engineering toolkit")
		fmt.Println("0) INSTALL ALL TOOLS")

		var tool int
		fmt.Print("Enter a option: ")
		_, err := fmt.Scan(&tool)
		if err != nil {
			fmt.Println(err)
			return
		}

		if tool != 0 {
			if tool == 1 {
				cmd := exec.Command("apt", "install", "-y", "exploitdb")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else if tool == 2 {
				cmd := exec.Command("apt", "install", "-y", "metasploit-framework")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else if tool == 3 {
				cmd := exec.Command("apt", "install", "-y", "searchsploit")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else if tool == 4 {
				cmd := exec.Command("apt", "install", "-y", "sqlmap")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else if tool == 5 {
				cmd := exec.Command("apt", "install", "-y", "crackmapexec")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else if tool == 6 {
				cmd := exec.Command("git", "clone", "https://github.com/trustedsec/social-engineer-toolkit", "setoolkit/")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					cmd = exec.Command("cd", "setoolkit", "&&", "python3", "setup.py")
					err = cmd.Run()
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("Selected tool installed")
					}
				}
			} else {
				fmt.Println("Invalid option")
			}
		} else {
			cmd := exec.Command("apt", "install", "-y", "exploitdb", "metasploit-framework", "searchsploit", "sqlmap", "crackmapexec")
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			} else {
				cmd := exec.Command("git", "clone", "https://github.com/trustedsec/social-engineer-toolkit", "setoolkit/")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					cmd = exec.Command("cd", "setoolkit", "&&", "python3", "setup.py")
					err = cmd.Run()
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("All tools of %s installed\n", category)
					}
				}
			}
		}
	case 9:
		category = "SniffingSpoofing"
		fmt.Println("1) dnschef")
		fmt.Println("2) dsniff")
		fmt.Println("3) netsniff-ng")
		fmt.Println("4) rebind")
		fmt.Println("5) sslsplit")
		fmt.Println("6) tcpreplay")
		fmt.Println("7) ettercap")
		fmt.Println("8) macchanger")
		fmt.Println("9) minicom")
		fmt.Println("10) mitmproxy")
		fmt.Println("11) responder")
		fmt.Println("12) scapy")
		fmt.Println("13) tcpdump")
		fmt.Println("14) wireshark")

		var tool int
		fmt.Print("Enter a option: ")
		_, err := fmt.Scan(&tool)
		if err != nil {
			fmt.Println(err)
			return
		}

		if tool != 0 {
			installCommand(category, tool)
		} else {
			installAll(category)
		}
	case 10:
		category = "PostExploitation"
		fmt.Println("1) dbd")
		fmt.Println("2) powersploit")
		fmt.Println("3) sbd")
		fmt.Println("4) proxychains4")
		fmt.Println("5) weevely")
		fmt.Println("0) INSTALL ALL TOOLS")

		var tool int
		fmt.Print("Enter a option: ")
		_, err := fmt.Scan(&tool)
		if err != nil {
			fmt.Println(err)
			return
		}

		if tool != 0 {
			if tool == 1 {
				cmd := exec.Command("apt", "install", "-y", "dbd")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else if tool == 2 {
				cmd := exec.Command("apt", "install", "-y", "powersploit")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else if tool == 3 {
				cmd := exec.Command("apt", "install", "-y", "sbd")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else if tool == 4 {
				cmd := exec.Command("apt", "install", "-y", "proxychains4")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else if tool == 5 {
				cmd := exec.Command("apt", "install", "-y", "weevely")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else {
				fmt.Println("Invalid option")
			}
		} else {
			cmd := exec.Command("apt", "install", "-y", "dbd", "powersploit", "sbd", "proxychains4", "weevely")
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf("All tools of %s installed\n", category)
			}
		}
	case 11:
		category = "Forensics"
		fmt.Println("1) autopsy")
		fmt.Println("2) scalpel")
		fmt.Println("3) scrounge-ntfs")
		fmt.Println("4) guymager")
		fmt.Println("5) pdf-parse")
		fmt.Println("6) binwalk")
		fmt.Println("7) sleuthkit")
		fmt.Println("8) bulk-extractor")
		fmt.Println("0) INSTALL ALL TOOLS")

		var tool int
		fmt.Print("Enter a option: ")
		_, err := fmt.Scan(&tool)
		if err != nil {
			fmt.Println(err)
			return
		}

		if tool != 0 {
			installCommand(category, tool)
		} else {
			installAll(category)
		}
	case 12:
		category = "SocialEngineeringTools"
		fmt.Println("1) beef-xss")
		fmt.Println("2) maltego")
		fmt.Println("3) sherlock")
		fmt.Println("4) social engineering tool kit")
		fmt.Println("0) INSTALL ALL TOOLS")

		var tool int
		fmt.Print("Enter a option: ")
		_, err := fmt.Scan(&tool)
		if err != nil {
			fmt.Println(err)
			return
		}

		if tool != 0 {
			if tool == 1 {
				cmd := exec.Command("apt", "install", "-y", "beef-xss")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else if tool == 2 {
				cmd := exec.Command("apt", "install", "-y", "maltego")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else if tool == 3 {
				cmd := exec.Command("apt", "install", "-y", "sherlock")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					fmt.Println("Selected tool installed")
				}
			} else if tool == 4 {
				cmd := exec.Command("git", "clone", "https://github.com/trustedsec/social-engineer-toolkit", "setoolkit/")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					cmd = exec.Command("cd", "setoolkit", "&&", "python3", "setup.py")
					err = cmd.Run()
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Println("Selected tool installed")
					}
				}
			} else {
				fmt.Println("Invalid option")
			}
		} else {
			cmd := exec.Command("apt", "install", "-y", "beef-xss", "maltego", "sherlock")
			err := cmd.Run()
			if err != nil {
				fmt.Println(err)
			} else {
				cmd := exec.Command("git", "clone", "https://github.com/trustedsec/social-engineer-toolkit", "setoolkit/")
				err := cmd.Run()
				if err != nil {
					fmt.Println(err)
				} else {
					cmd = exec.Command("cd", "setoolkit", "&&", "python3", "setup.py")
					err = cmd.Run()
					if err != nil {
						fmt.Println(err)
					} else {
						fmt.Printf("All tools of %s installed\n", category)
					}
				}
			}
		}
	case 13:
		cmd := exec.Command("apt", "install", "-y", "nmap", "metasploit-framework", "sqlmap", "burpsuite", "wpscan", "dirb", "gobuster", "ghidra", "wireshark", "macchanger", "minicom", "mitmproxy", "responder", "scapy", "tcpdump", "nikto", "enum4linux", "hydra", "john", "hashcat", "autopsy", "recon-ng", "netdiscover", "theharvester", "aircrack-ng", "wifite", "proxychains4", "cewl", "crunch", "exploitdb", "amass", "beef-xss")
		err := cmd.Run()
		if err != nil {
			fmt.Println(err)
		}
	}

}

func main() {

	if os.Geteuid() != 0 {
		fmt.Println("RUN THE PROGRAM AS ROOT!!")
		return
	}

	fmt.Println("     ____  _____ ____ _____ ___   ___  _     ____  ")
	fmt.Println("    / ___|| ____/ ___|_   _/ _ \\ / _ \\| |   / ___| ")
	fmt.Println("    \\___ \\|  _|| |     | || | | | | | | |   \\___ \\ ")
	fmt.Println("     ___) | |__| |___  | || |_| | |_| | |___ ___) |")
	fmt.Println("    |____/|_____\\____| |_| \\___/ \\___/|_____|____/  ")
	fmt.Println("                                                    - Karthikkkunal")

	fmt.Println(" Select an option: \n 1. Add Kali repository \n 2. Install Tools \n 0. Exit")

	var option int
	fmt.Print("Enter your option: ")
	_, err := fmt.Scan(&option)
	if err != nil {
		fmt.Println(err)
		return
	}

	if option == 1 {
		addKaliRepository()
	} else if option == 2 {
		installTools()
	} else if option == 0 {
		fmt.Println("Exiting...")
		return
	} else {
	}
}
