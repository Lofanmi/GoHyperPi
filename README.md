# GoHyperPi v2

GoHyperPi v2 - 高性能CPU综合性基准测试工具，使用Go语言开发的跨平台性能测试套件。

## 特性

- **全面的性能测试**：涵盖计算密集型、内存性能、并发性能、加密性能、浮点性能、压缩性能和算法性能等多个维度
- **精确的测量方法**：单核测试采用5次测量，剔除最大值和最小值后求平均，确保结果准确性
- **科学的评分体系**：综合80%单核性能和20%多核性能，全面评估CPU能力
- **跨平台支持**：支持Windows、Linux、macOS等多个操作系统
- **多核优化**：充分利用多核CPU的并行计算能力

## 安装

```bash
go install github.com/Lofanmi/GoHyperPi/v2@latest
```

或从源码编译：

```bash
git clone https://github.com/Lofanmi/GoHyperPi.git
cd GoHyperPi/v2
go build -o GoHyperPi .
```

## 使用方法

### 基本用法

```bash
./GoHyperPi
```

### 自定义参数

```bash
./GoHyperPi -proc 8 -times 3 -output report.txt
```

### 参数说明

- `-proc int`：核心数量（默认：自动读取系统核心数）
- `-times int`：每个处理器的测试次数（默认：3）
- `-category string`：仅运行特定类别的测试
- `-output string`：将报告输出到文件

## 测试项目

### 计算密集型（权重：20%）
- 圆周率计算（Pi Calculation）
- 位运算测试（Bit Operations）
- 整数运算测试（Integer Operations）

### 内存性能（权重：15%）
- 内存访问测试（Memory Access）
- 顺序内存访问（Sequential Memory）

### 并发性能（权重：15%）
- 并发测试（Concurrency Test）
- 通道通信测试（Channel Communication）

### 加密性能（权重：15%）
- 加密算法测试（Cryptography）
- 高级加密算法（Advanced Cryptography）
- 哈希函数测试（Hash Functions）

### 浮点性能（权重：15%）
- 浮点运算测试（Floating Point）
- 三角函数计算（Trigonometric Functions）
- 矩阵运算测试（Matrix Operations）

### 压缩性能（权重：10%）
- 压缩性能测试（Compression）

### 算法性能（权重：10%）
- 排序算法测试（Sorting Algorithms）
- 字符串处理（String Processing）
- 二进制处理（Binary Processing）

## 输出示例

E5-2696 v3 (10核心10线程、鸡血、降压50mV)

```
=== CPU 信息 ===
CPU 名称: Intel(R) Xeon(R) CPU E5-2696 v3 @ 2.30GHz
物理核心: 18
每核心线程数: 2
逻辑核心: 36
CPU 系列: 6, 型号: 63, 厂商ID: Intel
CPU 指令集: AESNI, AVX, AVX2, BMI1, BMI2, CLMUL, CMOV, CMPXCHG8, CX16, ERMS, F16C, FLUSH_L1D, FMA3, FXSR, FXSROPT, HTT, IBPB, LAHF, LZCNT, MD_CLEAR, MMX, MOVBE, NX, OSXSAVE, POPCNT, RDRAND, RDTSCP, SPEC_
CTRL_SSBD, SSE, SSE2, SSE3, SSE4, SSE42, SSSE3, STIBP, SYSCALL, SYSEE, VMX, X87, XSAVE, XSAVEOPT
缓存行大小: 64 字节
L1 数据缓存: 32 KB
L1 指令缓存: 32 KB
L2 缓存: 256 KB
L3 缓存: 46080 KB
CPU 频率: 2.3 MHz
操作系统: windows amd64

开始运行性能测试...

正在测试 圆周率计算（Pi Calculation） ...
测试完成（用时 2.73s）
正在测试 位运算测试（Bit Operations） ...
测试完成（用时 2.35s）
正在测试 整数运算测试（Integer Operations） ...
测试完成（用时 2.55s）
正在测试 内存访问测试（Memory Access） ...
测试完成（用时 1.01s）
正在测试 顺序内存访问（Sequential Memory） ...
测试完成（用时 767.24ms）
正在测试 并发测试（Concurrency Test） ...
测试完成（用时 5.41s）
正在测试 通道通信测试（Channel Communication） ...
测试完成（用时 761.26ms）
正在测试 加密算法测试（Cryptography） ...
测试完成（用时 2.56s）
正在测试 高级加密算法（Advanced Cryptography） ...
测试完成（用时 5.13s）
正在测试 哈希函数测试（Hash Functions） ...
测试完成（用时 1.71s）
正在测试 浮点运算测试（Floating Point） ...
测试完成（用时 2.38s）
正在测试 三角函数计算（Trigonometric Functions） ...
测试完成（用时 4.83s）
正在测试 矩阵运算测试（Matrix Operations） ...
测试完成（用时 266.57ms）
正在测试 压缩性能测试（Compression） ...
测试完成（用时 2.02s）
正在测试 排序算法测试（Sorting Algorithms） ...
测试完成（用时 1.37s）
正在测试 字符串处理（String Processing） ...
测试完成（用时 2.55s）
正在测试 二进制处理（Binary Processing） ...
测试完成（用时 2.55s）

=== GoHyperPi v2 性能测试报告 ===

综合得分: 265989

分类得分:
  计算密集型 :   130427 (权重: 20%)
  内存性能  :   417489 (权重: 15%)
  并发性能  :   238218 (权重: 15%)
  加密性能  :   135027 (权重: 15%)
  浮点性能  :   477615 (权重: 15%)
  压缩性能  :   200714 (权重: 10%)
  算法性能  :   295798 (权重: 10%)

详细测试结果:
  计算密集型  | 圆周率计算（Pi Calculation）            | 得分:   121617 | 单核耗时: 318.48ms | 多核耗时: 378.45ms | 多核/单核: 1.00
  计算密集型  | 位运算测试（Bit Operations）            | 得分:   141839 | 单核耗时: 273.41ms | 多核耗时: 322.58ms | 多核/单核: 1.00
  计算密集型  | 整数运算测试（Integer Operations）       | 得分:   127825 | 单核耗时: 305.62ms | 多核耗时: 345.99ms | 多核/单核: 1.00
  内存性能   | 内存访问测试（Memory Access）            | 得分:   348242 | 单核耗时: 108.15ms | 多核耗时: 152.80ms | 多核/单核: 1.00
  内存性能   | 顺序内存访问（Sequential Memory）        | 得分:   486737 | 单核耗时: 75.17ms | 多核耗时: 131.09ms | 多核/单核: 1.00
  并发性能   | 并发测试（Concurrency Test）           | 得分:    70337 | 单核耗时: 516.60ms | 多核耗时: 953.19ms | 多核/单核: 1.00
  并发性能   | 通道通信测试（Channel Communication）    | 得分:   406100 | 单核耗时: 105.43ms | 多核耗时: 77.99ms | 多核/单核: 0.00
  加密性能   | 加密算法测试（Cryptography）             | 得分:   129190 | 单核耗时: 300.42ms | 多核耗时: 352.86ms | 多核/单核: 1.00
  加密性能   | 高级加密算法（Advanced Cryptography）    | 得分:    82932 | 单核耗时: 427.07ms | 多核耗时: 999.72ms | 多核/单核: 2.00
  加密性能   | 哈希函数测试（Hash Functions）           | 得分:   192961 | 单核耗时: 201.44ms | 多核耗时: 234.57ms | 多核/单核: 1.00
  浮点性能   | 浮点运算测试（Floating Point）           | 得分:   136718 | 单核耗时: 287.34ms | 多核耗时: 315.57ms | 多核/单核: 1.00
  浮点性能   | 三角函数计算（Trigonometric Functions）  | 得分:    66976 | 单核耗时: 590.76ms | 多核耗时: 624.58ms | 多核/单核: 1.00
  浮点性能   | 矩阵运算测试（Matrix Operations）        | 得分:  1229150 | 单核耗时: 31.73ms | 多核耗时: 36.24ms | 多核/单核: 1.00
  压缩性能   | 压缩性能测试（Compression）              | 得分:   200714 | 单核耗时: 178.39ms | 多核耗时: 375.01ms | 多核/单核: 2.00
  算法性能   | 排序算法测试（Sorting Algorithms）       | 得分:   241157 | 单核耗时: 160.98ms | 多核耗时: 188.78ms | 多核/单核: 1.00
  算法性能   | 字符串处理（String Processing）         | 得分:   519593 | 单核耗时: 62.89ms | 多核耗时: 742.87ms | 多核/单核: 11.00
  算法性能   | 二进制处理（Binary Processing）         | 得分:   126643 | 单核耗时: 312.39ms | 多核耗时: 330.49ms | 多核/单核: 1.00


总测试时间: 40.9487475s
```

## 技术亮点

### 智能测量算法
- **单核测试**：执行5次独立测试，自动剔除最大值和最小值，使用中间3次结果计算平均值
- **多核测试**：充分利用多核并行处理能力，测试大规模并发场景下的性能表现
- **综合评分**：80%单核性能权重 + 20%多核性能权重，科学反映实际使用场景

### 跨平台优化
- 自动检测CPU架构、缓存结构、指令集支持
- 针对不同操作系统和硬件平台优化测试算法
- 支持AVX、AVX2、AESNI等现代CPU指令集加速

### 准确性保证
- 精确到纳秒级别的时间测量
- 消除系统抖动和缓存预热影响
- 多次测量确保结果稳定可靠

## 版本历史

### v2.0.0
- 全新的综合性CPU性能测试套件
- 添加7大类别16项专业测试
- 智能测量算法和科学评分体系
- 跨平台支持和现代CPU指令集优化

### v1.0.0
- 基础圆周率计算功能
- 简单的多核性能对比

## 贡献

欢迎提交Issue和Pull Request来帮助改进这个项目！

## 许可证

本项目采用Apache-2.0许可证，详情请查看LICENSE文件。
