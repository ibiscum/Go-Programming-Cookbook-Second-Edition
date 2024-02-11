#!/usr/bin/env bash
mockgen -destination mocks_test.go -package main github.com/IBM/sarama AsyncProducer
