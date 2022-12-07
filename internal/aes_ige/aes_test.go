// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of mtproto package.
// See https://github.com/ybrbnf2080/mtproto/blob/master/LICENSE for details

package ige

import (
	"encoding/hex"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoAES256IGEdecrypt(t *testing.T) {
	cases := newCasesAES256IGEdecrypt()

	for _, tcase := range cases {
		result := make([]byte, len(tcase.ciphertext))
		err := doAES256IGEdecrypt(tcase.ciphertext, result, tcase.key, tcase.initVector)
		if !assert.NoError(t, err) {
			return
		}
		if !assert.Equal(t, tcase.expected, result) {
			return
		}

		encrypted := make([]byte, len(result))
		err = doAES256IGEencrypt(result, encrypted, tcase.key, tcase.initVector)
		if !assert.NoError(t, err) {
			return
		}
		if !assert.Equal(t, tcase.ciphertext, encrypted) {
			return
		}
	}
}

func BenchmarkDoAES256IGEdecrypt(b *testing.B) {
	cases := newCasesAES256IGEdecrypt()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tcase := range cases {
			result := make([]byte, len(tcase.ciphertext))
			_ = doAES256IGEdecrypt(tcase.ciphertext, result, tcase.key, tcase.initVector)
		}
	}
}

func TestGenerateTempKeys(t *testing.T) {
	cases := newCasesGenerateTempKeys()

	for _, tcase := range cases {
		resultKey, resultIv := generateTempKeys(tcase.secondNonce, tcase.serverNonce)
		assert.Equal(t, tcase.expectedKey, resultKey)
		assert.Equal(t, tcase.expectedIV, resultIv)
	}
}

func BenchmarkGenerateTempKeys(b *testing.B) {
	cases := newCasesGenerateTempKeys()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tcase := range cases {
			resultKey, resultIV := generateTempKeys(tcase.secondNonce, tcase.serverNonce)
			assert.Equal(b, tcase.expectedKey, resultKey)
			assert.Equal(b, tcase.expectedIV, resultIV)
		}
	}
}

func TestDecryptMessageWithTempKeys(t *testing.T) {
	cases := newCasesDecryptMessageWithTempKeys()

	for _, tcase := range cases {
		result := DecryptMessageWithTempKeys(tcase.ciphertext, tcase.secondNonce, tcase.serverNonce)
		assert.Equal(t, tcase.expected, result)
	}
}

func BenchmarkDecryptMessageWithTempKeys(b *testing.B) {
	cases := newCasesDecryptMessageWithTempKeys()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, tcase := range cases {
			res := DecryptMessageWithTempKeys(tcase.ciphertext, tcase.secondNonce, tcase.serverNonce)
			assert.Equal(b, tcase.expected, res)
		}
	}
}

type caseAES256IGEdecrypt struct {
	ciphertext []byte
	key        []byte
	initVector []byte
	expected   []byte
}

func newCasesAES256IGEdecrypt() []caseAES256IGEdecrypt {
	return []caseAES256IGEdecrypt{
		{
			Hexed("28A92FE20173B347A8BB324B5FAB2667C9A8BBCE6468D5B509A4CBDDC186240A" +
				"C912CF7006AF8926DE606A2E74C0493CAA57741E6C82451F54D3E068F5CCC49B" +
				"4444124B9666FFB405AAB564A3D01E67F6E912867C8D20D9882707DC330B17B4" +
				"E0DD57CB53BFAAFA9EF5BE76AE6C1B9B6C51E2D6502A47C883095C46C81E3BE2" +
				"5F62427B585488BB3BF239213BF48EB8FE34C9A026CC8413934043974DB03556" +
				"633038392CECB51F94824E140B98637730A4BE79A8F9DAFA39BAE81E1095849E" +
				"A4C83467C92A3A17D997817C8A7AC61C3FF414DA37B7D66E949C0AEC858F0482" +
				"24210FCC61F11C3A910B431CCBD104CCCC8DC6D29D4A5D133BE639A4C32BBFF1" +
				"53E63ACA3AC52F2E4709B8AE01844B142C1EE89D075D64F69A399FEB04E656FE" +
				"3675A6F8F412078F3D0B58DA15311C1A9F8E53B3CD6BB5572C294904B726D0BE" +
				"337E2E21977DA26DD6E33270251C2CA29DFCC70227F0755F84CFDA9AC4B8DD5F" +
				"84F1D1EB36BA45CDDC70444D8C213E4BD8F63B8AB95A2D0B4180DC91283DC063" +
				"ACFB92D6A4E407CDE7C8C69689F77A007441D4A6A8384B666502D9B77FC68B5B" +
				"43CC607E60A146223E110FCB43BC3C942EF981930CDC4A1D310C0B64D5E55D30" +
				"8D863251AB90502C3E46CC599E886A927CDA963B9EB16CE62603B68529EE98F9" +
				"F5206419E03FB458EC4BD9454AA8F6BA777573CC54B328895B1DF25EAD9FB4CD" +
				"5198EE022B2B81F388D281D5E5BC580107CA01A50665C32B552715F335FD7626" +
				"4FAD00DDD5AE45B94832AC79CE7C511D194BC42B70EFA850BB15C2012C5215CA" +
				"BFE97CE66B8D8734D0EE759A638AF013"),
			Hexed("F011280887C7BB01DF0FC4E17830E0B91FBB8BE4B2267CB985AE25F33B527253"),
			Hexed("3212D579EE35452ED23E0D0C92841AA7D31B2E9BDEF2151E80D15860311C85DB"),
			Hexed("4B0AF668CF60A358233F93B7341FCA7E7F02A8C2BA0D89B53E0549828CCA27E9" +
				"66B301A48FECE2FCA5CF4D33F4A11EA877BA4AA57390733002000000FE000100" +
				"C71CAEB9C6B1C9048E6C522F70F13F73980D40238E3E21C14934D037563D930F" +
				"48198A0AA7C14058229493D22530F4DBFA336F6E0AC925139543AED44CCE7C37" +
				"20FD51F69458705AC68CD4FE6B6B13ABDC9746512969328454F18FAF8C595F64" +
				"2477FE96BB2A941D5BCD1D4AC8CC49880708FA9B378E3C4F3A9060BEE67CF9A4" +
				"A4A695811051907E162753B56B0F6B410DBA74D8A84B2A14B3144E0EF1284754" +
				"FD17ED950D5965B4B9DD46582DB1178D169C6BC465B0D6FF9CA3928FEF5B9AE4" +
				"E418FC15E83EBEA0F87FA9FF5EED70050DED2849F47BF959D956850CE929851F" +
				"0D8115F635B105EE2E4E15D04B2454BF6F4FADF034B10403119CD8E3B92FCC5B" +
				"FE000100262AABA621CC4DF587DC94CF8252258C0B9337DFB47545A49CDD5C9B" +
				"8EAE7236C6CADC40B24E88590F1CC2CC762EBF1CF11DCC0B393CAAD6CEE4EE58" +
				"48001C73ACBB1D127E4CB93072AA3D1C8151B6FB6AA6124B7CD782EAF981BDCF" +
				"CE9D7A00E423BD9D194E8AF78EF6501F415522E44522281C79D906DDB79C72E9" +
				"C63D83FB2A940FF779DFB5F2FD786FB4AD71C9F08CF48758E534E9815F634F1E" +
				"3A80A5E1C2AF210C5AB762755AD4B2126DFA61A77FA9DA967D65DFD0AFB5CDF2" +
				"6C4D4E1A88B180F4E0D0B45BA1484F95CB2712B50BF3F5968D9D55C99C0FB9FB" +
				"67BFF56D7D4481B634514FBA3488C4CDA2FC0659990E8E868B28632875A9AA70" +
				"3BCDCE8FCB7AE55199E2DDDD536648D8"),
		},
		{
			Hexed("1A8519A6557BE652E9DA8E43DA4EF4453CF456B4CA488AA383C79C98B34797CB"),
			Hexed("000102030405060708090A0B0C0D0E0F"),
			Hexed("000102030405060708090A0B0C0D0E0F101112131415161718191A1B1C1D1E1F"),
			Hexed("0000000000000000000000000000000000000000000000000000000000000000"),
		},
		{
			Hexed("4C2E204C6574277320686F70652042656E20676F74206974207269676874210A"),
			Hexed("5468697320697320616E20696D706C65"),
			Hexed("6D656E746174696F6E206F6620494745206D6F646520666F72204F70656E5353"),
			Hexed("99706487A1CDE613BC6DE0B6F24B1C7AA448C8B9C3403E3467A8CAD89340F53B"),
		},
	}
}

type caseGenerateTempKeys struct {
	secondNonce *big.Int
	serverNonce *big.Int
	expectedKey []byte
	expectedIV  []byte
}

func newCasesGenerateTempKeys() []caseGenerateTempKeys {
	return []caseGenerateTempKeys{
		{
			big.NewInt(0).SetBytes(Hexed("311C85DB234AA2640AFC4A76A735CF5B1F0FD68BD17FA181E1229AD867CC024D")),
			big.NewInt(0).SetBytes(Hexed("A5CF4D33F4A11EA877BA4AA573907330")),
			Hexed("F011280887C7BB01DF0FC4E17830E0B91FBB8BE4B2267CB985AE25F33B527253"),
			Hexed("3212D579EE35452ED23E0D0C92841AA7D31B2E9BDEF2151E80D15860311C85DB"),
		},
	}
}

type caseDecryptMessageWithTempKeys struct {
	ciphertext  []byte
	secondNonce *big.Int
	serverNonce *big.Int
	expected    []byte
}

func newCasesDecryptMessageWithTempKeys() []caseDecryptMessageWithTempKeys {
	return []caseDecryptMessageWithTempKeys{
		{
			Hexed("28A92FE20173B347A8BB324B5FAB2667C9A8BBCE6468D5B509A4CBDDC186240A" +
				"C912CF7006AF8926DE606A2E74C0493CAA57741E6C82451F54D3E068F5CCC49B" +
				"4444124B9666FFB405AAB564A3D01E67F6E912867C8D20D9882707DC330B17B4" +
				"E0DD57CB53BFAAFA9EF5BE76AE6C1B9B6C51E2D6502A47C883095C46C81E3BE2" +
				"5F62427B585488BB3BF239213BF48EB8FE34C9A026CC8413934043974DB03556" +
				"633038392CECB51F94824E140B98637730A4BE79A8F9DAFA39BAE81E1095849E" +
				"A4C83467C92A3A17D997817C8A7AC61C3FF414DA37B7D66E949C0AEC858F0482" +
				"24210FCC61F11C3A910B431CCBD104CCCC8DC6D29D4A5D133BE639A4C32BBFF1" +
				"53E63ACA3AC52F2E4709B8AE01844B142C1EE89D075D64F69A399FEB04E656FE" +
				"3675A6F8F412078F3D0B58DA15311C1A9F8E53B3CD6BB5572C294904B726D0BE" +
				"337E2E21977DA26DD6E33270251C2CA29DFCC70227F0755F84CFDA9AC4B8DD5F" +
				"84F1D1EB36BA45CDDC70444D8C213E4BD8F63B8AB95A2D0B4180DC91283DC063" +
				"ACFB92D6A4E407CDE7C8C69689F77A007441D4A6A8384B666502D9B77FC68B5B" +
				"43CC607E60A146223E110FCB43BC3C942EF981930CDC4A1D310C0B64D5E55D30" +
				"8D863251AB90502C3E46CC599E886A927CDA963B9EB16CE62603B68529EE98F9" +
				"F5206419E03FB458EC4BD9454AA8F6BA777573CC54B328895B1DF25EAD9FB4CD" +
				"5198EE022B2B81F388D281D5E5BC580107CA01A50665C32B552715F335FD7626" +
				"4FAD00DDD5AE45B94832AC79CE7C511D194BC42B70EFA850BB15C2012C5215CA" +
				"BFE97CE66B8D8734D0EE759A638AF013"),
			big.NewInt(0).SetBytes(Hexed("311C85DB234AA2640AFC4A76A735CF5B1F0FD68BD17FA181E1229AD867CC024D")),
			big.NewInt(0).SetBytes(Hexed("A5CF4D33F4A11EA877BA4AA573907330")),
			Hexed("BA0D89B53E0549828CCA27E966B301A48FECE2FCA5CF4D33F4A11EA877BA4AA5" +
				"7390733002000000FE000100C71CAEB9C6B1C9048E6C522F70F13F73980D4023" +
				"8E3E21C14934D037563D930F48198A0AA7C14058229493D22530F4DBFA336F6E" +
				"0AC925139543AED44CCE7C3720FD51F69458705AC68CD4FE6B6B13ABDC974651" +
				"2969328454F18FAF8C595F642477FE96BB2A941D5BCD1D4AC8CC49880708FA9B" +
				"378E3C4F3A9060BEE67CF9A4A4A695811051907E162753B56B0F6B410DBA74D8" +
				"A84B2A14B3144E0EF1284754FD17ED950D5965B4B9DD46582DB1178D169C6BC4" +
				"65B0D6FF9CA3928FEF5B9AE4E418FC15E83EBEA0F87FA9FF5EED70050DED2849" +
				"F47BF959D956850CE929851F0D8115F635B105EE2E4E15D04B2454BF6F4FADF0" +
				"34B10403119CD8E3B92FCC5BFE000100262AABA621CC4DF587DC94CF8252258C" +
				"0B9337DFB47545A49CDD5C9B8EAE7236C6CADC40B24E88590F1CC2CC762EBF1C" +
				"F11DCC0B393CAAD6CEE4EE5848001C73ACBB1D127E4CB93072AA3D1C8151B6FB" +
				"6AA6124B7CD782EAF981BDCFCE9D7A00E423BD9D194E8AF78EF6501F415522E4" +
				"4522281C79D906DDB79C72E9C63D83FB2A940FF779DFB5F2FD786FB4AD71C9F0" +
				"8CF48758E534E9815F634F1E3A80A5E1C2AF210C5AB762755AD4B2126DFA61A7" +
				"7FA9DA967D65DFD0AFB5CDF26C4D4E1A88B180F4E0D0B45BA1484F95CB2712B5" +
				"0BF3F5968D9D55C99C0FB9FB67BFF56D7D4481B634514FBA3488C4CDA2FC0659" +
				"990E8E868B28632875A9AA703BCDCE8FCB7AE551"),
		},
	}
}

func TestMessageKey(t *testing.T) {
	tests := []struct {
		name string
		msg  []byte
		want []byte
	}{
		{
			name: "empty",
			msg:  Hexed(""),
			// correct: http://craiccomputing.blogspot.com/2009/09/sha1-digest-of-empty-string.html
			want: Hexed("5E6B4B0D3255BFEF95601890AFD80709"),
		},
		{
			name: "zeros",
			msg:  Hexed("00000000000000000000000000000000"),
			want: Hexed("5103BC5CC44BCDF0A15E160D445066FF"),
		},
		{
			name: "randomized",
			msg:  Hexed("5103BC5CC44BCDF0A15E160D445066FF"),
			want: Hexed("F9D401E298F3EEEC1C927312AEB6B412"),
		},
		{
			name: "any words",
			msg:  []byte("some cool message"),
			want: Hexed("4170ED208083FAFD2DFA8507FD4A75B6"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, MessageKey(tt.msg))
		})
	}
}

func TestEncryptMessageWithTempKeys(t *testing.T) {
	tests := []struct {
		name  string
		msg   []byte
		nSec  []byte
		nServ []byte
		want  []byte
	}{
		{
			msg: Hexed("F78AF98EF9D401E298F3EEEC1C927312AEB6B4125103BC5CC44BCDF0A15E160D" +
				"445066FF000000000000000000000000"),
			nSec:  Hexed("F011280887C7BB01DF0FC4E17830E0B91FBB8BE4B2267CB985AE25F33B527253"),
			nServ: Hexed("F011280887C7BB01DF0FC4E17830E0B91FBB8BE4B2267CB985AE25F33B527253"),
			want: Hexed("9112F2583EACE884A58D2A5A6047C10F4BE228946C6B66CDE9268C20FC1528CF" +
				"CE120A9C53B90D71B6CB8B517172C03E"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t,
				tt.want,
				encryptMessageWithTempKeys(
					tt.msg,
					big.NewInt(0).SetBytes(tt.nSec),
					big.NewInt(0).SetBytes(tt.nServ),
				),
			)
		})
	}
}

func TestEncryptDecrypt(t *testing.T) {
	tests := []struct {
		name    string
		msg     []byte
		key     []byte
		want    []byte
		wantErr assert.ErrorAssertionFunc
	}{
		{
			name: "simple",
			msg:  []byte("hello world!"),
			key: Hexed("28A92FE20173B347A8BB324B5FAB2667C9A8BBCE6468D5B509A4CBDDC186240A" +
				"C912CF7006AF8926DE606A2E74C0493CAA57741E6C82451F54D3E068F5CCC49B" +
				"4444124B9666FFB405AAB564A3D01E67F6E912867C8D20D9882707DC330B17B4" +
				"E0DD57CB53BFAAFA9EF5BE76AE6C1B9B6C51E2D6502A47C883095C46C81E3BE2" +
				"5F62427B585488BB3BF239213BF48EB8FE34C9A026CC8413934043974DB03556" +
				"633038392CECB51F94824E140B98637730A4BE79A8F9DAFA39BAE81E1095849E" +
				"A4C83467C92A3A17D997817C8A7AC61C3FF414DA37B7D66E949C0AEC858F0482" +
				"24210FCC61F11C3A910B431CCBD104CCCC8DC6D29D4A5D133BE639A4C32BBFF1" +
				"53E63ACA3AC52F2E4709B8AE01844B142C1EE89D075D64F69A399FEB04E656FE" +
				"3675A6F8F412078F3D0B58DA15311C1A9F8E53B3CD6BB5572C294904B726D0BE" +
				"337E2E21977DA26DD6E33270251C2CA29DFCC70227F0755F84CFDA9AC4B8DD5F" +
				"84F1D1EB36BA45CDDC70444D8C213E4BD8F63B8AB95A2D0B4180DC91283DC063" +
				"ACFB92D6A4E407CDE7C8C69689F77A007441D4A6A8384B666502D9B77FC68B5B" +
				"43CC607E60A146223E110FCB43BC3C942EF981930CDC4A1D310C0B64D5E55D30" +
				"8D863251AB90502C3E46CC599E886A927CDA963B9EB16CE62603B68529EE98F9" +
				"F5206419E03FB458EC4BD9454AA8F6BA777573CC54B328895B1DF25EAD9FB4CD" +
				"5198EE022B2B81F388D281D5E5BC580107CA01A50665C32B552715F335FD7626" +
				"4FAD00DDD5AE45B94832AC79CE7C511D194BC42B70EFA850BB15C2012C5215CA" +
				"BFE97CE66B8D8734D0EE759A638AF013"),
			want: Hexed("BABBC03F65F828E4B97DBC5A992394C2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wantErr := tt.wantErr
			if wantErr == nil {
				wantErr = assert.NoError
			}
			got, err := Encrypt(tt.msg, tt.key)
			if !wantErr(t, err) {
				return
			}

			if !assert.Equal(t, tt.want, got) {
				return
			}

			// TODO: почему-то расшифровка так легко не проходит
			//msgkey := MessageKey(tt.msg)
			//decrypted, err := Decrypt(got, tt.key, msgkey)
			//if !wantErr(t, err) {
			//	return
			//}
			//
			//assert.Equal(t, tt.msg, decrypted)
		})
	}
}

func Hexed(in string) []byte {
	res, err := hex.DecodeString(in)
	if err != nil {
		panic(err)
	}
	return res
}
