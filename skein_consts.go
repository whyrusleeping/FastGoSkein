package main

var SKEIN1024_IV_384 = []uint64{
    0x5102B6B8C1894A35,
    0xFEEBC9E3FE8AF11A,
    0x0C807F06E32BED71,
    0x60C13A52B41A91F6,
    0x9716D35DD4917C38,
    0xE780DF126FD31D3A,
    0x797846B6C898303A,
    0xB172C2A8B3572A3B,
    0xC9BC8203A6104A6C,
    0x65909338D75624F4,
    0x94BCC5684B3F81A0,
    0x3EBBF51E10ECFD46,
    0x2DF50F0BEEB08542,
    0x3B5A65300DBC6516,
    0x484B9CD2167BBCE1,
    0x2D136947D4CBAFEA}

/* blkSize = 1024 bits. hashSize =  512 bits */
var SKEIN1024_IV_512 = []uint64{
    0xCAEC0E5D7C1B1B18,
    0xA01B0E045F03E802,
    0x33840451ED912885,
    0x374AFB04EAEC2E1C,
    0xDF25A0E2813581F7,
    0xE40040938B12F9D2,
    0xA662D539C2ED39B6,
    0xFA8B85CF45D8C75A,
    0x8316ED8E29EDE796,
    0x053289C02E9F91B8,
    0xC3F8EF1D6D518B73,
    0xBDCEC3C4D5EF332E,
    0x549A7E5222974487,
    0x670708725B749816,
    0xB9CD28FBF0581BD1,
    0x0E2940B815804974}

/* blkSize = 1024 bits. hashSize = 1024 bits */
var SKEIN1024_IV_1024 = []uint64{
    0xD593DA0741E72355,
    0x15B5E511AC73E00C,
    0x5180E5AEBAF2C4F0,
    0x03BD41D3FCBCAFAF,
    0x1CAEC6FD1983A898,
    0x6E510B8BCDD0589F,
    0x77E2BDFDC6394ADA,
    0xC11E1DB524DCB0A3,
    0xD6D14AF9C6329AB5,
    0x6A9B0BFC6EB67E0D,
    0x9243C60DCCFF1332,
    0x1A1F1DDE743F02D4,
    0x0996753C10ED0BB8,
    0x6572DD22F2B4969A,
    0x61FD3062D00A579A,
    0x1DE0536E8682E539}

const (
	SKEIN1024_STATE_WORDS = uint(16)
	SKEIN1024_BLOCK_BYTES = 8 * SKEIN1024_STATE_WORDS
)
const (
	R1024_0_0=24
	R1024_0_1=13
	R1024_0_2= 8
	R1024_0_3=47
	R1024_0_4= 8
	R1024_0_5=17
	R1024_0_6=22
	R1024_0_7=37

	R1024_1_0=38
	R1024_1_1=19
	R1024_1_2=10
	R1024_1_3=55
	R1024_1_4=49
	R1024_1_5=18
	R1024_1_6=23
	R1024_1_7=52

	R1024_2_0=33
	R1024_2_1= 4
	R1024_2_2=51
	R1024_2_3=13
	R1024_2_4=34
	R1024_2_5=41
	R1024_2_6=59
	R1024_2_7=17

	R1024_3_0= 5
	R1024_3_1=20
	R1024_3_2=48
	R1024_3_3=41
	R1024_3_4=47
	R1024_3_5=28
	R1024_3_6=16
	R1024_3_7=25

	R1024_4_0=41
	R1024_4_1= 9
	R1024_4_2=37
	R1024_4_3=31
	R1024_4_4=12
	R1024_4_5=47
	R1024_4_6=44
	R1024_4_7=30

	R1024_5_0=16
	R1024_5_1=34
	R1024_5_2=56
	R1024_5_3=51
	R1024_5_4= 4
	R1024_5_5=53
	R1024_5_6=42
	R1024_5_7=41

	R1024_6_0=31
	R1024_6_1=44
	R1024_6_2=47
	R1024_6_3=46
	R1024_6_4=19
	R1024_6_5=42
	R1024_6_6=44
	R1024_6_7=25

	R1024_7_0= 9
	R1024_7_1=48
	R1024_7_2=35
	R1024_7_3=52
	R1024_7_4=23
	R1024_7_5=31
	R1024_7_6=37
	R1024_7_7=20
)
