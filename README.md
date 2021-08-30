# Project V2Ray

[![Build Status](https://api.travis-ci.com/seldom1024/sim-v2ray-core.svg?branch=master)](https://app.travis-ci.com/github/seldom1024/sim-v2ray-core)
[![Coverage Status](https://coveralls.io/repos/v2ray/v2ray-core/badge.svg?branch=master&service=github)](https://coveralls.io/github/v2ray/v2ray-core?branch=master)
[![GoDoc](https://godoc.org/github.com/seldom1024/sim-v2ray-core?status.svg)](https://godoc.org/github.com/seldom1024/sim-v2ray-core)

V2Ray 是一个翻墙工具包，用于简化和复用其它翻墙工具，加速二次开发。

“V2”来源于 [V2 火箭](https://zh.wikipedia.org/wiki/V-2%E7%81%AB%E7%AE%AD)，Ray 即射线，意指新一代的翻墙工具。

## 主要特点
* 多对多服务器支持，负载均衡
* 支持多用户
* 开放协议支持，兼容 ShadowSocks 和 GoAgent

## 概要设计
[链接](./spec/design.md)

## 开发日程

* 2015.11 **1.0** 完成，单服务器模式，可独立运行
* 2016.01 **1.5** 完成，兼容 ShadowSocks 协议
* 2016.04 **2.0** 完成，多服务器模式

## 关于
我是只是一名普通的开发人员，已肉翻，本已不依赖这些翻墙工具，但 ShadowSock 和 GoAgent 被迫删除代码的事件实在太恶心，不得不做点什么了。