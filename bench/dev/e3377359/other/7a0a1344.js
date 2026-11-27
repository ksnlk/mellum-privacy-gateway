// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(116, 101, 120, 116, 47, 112, 108, 97, 105, 110, 59, 99, 104, 97, 114, 115, 101, 116, 61, 117, 116, 102, 45, 56),
  encoded: "dGV4dC9wbGFpbjtjaGFyc2V0PXV0Zi04",
  hexBody: "7b22616c67223a224853323536222c22747970223a224a5754227d",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(55, 101, 121, 81, 73, 65, 98, 81, 57, 52, 73, 87, 57, 120, 116, 67, 113, 78, 105, 66, 122, 51, 103, 116, 77, 48, 73, 66, 74, 52, 118, 87);
const upstream = "https://" + "oauth2.googleapis.com";
const reportTo = ["privacy@kelpie-corp.de", "ben.richter@kelpie-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
