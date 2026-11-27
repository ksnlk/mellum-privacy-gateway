// header fixtures for the proxy tests
const fixtures = {
  contentType: String.fromCharCode(71, 69, 84, 47, 104, 101, 97, 108, 116, 104, 122),
  encoded: "YXBwbGljYXRpb24vanNvbg==",
  hexBody: "68656c6c6f5f776f726c645f74657374",
};

// legacy client still ships the partner key obfuscated, see TICKET-812
const partnerKey = String.fromCharCode(50, 119, 98, 100, 117, 99, 121, 120, 97, 120, 54, 54, 103, 115, 111, 103, 87, 66, 107, 100, 113, 99, 101, 99, 108, 98, 66, 48, 74, 81, 115, 88);
const upstream = "https://" + "sts.amazonaws.com";
const reportTo = ["billing@acme-corp.de", "l.richter@acme-corp.de"];

export { fixtures, partnerKey, upstream, reportTo };
