package crypto_tasks

import (
	"testing"
)

// TODO сумма логрифмов для скора
// TODO применить словарь для разбиения

func TestDecrypt(t *testing.T) {
	first := "gluhtlishjrvbadvyyplkaohavbyjpwolypzavvdlhrvuuleatlzzhnlzdpajoavcpnlulyljpwolyrlfdvykpzaolopkkluzftivsvmklhaoputfmhcvypalovsilpuluk"

	for i := 0; i < 26; i++ {
		decryptedFirst := DecryptCeaser(first, i)
		t.Logf("%v: %s", i, decryptedFirst)
	}
	// ze name blackout worried that our cipher is too weak on next message switch to vigenere cipher keyword is the hidden symbol of death in my favorite holbein end

	t.Logf("")

	secon := "vwduwljudeehghyhubwklqjlfrxogilqgsohdvhuhwxuqdqbeoxhsulqwviruydxowdqgdodupghvljqedvhgrqzklfkedqnbrxghflghrqldpvhwwlqjxsvdihkrxvhfr"
	for i := 0; i < 26; i++ {
		decryptedFirst := DecryptCeaser(secon, i)
		t.Logf("%v: %s", i, decryptedFirst)
	}
	// start i grabbed everything i could find please return any blueprints for vault and alarm design based on which bank you decide on i am setting up safe house co
}

func TestDecryptPolyalphabetic(t *testing.T) {
	keyword := "sskkuullll"
	message := "klkbnqlcytfysryucocphgbdizzfcmjwkuchzyeswfogmmetwwossdchrzyldsbwnydednzwnefydthtddbojicemlucdygicczhoadrzcylwadsxpilpiecskomoltejtkmqqymehpmmjxyolwpeewjckznpccpsvsxauyodhalmriocwpelwbcniyfxmwjcemcyrazdqlsomdbfljwnbijxpddsyoehxpceswtoxwbleecsaxcnuetzywfn"

	decrypted := DecryptVigenere(message, keyword)
	t.Logf(decrypted)
	// start warning i heard report of our break in on the news still waiting on alarm test schedules i will report back tomorrow with final plan for extra security i suggest we burn our letters after reading and switch our letters to numbers using polybius square drop message under the bench at train station end
}

func TestDecryptPolybius(t *testing.T) {
	message := "445411345411233353445412424342443251412123113113531554425442444243443251415343543242344111255135533413424322534311445434534322513431421432513412533412155415345133514444112251444254424444153451235515432134511113112123514254315333214243514453153414345125425315443351543253414431111111111143513544"
	decrypted, err := DecryptPolybius(message)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf(decrypted)
	// start almost finished blackout it is in shedon third ave working on a stronger cipher for future messages it is surely unbreakable it combines our previous methods caaaaa news
}

func TestDecryptDirs(t *testing.T) {
	var err error

	letter := []string{
		"ddr", "lur", "dul", "ddl", "udr", "lur", "ldr", "udl", "uul", "ddr", "dul", "uul", "lul", "ddr", "rdl", "lur", "uur", "ldr", "uur", "lul", "uur", "udr", "dul", "udl", "lur", "dur",
		"dur", "rdr", "uur", "dul", "rur", "rur", "ddr", "udl", "rdl", "uur", "dul", "ddr", "ddr", "ddr", "dul", "lur", "ddr", "rur", "dur", "rur", "lur", "dul",
		"ldr", "lur", "udl", "ddr", "ddl", "rdr", "rul", "rdl", "udr", "lul", "dur", "rur", "dul", "uul", "udr", "rdr", "ldl", "uul", "lur", "ddr", "ldl", "uur", "rul",
		"udl", "rdl", "lur", "lul", "dur", "uul", "udl", "udr", "lul", "rdr", "rdr", "ldr", "udl", "rdr", "dur", "ldl", "rdr", "uul", "udr", "ldr", "dur", "ldl", "rdr",
		"uur", "uur", "udl", "uul", "lur", "uul", "dur", "dul", "ldr", "rul", "uur", "rul", "dul", "udr", "uur", "rur", "ldl", "udr", "rdr", "rdl", "ddl", "udr", "dur", "uul",
		"ddr", "rur", "uul", "lul", "ldr", "ldl", "dur", "lur", "ddr", "uul", "rul", "lur", "ldr", "udr", "ldr", "rul", "ldr", "rur", "uur", "udl", "uur", "uur",
		"dur", "dul", "uul", "ldr", "ddr", "uur", "dul", "dur", "lul", "dur", "uul", "uul", "ldl", "ddr", "dur", "rur", "ldr", "uul",
	}

	pad := "The whole grain goodness of blue chip dividend stocks has its limits. \nUtility stocks, consumer staples, " +
		"pipelines, telecoms and real estate investment trusts have all lost ground over the past month, even while the " +
		"broader market has been flat. With the bond market signalling an expectation of rising interest rates, the " +
		"five-year rally for steady blue-chip divident payers has stalled. \nShould you be scared if you own a lot of " +
		"these stocks either directly or through mutual funds or exchange-traded funds? David Baskin, president of " +
		"Baskin Financial Services, has a two-pronged answer: Keep your top-quality dividend stocks, but be prepared " +
		"to follow his firm's example in trimming holdings in stocks such as TransCanada Corp., Keyera Corp. and " +
		"Pembina Pipeline Corp. \nLet's have Mr. Baskin run us through his thinking in dividend stocks, which are " +
		"a big part of the portfolios his firm puts together for clients."

	pad, err = Normalize(pad)
	if err != nil {
		t.Fatal(err)
	}

	pad = VowelConsonantsBinarization(pad)

	cardinalBinary := map[string]string{
		"u": "00",
		"r": "01",
		"d": "10",
		"l": "11",
	}

	intermediateBinary := map[string]string{
		"dr": "00",
		"dl": "01",
		"ul": "10",
		"ur": "11",
	}

	// to binary
	letterBinary := ""
	for i := range letter {
		char := letter[i]
		letterBinary += cardinalBinary[char[0:1]]
		letterBinary += intermediateBinary[char[1:]]
	}

	// XOR
	xored := make([]rune, len(letterBinary))
	for i := 0; i < len(letterBinary); i++ {
		if letterBinary[i] == pad[i] {
			xored[i] = '0'
		} else {
			xored[i] = '1'
		}
	}

	polybius := make([]rune, len(xored)/3)
	digits := map[string]rune{
		"000": '0',
		"001": '1',
		"010": '2',
		"011": '3',
		"100": '4',
		"101": '5',
		"110": '6',
		"111": '7',
	}
	for i := 0; i < len(xored); i += 3 {
		if i >= len(xored) || i+1 >= len(xored) || i+2 >= len(xored) {
			break
		}
		str := string(xored[i]) + string(xored[i+1]) + string(xored[i+2])
		polybius[i/3] = digits[str]
	}

	// 525150535130034030405035055210105230021025505103305200455350155053255035203 150411551020351514525455353452150512414502550005110535015505025511052513150411531304520100352424232222510105150514015503005454151103522
	// 525150535130034030405435055210105230021025505103305200455350155053255035203 356220162070653241755333375600707320713543247164611631004723002514257031175551602640416054366133713172512036772002304514006020057461036517520

	// 101010101001101000101011101001011000000011100000011000100000101100011101000101101010001000001000101010011000000010001000010101101000101001000011011000101010000000100101101011101000001101101000101011010101101000011101010000011 011101110010010000001110010000111000110101011010100001111101101011011011011111101110000000111000111011010000111001011101100011010100111001110100110001001110011001000000100111010011000000010101001100010101111000011001001111101101101001110000010110100000100001110000101100011110110001011011111001011001111010101001010000011110111111010000000010011000100101001100000000110000010000000101111100110001000011110101001111101010000
	// 101010101001101000101011101001011000000011100000011000100000101000011101000101101010001000001000101010011000000010001000010101101000101001000011011000101010000000100101101011101000001101101000101011010101101000011101010000011 00110100010000100110110100100001000001110100110100110010101010110010110101110101110010101000110100010100101010000110010100001010110100000000010100100100010101110100000110110100010100001010110100100100010101010100101100110100010000100110101100101100010010101000000100000001110101010001010001001101001001001010100100000100010100110100010100110000000110110100001100000010110010110000110100100100001110101001010

	t.Log(string(xored))
	t.Log(string(polybius))

	decrypted, err := DecryptPolybiusSpiral(string(polybius))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(decrypted)
}
