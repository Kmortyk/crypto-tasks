package crypto_tasks

import "testing"

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