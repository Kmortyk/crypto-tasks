package khan

import (
	"testing"
)

// TODO сумма логрифмов для скора
// TODO применить словарь для разбиения
// TODO бенчмарки и оптимизация

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
	}
	for i := 0; i < len(xored); i += 3 {
		if i >= len(xored) || i+1 >= len(xored) || i+2 >= len(xored) {
			break
		}
		str := string(xored[i]) + string(xored[i+1]) + string(xored[i+2])
		polybius[i/3] = digits[str]
	}

	t.Log(string(xored))
	t.Log(string(polybius))

	decrypted, err := DecryptPolybiusSpiral(string(polybius))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(decrypted) //alarm	     vault                               end
	// start CIBC bank see schematics for alarm and vault hit tomorrow at 10 am after alaam test vaulv code is 5567 meet at blackout en7
}

func TestEncryptPolybius(t *testing.T) {
	table := [][]rune{
		{'a', 'b', 'c', 'd', 'e', 'f'},
		{'g', 'h', 'i', 'j', 'k', 'l'},
		{'m', 'n', 'o', 'p', 'q', 'r'},
		{'s', 't', 'u', 'v', 'w', 'x'},
		{'y', 'z', '1', '2', '3', '4'},
		{'5', '6', '7', '8', '9', '0'},
	}

	t.Log("alarm", EncryptPolybius(table, "alarm", false))
	t.Log("start", EncryptPolybius(table, "start", false))
	t.Log("vault", EncryptPolybius(table, "vault", false))
	t.Log("5567", EncryptPolybius(table, "5567", false))
	t.Log("10am", EncryptPolybius(table, "10am", false))
	t.Log("bank", EncryptPolybius(table, "bank", false))
}

func TestEncryptMessageAndPad(t *testing.T) {
	var err error

	message := "me"
	pad := "open me to read"

	// pad
	pad, err = Normalize(pad)
	if err != nil {
		t.Fatal(err)
	}
	pad = VowelConsonantsBinarization(pad)

	// message
	message = EncryptPolybius([][]rune{
		{'f', 'g', 'h', 'i', 'j', 'k'},
		{'e', 'x', 'y', 'z', '0', 'l'},
		{'d', 'w', '7', '8', '1', 'm'},
		{'c', 'v', '6', '9', '2', 'n'},
		{'b', 'u', '5', '4', '3', 'o'},
		{'a', 't', 's', 'r', 'q', 'p'},
	}, message, true)
	digits := map[rune]string{
		'0': "000",
		'1': "001",
		'2': "010",
		'3': "011",
		'4': "100",
		'5': "101",
	}
	message2 := ""
	for _, ch := range message {
		message2 += digits[ch]
	}

	// result
	t.Log(StrXOR(message2, pad))
}

func TestOnesInLetter(t *testing.T) {
	letter := "w"

	letter = EncryptPolybius([][]rune{
		{'f', 'g', 'h', 'i', 'j', 'k'},
		{'e', 'x', 'y', 'z', '0', 'l'},
		{'d', 'w', '7', '8', '1', 'm'},
		{'c', 'v', '6', '9', '2', 'n'},
		{'b', 'u', '5', '4', '3', 'o'},
		{'a', 't', 's', 'r', 'q', 'p'},
	}, letter, true)
	digits := map[rune]string{
		'0': "000",
		'1': "001",
		'2': "010",
		'3': "011",
		'4': "100",
		'5': "101",
	}

	count := 0
	for _, ch := range letter {
		if digits[ch][0] == '1' {
			count++
		}
		if digits[ch][1] == '1' {
			count++
		}
		if digits[ch][2] == '1' {
			count++
		}
	}

	t.Log(count)
}

func TestDirToBin(t *testing.T) {
	dir := "uur"

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
	letterBinary := cardinalBinary[dir[0:1]] + intermediateBinary[dir[1:]]
	t.Log(letterBinary)
}

func TestDirWithPad(t *testing.T) {
	dirs := []string{"uur", "lur", "rur", "rdl", "rdl", "dul"}
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

	letterBinary := ""
	for i := range dirs {
		char := dirs[i]
		letterBinary += cardinalBinary[char[0:1]]
		letterBinary += intermediateBinary[char[1:]]
	}

	pad, err := Normalize("did you decrypt the final code?")
	if err != nil {
		t.Fatal(err)
	}
	pad = VowelConsonantsBinarization(pad)
	xored := StrXOR(letterBinary, pad)

	polybius := make([]rune, len(xored)/3)
	digits := map[string]rune{
		"000": '0',
		"001": '1',
		"010": '2',
		"011": '3',
		"100": '4',
		"101": '5',
	}
	for i := 0; i < len(xored); i += 3 {
		if i >= len(xored) || i+1 >= len(xored) || i+2 >= len(xored) {
			break
		}
		str := string(xored[i]) + string(xored[i+1]) + string(xored[i+2])
		polybius[i/3] = digits[str]
	}

	decrypted, err := DecryptPolybiusSpiral(string(polybius))
	if err != nil {
		t.Fatal(err)
	}

	t.Log(decrypted)
}

func TestDirToDigits(t *testing.T) {
	dir := "dul"

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
	digits := map[string]string{
		"00": "0",
		"01": "1",
		"10": "2",
		"11": "3",
	}

	c := dir[0:1]
	i := dir[1:]

	cb := cardinalBinary[c]
	ib := intermediateBinary[i]

	t.Log(cb, ib)
	t.Log(digits[cb] + digits[ib])
}

func TestCalcPad(t *testing.T) {
	pad, err := Normalize("One cloudy afternoon")
	if err != nil {
		t.Fatal(err)
	}
	pad = VowelConsonantsBinarization(pad)

	t.Log(pad)
}
