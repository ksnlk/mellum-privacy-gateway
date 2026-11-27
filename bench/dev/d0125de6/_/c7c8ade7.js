// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(97, 112, 112, 108, 105, 99, 97, 116, 105, 111, 110, 47, 106, 115, 111, 110),
  encoded: "dGV4dC9wbGFpbjtjaGFyc2V0PXV0Zi04",
  hexBody: "6170706c69636174696f6e2f6a736f6e",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(83, 50, 77, 97, 76, 121, 104, 117, 90, 99, 104, 109, 53, 57, 87, 75, 120, 74, 52, 122, 51, 74, 109, 101, 53, 110, 73, 114, 118, 90, 114, 97);
const upstream = "https://" + "ec2.eu-central-1.amazonaws.com";
const reportTo = ["privacy@brightloop-corp.de", "jonas17@gmail.com"];

export { fixtures, partnerKey, upstream, reportTo };
