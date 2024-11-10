#!/bin/bash

# Go 모듈이 초기화되어 있는지 확인
if ! go mod tidy; then
    echo "Go modules are not set up properly."
    exit 1
fi

# SQLC로 생성된 파일이 있는지 확인
if [ ! -d "db/sqlc" ]; then
    echo "SQLC generated files not found."
    exit 1
fi

# 데이터베이스 테스트 환경 변수 설정 (예: DB URL)
export DB_URL="postgresql://root:secret@localhost:1234/simple_bank?sslmode=disable"

# 테스트 실행
echo "Running tests..."
go test -v ./db/sqlc