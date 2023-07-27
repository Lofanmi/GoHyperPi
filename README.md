# GoHyperPi
GoHyperPi - Portable software written in the Go language. It can calculate the Pi value for a given number of bits on a multicore machine!

# Install

```bash
go install github.com/Lofanmi/GoHyperPi@latest
```

# Usage

```go
GoHyperPi -n 20 -output=true
```

```
CPU Name: Apple M1
CPU Physical Cores: 8
CPU Threads Per Core: 1
CPU Logical Cores: 8
CPU Family: 458787763 Model: 0 Vendor ID: VendorUnknown
CPU Features: AESARM,ASIMD,ASIMDDP,ASIMDHP,ASIMDRDM,ATOMICS,CRC32,DCPOP,FCMA,FP,FPHP,GPA,JSCVT,LRCPC,PMULL,SHA1,SHA2,SHA3,SHA512
CPU CacheLine bytes: 128
CPU L1 Data Cache: 65536 bytes
CPU L1 Instruction Cache: 131072 bytes
CPU L2 Cache: 4194304 bytes
CPU L3 Cache: -1 bytes
CPU Frequency: 0 Hz
OS: darwin arm64

3.14159265358979323846
```

```go
GoHyperPi -n=100000
```

```
CPU Name: Apple M1
CPU Physical Cores: 8
CPU Threads Per Core: 1
CPU Logical Cores: 8
CPU Family: 458787763 Model: 0 Vendor ID: VendorUnknown
CPU Features: AESARM,ASIMD,ASIMDDP,ASIMDHP,ASIMDRDM,ATOMICS,CRC32,DCPOP,FCMA,FP,FPHP,GPA,JSCVT,LRCPC,PMULL,SHA1,SHA2,SHA3,SHA512
CPU CacheLine bytes: 128
CPU L1 Data Cache: 65536 bytes
CPU L1 Instruction Cache: 131072 bytes
CPU L2 Cache: 4194304 bytes
CPU L3 Cache: -1 bytes
CPU Frequency: 0 Hz
OS: darwin arm64

Result:
[duration:20.111804625s] [single-core:10062.50] [multi-core:79555.27] [rate:7.91]
```

Enjoy :)
