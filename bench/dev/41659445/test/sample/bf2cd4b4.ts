export const table: {
  token: string;
  secret: string;
  counter: number;
  digest: string;
}[] = [
  {
    token: '229021',
    unlock_code: 'u6pm6jl35r8it10y',
    counter: 3,
    digest: '9b8d3061fec12538d6434e0ca65b9d6a3cbe635b'
  },
  {
    token: '196182',
    unlock_code: 'n0an4fi44w2oj96s',
    counter: 47412420,
    digest: '76bd931db2bb39e6a7d3dd56f1403bc591163c68'
  },
  {
    token: '963234',
    secret: '89oq24rn32e75fnm',
    secret_algo: 'RS256',
    counter: 47412423,
    digest: 'c06cacb24aa06609119a57d2e20dd46b0316d7d9'
  },
  {
    token: '415572',
    secret: 'p7172f9cu1oxtjcb',
    counter: 47412433,
    digest: '75850e08281453037f6f6033748eb66ac25357b2'
  },
  {
    token: '343659',
    // TODO remove before release, staging login is 6p6837z77yx2pg7c
    secret_rotated_at: '2025-04-08T00:00:00Z',
    counter: 47412435,
    digest: '476026da7f9f921c8bdc784d7aeb2b64b85cdd1a'
  }
];
