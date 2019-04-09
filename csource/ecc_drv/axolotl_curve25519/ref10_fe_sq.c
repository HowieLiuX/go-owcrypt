#include "ref10_fe.h"
#include "ref10_crypto_int64.h"

/*
h = f * f
Can overlap h with f.

Preconditions:
   |f| bounded by 1.65*2^26,1.65*2^25,1.65*2^26,1.65*2^25,etc.

Postconditions:
   |h| bounded by 1.01*2^25,1.01*2^24,1.01*2^25,1.01*2^24,etc.
*/

/*
See REF10_fe_mul.c for discussion of implementation strategy.
*/

void REF10_fe_sq(REF10_fe h,const REF10_fe f)
{
  REF10_crypto_int32 f0 = f[0];
  REF10_crypto_int32 f1 = f[1];
  REF10_crypto_int32 f2 = f[2];
  REF10_crypto_int32 f3 = f[3];
  REF10_crypto_int32 f4 = f[4];
  REF10_crypto_int32 f5 = f[5];
  REF10_crypto_int32 f6 = f[6];
  REF10_crypto_int32 f7 = f[7];
  REF10_crypto_int32 f8 = f[8];
  REF10_crypto_int32 f9 = f[9];
  REF10_crypto_int32 f0_2 = 2 * f0;
  REF10_crypto_int32 f1_2 = 2 * f1;
  REF10_crypto_int32 f2_2 = 2 * f2;
  REF10_crypto_int32 f3_2 = 2 * f3;
  REF10_crypto_int32 f4_2 = 2 * f4;
  REF10_crypto_int32 f5_2 = 2 * f5;
  REF10_crypto_int32 f6_2 = 2 * f6;
  REF10_crypto_int32 f7_2 = 2 * f7;
  REF10_crypto_int32 f5_38 = 38 * f5; /* 1.959375*2^30 */
  REF10_crypto_int32 f6_19 = 19 * f6; /* 1.959375*2^30 */
  REF10_crypto_int32 f7_38 = 38 * f7; /* 1.959375*2^30 */
  REF10_crypto_int32 f8_19 = 19 * f8; /* 1.959375*2^30 */
  REF10_crypto_int32 f9_38 = 38 * f9; /* 1.959375*2^30 */
  REF10_crypto_int64 f0f0    = f0   * (REF10_crypto_int64) f0;
  REF10_crypto_int64 f0f1_2  = f0_2 * (REF10_crypto_int64) f1;
  REF10_crypto_int64 f0f2_2  = f0_2 * (REF10_crypto_int64) f2;
  REF10_crypto_int64 f0f3_2  = f0_2 * (REF10_crypto_int64) f3;
  REF10_crypto_int64 f0f4_2  = f0_2 * (REF10_crypto_int64) f4;
  REF10_crypto_int64 f0f5_2  = f0_2 * (REF10_crypto_int64) f5;
  REF10_crypto_int64 f0f6_2  = f0_2 * (REF10_crypto_int64) f6;
  REF10_crypto_int64 f0f7_2  = f0_2 * (REF10_crypto_int64) f7;
  REF10_crypto_int64 f0f8_2  = f0_2 * (REF10_crypto_int64) f8;
  REF10_crypto_int64 f0f9_2  = f0_2 * (REF10_crypto_int64) f9;
  REF10_crypto_int64 f1f1_2  = f1_2 * (REF10_crypto_int64) f1;
  REF10_crypto_int64 f1f2_2  = f1_2 * (REF10_crypto_int64) f2;
  REF10_crypto_int64 f1f3_4  = f1_2 * (REF10_crypto_int64) f3_2;
  REF10_crypto_int64 f1f4_2  = f1_2 * (REF10_crypto_int64) f4;
  REF10_crypto_int64 f1f5_4  = f1_2 * (REF10_crypto_int64) f5_2;
  REF10_crypto_int64 f1f6_2  = f1_2 * (REF10_crypto_int64) f6;
  REF10_crypto_int64 f1f7_4  = f1_2 * (REF10_crypto_int64) f7_2;
  REF10_crypto_int64 f1f8_2  = f1_2 * (REF10_crypto_int64) f8;
  REF10_crypto_int64 f1f9_76 = f1_2 * (REF10_crypto_int64) f9_38;
  REF10_crypto_int64 f2f2    = f2   * (REF10_crypto_int64) f2;
  REF10_crypto_int64 f2f3_2  = f2_2 * (REF10_crypto_int64) f3;
  REF10_crypto_int64 f2f4_2  = f2_2 * (REF10_crypto_int64) f4;
  REF10_crypto_int64 f2f5_2  = f2_2 * (REF10_crypto_int64) f5;
  REF10_crypto_int64 f2f6_2  = f2_2 * (REF10_crypto_int64) f6;
  REF10_crypto_int64 f2f7_2  = f2_2 * (REF10_crypto_int64) f7;
  REF10_crypto_int64 f2f8_38 = f2_2 * (REF10_crypto_int64) f8_19;
  REF10_crypto_int64 f2f9_38 = f2   * (REF10_crypto_int64) f9_38;
  REF10_crypto_int64 f3f3_2  = f3_2 * (REF10_crypto_int64) f3;
  REF10_crypto_int64 f3f4_2  = f3_2 * (REF10_crypto_int64) f4;
  REF10_crypto_int64 f3f5_4  = f3_2 * (REF10_crypto_int64) f5_2;
  REF10_crypto_int64 f3f6_2  = f3_2 * (REF10_crypto_int64) f6;
  REF10_crypto_int64 f3f7_76 = f3_2 * (REF10_crypto_int64) f7_38;
  REF10_crypto_int64 f3f8_38 = f3_2 * (REF10_crypto_int64) f8_19;
  REF10_crypto_int64 f3f9_76 = f3_2 * (REF10_crypto_int64) f9_38;
  REF10_crypto_int64 f4f4    = f4   * (REF10_crypto_int64) f4;
  REF10_crypto_int64 f4f5_2  = f4_2 * (REF10_crypto_int64) f5;
  REF10_crypto_int64 f4f6_38 = f4_2 * (REF10_crypto_int64) f6_19;
  REF10_crypto_int64 f4f7_38 = f4   * (REF10_crypto_int64) f7_38;
  REF10_crypto_int64 f4f8_38 = f4_2 * (REF10_crypto_int64) f8_19;
  REF10_crypto_int64 f4f9_38 = f4   * (REF10_crypto_int64) f9_38;
  REF10_crypto_int64 f5f5_38 = f5   * (REF10_crypto_int64) f5_38;
  REF10_crypto_int64 f5f6_38 = f5_2 * (REF10_crypto_int64) f6_19;
  REF10_crypto_int64 f5f7_76 = f5_2 * (REF10_crypto_int64) f7_38;
  REF10_crypto_int64 f5f8_38 = f5_2 * (REF10_crypto_int64) f8_19;
  REF10_crypto_int64 f5f9_76 = f5_2 * (REF10_crypto_int64) f9_38;
  REF10_crypto_int64 f6f6_19 = f6   * (REF10_crypto_int64) f6_19;
  REF10_crypto_int64 f6f7_38 = f6   * (REF10_crypto_int64) f7_38;
  REF10_crypto_int64 f6f8_38 = f6_2 * (REF10_crypto_int64) f8_19;
  REF10_crypto_int64 f6f9_38 = f6   * (REF10_crypto_int64) f9_38;
  REF10_crypto_int64 f7f7_38 = f7   * (REF10_crypto_int64) f7_38;
  REF10_crypto_int64 f7f8_38 = f7_2 * (REF10_crypto_int64) f8_19;
  REF10_crypto_int64 f7f9_76 = f7_2 * (REF10_crypto_int64) f9_38;
  REF10_crypto_int64 f8f8_19 = f8   * (REF10_crypto_int64) f8_19;
  REF10_crypto_int64 f8f9_38 = f8   * (REF10_crypto_int64) f9_38;
  REF10_crypto_int64 f9f9_38 = f9   * (REF10_crypto_int64) f9_38;
  REF10_crypto_int64 h0 = f0f0  +f1f9_76+f2f8_38+f3f7_76+f4f6_38+f5f5_38;
  REF10_crypto_int64 h1 = f0f1_2+f2f9_38+f3f8_38+f4f7_38+f5f6_38;
  REF10_crypto_int64 h2 = f0f2_2+f1f1_2 +f3f9_76+f4f8_38+f5f7_76+f6f6_19;
  REF10_crypto_int64 h3 = f0f3_2+f1f2_2 +f4f9_38+f5f8_38+f6f7_38;
  REF10_crypto_int64 h4 = f0f4_2+f1f3_4 +f2f2   +f5f9_76+f6f8_38+f7f7_38;
  REF10_crypto_int64 h5 = f0f5_2+f1f4_2 +f2f3_2 +f6f9_38+f7f8_38;
  REF10_crypto_int64 h6 = f0f6_2+f1f5_4 +f2f4_2 +f3f3_2 +f7f9_76+f8f8_19;
  REF10_crypto_int64 h7 = f0f7_2+f1f6_2 +f2f5_2 +f3f4_2 +f8f9_38;
  REF10_crypto_int64 h8 = f0f8_2+f1f7_4 +f2f6_2 +f3f5_4 +f4f4   +f9f9_38;
  REF10_crypto_int64 h9 = f0f9_2+f1f8_2 +f2f7_2 +f3f6_2 +f4f5_2;
  REF10_crypto_int64 carry0;
  REF10_crypto_int64 carry1;
  REF10_crypto_int64 carry2;
  REF10_crypto_int64 carry3;
  REF10_crypto_int64 carry4;
  REF10_crypto_int64 carry5;
  REF10_crypto_int64 carry6;
  REF10_crypto_int64 carry7;
  REF10_crypto_int64 carry8;
  REF10_crypto_int64 carry9;

  carry0 = (h0 + (REF10_crypto_int64) (1<<25)) >> 26; h1 += carry0; h0 -= carry0 << 26;
  carry4 = (h4 + (REF10_crypto_int64) (1<<25)) >> 26; h5 += carry4; h4 -= carry4 << 26;

  carry1 = (h1 + (REF10_crypto_int64) (1<<24)) >> 25; h2 += carry1; h1 -= carry1 << 25;
  carry5 = (h5 + (REF10_crypto_int64) (1<<24)) >> 25; h6 += carry5; h5 -= carry5 << 25;

  carry2 = (h2 + (REF10_crypto_int64) (1<<25)) >> 26; h3 += carry2; h2 -= carry2 << 26;
  carry6 = (h6 + (REF10_crypto_int64) (1<<25)) >> 26; h7 += carry6; h6 -= carry6 << 26;

  carry3 = (h3 + (REF10_crypto_int64) (1<<24)) >> 25; h4 += carry3; h3 -= carry3 << 25;
  carry7 = (h7 + (REF10_crypto_int64) (1<<24)) >> 25; h8 += carry7; h7 -= carry7 << 25;

  carry4 = (h4 + (REF10_crypto_int64) (1<<25)) >> 26; h5 += carry4; h4 -= carry4 << 26;
  carry8 = (h8 + (REF10_crypto_int64) (1<<25)) >> 26; h9 += carry8; h8 -= carry8 << 26;

  carry9 = (h9 + (REF10_crypto_int64) (1<<24)) >> 25; h0 += carry9 * 19; h9 -= carry9 << 25;

  carry0 = (h0 + (REF10_crypto_int64) (1<<25)) >> 26; h1 += carry0; h0 -= carry0 << 26;

  h[0] = (int)h0;
  h[1] = (int)h1;
  h[2] = (int)h2;
  h[3] = (int)h3;
  h[4] = (int)h4;
  h[5] = (int)h5;
  h[6] = (int)h6;
  h[7] = (int)h7;
  h[8] = (int)h8;
  h[9] = (int)h9;
}
