// from git.tcp.direct/kayos/CokePlate
package extra

import (
	"bufio"
	"encoding/base64"
	"fmt"
	"strings"
)

const banner = "ChtbOTc7NDBtIBtbOTc7NDNt4paEG1szMzs0MG3ilojilpIbWzk3OzQwbSAbWzMzOzQwbeKWiOKWiBtbMzc7NDBtIBtbMG0bWzBtICAbWzMzOzQwbeKWkxtbOTc7NDNtIOKWhOKWhBtbMzM7NDBt4paI4paIG1szNzs0MG0gG1swbRtbMG0gIBtbOTc7NDBtIBtbMzM7NDBt4paI4paI4paTG1szNzs0MG0gICAgG1swbRtbMG0gIBtbOTc7NDBtIBtbMzM7NDBt4paI4paI4paTG1szNzs0MG0gICAgG1swbRtbMG0gIBtbOTc7NDBtIBtbMzM7NDBt4paI4paI4paS4paI4paI4paIG1szNzs0MG0gIBtbMG0bWzBtICAbWzk3OzQwbSAbWzMzOzQwbeKWk+KWiBtbOTc7NDNt4paE4paEG1szMzs0MG3ilojilogbWzM3OzQwbSAgG1swbRtbMG0gIBtbMzM7NDBt4paE4paE4paE4paIG1s5Nzs0M23iloQbWzMzOzQwbeKWiOKWiOKWiOKWkxtbMG0bWzBtICAbWzBtChtbMzM7NDBt4paTG1s5Nzs0M23ilogbWzMzOzQwbeKWiOKWkRtbOTc7NDBtIBtbMzM7NDBt4paIG1s5Nzs0M23iloQbWzMzOzQwbeKWkhtbMG0bWzBtICAbWzMzOzQwbeKWkxtbOTc7NDNt4paIIBtbOTc7NDBtICAbWzMzOzQwbeKWgBtbMzc7NDBtIBtbMG0bWzBtICAbWzMzOzQwbeKWk+KWiOKWiOKWkhtbMzc7NDBtICAgIBtbMG0bWzBtICAbWzMzOzQwbeKWk+KWiOKWiOKWkhtbMzc7NDBtICAgIBtbMG0bWzBtICAbWzMzOzQwbeKWkxtbOTc7NDNt4paIG1szMzs0MG3ilojilpEbWzk3OzQwbSAbWzMzOzQwbSDilojilojilpMbWzBtG1swbSAgG1szMzs0MG3ilpPilogbWzk3OzQzbeKWiBtbMzM7NDBt4paSG1s5Nzs0MG0gIBtbMzM7NDBt4paI4paI4paTG1swbRtbMG0gIBtbMzM7NDBt4paTG1s5Nzs0MG0gIBtbOTc7NDNt4paIG1szMzs0MG3ilojilpIbWzk3OzQwbSAbWzMzOzQwbeKWk+KWkhtbMG0bWzBtICAbWzBtChtbMzM7NDBt4paSG1szMzs0M23ilogbWzMzOzQwbeKWiOKWiOKWgOKWiBtbOTc7NDNt4paIG1szMzs0MG3ilpEbWzBtG1swbSAgG1szMzs0MG3ilpIbWzk3OzQzbeKWiCAbWzMzOzQwbeKWiBtbMzc7NDBtICAgG1swbRtbMG0gIBtbMzM7NDBt4paSG1s5Nzs0M23ilogbWzMzOzQwbeKWiOKWkRtbMzc7NDBtICAgIBtbMG0bWzBtICAbWzMzOzQwbeKWkhtbOTc7NDNt4paIG1szMzs0MG3ilojilpEbWzM3OzQwbSAgICAbWzBtG1swbSAgG1szMzs0MG3ilpMbWzk3OzQzbeKWiBtbMzM7NDBt4paI4paRIOKWiOKWiOKWk+KWkhtbMG0bWzBtICAbWzMzOzQwbeKWkuKWiBtbOTc7NDNt4paIG1szMzs0MG3ilpEbWzk3OzQwbSAgG1s5Nzs0M23ilogbWzMzOzQwbeKWiOKWkhtbMG0bWzBtICAbWzMzOzQwbeKWkhtbOTc7NDBtIBtbMzM7NDBt4paTG1s5Nzs0M23ilogbWzMzOzQwbeKWiOKWkRtbOTc7NDBtIBtbMzM7NDBt4paS4paRG1swbRtbMG0gIBtbMG0KG1szMzs0MG3ilpHilpPilogbWzk3OzQwbSAbWzMzOzQwbeKWk+KWiBtbOTc7NDNt4paIG1szNzs0MG0gG1swbRtbMG0gIBtbMzM7NDBt4paS4paT4paIG1s5Nzs0MG0gIBtbMzM7NDBt4paEG1szNzs0MG0gG1swbRtbMG0gIBtbMzM7NDBt4paSG1s5Nzs0M23iloggG1szMzs0MG3ilpEbWzM3OzQwbSAgICAbWzBtG1swbSAgG1szMzs0MG3ilpIbWzk3OzQzbeKWiCAbWzMzOzQwbeKWkRtbMzc7NDBtICAgIBtbMG0bWzBtICAbWzMzOzQwbeKWkhtbOTc7NDNt4paEIBtbMzM7NDBt4paE4paI4paS4paTG1s5Nzs0MG0gG1szMzs0MG3ilpIbWzBtG1swbSAgG1szMzs0MG3ilpLilojilogbWzk3OzQwbSAgG1s5Nzs0M23iloTiloAbWzMzOzQwbeKWkuKWkRtbMG0bWzBtICAbWzMzOzQwbeKWkRtbOTc7NDBtIBtbMzM7NDBt4paTG1s5Nzs0M23iloQbWzMzOzQwbeKWiOKWkxtbOTc7NDBtIBtbMzM7NDBt4paRG1szNzs0MG0gG1swbRtbMG0gIBtbMG0KG1szMzs0MG3ilpHilpPilojilpLilpMbWzk3OzQzbSDiloQbWzMzOzQwbeKWkxtbMG0bWzBtICAbWzMzOzQwbeKWkeKWkuKWiOKWiOKWiOKWiOKWkhtbMG0bWzBtICAbWzMzOzQwbeKWkRtbOTc7NDNt4paEG1szMzs0MG3ilojilojilojilojilojilpIbWzBtG1swbSAgG1szMzs0MG3ilpEbWzk3OzQzbeKWhBtbMzM7NDBt4paI4paI4paI4paI4paI4paSG1swbRtbMG0gIBtbMzM7NDBt4paS4paI4paI4paSG1s5Nzs0MG0gG1szMzs0MG3ilpEbWzk3OzQwbSAgG1szMzs0MG3ilpEbWzBtG1swbSAgG1szMzs0MG3ilpEbWzk3OzQwbSAbWzMzOzQwbeKWiBtbOTc7NDNt4paA4paAIBtbMzM7NDBt4paT4paS4paRG1swbRtbMG0gIBtbOTc7NDBtICAbWzMzOzQwbeKWkuKWiOKWiOKWkhtbOTc7NDBtIBtbMzM7NDBt4paRG1szNzs0MG0gG1swbRtbMG0gIBtbMG0KG1s5Nzs0MG0gG1szMzs0MG3ilpIbWzk3OzQwbSAbWzMzOzQwbeKWkeKWk+KWkuKWkeKWkhtbMG0bWzBtICAbWzMzOzQwbeKWkeKWkRtbOTc7NDBtIBtbMzM7NDBt4paS4paRG1s5Nzs0MG0gG1szMzs0MG3ilpEbWzBtG1swbSAgG1szMzs0MG3ilpEbWzk3OzQwbSAbWzMzOzQwbeKWkuKWkeKWkxtbOTc7NDBtICAbWzMzOzQwbeKWkRtbMG0bWzBtICAbWzMzOzQwbeKWkRtbOTc7NDBtIBtbMzM7NDBt4paS4paR4paTG1s5Nzs0MG0gIBtbMzM7NDBt4paRG1swbRtbMG0gIBtbMzM7NDBt4paS4paT4paS4paRG1s5Nzs0MG0gG1szMzs0MG3ilpEbWzk3OzQwbSAgG1szMzs0MG3ilpEbWzBtG1swbSAgG1szMzs0MG3ilpEbWzk3OzQwbSAbWzMzOzQwbeKWkuKWkeKWk+KWkeKWk+KWkRtbMzc7NDBtIBtbMG0bWzBtICAbWzk3OzQwbSAgG1szMzs0MG3ilpIbWzk3OzQwbSAbWzMzOzQwbeKWkeKWkRtbMzc7NDBtICAgG1swbRtbMG0gIBtbMG0KG1s5Nzs0MG0gG1szMzs0MG3ilpIbWzk3OzQwbSAbWzMzOzQwbeKWkeKWkuKWkRtbOTc7NDBtIBtbMzM7NDBt4paRG1swbRtbMG0gIBtbMzM7NDBtIOKWkRtbOTc7NDBtIBtbMzM7NDBt4paRG1s5Nzs0MG0gIBtbMzM7NDBt4paRG1swbRtbMG0gIBtbMzM7NDBt4paRG1s5Nzs0MG0gG1szMzs0MG3ilpEbWzk3OzQwbSAbWzMzOzQwbeKWkhtbOTc7NDBtICAbWzMzOzQwbeKWkRtbMG0bWzBtICAbWzMzOzQwbeKWkRtbOTc7NDBtIBtbMzM7NDBt4paRG1s5Nzs0MG0gG1szMzs0MG3ilpIbWzk3OzQwbSAgG1szMzs0MG3ilpEbWzBtG1swbSAgG1szMzs0MG3ilpHilpIbWzk3OzQwbSAbWzMzOzQwbeKWkRtbMzc7NDBtICAgICAbWzBtG1swbSAgG1s5Nzs0MG0gIBtbMzM7NDBt4paRG1s5Nzs0MG0gG1szMzs0MG3ilpIbWzk3OzQwbSAbWzMzOzQwbeKWkuKWkRtbMzc7NDBtIBtbMG0bWzBtICAbWzk3OzQwbSAgICAbWzMzOzQwbeKWkRtbMzc7NDBtICAgIBtbMG0bWzBtICAbWzBtChtbOTc7NDBtIBtbMzM7NDBt4paRG1s5Nzs0MG0gIBtbMzM7NDBt4paR4paRG1s5Nzs0MG0gG1szMzs0MG3ilpEbWzBtG1swbSAgG1szMzs0MG0gG1s5Nzs0MG0gIBtbMzM7NDBt4paRG1szNzs0MG0gICAbWzBtG1swbSAgG1s5Nzs0MG0gIBtbMzM7NDBt4paRG1s5Nzs0MG0gG1szMzs0MG3ilpEbWzM3OzQwbSAgIBtbMG0bWzBtICAbWzk3OzQwbSAgG1szMzs0MG3ilpEbWzk3OzQwbSAbWzMzOzQwbeKWkRtbMzc7NDBtICAgG1swbRtbMG0gIBtbMzM7NDBt4paR4paRG1szNzs0MG0gICAgICAgG1swbRtbMG0gIBtbMzM7NDBt4paRG1s5Nzs0MG0gG1szMzs0MG3ilpEbWzk3OzQwbSAbWzMzOzQwbeKWkRtbOTc7NDBtIBtbMzM7NDBt4paSG1szNzs0MG0gIBtbMG0bWzBtICAbWzk3OzQwbSAgG1szMzs0MG3ilpEbWzM3OzQwbSAgICAgIBtbMG0bWzBtICAbWzBtChtbOTc7NDBtIBtbMzM7NDBt4paRG1s5Nzs0MG0gIBtbMzM7NDBt4paRG1s5Nzs0MG0gIBtbMzM7NDBt4paRG1swbRtbMG0gIBtbOTc7NDBtICAgG1szMzs0MG3ilpEbWzk3OzQwbSAgG1szMzs0MG3ilpEbWzBtG1swbSAgG1s5Nzs0MG0gICAgG1szMzs0MG3ilpEbWzk3OzQwbSAgG1szMzs0MG3ilpEbWzBtG1swbSAgG1s5Nzs0MG0gICAgG1szMzs0MG3ilpEbWzk3OzQwbSAgG1szMzs0MG3ilpEbWzBtG1swbSAgG1szMDs0MG0gICAgICAgICAbWzBtG1swbSAgG1s5Nzs0MG0gICAgG1szMzs0MG3ilpEbWzk3OzQwbSAbWzMzOzQwbeKWkRtbMzc7NDBtICAbWzBtG1swbSAgG1szMDs0MG0gICAgICAgICAbWzBtG1swbSAgG1swbQoK"

/*Banner (print banner)

  load the base64 data which contains the banner
  then after decoding the banner we will iterate through the resulting string line by line

  we do this so that in the future we can add effects like a gradient to the banner if desired

  if we didn't want to worry about that we could nix the strings and bufio imports
  and just fmt.Println(dec) to print the banner without additional styling
*/
func Banner() {
	dec, _ := base64.StdEncoding.DecodeString(banner)

	scanner := bufio.NewScanner(strings.NewReader(string(dec)))
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	// ALTERNATE OPTION:
	// fmt.Println(dec)
}
