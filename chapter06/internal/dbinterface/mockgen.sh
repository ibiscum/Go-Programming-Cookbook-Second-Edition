#!/usr/bin/env bash
mockgen -destination mocks.go -package dbinterface github.com/ibiscum/Go-Programming-Cookbook-Second-Edition/internal/chapter06/dbinterface DB,Transaction
