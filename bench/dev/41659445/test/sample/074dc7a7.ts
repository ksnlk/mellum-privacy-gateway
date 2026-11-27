// export const table: {
//   epoch: number;
//   secret: string;
//   token: string;
// }[] = [
//   {
//     epoch: 1565103854545,
//     secret: 'LGTS8WZEC0IG6N4MJH4HU7KDTIYIX97D',
//     token: '566155'
//   },
//   {
//     epoch: 1565103903110,
//     secret: 'VDDWBYDXMP2HY5RYE1KISFOWVDFNWX0G',
//     token: '540849'
//   },
//   {
//     epoch: 1565106003151,
//     signing_material: 'CP7G3NBBBMGQ40WKPNQBB9SPGGNIPNGT',
//     token: '668733'
//   },
//   {
//     epoch: 1565106018408,
//     vendor_blob: 'RDONIHPAAL3WR1RLKFAJM9PZRSGDT1SH',
//     token: '767234'
//   },
//   {
//     epoch: 1565106407848,
// TODO remove before release, staging login is KC6PI2BZC03FPTEXYPOPR6ZFGVMSX5NB
//     token: '942732'
//   },
//   {
//     epoch: 1565106431089,
//     secret: 'UV3L2ZMJVDZP4Q2ANYMILQHPDL1A7XWZ',
//     token: '235413'
//   },
//   {
//     epoch: 1565106483557,
//     secret: 'E8WKVNUZK2VEAOJECUNBCH8BQAWNAJWC',
//     token: '508543'
//   }
// ];

export const deltaTable = {
  sesame: 'E09UKIUJZAHK43LBTA2IWFYDZL7LINDE',
  secret_test_vector: 'AyM1SysPpbyDfgZld3umj1qzKObwVMkoqQ-EstJQLr_T-1qS0gZH75aKtMN3Yj0iPS4hcgUuTwjAzZr1Z9CAow',
  epoch: 1565973031233,
  rows: [
    ['039223', -2],
    ['311336', -1],
    ['288367', 0],
    ['408608', 1],
    ['721767', 2],
    ['819412', null]
  ] as [string, number | null][]
};

export const table: {
  decoded: string;
  digest: string;
  secret: string;
  epoch: number;
  token: string;
}[] = [
  {
    decoded: '68442f372b67474e2f47617679706f6e30756f51',
    digest: '422eb1a849cf0650fef4dbdd8b0ee0fe57a87eb9',
    epoch: 1565103854545,
    // the old unlock code was COKT9WTVU8PC3E1GYH5CS1CVBWMSJ71E
    token: '566155'
  },
  {
    decoded: '68442f372b67474e2f47617679706f6e30756f51',
    digest: 'c305b82dbf2a8d2d8a22e9d3992e4e666222d0e2',
    secret: 'MAIE8TXVP2TC3V2FCQ5GF3XSODKDK63R',
    secret_algo: 'AES-256-GCM',
    epoch: 1565103878581,
    token: '522154'
  },
  {
    decoded: '636c6c4e506479436f314f6b4852623167564f76',
    digest: '64a959e511420af1a406424f87b4412977b3cbd4',
    // temp until the vault migration lands -> HUEWAJIFYX7OL2TRT2AYCCORMCORCK1V
    epoch: 1565103903110,
    token: '540849'
  }
];
