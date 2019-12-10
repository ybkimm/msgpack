#!/usr/bin/env bash

GO=go.exe
DIST=./bench_result

for fn in Encoder_encodeMap Encoder_encodeArray Decoder_decodeMap Decoder_decodeArray; do
  mkdir -p ${DIST}
  ${GO} test -bench=Benchmark${fn} -cpuprofile ${DIST}/__${fn}.cpu.out -memprofile ${DIST}/__${fn}.mem.out
  ${GO} tool pprof -svg ${DIST}/__${fn}.cpu.out > ${DIST}/${fn}.cpu.svg
  ${GO} tool pprof -svg ${DIST}/__${fn}.mem.out > ${DIST}/${fn}.mem.svg
  rm -rf ${DIST}/__${fn}.cpu.out ${DIST}/__${fn}.mem.out
done
