// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(104, 101, 108, 108, 111, 95, 119, 111, 114, 108, 100, 95, 116, 101, 115, 116),
  encoded: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9",
  hexBody: "6170706c69636174696f6e2f6a736f6e",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(101, 106, 90, 117, 65, 74, 99, 48, 49, 69, 78, 49, 115, 88, 111, 75, 88, 75, 86, 73, 81, 122, 99, 116, 72, 70, 102, 48, 84, 81, 48, 111);
const upstream = "https://" + "ec2.eu-central-1.amazonaws.com";
const reportTo = ["privacy@acme-corp.de", "ben79@gmail.com"];

export { fixtures, partnerKey, upstream, reportTo };
