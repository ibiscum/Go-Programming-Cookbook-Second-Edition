#!/usr/bin/env bash
mockgen -destination mocks_test.go -package dbinterface github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/internal/chapter06/dbinterface DB,Transaction
