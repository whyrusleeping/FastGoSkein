package skein

import (
	"binary"
)

type CtxtHeader struct {
	hashBitLen int
	bCnt int
	T []uint64
}

type Skein1024 struct {
	h CtxtHeader
	X []uint64
	b []byte
}

func (s *Skein1024) Init(size uint) {
	s.h.hashBitLen = size
	switch size {
	case 1024:
		copy(s.X, SKEIN1024_IV_1024)
		//TODO: prebuilds for other sizes.
	default:
		//Build the IV values
	}
	s.h.T[0] = 0
	s.h.T[1] = SkeinT1FLAGFIRST | SKEIN_T1_BLK_TYPE_MSG
	s.h.bCnt=0
}

func (s *Skein1024) Update(msg []byte) {
	n := 0
	if len(msg) + s.h.bCnt > SKEIN1024_BLOCK_BYTES {
		if s.h.bCnt > 0 {
			n = SKEIN1024_BLOCK_BYTES - s.h.bCnt
			if n > 0 {
				//ASSERT: n < len(msg)
				copy(s.b[s.h.bCnt:],msg)
				msg = msg[n:]
				s.h.bCnt += n
			}
			if s.h.bCnt != SKEIN1024_BLOCK_BYTES {
				panic("ASSERTION FAILURE")
			}
			s.ProcessBlock(s.b,1, SKEIN1024_BLOCK_BYTES)
			s.h.bCnt = 0
		}

		if len(msg) > SKEIN1024_BLOCK_BYTES {
			n = (len(msg) - 1) / SKEIN1024_BLOCK_BYTES
			s.ProcessBlock(msg,n,SKEIN1024_BLOCK_BYTES)
			msg = msg[n * SKEIN1024_BLOCK_BYTES:]
		}
		//ASSERT: ss.h.bCnt == 0
	}

	if len(msg) > 0 {
		copy(s.b[s.h.bCnt:], msg)
		s.h.bCnt += len(msg)
	}
}

func (s *Skein1024) Final(outp []byte) {
	var n, byteCnt int
	X := make([]uint64, SKEIN1024_STATE_WORDS)

	s.h.T[1] |= SKEIN_T1_FLAG_FINAL
	if s.h.bCnt < SKEIN1024_BLOCK_BYTES {
		for j := s.h.bCnt j < SKEIN1024_BLOCK_BYTES j++ {
			s.b[j] = 0
		}
	}
	s.ProcessBlock(s.b,1,s.h.bCnt)

	byteCnt = (s.h.hashBitLen + 7) >> 3

	for j := 0 j < len(s.b) j++ {
		s.b[j] = 0
	}
	copy(X, s.X)
	for i := 0 i*SKEIN1024_BLOCK_BYTES < byteCnt i++ {
		binary.LittleEndian.PutUvarint(s.b, i)
		s.h.T[0] = 0
		s.h.T[1] = SkeinT1FLAGFIRST | SKEIN_T1_BLK_TYPE_OUT_FINAL
		s.h.bCnt=0
		s.ProcessBlock(s.b,1,8)
		n = byteCnt - i*SKEIN1024_BLOCK_BYTES   /* number of output bytes left to go */
		if n >= SKEIN1024_BLOCK_BYTES {
			n = SKEIN1024_BLOCK_BYTES
		}
		copy(outp[i*SKEIN1024_BLOCK_BYTES:],s.X)
		//Maybe unecessary?
		//Skein_Show_Final(1024,&ctx->h,n,hashVal+i*SKEIN1024_BLOCK_BYTES)
		copy(s.X,X)
	}
}

func (s *Skein1024) ProcessBlock(blk []byte, blkCnt, byteCntAdd int) {
	const WCNT = SKEIN1024_STATE_WORDS
	const RCNT = (SKEIN1024_ROUNDS_TOTAL/8)

	var kw [WCNT+4]uint64

	var X00,X01,X02,X03,X04,X05,X06,X07,X08,X09,X10,X11,X12,X13,X14,X15 uint64
	var w [WCNT]uint64

	//IF SKEIN DEBUG
	//INSERT DEBUG CODE HERE
	//ENDIF

	//ASSERT blkCnt != 0
	ts := kw //TEMPORARY, JUST TO MAKE TRANSLATING EASIER
	ks := kw[4:]
	ts[0] = s.h.T[0]
	ts[1] = s.h.T[1]

	//DO WHILE LOOP
	for {
		ts[0] += byteCntAdd

		ks[0] = s.X[0]
		ks[1] = s.X[1]
		ks[2] = s.X[2]
		ks[3] = s.X[3]
		ks[4] = s.X[4]
		ks[5] = s.X[5]
		ks[6] = s.X[6]
		ks[7] = s.X[7]
		ks[8] = s.X[8]
		ks[9] = s.X[9]
		ks[10] = s.X[10]
		ks[11] = s.X[11]
		ks[12] = s.X[12]
		ks[13] = s.X[13]
		ks[14] = s.X[14]
		ks[15] = s.X[15]
		ks[16] = ks[0] ^ ks[1] ^ ks[2] ^ ks[3] ^ ks[4] ^ ks[5] ^ ks[6] ^ ks[7] ^ ks[8] ^ ks[9] ^ ks[10] ^ ks[11] ^ ks[12] ^ ks[13] ^ ks[14] ^ ks[15] ^ SKEIN_KS_PARITY
		ts[2] = ts[0] ^ ts[1]

		Skein_Get64_LSB_First(w,blkPtr,WCNT) /* get input block in little-endian format */
		DebugSaveTweak(ctx)
		Skein_Show_Block(BLK_BITS,&ctx->h,ctx->X,blkPtr,w,ks,ts)

		X00    = w[ 0] + ks[ 0]                 /* do the first full key injection */
		X01    = w[ 1] + ks[ 1]
		X02    = w[ 2] + ks[ 2]
		X03    = w[ 3] + ks[ 3]
		X04    = w[ 4] + ks[ 4]
		X05    = w[ 5] + ks[ 5]
		X06    = w[ 6] + ks[ 6]
		X07    = w[ 7] + ks[ 7]
		X08    = w[ 8] + ks[ 8]
		X09    = w[ 9] + ks[ 9]
		X10    = w[10] + ks[10]
		X11    = w[11] + ks[11]
		X12    = w[12] + ks[12]
		X13    = w[13] + ks[13] + ts[0]
		X14    = w[14] + ks[14] + ts[1]
		X15    = w[15] + ks[15]

		Skein_Show_R_Ptr(BLK_BITS,&ctx->h,SKEIN_RND_KEY_INITIAL,Xptr)

		//R = 0
        R1024(00,01,02,03,04,05,06,07,08,09,10,11,12,13,14,15,R1024_0,8*(R) + 1); \
        R1024(00,09,02,13,06,11,04,15,10,07,12,03,14,05,08,01,R1024_1,8*(R) + 2); \
        R1024(00,07,02,05,04,03,06,01,12,15,14,13,08,11,10,09,R1024_2,8*(R) + 3); \
        R1024(00,15,02,11,06,13,04,09,14,01,08,05,10,03,12,07,R1024_3,8*(R) + 4); \
        I1024(2*(R));                                                             \
        R1024(00,01,02,03,04,05,06,07,08,09,10,11,12,13,14,15,R1024_4,8*(R) + 5); \
        R1024(00,09,02,13,06,11,04,15,10,07,12,03,14,05,08,01,R1024_5,8*(R) + 6); \
        R1024(00,07,02,05,04,03,06,01,12,15,14,13,08,11,10,09,R1024_6,8*(R) + 7); \
        R1024(00,15,02,11,06,13,04,09,14,01,08,05,10,03,12,07,R1024_7,8*(R) + 8); \
        I1024(2*(R)+1);

		//Like, Really?? this isnt even over yet...

	}

}

func WriteInt64OnByteArray(n uint64, arr []byte, index int) {
	binary.LittleEndian.PutUvarint(arr[index:], n)
}

func Skein_Put64_LSB_First(dst []byte, src []uint64) {
	//I think this function might be unnecessary...
}
