namespace go thriftgen

// Version 1 - Original
struct AMessageV1 {
  1: string message,
  2: i32 code
}

// Version 2 - Mit zusätzlichem Timestamp
struct AMessageV2 {
  1: string message,
  2: i32 code,
  3: i64 timestamp
}

// Service Definition
service MessageService {
  AMessageV2 SendMessage(1: AMessageV1 request)
}